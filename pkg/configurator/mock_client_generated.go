// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/flomesh-io/fsm/pkg/configurator (interfaces: Configurator)

// Package configurator is a generated GoMock package.
package configurator

import (
	reflect "reflect"
	time "time"

	v1alpha3 "github.com/flomesh-io/fsm/pkg/apis/config/v1alpha3"
	auth "github.com/flomesh-io/fsm/pkg/auth"
	trafficpolicy "github.com/flomesh-io/fsm/pkg/trafficpolicy"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockConfigurator is a mock of Configurator interface.
type MockConfigurator struct {
	ctrl     *gomock.Controller
	recorder *MockConfiguratorMockRecorder
}

// MockConfiguratorMockRecorder is the mock recorder for MockConfigurator.
type MockConfiguratorMockRecorder struct {
	mock *MockConfigurator
}

// NewMockConfigurator creates a new mock instance.
func NewMockConfigurator(ctrl *gomock.Controller) *MockConfigurator {
	mock := &MockConfigurator{ctrl: ctrl}
	mock.recorder = &MockConfiguratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurator) EXPECT() *MockConfiguratorMockRecorder {
	return m.recorder
}

// GetCertKeyBitSize mocks base method.
func (m *MockConfigurator) GetCertKeyBitSize() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertKeyBitSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetCertKeyBitSize indicates an expected call of GetCertKeyBitSize.
func (mr *MockConfiguratorMockRecorder) GetCertKeyBitSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertKeyBitSize", reflect.TypeOf((*MockConfigurator)(nil).GetCertKeyBitSize))
}

// GetClusterUID mocks base method.
func (m *MockConfigurator) GetClusterUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClusterUID indicates an expected call of GetClusterUID.
func (mr *MockConfiguratorMockRecorder) GetClusterUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterUID", reflect.TypeOf((*MockConfigurator)(nil).GetClusterUID))
}

// GetConfigResyncInterval mocks base method.
func (m *MockConfigurator) GetConfigResyncInterval() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigResyncInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetConfigResyncInterval indicates an expected call of GetConfigResyncInterval.
func (mr *MockConfiguratorMockRecorder) GetConfigResyncInterval() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigResyncInterval", reflect.TypeOf((*MockConfigurator)(nil).GetConfigResyncInterval))
}

// GetFLBSecretName mocks base method.
func (m *MockConfigurator) GetFLBSecretName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFLBSecretName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFLBSecretName indicates an expected call of GetFLBSecretName.
func (mr *MockConfiguratorMockRecorder) GetFLBSecretName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFLBSecretName", reflect.TypeOf((*MockConfigurator)(nil).GetFLBSecretName))
}

// GetFSMGatewayLogLevel mocks base method.
func (m *MockConfigurator) GetFSMGatewayLogLevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFSMGatewayLogLevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFSMGatewayLogLevel indicates an expected call of GetFSMGatewayLogLevel.
func (mr *MockConfiguratorMockRecorder) GetFSMGatewayLogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFSMGatewayLogLevel", reflect.TypeOf((*MockConfigurator)(nil).GetFSMGatewayLogLevel))
}

// GetFSMIngressLogLevel mocks base method.
func (m *MockConfigurator) GetFSMIngressLogLevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFSMIngressLogLevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFSMIngressLogLevel indicates an expected call of GetFSMIngressLogLevel.
func (mr *MockConfiguratorMockRecorder) GetFSMIngressLogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFSMIngressLogLevel", reflect.TypeOf((*MockConfigurator)(nil).GetFSMIngressLogLevel))
}

// GetFSMLogLevel mocks base method.
func (m *MockConfigurator) GetFSMLogLevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFSMLogLevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFSMLogLevel indicates an expected call of GetFSMLogLevel.
func (mr *MockConfiguratorMockRecorder) GetFSMLogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFSMLogLevel", reflect.TypeOf((*MockConfigurator)(nil).GetFSMLogLevel))
}

// GetFSMNamespace mocks base method.
func (m *MockConfigurator) GetFSMNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFSMNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFSMNamespace indicates an expected call of GetFSMNamespace.
func (mr *MockConfiguratorMockRecorder) GetFSMNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFSMNamespace", reflect.TypeOf((*MockConfigurator)(nil).GetFSMNamespace))
}

// GetFeatureFlags mocks base method.
func (m *MockConfigurator) GetFeatureFlags() v1alpha3.FeatureFlags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatureFlags")
	ret0, _ := ret[0].(v1alpha3.FeatureFlags)
	return ret0
}

// GetFeatureFlags indicates an expected call of GetFeatureFlags.
func (mr *MockConfiguratorMockRecorder) GetFeatureFlags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatureFlags", reflect.TypeOf((*MockConfigurator)(nil).GetFeatureFlags))
}

// GetGlobalPluginChains mocks base method.
func (m *MockConfigurator) GetGlobalPluginChains() map[string][]trafficpolicy.Plugin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGlobalPluginChains")
	ret0, _ := ret[0].(map[string][]trafficpolicy.Plugin)
	return ret0
}

// GetGlobalPluginChains indicates an expected call of GetGlobalPluginChains.
func (mr *MockConfiguratorMockRecorder) GetGlobalPluginChains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGlobalPluginChains", reflect.TypeOf((*MockConfigurator)(nil).GetGlobalPluginChains))
}

// GetImagePullPolicy mocks base method.
func (m *MockConfigurator) GetImagePullPolicy() v1.PullPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImagePullPolicy")
	ret0, _ := ret[0].(v1.PullPolicy)
	return ret0
}

// GetImagePullPolicy indicates an expected call of GetImagePullPolicy.
func (mr *MockConfiguratorMockRecorder) GetImagePullPolicy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImagePullPolicy", reflect.TypeOf((*MockConfigurator)(nil).GetImagePullPolicy))
}

// GetImageRegistry mocks base method.
func (m *MockConfigurator) GetImageRegistry() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageRegistry")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetImageRegistry indicates an expected call of GetImageRegistry.
func (mr *MockConfiguratorMockRecorder) GetImageRegistry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageRegistry", reflect.TypeOf((*MockConfigurator)(nil).GetImageRegistry))
}

// GetImageTag mocks base method.
func (m *MockConfigurator) GetImageTag() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageTag")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetImageTag indicates an expected call of GetImageTag.
func (mr *MockConfiguratorMockRecorder) GetImageTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageTag", reflect.TypeOf((*MockConfigurator)(nil).GetImageTag))
}

// GetInboundExternalAuthConfig mocks base method.
func (m *MockConfigurator) GetInboundExternalAuthConfig() auth.ExtAuthConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInboundExternalAuthConfig")
	ret0, _ := ret[0].(auth.ExtAuthConfig)
	return ret0
}

// GetInboundExternalAuthConfig indicates an expected call of GetInboundExternalAuthConfig.
func (mr *MockConfiguratorMockRecorder) GetInboundExternalAuthConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInboundExternalAuthConfig", reflect.TypeOf((*MockConfigurator)(nil).GetInboundExternalAuthConfig))
}

// GetIngressGatewayCertValidityPeriod mocks base method.
func (m *MockConfigurator) GetIngressGatewayCertValidityPeriod() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressGatewayCertValidityPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetIngressGatewayCertValidityPeriod indicates an expected call of GetIngressGatewayCertValidityPeriod.
func (mr *MockConfiguratorMockRecorder) GetIngressGatewayCertValidityPeriod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressGatewayCertValidityPeriod", reflect.TypeOf((*MockConfigurator)(nil).GetIngressGatewayCertValidityPeriod))
}

// GetIngressHTTPListenPort mocks base method.
func (m *MockConfigurator) GetIngressHTTPListenPort() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressHTTPListenPort")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetIngressHTTPListenPort indicates an expected call of GetIngressHTTPListenPort.
func (mr *MockConfiguratorMockRecorder) GetIngressHTTPListenPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressHTTPListenPort", reflect.TypeOf((*MockConfigurator)(nil).GetIngressHTTPListenPort))
}

// GetIngressSSLPassthroughUpstreamPort mocks base method.
func (m *MockConfigurator) GetIngressSSLPassthroughUpstreamPort() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressSSLPassthroughUpstreamPort")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetIngressSSLPassthroughUpstreamPort indicates an expected call of GetIngressSSLPassthroughUpstreamPort.
func (mr *MockConfiguratorMockRecorder) GetIngressSSLPassthroughUpstreamPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressSSLPassthroughUpstreamPort", reflect.TypeOf((*MockConfigurator)(nil).GetIngressSSLPassthroughUpstreamPort))
}

// GetIngressTLSListenPort mocks base method.
func (m *MockConfigurator) GetIngressTLSListenPort() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressTLSListenPort")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetIngressTLSListenPort indicates an expected call of GetIngressTLSListenPort.
func (mr *MockConfiguratorMockRecorder) GetIngressTLSListenPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressTLSListenPort", reflect.TypeOf((*MockConfigurator)(nil).GetIngressTLSListenPort))
}

// GetInitContainerImage mocks base method.
func (m *MockConfigurator) GetInitContainerImage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitContainerImage")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetInitContainerImage indicates an expected call of GetInitContainerImage.
func (mr *MockConfiguratorMockRecorder) GetInitContainerImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitContainerImage", reflect.TypeOf((*MockConfigurator)(nil).GetInitContainerImage))
}

// GetLocalDNSProxyPrimaryUpstream mocks base method.
func (m *MockConfigurator) GetLocalDNSProxyPrimaryUpstream() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocalDNSProxyPrimaryUpstream")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLocalDNSProxyPrimaryUpstream indicates an expected call of GetLocalDNSProxyPrimaryUpstream.
func (mr *MockConfiguratorMockRecorder) GetLocalDNSProxyPrimaryUpstream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocalDNSProxyPrimaryUpstream", reflect.TypeOf((*MockConfigurator)(nil).GetLocalDNSProxyPrimaryUpstream))
}

// GetLocalDNSProxySecondaryUpstream mocks base method.
func (m *MockConfigurator) GetLocalDNSProxySecondaryUpstream() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocalDNSProxySecondaryUpstream")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLocalDNSProxySecondaryUpstream indicates an expected call of GetLocalDNSProxySecondaryUpstream.
func (mr *MockConfiguratorMockRecorder) GetLocalDNSProxySecondaryUpstream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocalDNSProxySecondaryUpstream", reflect.TypeOf((*MockConfigurator)(nil).GetLocalDNSProxySecondaryUpstream))
}

// GetMaxDataPlaneConnections mocks base method.
func (m *MockConfigurator) GetMaxDataPlaneConnections() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxDataPlaneConnections")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMaxDataPlaneConnections indicates an expected call of GetMaxDataPlaneConnections.
func (mr *MockConfiguratorMockRecorder) GetMaxDataPlaneConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxDataPlaneConnections", reflect.TypeOf((*MockConfigurator)(nil).GetMaxDataPlaneConnections))
}

// GetMeshConfig mocks base method.
func (m *MockConfigurator) GetMeshConfig() v1alpha3.MeshConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshConfig")
	ret0, _ := ret[0].(v1alpha3.MeshConfig)
	return ret0
}

// GetMeshConfig indicates an expected call of GetMeshConfig.
func (mr *MockConfiguratorMockRecorder) GetMeshConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshConfig", reflect.TypeOf((*MockConfigurator)(nil).GetMeshConfig))
}

// GetMeshConfigJSON mocks base method.
func (m *MockConfigurator) GetMeshConfigJSON() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshConfigJSON")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeshConfigJSON indicates an expected call of GetMeshConfigJSON.
func (mr *MockConfiguratorMockRecorder) GetMeshConfigJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshConfigJSON", reflect.TypeOf((*MockConfigurator)(nil).GetMeshConfigJSON))
}

// GetMultiClusterControlPlaneUID mocks base method.
func (m *MockConfigurator) GetMultiClusterControlPlaneUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMultiClusterControlPlaneUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMultiClusterControlPlaneUID indicates an expected call of GetMultiClusterControlPlaneUID.
func (mr *MockConfiguratorMockRecorder) GetMultiClusterControlPlaneUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiClusterControlPlaneUID", reflect.TypeOf((*MockConfigurator)(nil).GetMultiClusterControlPlaneUID))
}

// GetProxyResources mocks base method.
func (m *MockConfigurator) GetProxyResources() v1.ResourceRequirements {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProxyResources")
	ret0, _ := ret[0].(v1.ResourceRequirements)
	return ret0
}

// GetProxyResources indicates an expected call of GetProxyResources.
func (mr *MockConfiguratorMockRecorder) GetProxyResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProxyResources", reflect.TypeOf((*MockConfigurator)(nil).GetProxyResources))
}

// GetProxyServerPort mocks base method.
func (m *MockConfigurator) GetProxyServerPort() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProxyServerPort")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetProxyServerPort indicates an expected call of GetProxyServerPort.
func (mr *MockConfiguratorMockRecorder) GetProxyServerPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProxyServerPort", reflect.TypeOf((*MockConfigurator)(nil).GetProxyServerPort))
}

// GetRemoteLoggingAuthorization mocks base method.
func (m *MockConfigurator) GetRemoteLoggingAuthorization() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteLoggingAuthorization")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRemoteLoggingAuthorization indicates an expected call of GetRemoteLoggingAuthorization.
func (mr *MockConfiguratorMockRecorder) GetRemoteLoggingAuthorization() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteLoggingAuthorization", reflect.TypeOf((*MockConfigurator)(nil).GetRemoteLoggingAuthorization))
}

// GetRemoteLoggingEndpoint mocks base method.
func (m *MockConfigurator) GetRemoteLoggingEndpoint() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteLoggingEndpoint")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRemoteLoggingEndpoint indicates an expected call of GetRemoteLoggingEndpoint.
func (mr *MockConfiguratorMockRecorder) GetRemoteLoggingEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteLoggingEndpoint", reflect.TypeOf((*MockConfigurator)(nil).GetRemoteLoggingEndpoint))
}

// GetRemoteLoggingHost mocks base method.
func (m *MockConfigurator) GetRemoteLoggingHost() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteLoggingHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRemoteLoggingHost indicates an expected call of GetRemoteLoggingHost.
func (mr *MockConfiguratorMockRecorder) GetRemoteLoggingHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteLoggingHost", reflect.TypeOf((*MockConfigurator)(nil).GetRemoteLoggingHost))
}

// GetRemoteLoggingPort mocks base method.
func (m *MockConfigurator) GetRemoteLoggingPort() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteLoggingPort")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetRemoteLoggingPort indicates an expected call of GetRemoteLoggingPort.
func (mr *MockConfiguratorMockRecorder) GetRemoteLoggingPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteLoggingPort", reflect.TypeOf((*MockConfigurator)(nil).GetRemoteLoggingPort))
}

// GetRemoteLoggingSampledFraction mocks base method.
func (m *MockConfigurator) GetRemoteLoggingSampledFraction() float32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteLoggingSampledFraction")
	ret0, _ := ret[0].(float32)
	return ret0
}

// GetRemoteLoggingSampledFraction indicates an expected call of GetRemoteLoggingSampledFraction.
func (mr *MockConfiguratorMockRecorder) GetRemoteLoggingSampledFraction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteLoggingSampledFraction", reflect.TypeOf((*MockConfigurator)(nil).GetRemoteLoggingSampledFraction))
}

// GetRemoteLoggingSecretName mocks base method.
func (m *MockConfigurator) GetRemoteLoggingSecretName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteLoggingSecretName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRemoteLoggingSecretName indicates an expected call of GetRemoteLoggingSecretName.
func (mr *MockConfiguratorMockRecorder) GetRemoteLoggingSecretName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteLoggingSecretName", reflect.TypeOf((*MockConfigurator)(nil).GetRemoteLoggingSecretName))
}

// GetRepoServerCodebase mocks base method.
func (m *MockConfigurator) GetRepoServerCodebase() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepoServerCodebase")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRepoServerCodebase indicates an expected call of GetRepoServerCodebase.
func (mr *MockConfiguratorMockRecorder) GetRepoServerCodebase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepoServerCodebase", reflect.TypeOf((*MockConfigurator)(nil).GetRepoServerCodebase))
}

// GetRepoServerIPAddr mocks base method.
func (m *MockConfigurator) GetRepoServerIPAddr() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepoServerIPAddr")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRepoServerIPAddr indicates an expected call of GetRepoServerIPAddr.
func (mr *MockConfiguratorMockRecorder) GetRepoServerIPAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepoServerIPAddr", reflect.TypeOf((*MockConfigurator)(nil).GetRepoServerIPAddr))
}

// GetServiceAccessMode mocks base method.
func (m *MockConfigurator) GetServiceAccessMode() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceAccessMode")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceAccessMode indicates an expected call of GetServiceAccessMode.
func (mr *MockConfiguratorMockRecorder) GetServiceAccessMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceAccessMode", reflect.TypeOf((*MockConfigurator)(nil).GetServiceAccessMode))
}

// GetServiceCertValidityPeriod mocks base method.
func (m *MockConfigurator) GetServiceCertValidityPeriod() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceCertValidityPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetServiceCertValidityPeriod indicates an expected call of GetServiceCertValidityPeriod.
func (mr *MockConfiguratorMockRecorder) GetServiceCertValidityPeriod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceCertValidityPeriod", reflect.TypeOf((*MockConfigurator)(nil).GetServiceCertValidityPeriod))
}

// GetSidecarClass mocks base method.
func (m *MockConfigurator) GetSidecarClass() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSidecarClass")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSidecarClass indicates an expected call of GetSidecarClass.
func (mr *MockConfiguratorMockRecorder) GetSidecarClass() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSidecarClass", reflect.TypeOf((*MockConfigurator)(nil).GetSidecarClass))
}

// GetSidecarDisabledMTLS mocks base method.
func (m *MockConfigurator) GetSidecarDisabledMTLS() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSidecarDisabledMTLS")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetSidecarDisabledMTLS indicates an expected call of GetSidecarDisabledMTLS.
func (mr *MockConfiguratorMockRecorder) GetSidecarDisabledMTLS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSidecarDisabledMTLS", reflect.TypeOf((*MockConfigurator)(nil).GetSidecarDisabledMTLS))
}

// GetSidecarImage mocks base method.
func (m *MockConfigurator) GetSidecarImage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSidecarImage")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSidecarImage indicates an expected call of GetSidecarImage.
func (mr *MockConfiguratorMockRecorder) GetSidecarImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSidecarImage", reflect.TypeOf((*MockConfigurator)(nil).GetSidecarImage))
}

// GetSidecarLogLevel mocks base method.
func (m *MockConfigurator) GetSidecarLogLevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSidecarLogLevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSidecarLogLevel indicates an expected call of GetSidecarLogLevel.
func (mr *MockConfiguratorMockRecorder) GetSidecarLogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSidecarLogLevel", reflect.TypeOf((*MockConfigurator)(nil).GetSidecarLogLevel))
}

// GetTracingEndpoint mocks base method.
func (m *MockConfigurator) GetTracingEndpoint() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracingEndpoint")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTracingEndpoint indicates an expected call of GetTracingEndpoint.
func (mr *MockConfiguratorMockRecorder) GetTracingEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracingEndpoint", reflect.TypeOf((*MockConfigurator)(nil).GetTracingEndpoint))
}

// GetTracingHost mocks base method.
func (m *MockConfigurator) GetTracingHost() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracingHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTracingHost indicates an expected call of GetTracingHost.
func (mr *MockConfiguratorMockRecorder) GetTracingHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracingHost", reflect.TypeOf((*MockConfigurator)(nil).GetTracingHost))
}

// GetTracingPort mocks base method.
func (m *MockConfigurator) GetTracingPort() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracingPort")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetTracingPort indicates an expected call of GetTracingPort.
func (mr *MockConfiguratorMockRecorder) GetTracingPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracingPort", reflect.TypeOf((*MockConfigurator)(nil).GetTracingPort))
}

// GetTracingSampledFraction mocks base method.
func (m *MockConfigurator) GetTracingSampledFraction() float32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracingSampledFraction")
	ret0, _ := ret[0].(float32)
	return ret0
}

// GetTracingSampledFraction indicates an expected call of GetTracingSampledFraction.
func (mr *MockConfiguratorMockRecorder) GetTracingSampledFraction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracingSampledFraction", reflect.TypeOf((*MockConfigurator)(nil).GetTracingSampledFraction))
}

// GetTrafficInterceptionMode mocks base method.
func (m *MockConfigurator) GetTrafficInterceptionMode() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrafficInterceptionMode")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTrafficInterceptionMode indicates an expected call of GetTrafficInterceptionMode.
func (mr *MockConfiguratorMockRecorder) GetTrafficInterceptionMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrafficInterceptionMode", reflect.TypeOf((*MockConfigurator)(nil).GetTrafficInterceptionMode))
}

// IsDebugServerEnabled mocks base method.
func (m *MockConfigurator) IsDebugServerEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDebugServerEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDebugServerEnabled indicates an expected call of IsDebugServerEnabled.
func (mr *MockConfiguratorMockRecorder) IsDebugServerEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDebugServerEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsDebugServerEnabled))
}

// IsEgressEnabled mocks base method.
func (m *MockConfigurator) IsEgressEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEgressEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEgressEnabled indicates an expected call of IsEgressEnabled.
func (mr *MockConfiguratorMockRecorder) IsEgressEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEgressEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsEgressEnabled))
}

// IsFLBEnabled mocks base method.
func (m *MockConfigurator) IsFLBEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFLBEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFLBEnabled indicates an expected call of IsFLBEnabled.
func (mr *MockConfiguratorMockRecorder) IsFLBEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFLBEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsFLBEnabled))
}

// IsFLBStrictModeEnabled mocks base method.
func (m *MockConfigurator) IsFLBStrictModeEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFLBStrictModeEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFLBStrictModeEnabled indicates an expected call of IsFLBStrictModeEnabled.
func (mr *MockConfiguratorMockRecorder) IsFLBStrictModeEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFLBStrictModeEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsFLBStrictModeEnabled))
}

// IsGatewayAPIEnabled mocks base method.
func (m *MockConfigurator) IsGatewayAPIEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsGatewayAPIEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsGatewayAPIEnabled indicates an expected call of IsGatewayAPIEnabled.
func (mr *MockConfiguratorMockRecorder) IsGatewayAPIEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsGatewayAPIEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsGatewayAPIEnabled))
}

// IsIngressEnabled mocks base method.
func (m *MockConfigurator) IsIngressEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIngressEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsIngressEnabled indicates an expected call of IsIngressEnabled.
func (mr *MockConfiguratorMockRecorder) IsIngressEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIngressEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsIngressEnabled))
}

// IsIngressHTTPEnabled mocks base method.
func (m *MockConfigurator) IsIngressHTTPEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIngressHTTPEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsIngressHTTPEnabled indicates an expected call of IsIngressHTTPEnabled.
func (mr *MockConfiguratorMockRecorder) IsIngressHTTPEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIngressHTTPEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsIngressHTTPEnabled))
}

// IsIngressMTLSEnabled mocks base method.
func (m *MockConfigurator) IsIngressMTLSEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIngressMTLSEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsIngressMTLSEnabled indicates an expected call of IsIngressMTLSEnabled.
func (mr *MockConfiguratorMockRecorder) IsIngressMTLSEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIngressMTLSEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsIngressMTLSEnabled))
}

// IsIngressSSLPassthroughEnabled mocks base method.
func (m *MockConfigurator) IsIngressSSLPassthroughEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIngressSSLPassthroughEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsIngressSSLPassthroughEnabled indicates an expected call of IsIngressSSLPassthroughEnabled.
func (mr *MockConfiguratorMockRecorder) IsIngressSSLPassthroughEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIngressSSLPassthroughEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsIngressSSLPassthroughEnabled))
}

// IsIngressTLSEnabled mocks base method.
func (m *MockConfigurator) IsIngressTLSEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIngressTLSEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsIngressTLSEnabled indicates an expected call of IsIngressTLSEnabled.
func (mr *MockConfiguratorMockRecorder) IsIngressTLSEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIngressTLSEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsIngressTLSEnabled))
}

// IsLocalDNSProxyEnabled mocks base method.
func (m *MockConfigurator) IsLocalDNSProxyEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLocalDNSProxyEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLocalDNSProxyEnabled indicates an expected call of IsLocalDNSProxyEnabled.
func (mr *MockConfiguratorMockRecorder) IsLocalDNSProxyEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLocalDNSProxyEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsLocalDNSProxyEnabled))
}

// IsManaged mocks base method.
func (m *MockConfigurator) IsManaged() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsManaged")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsManaged indicates an expected call of IsManaged.
func (mr *MockConfiguratorMockRecorder) IsManaged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsManaged", reflect.TypeOf((*MockConfigurator)(nil).IsManaged))
}

// IsMultiClusterControlPlane mocks base method.
func (m *MockConfigurator) IsMultiClusterControlPlane() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMultiClusterControlPlane")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMultiClusterControlPlane indicates an expected call of IsMultiClusterControlPlane.
func (mr *MockConfiguratorMockRecorder) IsMultiClusterControlPlane() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMultiClusterControlPlane", reflect.TypeOf((*MockConfigurator)(nil).IsMultiClusterControlPlane))
}

// IsNamespacedIngressEnabled mocks base method.
func (m *MockConfigurator) IsNamespacedIngressEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNamespacedIngressEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNamespacedIngressEnabled indicates an expected call of IsNamespacedIngressEnabled.
func (mr *MockConfiguratorMockRecorder) IsNamespacedIngressEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNamespacedIngressEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsNamespacedIngressEnabled))
}

// IsPermissiveTrafficPolicyMode mocks base method.
func (m *MockConfigurator) IsPermissiveTrafficPolicyMode() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPermissiveTrafficPolicyMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPermissiveTrafficPolicyMode indicates an expected call of IsPermissiveTrafficPolicyMode.
func (mr *MockConfiguratorMockRecorder) IsPermissiveTrafficPolicyMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPermissiveTrafficPolicyMode", reflect.TypeOf((*MockConfigurator)(nil).IsPermissiveTrafficPolicyMode))
}

// IsPrivilegedInitContainer mocks base method.
func (m *MockConfigurator) IsPrivilegedInitContainer() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPrivilegedInitContainer")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPrivilegedInitContainer indicates an expected call of IsPrivilegedInitContainer.
func (mr *MockConfiguratorMockRecorder) IsPrivilegedInitContainer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPrivilegedInitContainer", reflect.TypeOf((*MockConfigurator)(nil).IsPrivilegedInitContainer))
}

// IsRemoteLoggingEnabled mocks base method.
func (m *MockConfigurator) IsRemoteLoggingEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRemoteLoggingEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRemoteLoggingEnabled indicates an expected call of IsRemoteLoggingEnabled.
func (mr *MockConfiguratorMockRecorder) IsRemoteLoggingEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRemoteLoggingEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsRemoteLoggingEnabled))
}

// IsServiceLBEnabled mocks base method.
func (m *MockConfigurator) IsServiceLBEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsServiceLBEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsServiceLBEnabled indicates an expected call of IsServiceLBEnabled.
func (mr *MockConfiguratorMockRecorder) IsServiceLBEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsServiceLBEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsServiceLBEnabled))
}

// IsTracingEnabled mocks base method.
func (m *MockConfigurator) IsTracingEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTracingEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTracingEnabled indicates an expected call of IsTracingEnabled.
func (mr *MockConfiguratorMockRecorder) IsTracingEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTracingEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsTracingEnabled))
}

// ServiceLBImage mocks base method.
func (m *MockConfigurator) ServiceLBImage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceLBImage")
	ret0, _ := ret[0].(string)
	return ret0
}

// ServiceLBImage indicates an expected call of ServiceLBImage.
func (mr *MockConfiguratorMockRecorder) ServiceLBImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceLBImage", reflect.TypeOf((*MockConfigurator)(nil).ServiceLBImage))
}