package cache

import (
	gwv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"

	"github.com/flomesh-io/fsm/pkg/gateway/routecfg"
	gwtypes "github.com/flomesh-io/fsm/pkg/gateway/types"
	gwutils "github.com/flomesh-io/fsm/pkg/gateway/utils"
)

func processHTTPRoute(gw *gwv1beta1.Gateway, validListeners []gwtypes.Listener, httpRoute *gwv1beta1.HTTPRoute, policies globalPolicyAttachments, rules map[int32]routecfg.RouteRule, services map[string]serviceInfo) {
	routePolicies := filterPoliciesByRoute(policies, httpRoute)
	log.Debug().Msgf("[GW-CACHE] routePolicies: %v", routePolicies)
	hostnameEnrichers := getHostnamePolicyEnrichers(routePolicies)

	for _, ref := range httpRoute.Spec.ParentRefs {
		if !gwutils.IsRefToGateway(ref, gwutils.ObjectKey(gw)) {
			continue
		}

		allowedListeners := allowedListeners(ref, httpRoute.GroupVersionKind(), validListeners)
		log.Debug().Msgf("allowedListeners: %v", allowedListeners)
		if len(allowedListeners) == 0 {
			continue
		}

		for _, listener := range allowedListeners {
			hostnames := gwutils.GetValidHostnames(listener.Hostname, httpRoute.Spec.Hostnames)
			log.Debug().Msgf("hostnames: %v", hostnames)

			if len(hostnames) == 0 {
				// no valid hostnames, should ignore it
				continue
			}

			httpRule := routecfg.L7RouteRule{}
			for _, hostname := range hostnames {
				r := generateHTTPRouteConfig(httpRoute, routePolicies, services)

				for _, enricher := range hostnameEnrichers {
					enricher.Enrich(hostname, r)
				}

				httpRule[hostname] = r
			}

			port := int32(listener.Port)
			if rule, exists := rules[port]; exists {
				if l7Rule, ok := rule.(routecfg.L7RouteRule); ok {
					rules[port] = mergeL7RouteRule(l7Rule, httpRule)
				}
			} else {
				rules[port] = httpRule
			}
		}
	}
}

func generateHTTPRouteConfig(httpRoute *gwv1beta1.HTTPRoute, routePolicies routePolicies, services map[string]serviceInfo) *routecfg.HTTPRouteRuleSpec {
	httpSpec := &routecfg.HTTPRouteRuleSpec{
		RouteType: routecfg.L7RouteTypeHTTP,
		Matches:   make([]routecfg.HTTPTrafficMatch, 0),
	}
	enrichers := getHTTPRoutePolicyEnrichers(routePolicies)

	for _, rule := range httpRoute.Spec.Rules {
		backends := map[string]routecfg.BackendServiceConfig{}

		for _, bk := range rule.BackendRefs {
			if svcPort := backendRefToServicePortName(bk.BackendRef.BackendObjectReference, httpRoute.Namespace); svcPort != nil {
				svcLevelFilters := make([]routecfg.Filter, 0)
				for _, filter := range bk.Filters {
					svcLevelFilters = append(svcLevelFilters, toFSMHTTPRouteFilter(filter, httpRoute.Namespace, services))
				}

				backends[svcPort.String()] = routecfg.BackendServiceConfig{
					Weight:  backendWeight(bk.BackendRef),
					Filters: svcLevelFilters,
				}

				services[svcPort.String()] = serviceInfo{
					svcPortName: *svcPort,
				}
			}
		}

		ruleLevelFilters := make([]routecfg.Filter, 0)
		for _, ruleFilter := range rule.Filters {
			ruleLevelFilters = append(ruleLevelFilters, toFSMHTTPRouteFilter(ruleFilter, httpRoute.Namespace, services))
		}

		for _, m := range rule.Matches {
			match := &routecfg.HTTPTrafficMatch{
				BackendService: backends,
				Filters:        ruleLevelFilters,
			}

			if m.Path != nil {
				match.Path = &routecfg.Path{
					MatchType: httpPathMatchType(m.Path.Type),
					Path:      httpPath(m.Path.Value),
				}
			}

			if m.Method != nil {
				match.Methods = []string{string(*m.Method)}
			}

			if len(m.Headers) > 0 {
				match.Headers = httpMatchHeaders(m)
			}

			if len(m.QueryParams) > 0 {
				match.RequestParams = httpMatchQueryParams(m)
			}

			for _, enricher := range enrichers {
				enricher.Enrich(m, match)
			}

			httpSpec.Matches = append(httpSpec.Matches, *match)
		}
	}
	return httpSpec
}