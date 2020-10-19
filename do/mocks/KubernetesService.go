// Code generated by MockGen. DO NOT EDIT.
// Source: kubernetes.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "github.com/digitalocean/doctl/do"
	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockKubernetesService is a mock of KubernetesService interface.
type MockKubernetesService struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesServiceMockRecorder
}

// MockKubernetesServiceMockRecorder is the mock recorder for MockKubernetesService.
type MockKubernetesServiceMockRecorder struct {
	mock *MockKubernetesService
}

// NewMockKubernetesService creates a new mock instance.
func NewMockKubernetesService(ctrl *gomock.Controller) *MockKubernetesService {
	mock := &MockKubernetesService{ctrl: ctrl}
	mock.recorder = &MockKubernetesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesService) EXPECT() *MockKubernetesServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockKubernetesService) Get(clusterID string) (*do.KubernetesCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", clusterID)
	ret0, _ := ret[0].(*do.KubernetesCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKubernetesServiceMockRecorder) Get(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKubernetesService)(nil).Get), clusterID)
}

// GetKubeConfig mocks base method.
func (m *MockKubernetesService) GetKubeConfig(clusterID string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKubeConfig", clusterID)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKubeConfig indicates an expected call of GetKubeConfig.
func (mr *MockKubernetesServiceMockRecorder) GetKubeConfig(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubeConfig", reflect.TypeOf((*MockKubernetesService)(nil).GetKubeConfig), clusterID)
}

// GetCredentials mocks base method.
func (m *MockKubernetesService) GetCredentials(clusterID string) (*do.KubernetesClusterCredentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentials", clusterID)
	ret0, _ := ret[0].(*do.KubernetesClusterCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentials indicates an expected call of GetCredentials.
func (mr *MockKubernetesServiceMockRecorder) GetCredentials(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentials", reflect.TypeOf((*MockKubernetesService)(nil).GetCredentials), clusterID)
}

// GetUpgrades mocks base method.
func (m *MockKubernetesService) GetUpgrades(clusterID string) (do.KubernetesVersions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpgrades", clusterID)
	ret0, _ := ret[0].(do.KubernetesVersions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpgrades indicates an expected call of GetUpgrades.
func (mr *MockKubernetesServiceMockRecorder) GetUpgrades(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpgrades", reflect.TypeOf((*MockKubernetesService)(nil).GetUpgrades), clusterID)
}

// List mocks base method.
func (m *MockKubernetesService) List() (do.KubernetesClusters, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(do.KubernetesClusters)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockKubernetesServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockKubernetesService)(nil).List))
}

// Create mocks base method.
func (m *MockKubernetesService) Create(create *godo.KubernetesClusterCreateRequest) (*do.KubernetesCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", create)
	ret0, _ := ret[0].(*do.KubernetesCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockKubernetesServiceMockRecorder) Create(create interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockKubernetesService)(nil).Create), create)
}

// Update mocks base method.
func (m *MockKubernetesService) Update(clusterID string, update *godo.KubernetesClusterUpdateRequest) (*do.KubernetesCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", clusterID, update)
	ret0, _ := ret[0].(*do.KubernetesCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockKubernetesServiceMockRecorder) Update(clusterID, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockKubernetesService)(nil).Update), clusterID, update)
}

// Upgrade mocks base method.
func (m *MockKubernetesService) Upgrade(clusterID, versionSlug string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", clusterID, versionSlug)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockKubernetesServiceMockRecorder) Upgrade(clusterID, versionSlug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockKubernetesService)(nil).Upgrade), clusterID, versionSlug)
}

// Delete mocks base method.
func (m *MockKubernetesService) Delete(clusterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", clusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockKubernetesServiceMockRecorder) Delete(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKubernetesService)(nil).Delete), clusterID)
}

// CreateNodePool mocks base method.
func (m *MockKubernetesService) CreateNodePool(clusterID string, req *godo.KubernetesNodePoolCreateRequest) (*do.KubernetesNodePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNodePool", clusterID, req)
	ret0, _ := ret[0].(*do.KubernetesNodePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNodePool indicates an expected call of CreateNodePool.
func (mr *MockKubernetesServiceMockRecorder) CreateNodePool(clusterID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNodePool", reflect.TypeOf((*MockKubernetesService)(nil).CreateNodePool), clusterID, req)
}

// GetNodePool mocks base method.
func (m *MockKubernetesService) GetNodePool(clusterID, poolID string) (*do.KubernetesNodePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodePool", clusterID, poolID)
	ret0, _ := ret[0].(*do.KubernetesNodePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodePool indicates an expected call of GetNodePool.
func (mr *MockKubernetesServiceMockRecorder) GetNodePool(clusterID, poolID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodePool", reflect.TypeOf((*MockKubernetesService)(nil).GetNodePool), clusterID, poolID)
}

// ListNodePools mocks base method.
func (m *MockKubernetesService) ListNodePools(clusterID string) (do.KubernetesNodePools, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNodePools", clusterID)
	ret0, _ := ret[0].(do.KubernetesNodePools)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNodePools indicates an expected call of ListNodePools.
func (mr *MockKubernetesServiceMockRecorder) ListNodePools(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodePools", reflect.TypeOf((*MockKubernetesService)(nil).ListNodePools), clusterID)
}

// UpdateNodePool mocks base method.
func (m *MockKubernetesService) UpdateNodePool(clusterID, poolID string, req *godo.KubernetesNodePoolUpdateRequest) (*do.KubernetesNodePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNodePool", clusterID, poolID, req)
	ret0, _ := ret[0].(*do.KubernetesNodePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNodePool indicates an expected call of UpdateNodePool.
func (mr *MockKubernetesServiceMockRecorder) UpdateNodePool(clusterID, poolID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNodePool", reflect.TypeOf((*MockKubernetesService)(nil).UpdateNodePool), clusterID, poolID, req)
}

// RecycleNodePoolNodes mocks base method.
func (m *MockKubernetesService) RecycleNodePoolNodes(clusterID, poolID string, req *godo.KubernetesNodePoolRecycleNodesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecycleNodePoolNodes", clusterID, poolID, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecycleNodePoolNodes indicates an expected call of RecycleNodePoolNodes.
func (mr *MockKubernetesServiceMockRecorder) RecycleNodePoolNodes(clusterID, poolID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecycleNodePoolNodes", reflect.TypeOf((*MockKubernetesService)(nil).RecycleNodePoolNodes), clusterID, poolID, req)
}

// DeleteNodePool mocks base method.
func (m *MockKubernetesService) DeleteNodePool(clusterID, poolID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNodePool", clusterID, poolID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNodePool indicates an expected call of DeleteNodePool.
func (mr *MockKubernetesServiceMockRecorder) DeleteNodePool(clusterID, poolID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNodePool", reflect.TypeOf((*MockKubernetesService)(nil).DeleteNodePool), clusterID, poolID)
}

// DeleteNode mocks base method.
func (m *MockKubernetesService) DeleteNode(clusterID, poolID, nodeID string, req *godo.KubernetesNodeDeleteRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNode", clusterID, poolID, nodeID, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNode indicates an expected call of DeleteNode.
func (mr *MockKubernetesServiceMockRecorder) DeleteNode(clusterID, poolID, nodeID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNode", reflect.TypeOf((*MockKubernetesService)(nil).DeleteNode), clusterID, poolID, nodeID, req)
}

// GetVersions mocks base method.
func (m *MockKubernetesService) GetVersions() (do.KubernetesVersions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersions")
	ret0, _ := ret[0].(do.KubernetesVersions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersions indicates an expected call of GetVersions.
func (mr *MockKubernetesServiceMockRecorder) GetVersions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersions", reflect.TypeOf((*MockKubernetesService)(nil).GetVersions))
}

// GetRegions mocks base method.
func (m *MockKubernetesService) GetRegions() (do.KubernetesRegions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegions")
	ret0, _ := ret[0].(do.KubernetesRegions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegions indicates an expected call of GetRegions.
func (mr *MockKubernetesServiceMockRecorder) GetRegions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegions", reflect.TypeOf((*MockKubernetesService)(nil).GetRegions))
}

// GetNodeSizes mocks base method.
func (m *MockKubernetesService) GetNodeSizes() (do.KubernetesNodeSizes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeSizes")
	ret0, _ := ret[0].(do.KubernetesNodeSizes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeSizes indicates an expected call of GetNodeSizes.
func (mr *MockKubernetesServiceMockRecorder) GetNodeSizes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeSizes", reflect.TypeOf((*MockKubernetesService)(nil).GetNodeSizes))
}

// AddRegistry mocks base method.
func (m *MockKubernetesService) AddRegistry(req *godo.KubernetesClusterRegistryRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRegistry", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRegistry indicates an expected call of AddRegistry.
func (mr *MockKubernetesServiceMockRecorder) AddRegistry(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRegistry", reflect.TypeOf((*MockKubernetesService)(nil).AddRegistry), req)
}

// RemoveRegistry mocks base method.
func (m *MockKubernetesService) RemoveRegistry(req *godo.KubernetesClusterRegistryRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRegistry", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRegistry indicates an expected call of RemoveRegistry.
func (mr *MockKubernetesServiceMockRecorder) RemoveRegistry(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRegistry", reflect.TypeOf((*MockKubernetesService)(nil).RemoveRegistry), req)
}
