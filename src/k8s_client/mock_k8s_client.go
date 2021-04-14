// Code generated by MockGen. DO NOT EDIT.
// Source: k8s_client.go

// Package k8s_client is a generated GoMock package.
package k8s_client

import (
	bytes "bytes"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/metal3-io/baremetal-operator/pkg/apis/metal3/v1alpha1"
	v1 "github.com/openshift/api/config/v1"
	ops "github.com/openshift/assisted-installer/src/ops"
	v1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	v1alpha10 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	v10 "k8s.io/api/certificates/v1"
	v11 "k8s.io/api/core/v1"
)

// MockK8SClient is a mock of K8SClient interface
type MockK8SClient struct {
	ctrl     *gomock.Controller
	recorder *MockK8SClientMockRecorder
}

// MockK8SClientMockRecorder is the mock recorder for MockK8SClient
type MockK8SClientMockRecorder struct {
	mock *MockK8SClient
}

// NewMockK8SClient creates a new mock instance
func NewMockK8SClient(ctrl *gomock.Controller) *MockK8SClient {
	mock := &MockK8SClient{ctrl: ctrl}
	mock.recorder = &MockK8SClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockK8SClient) EXPECT() *MockK8SClientMockRecorder {
	return m.recorder
}

// ListMasterNodes mocks base method
func (m *MockK8SClient) ListMasterNodes() (*v11.NodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMasterNodes")
	ret0, _ := ret[0].(*v11.NodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMasterNodes indicates an expected call of ListMasterNodes
func (mr *MockK8SClientMockRecorder) ListMasterNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMasterNodes", reflect.TypeOf((*MockK8SClient)(nil).ListMasterNodes))
}

// PatchEtcd mocks base method
func (m *MockK8SClient) PatchEtcd() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchEtcd")
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchEtcd indicates an expected call of PatchEtcd
func (mr *MockK8SClientMockRecorder) PatchEtcd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchEtcd", reflect.TypeOf((*MockK8SClient)(nil).PatchEtcd))
}

// UnPatchEtcd mocks base method
func (m *MockK8SClient) UnPatchEtcd() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnPatchEtcd")
	ret0, _ := ret[0].(error)
	return ret0
}

// UnPatchEtcd indicates an expected call of UnPatchEtcd
func (mr *MockK8SClientMockRecorder) UnPatchEtcd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnPatchEtcd", reflect.TypeOf((*MockK8SClient)(nil).UnPatchEtcd))
}

// PatchControlPlaneReplicas mocks base method
func (m *MockK8SClient) PatchControlPlaneReplicas() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchControlPlaneReplicas")
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchControlPlaneReplicas indicates an expected call of PatchControlPlaneReplicas
func (mr *MockK8SClientMockRecorder) PatchControlPlaneReplicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchControlPlaneReplicas", reflect.TypeOf((*MockK8SClient)(nil).PatchControlPlaneReplicas))
}

// UnPatchControlPlaneReplicas mocks base method
func (m *MockK8SClient) UnPatchControlPlaneReplicas() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnPatchControlPlaneReplicas")
	ret0, _ := ret[0].(error)
	return ret0
}

// UnPatchControlPlaneReplicas indicates an expected call of UnPatchControlPlaneReplicas
func (mr *MockK8SClientMockRecorder) UnPatchControlPlaneReplicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnPatchControlPlaneReplicas", reflect.TypeOf((*MockK8SClient)(nil).UnPatchControlPlaneReplicas))
}

// ListNodes mocks base method
func (m *MockK8SClient) ListNodes() (*v11.NodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNodes")
	ret0, _ := ret[0].(*v11.NodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNodes indicates an expected call of ListNodes
func (mr *MockK8SClientMockRecorder) ListNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodes", reflect.TypeOf((*MockK8SClient)(nil).ListNodes))
}

// ListMachines mocks base method
func (m *MockK8SClient) ListMachines() (*v1beta1.MachineList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMachines")
	ret0, _ := ret[0].(*v1beta1.MachineList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMachines indicates an expected call of ListMachines
func (mr *MockK8SClientMockRecorder) ListMachines() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMachines", reflect.TypeOf((*MockK8SClient)(nil).ListMachines))
}

// RunOCctlCommand mocks base method
func (m *MockK8SClient) RunOCctlCommand(args []string, kubeconfigPath string, o ops.Ops) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunOCctlCommand", args, kubeconfigPath, o)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunOCctlCommand indicates an expected call of RunOCctlCommand
func (mr *MockK8SClientMockRecorder) RunOCctlCommand(args, kubeconfigPath, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunOCctlCommand", reflect.TypeOf((*MockK8SClient)(nil).RunOCctlCommand), args, kubeconfigPath, o)
}

// ApproveCsr mocks base method
func (m *MockK8SClient) ApproveCsr(csr *v10.CertificateSigningRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveCsr", csr)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApproveCsr indicates an expected call of ApproveCsr
func (mr *MockK8SClientMockRecorder) ApproveCsr(csr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveCsr", reflect.TypeOf((*MockK8SClient)(nil).ApproveCsr), csr)
}

// ListCsrs mocks base method
func (m *MockK8SClient) ListCsrs() (*v10.CertificateSigningRequestList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCsrs")
	ret0, _ := ret[0].(*v10.CertificateSigningRequestList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCsrs indicates an expected call of ListCsrs
func (mr *MockK8SClientMockRecorder) ListCsrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCsrs", reflect.TypeOf((*MockK8SClient)(nil).ListCsrs))
}

// GetConfigMap mocks base method
func (m *MockK8SClient) GetConfigMap(namespace, name string) (*v11.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigMap", namespace, name)
	ret0, _ := ret[0].(*v11.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigMap indicates an expected call of GetConfigMap
func (mr *MockK8SClientMockRecorder) GetConfigMap(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigMap", reflect.TypeOf((*MockK8SClient)(nil).GetConfigMap), namespace, name)
}

// GetPodLogs mocks base method
func (m *MockK8SClient) GetPodLogs(namespace, podName string, sinceSeconds int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodLogs", namespace, podName, sinceSeconds)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodLogs indicates an expected call of GetPodLogs
func (mr *MockK8SClientMockRecorder) GetPodLogs(namespace, podName, sinceSeconds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodLogs", reflect.TypeOf((*MockK8SClient)(nil).GetPodLogs), namespace, podName, sinceSeconds)
}

// GetPodLogsAsBuffer mocks base method
func (m *MockK8SClient) GetPodLogsAsBuffer(namespace, podName string, sinceSeconds int64) (*bytes.Buffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodLogsAsBuffer", namespace, podName, sinceSeconds)
	ret0, _ := ret[0].(*bytes.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodLogsAsBuffer indicates an expected call of GetPodLogsAsBuffer
func (mr *MockK8SClientMockRecorder) GetPodLogsAsBuffer(namespace, podName, sinceSeconds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodLogsAsBuffer", reflect.TypeOf((*MockK8SClient)(nil).GetPodLogsAsBuffer), namespace, podName, sinceSeconds)
}

// GetPods mocks base method
func (m *MockK8SClient) GetPods(namespace string, labelMatch map[string]string, fieldSelector string) ([]v11.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPods", namespace, labelMatch, fieldSelector)
	ret0, _ := ret[0].([]v11.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPods indicates an expected call of GetPods
func (mr *MockK8SClientMockRecorder) GetPods(namespace, labelMatch, fieldSelector interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPods", reflect.TypeOf((*MockK8SClient)(nil).GetPods), namespace, labelMatch, fieldSelector)
}

// GetCSV mocks base method
func (m *MockK8SClient) GetCSV(namespace, name string) (*v1alpha10.ClusterServiceVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCSV", namespace, name)
	ret0, _ := ret[0].(*v1alpha10.ClusterServiceVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCSV indicates an expected call of GetCSV
func (mr *MockK8SClientMockRecorder) GetCSV(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCSV", reflect.TypeOf((*MockK8SClient)(nil).GetCSV), namespace, name)
}

// GetCSVFromSubscription mocks base method
func (m *MockK8SClient) GetCSVFromSubscription(namespace, name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCSVFromSubscription", namespace, name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCSVFromSubscription indicates an expected call of GetCSVFromSubscription
func (mr *MockK8SClientMockRecorder) GetCSVFromSubscription(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCSVFromSubscription", reflect.TypeOf((*MockK8SClient)(nil).GetCSVFromSubscription), namespace, name)
}

// IsMetalProvisioningExists mocks base method
func (m *MockK8SClient) IsMetalProvisioningExists() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMetalProvisioningExists")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMetalProvisioningExists indicates an expected call of IsMetalProvisioningExists
func (mr *MockK8SClientMockRecorder) IsMetalProvisioningExists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMetalProvisioningExists", reflect.TypeOf((*MockK8SClient)(nil).IsMetalProvisioningExists))
}

// ListBMHs mocks base method
func (m *MockK8SClient) ListBMHs() (v1alpha1.BareMetalHostList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBMHs")
	ret0, _ := ret[0].(v1alpha1.BareMetalHostList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBMHs indicates an expected call of ListBMHs
func (mr *MockK8SClientMockRecorder) ListBMHs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBMHs", reflect.TypeOf((*MockK8SClient)(nil).ListBMHs))
}

// GetBMH mocks base method
func (m *MockK8SClient) GetBMH(name string) (*v1alpha1.BareMetalHost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBMH", name)
	ret0, _ := ret[0].(*v1alpha1.BareMetalHost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBMH indicates an expected call of GetBMH
func (mr *MockK8SClientMockRecorder) GetBMH(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBMH", reflect.TypeOf((*MockK8SClient)(nil).GetBMH), name)
}

// UpdateBMHStatus mocks base method
func (m *MockK8SClient) UpdateBMHStatus(bmh *v1alpha1.BareMetalHost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBMHStatus", bmh)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBMHStatus indicates an expected call of UpdateBMHStatus
func (mr *MockK8SClientMockRecorder) UpdateBMHStatus(bmh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBMHStatus", reflect.TypeOf((*MockK8SClient)(nil).UpdateBMHStatus), bmh)
}

// UpdateBMH mocks base method
func (m *MockK8SClient) UpdateBMH(bmh *v1alpha1.BareMetalHost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBMH", bmh)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBMH indicates an expected call of UpdateBMH
func (mr *MockK8SClientMockRecorder) UpdateBMH(bmh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBMH", reflect.TypeOf((*MockK8SClient)(nil).UpdateBMH), bmh)
}

// SetProxyEnvVars mocks base method
func (m *MockK8SClient) SetProxyEnvVars() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProxyEnvVars")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProxyEnvVars indicates an expected call of SetProxyEnvVars
func (mr *MockK8SClientMockRecorder) SetProxyEnvVars() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProxyEnvVars", reflect.TypeOf((*MockK8SClient)(nil).SetProxyEnvVars))
}

// GetClusterVersion mocks base method
func (m *MockK8SClient) GetClusterVersion(name string) (*v1.ClusterVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterVersion", name)
	ret0, _ := ret[0].(*v1.ClusterVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterVersion indicates an expected call of GetClusterVersion
func (mr *MockK8SClientMockRecorder) GetClusterVersion(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterVersion", reflect.TypeOf((*MockK8SClient)(nil).GetClusterVersion), name)
}

// GetNetworkType mocks base method
func (m *MockK8SClient) GetNetworkType() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkType")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkType indicates an expected call of GetNetworkType
func (mr *MockK8SClientMockRecorder) GetNetworkType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkType", reflect.TypeOf((*MockK8SClient)(nil).GetNetworkType))
}

// GetControlPlaneReplicas mocks base method
func (m *MockK8SClient) GetControlPlaneReplicas() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControlPlaneReplicas")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetControlPlaneReplicas indicates an expected call of GetControlPlaneReplicas
func (mr *MockK8SClientMockRecorder) GetControlPlaneReplicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControlPlaneReplicas", reflect.TypeOf((*MockK8SClient)(nil).GetControlPlaneReplicas))
}

// ListEvents mocks base method
func (m *MockK8SClient) ListEvents(namespace string) (*v11.EventList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEvents", namespace)
	ret0, _ := ret[0].(*v11.EventList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEvents indicates an expected call of ListEvents
func (mr *MockK8SClientMockRecorder) ListEvents(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEvents", reflect.TypeOf((*MockK8SClient)(nil).ListEvents), namespace)
}

// ListClusterOperators mocks base method
func (m *MockK8SClient) ListClusterOperators() (*v1.ClusterOperatorList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusterOperators")
	ret0, _ := ret[0].(*v1.ClusterOperatorList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusterOperators indicates an expected call of ListClusterOperators
func (mr *MockK8SClientMockRecorder) ListClusterOperators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusterOperators", reflect.TypeOf((*MockK8SClient)(nil).ListClusterOperators))
}

// GetClusterOperator mocks base method
func (m *MockK8SClient) GetClusterOperator(name string) (*v1.ClusterOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterOperator", name)
	ret0, _ := ret[0].(*v1.ClusterOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterOperator indicates an expected call of GetClusterOperator
func (mr *MockK8SClientMockRecorder) GetClusterOperator(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterOperator", reflect.TypeOf((*MockK8SClient)(nil).GetClusterOperator), name)
}

// CreateEvent mocks base method
func (m *MockK8SClient) CreateEvent(namespace, name, message, component string) (*v11.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", namespace, name, message, component)
	ret0, _ := ret[0].(*v11.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent
func (mr *MockK8SClientMockRecorder) CreateEvent(namespace, name, message, component interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockK8SClient)(nil).CreateEvent), namespace, name, message, component)
}
