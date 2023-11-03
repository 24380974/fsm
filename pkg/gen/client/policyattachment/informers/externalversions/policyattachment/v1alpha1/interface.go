/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/flomesh-io/fsm/pkg/gen/client/policyattachment/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AccessControlPolicies returns a AccessControlPolicyInformer.
	AccessControlPolicies() AccessControlPolicyInformer
	// CircuitBreakingPolicies returns a CircuitBreakingPolicyInformer.
	CircuitBreakingPolicies() CircuitBreakingPolicyInformer
	// HealthCheckPolicies returns a HealthCheckPolicyInformer.
	HealthCheckPolicies() HealthCheckPolicyInformer
	// LoadBalancerPolicies returns a LoadBalancerPolicyInformer.
	LoadBalancerPolicies() LoadBalancerPolicyInformer
	// RateLimitPolicies returns a RateLimitPolicyInformer.
	RateLimitPolicies() RateLimitPolicyInformer
	// SessionStickyPolicies returns a SessionStickyPolicyInformer.
	SessionStickyPolicies() SessionStickyPolicyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AccessControlPolicies returns a AccessControlPolicyInformer.
func (v *version) AccessControlPolicies() AccessControlPolicyInformer {
	return &accessControlPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CircuitBreakingPolicies returns a CircuitBreakingPolicyInformer.
func (v *version) CircuitBreakingPolicies() CircuitBreakingPolicyInformer {
	return &circuitBreakingPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HealthCheckPolicies returns a HealthCheckPolicyInformer.
func (v *version) HealthCheckPolicies() HealthCheckPolicyInformer {
	return &healthCheckPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LoadBalancerPolicies returns a LoadBalancerPolicyInformer.
func (v *version) LoadBalancerPolicies() LoadBalancerPolicyInformer {
	return &loadBalancerPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RateLimitPolicies returns a RateLimitPolicyInformer.
func (v *version) RateLimitPolicies() RateLimitPolicyInformer {
	return &rateLimitPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SessionStickyPolicies returns a SessionStickyPolicyInformer.
func (v *version) SessionStickyPolicies() SessionStickyPolicyInformer {
	return &sessionStickyPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
