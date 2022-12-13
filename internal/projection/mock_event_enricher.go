// Code generated by MockGen. DO NOT EDIT.
// Source: event_enricher.go

// Package projection is a generated GoMock package.
package projection

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/openshift-assisted/assisted-events-streams/internal/types"
)

// MockEventEnricherInterface is a mock of EventEnricherInterface interface.
type MockEventEnricherInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEventEnricherInterfaceMockRecorder
}

// MockEventEnricherInterfaceMockRecorder is the mock recorder for MockEventEnricherInterface.
type MockEventEnricherInterfaceMockRecorder struct {
	mock *MockEventEnricherInterface
}

// NewMockEventEnricherInterface creates a new mock instance.
func NewMockEventEnricherInterface(ctrl *gomock.Controller) *MockEventEnricherInterface {
	mock := &MockEventEnricherInterface{ctrl: ctrl}
	mock.recorder = &MockEventEnricherInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventEnricherInterface) EXPECT() *MockEventEnricherInterfaceMockRecorder {
	return m.recorder
}

// GetEnrichedEvent mocks base method.
func (m *MockEventEnricherInterface) GetEnrichedEvent(event *types.Event, cluster map[string]interface{}, hosts, infraEnvs []map[string]interface{}) *types.EnrichedEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnrichedEvent", event, cluster, hosts, infraEnvs)
	ret0, _ := ret[0].(*types.EnrichedEvent)
	return ret0
}

// GetEnrichedEvent indicates an expected call of GetEnrichedEvent.
func (mr *MockEventEnricherInterfaceMockRecorder) GetEnrichedEvent(event, cluster, hosts, infraEnvs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnrichedEvent", reflect.TypeOf((*MockEventEnricherInterface)(nil).GetEnrichedEvent), event, cluster, hosts, infraEnvs)
}
