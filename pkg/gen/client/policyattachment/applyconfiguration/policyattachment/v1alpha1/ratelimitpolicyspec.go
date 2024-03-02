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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// RateLimitPolicySpecApplyConfiguration represents an declarative configuration of the RateLimitPolicySpec type for use
// with apply.
type RateLimitPolicySpecApplyConfiguration struct {
	TargetRef      *v1alpha2.PolicyTargetReference       `json:"targetRef,omitempty"`
	Ports          []PortRateLimitApplyConfiguration     `json:"ports,omitempty"`
	DefaultBPS     *int64                                `json:"bps,omitempty"`
	Hostnames      []HostnameRateLimitApplyConfiguration `json:"hostnames,omitempty"`
	HTTPRateLimits []HTTPRateLimitApplyConfiguration     `json:"http,omitempty"`
	GRPCRateLimits []GRPCRateLimitApplyConfiguration     `json:"grpc,omitempty"`
	DefaultConfig  *L7RateLimitApplyConfiguration        `json:"config,omitempty"`
}

// RateLimitPolicySpecApplyConfiguration constructs an declarative configuration of the RateLimitPolicySpec type for use with
// apply.
func RateLimitPolicySpec() *RateLimitPolicySpecApplyConfiguration {
	return &RateLimitPolicySpecApplyConfiguration{}
}

// WithTargetRef sets the TargetRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetRef field is set to the value of the last call.
func (b *RateLimitPolicySpecApplyConfiguration) WithTargetRef(value v1alpha2.PolicyTargetReference) *RateLimitPolicySpecApplyConfiguration {
	b.TargetRef = &value
	return b
}

// WithPorts adds the given value to the Ports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ports field.
func (b *RateLimitPolicySpecApplyConfiguration) WithPorts(values ...*PortRateLimitApplyConfiguration) *RateLimitPolicySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPorts")
		}
		b.Ports = append(b.Ports, *values[i])
	}
	return b
}

// WithDefaultBPS sets the DefaultBPS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefaultBPS field is set to the value of the last call.
func (b *RateLimitPolicySpecApplyConfiguration) WithDefaultBPS(value int64) *RateLimitPolicySpecApplyConfiguration {
	b.DefaultBPS = &value
	return b
}

// WithHostnames adds the given value to the Hostnames field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Hostnames field.
func (b *RateLimitPolicySpecApplyConfiguration) WithHostnames(values ...*HostnameRateLimitApplyConfiguration) *RateLimitPolicySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHostnames")
		}
		b.Hostnames = append(b.Hostnames, *values[i])
	}
	return b
}

// WithHTTPRateLimits adds the given value to the HTTPRateLimits field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the HTTPRateLimits field.
func (b *RateLimitPolicySpecApplyConfiguration) WithHTTPRateLimits(values ...*HTTPRateLimitApplyConfiguration) *RateLimitPolicySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHTTPRateLimits")
		}
		b.HTTPRateLimits = append(b.HTTPRateLimits, *values[i])
	}
	return b
}

// WithGRPCRateLimits adds the given value to the GRPCRateLimits field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the GRPCRateLimits field.
func (b *RateLimitPolicySpecApplyConfiguration) WithGRPCRateLimits(values ...*GRPCRateLimitApplyConfiguration) *RateLimitPolicySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGRPCRateLimits")
		}
		b.GRPCRateLimits = append(b.GRPCRateLimits, *values[i])
	}
	return b
}

// WithDefaultConfig sets the DefaultConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefaultConfig field is set to the value of the last call.
func (b *RateLimitPolicySpecApplyConfiguration) WithDefaultConfig(value *L7RateLimitApplyConfiguration) *RateLimitPolicySpecApplyConfiguration {
	b.DefaultConfig = value
	return b
}