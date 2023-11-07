package cache

import (
	gwpav1alpha1 "github.com/flomesh-io/fsm/pkg/apis/policyattachment/v1alpha1"

	"github.com/flomesh-io/fsm/pkg/gateway/utils"
)

// SessionStickyPoliciesTrigger is responsible for processing TLSRoute objects
type SessionStickyPoliciesTrigger struct {
}

// Insert adds a TLSRoute to the cache and returns true if the route is effective
func (p *SessionStickyPoliciesTrigger) Insert(obj interface{}, cache *GatewayCache) bool {
	policy, ok := obj.(*gwpav1alpha1.SessionStickyPolicy)
	if !ok {
		log.Error().Msgf("unexpected object type %T", obj)
		return false
	}

	cache.sessionstickies[utils.ObjectKey(policy)] = struct{}{}

	return cache.isRoutableTargetService(policy, policy.Spec.TargetRef)
}

// Delete removes a TLSRoute from the cache and returns true if the route was found
func (p *SessionStickyPoliciesTrigger) Delete(obj interface{}, cache *GatewayCache) bool {
	policy, ok := obj.(*gwpav1alpha1.SessionStickyPolicy)
	if !ok {
		log.Error().Msgf("unexpected object type %T", obj)
		return false
	}

	key := utils.ObjectKey(policy)
	_, found := cache.sessionstickies[key]
	delete(cache.sessionstickies, key)

	return found
}