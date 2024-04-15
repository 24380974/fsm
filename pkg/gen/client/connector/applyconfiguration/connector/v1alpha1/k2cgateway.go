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
	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/connector/v1alpha1"
)

// K2CGatewayApplyConfiguration represents an declarative configuration of the K2CGateway type for use
// with apply.
type K2CGatewayApplyConfiguration struct {
	Enable      *bool                     `json:"enable,omitempty"`
	GatewayMode *v1alpha1.WithGatewayMode `json:"gatewayMode,omitempty"`
}

// K2CGatewayApplyConfiguration constructs an declarative configuration of the K2CGateway type for use with
// apply.
func K2CGateway() *K2CGatewayApplyConfiguration {
	return &K2CGatewayApplyConfiguration{}
}

// WithEnable sets the Enable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enable field is set to the value of the last call.
func (b *K2CGatewayApplyConfiguration) WithEnable(value bool) *K2CGatewayApplyConfiguration {
	b.Enable = &value
	return b
}

// WithGatewayMode sets the GatewayMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GatewayMode field is set to the value of the last call.
func (b *K2CGatewayApplyConfiguration) WithGatewayMode(value v1alpha1.WithGatewayMode) *K2CGatewayApplyConfiguration {
	b.GatewayMode = &value
	return b
}
