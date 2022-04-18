// Code generated by MockGen. DO NOT EDIT.
// Source: release/client.go

// Package mock_release is a generated GoMock package.
package mock_release

import (
	context "context"
	reflect "reflect"

	release "github.com/forta-network/forta-core-go/release"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetReleaseManifest mocks base method.
func (m *MockClient) GetReleaseManifest(ctx context.Context, reference string) (*release.ReleaseManifest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReleaseManifest", ctx, reference)
	ret0, _ := ret[0].(*release.ReleaseManifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReleaseManifest indicates an expected call of GetReleaseManifest.
func (mr *MockClientMockRecorder) GetReleaseManifest(ctx, reference interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReleaseManifest", reflect.TypeOf((*MockClient)(nil).GetReleaseManifest), ctx, reference)
}
