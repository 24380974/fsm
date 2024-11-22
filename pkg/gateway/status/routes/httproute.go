package routes

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"

	"github.com/flomesh-io/fsm/pkg/constants"
	"github.com/flomesh-io/fsm/pkg/gateway/status"
)

func (p *RouteStatusProcessor) processHTTPRouteStatus(route *gwv1.HTTPRoute, parentRef gwv1.ParentReference, rps status.RouteParentStatusObject) bool {
	for _, rule := range route.Spec.Rules {
		if !p.processHTTPRouteRuleBackendRefs(route, parentRef, rule.BackendRefs, rps) {
			return false
		}
	}

	// All backend references of all rules have been resolved successfully for the parent
	p.addResolvedRefsCondition(route, rps, gwv1.RouteReasonResolvedRefs, "All backend references are resolved")

	return true
}

func (p *RouteStatusProcessor) processHTTPRouteRuleBackendRefs(route *gwv1.HTTPRoute, parentRef gwv1.ParentReference, backendRefs []gwv1.HTTPBackendRef, rps status.RouteParentStatusObject) bool {
	for _, bk := range backendRefs {
		if !p.processHTTPRouteBackend(route, parentRef, bk, rps) {
			return false
		}
	}

	return true
}

func (p *RouteStatusProcessor) processHTTPRouteBackend(route *gwv1.HTTPRoute, parentRef gwv1.ParentReference, bk gwv1.HTTPBackendRef, rps status.RouteParentStatusObject) bool {
	svcPort := p.backendRefToServicePortName(route, bk.BackendObjectReference, rps)
	if svcPort == nil {
		return false
	}

	if svcPort.AppProtocol != nil {
		switch *svcPort.AppProtocol {
		case constants.AppProtocolH2C, constants.AppProtocolWS, constants.AppProtocolWSS:
			log.Debug().Msgf("Backend Protocol: %q for service port %q", *svcPort.AppProtocol, svcPort.String())
			if svcPort.Protocol != corev1.ProtocolTCP {
				p.addNotResolvedRefsCondition(route, rps, gwv1.RouteReasonUnsupportedProtocol, fmt.Sprintf("Unsupported AppProtocol %q for protocol %q", *svcPort.AppProtocol, svcPort.Protocol))
				return false
			}
		default:
			p.addNotResolvedRefsCondition(route, rps, gwv1.RouteReasonUnsupportedProtocol, fmt.Sprintf("Unsupported AppProtocol %q", *svcPort.AppProtocol))
			return false
		}
	}

	if !func() bool {
		valid := true
		p.computeBackendTLSPolicyStatus(route, bk.BackendObjectReference, svcPort, parentRef, func(found bool) {
			if !found {
				if svcPort.AppProtocol != nil &&
					*svcPort.AppProtocol == constants.AppProtocolWSS &&
					svcPort.Protocol == corev1.ProtocolTCP {
					p.addNotResolvedRefsCondition(route, rps, gwv1.RouteReasonUnsupportedProtocol, fmt.Sprintf("No matching BackendTLSPolicy was found for the backend protocol %q and appProtocol %q", svcPort.Protocol, *svcPort.AppProtocol))
					valid = false
				}
			}
		})

		return valid
	}() {
		return false
	}

	p.computeBackendLBPolicyStatus(route, bk.BackendObjectReference, svcPort, parentRef)
	p.computeHealthCheckPolicyStatus(route, bk.BackendObjectReference, svcPort, parentRef)

	return true
}
