package cache

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
	gwv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"

	"github.com/flomesh-io/fsm/pkg/constants"
	"github.com/flomesh-io/fsm/pkg/gateway/routecfg"
	gwtypes "github.com/flomesh-io/fsm/pkg/gateway/types"
	"github.com/flomesh-io/fsm/pkg/utils"
)

func getSecretRefNamespace(gw *gwv1beta1.Gateway, secretRef gwv1beta1.SecretObjectReference) string {
	if secretRef.Namespace == nil {
		return gw.Namespace
	}

	return string(*secretRef.Namespace)
}

func generateHTTPRouteConfig(httpRoute *gwv1beta1.HTTPRoute) routecfg.HTTPRouteRuleSpec {
	httpSpec := routecfg.HTTPRouteRuleSpec{
		RouteType: routecfg.L7RouteTypeHTTP,
		Matches:   make([]routecfg.HTTPTrafficMatch, 0),
	}

	for _, rule := range httpRoute.Spec.Rules {
		backends := map[string]int32{}

		for _, bk := range rule.BackendRefs {
			if svcPort := backendRefToServicePortName(bk.BackendRef, httpRoute.Namespace); svcPort != nil {
				backends[svcPort.String()] = backendWeight(bk.BackendRef)
			}
		}

		for _, m := range rule.Matches {
			match := routecfg.HTTPTrafficMatch{
				BackendService: backends,
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

			httpSpec.Matches = append(httpSpec.Matches, match)
		}
	}
	return httpSpec
}

func httpPathMatchType(matchType *gwv1beta1.PathMatchType) routecfg.MatchType {
	if matchType == nil {
		return routecfg.MatchTypePrefix
	}

	switch *matchType {
	case gwv1beta1.PathMatchPathPrefix:
		return routecfg.MatchTypePrefix
	case gwv1beta1.PathMatchExact:
		return routecfg.MatchTypeExact
	case gwv1beta1.PathMatchRegularExpression:
		return routecfg.MatchTypeRegex
	default:
		return routecfg.MatchTypePrefix
	}
}

func httpPath(value *string) string {
	if value == nil {
		return "/"
	}

	return *value
}

func httpMatchHeaders(m gwv1beta1.HTTPRouteMatch) map[routecfg.MatchType]map[string]string {
	exact := make(map[string]string)
	regex := make(map[string]string)

	for _, header := range m.Headers {
		if header.Type == nil {
			exact[string(header.Name)] = header.Value
			continue
		}

		switch *header.Type {
		case gwv1beta1.HeaderMatchExact:
			exact[string(header.Name)] = header.Value
		case gwv1beta1.HeaderMatchRegularExpression:
			regex[string(header.Name)] = header.Value
		}
	}

	headers := make(map[routecfg.MatchType]map[string]string)

	if len(exact) > 0 {
		headers[routecfg.MatchTypeExact] = exact
	}

	if len(regex) > 0 {
		headers[routecfg.MatchTypeRegex] = regex
	}

	return headers
}

func httpMatchQueryParams(m gwv1beta1.HTTPRouteMatch) map[routecfg.MatchType]map[string]string {
	exact := make(map[string]string)
	regex := make(map[string]string)

	for _, param := range m.QueryParams {
		if param.Type == nil {
			exact[string(param.Name)] = param.Value
			continue
		}

		switch *param.Type {
		case gwv1beta1.QueryParamMatchExact:
			exact[string(param.Name)] = param.Value
		case gwv1beta1.QueryParamMatchRegularExpression:
			regex[string(param.Name)] = param.Value
		}
	}

	params := make(map[routecfg.MatchType]map[string]string)
	if len(exact) > 0 {
		params[routecfg.MatchTypeExact] = exact
	}

	if len(regex) > 0 {
		params[routecfg.MatchTypeRegex] = regex
	}

	return params
}

func generateGRPCRouteCfg(grpcRoute *gwv1alpha2.GRPCRoute) routecfg.GRPCRouteRuleSpec {
	grpcSpec := routecfg.GRPCRouteRuleSpec{
		RouteType: routecfg.L7RouteTypeGRPC,
		Matches:   make([]routecfg.GRPCTrafficMatch, 0),
	}

	for _, rule := range grpcRoute.Spec.Rules {
		backends := map[string]int32{}

		for _, bk := range rule.BackendRefs {
			if svcPort := backendRefToServicePortName(bk.BackendRef, grpcRoute.Namespace); svcPort != nil {
				backends[svcPort.String()] = backendWeight(bk.BackendRef)
			}
		}

		for _, m := range rule.Matches {
			match := routecfg.GRPCTrafficMatch{
				BackendService: backends,
			}

			if m.Method != nil {
				match.Method = &routecfg.GRPCMethod{
					MatchType: grpcMethodMatchType(m.Method.Type),
					Service:   m.Method.Service,
					Method:    m.Method.Method,
				}
			}

			if len(m.Headers) > 0 {
				match.Headers = grpcMatchHeaders(m)
			}

			grpcSpec.Matches = append(grpcSpec.Matches, match)
		}
	}

	return grpcSpec
}

func grpcMethodMatchType(matchType *gwv1alpha2.GRPCMethodMatchType) routecfg.MatchType {
	if matchType == nil {
		return routecfg.MatchTypeExact
	}

	switch *matchType {
	case gwv1alpha2.GRPCMethodMatchExact:
		return routecfg.MatchTypeExact
	case gwv1alpha2.GRPCMethodMatchRegularExpression:
		return routecfg.MatchTypeRegex
	default:
		return routecfg.MatchTypeExact
	}
}

func grpcMatchHeaders(m gwv1alpha2.GRPCRouteMatch) map[routecfg.MatchType]map[string]string {
	exact := make(map[string]string)
	regex := make(map[string]string)

	for _, header := range m.Headers {
		if header.Type == nil {
			exact[string(header.Name)] = header.Value
			continue
		}

		switch *header.Type {
		case gwv1beta1.HeaderMatchExact:
			exact[string(header.Name)] = header.Value
		case gwv1beta1.HeaderMatchRegularExpression:
			regex[string(header.Name)] = header.Value
		}
	}

	headers := make(map[routecfg.MatchType]map[string]string)

	if len(exact) > 0 {
		headers[routecfg.MatchTypeExact] = exact
	}

	if len(regex) > 0 {
		headers[routecfg.MatchTypeRegex] = regex
	}

	return headers
}

func generateTLSTerminateRouteCfg(tcpRoute *gwv1alpha2.TCPRoute) routecfg.TLSBackendService {
	backends := routecfg.TLSBackendService{}

	for _, rule := range tcpRoute.Spec.Rules {
		for _, bk := range rule.BackendRefs {
			if svcPort := backendRefToServicePortName(bk, tcpRoute.Namespace); svcPort != nil {
				backends[svcPort.String()] = backendWeight(bk)
			}
		}
	}

	return backends
}

func generateTLSPassthroughRouteCfg(tlsRoute *gwv1alpha2.TLSRoute) *string {
	for _, rule := range tlsRoute.Spec.Rules {
		for _, bk := range rule.BackendRefs {
			// return the first ONE
			return passthroughTarget(bk)
		}
	}

	return nil
}

func generateTCPRouteCfg(tcpRoute *gwv1alpha2.TCPRoute) routecfg.RouteRule {
	backends := routecfg.TCPRouteRule{}

	for _, rule := range tcpRoute.Spec.Rules {
		for _, bk := range rule.BackendRefs {
			if svcPort := backendRefToServicePortName(bk, tcpRoute.Namespace); svcPort != nil {
				backends[svcPort.String()] = backendWeight(bk)
			}
		}
	}

	return backends
}

func allowedListeners(
	parentRef gwv1beta1.ParentReference,
	routeGvk schema.GroupVersionKind,
	validListeners []gwtypes.Listener,
) []gwtypes.Listener {
	var selectedListeners []gwtypes.Listener
	for _, validListener := range validListeners {
		if (parentRef.SectionName == nil || *parentRef.SectionName == validListener.Name) &&
			(parentRef.Port == nil || *parentRef.Port == validListener.Port) {
			selectedListeners = append(selectedListeners, validListener)
		}
	}
	log.Debug().Msgf("[GW-CACHE] selectedListeners: %v", selectedListeners)

	if len(selectedListeners) == 0 {
		return nil
	}

	allowedListeners := make([]gwtypes.Listener, 0)
	for _, selectedListener := range selectedListeners {
		if !selectedListener.AllowsKind(routeGvk) {
			continue
		}

		allowedListeners = append(allowedListeners, selectedListener)
	}

	if len(allowedListeners) == 0 {
		return nil
	}

	return allowedListeners
}

func backendRefToServicePortName(ref gwv1beta1.BackendRef, defaultNs string) *routecfg.ServicePortName {
	// ONLY supports Service and ServiceImport backend now
	if (*ref.Kind == KindService && *ref.Group == GroupCore) || (*ref.Kind == KindServiceImport && *ref.Group == GroupFlomeshIo) {
		ns := defaultNs
		if ref.Namespace != nil {
			ns = string(*ref.Namespace)
		}

		return &routecfg.ServicePortName{
			NamespacedName: types.NamespacedName{
				Namespace: ns,
				Name:      string(ref.Name),
			},
			Port: pointer.Int32(int32(*ref.Port)),
		}
	}

	return nil
}

func passthroughTarget(ref gwv1beta1.BackendRef) *string {
	// ONLY supports service backend now
	if *ref.Kind == KindService && *ref.Group == GroupCore {
		port := int32(443)
		if ref.Port != nil {
			port = int32(*ref.Port)
		}

		target := fmt.Sprintf("%s:%d", ref.Name, port)

		return &target
	}

	return nil
}

func backendWeight(bk gwv1beta1.BackendRef) int32 {
	if bk.Weight != nil {
		return *bk.Weight
	}

	return 1
}

func mergeL7RouteRule(rule1 routecfg.L7RouteRule, rule2 routecfg.L7RouteRule) routecfg.L7RouteRule {
	mergedRule := routecfg.L7RouteRule{}

	for hostname, rule := range rule1 {
		mergedRule[hostname] = rule
	}

	for hostname, rule := range rule2 {
		if r1, exists := mergedRule[hostname]; exists {
			// can only merge same type of route into one hostname
			switch r1 := r1.(type) {
			case routecfg.GRPCRouteRuleSpec:
				switch r2 := rule.(type) {
				case routecfg.GRPCRouteRuleSpec:
					r1.Matches = append(r1.Matches, r2.Matches...)
					mergedRule[hostname] = r1
				default:
					log.Error().Msgf("%s has been already mapped to RouteRule[%s] %v, current RouteRule %v will be dropped.", hostname, r1.RouteType, r1, r2)
				}
			case routecfg.HTTPRouteRuleSpec:
				switch r2 := rule.(type) {
				case routecfg.HTTPRouteRuleSpec:
					r1.Matches = append(r1.Matches, r2.Matches...)
					mergedRule[hostname] = r1
				default:
					log.Error().Msgf("%s has been already mapped to RouteRule[%s] %v, current RouteRule %v will be dropped.", hostname, r1.RouteType, r1, r2)
				}
			}
		} else {
			mergedRule[hostname] = rule
		}
	}

	return mergedRule
}

func copyMap[K, V comparable](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}

func isEndpointReady(ep discoveryv1.Endpoint) bool {
	if ep.Conditions.Ready != nil && *ep.Conditions.Ready {
		return true
	}

	return false
}

func getServicePort(svc *corev1.Service, port *int32) (corev1.ServicePort, error) {
	if port == nil && len(svc.Spec.Ports) == 1 {
		return svc.Spec.Ports[0], nil
	}

	if port != nil {
		for _, p := range svc.Spec.Ports {
			if p.Port == *port {
				return p, nil
			}
		}
	}

	return corev1.ServicePort{}, fmt.Errorf("no matching port for Service %s and port %d", svc.Name, port)
}

func filterEndpointSliceList(endpointSliceList []*discoveryv1.EndpointSlice, port corev1.ServicePort) []*discoveryv1.EndpointSlice {
	filtered := make([]*discoveryv1.EndpointSlice, 0, len(endpointSliceList))

	for _, endpointSlice := range endpointSliceList {
		if !ignoreEndpointSlice(endpointSlice, port) {
			filtered = append(filtered, endpointSlice)
		}
	}

	return filtered
}

func ignoreEndpointSlice(endpointSlice *discoveryv1.EndpointSlice, port corev1.ServicePort) bool {
	if endpointSlice.AddressType != discoveryv1.AddressTypeIPv4 {
		return true
	}

	// ignore endpoint slices that don't have a matching port.
	return findPort(endpointSlice.Ports, port) == 0
}

func findPort(ports []discoveryv1.EndpointPort, svcPort corev1.ServicePort) int32 {
	portName := svcPort.Name
	for _, p := range ports {
		if p.Port == nil {
			return getDefaultPort(svcPort)
		}

		if p.Name != nil && *p.Name == portName {
			return *p.Port
		}
	}

	return 0
}

func getDefaultPort(svcPort corev1.ServicePort) int32 {
	switch svcPort.TargetPort.Type {
	case intstr.Int:
		if svcPort.TargetPort.IntVal != 0 {
			return svcPort.TargetPort.IntVal
		}
	}

	return svcPort.Port
}

func isMTLSEnabled(gw *gwv1beta1.Gateway) bool {
	if len(gw.Annotations) == 0 {
		return false
	}

	return utils.ParseEnabled(gw.Annotations[constants.GatewayMTLSAnnotation])
}