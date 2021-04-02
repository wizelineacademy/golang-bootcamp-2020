// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/interactor/infoclient.go

// Package mock_interactor is a generated GoMock package.
package mock_interactor

import (
	reflect "reflect"

	model "github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	resty "github.com/go-resty/resty/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockInfoClient is a mock of InfoClient interface.
type MockInfoClient struct {
	ctrl     *gomock.Controller
	recorder *MockInfoClientMockRecorder
}

// MockInfoClientMockRecorder is the mock recorder for MockInfoClient.
type MockInfoClientMockRecorder struct {
	mock *MockInfoClient
}

// NewMockInfoClient creates a new mock instance.
func NewMockInfoClient(ctrl *gomock.Controller) *MockInfoClient {
	mock := &MockInfoClient{ctrl: ctrl}
	mock.recorder = &MockInfoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInfoClient) EXPECT() *MockInfoClientMockRecorder {
	return m.recorder
}

// FetchBase mocks base method.
func (m *MockInfoClient) FetchBase(pokemon *model.Pokemon, client *resty.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBase", pokemon, client)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchBase indicates an expected call of FetchBase.
func (mr *MockInfoClientMockRecorder) FetchBase(pokemon, client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBase", reflect.TypeOf((*MockInfoClient)(nil).FetchBase), pokemon, client)
}

// FetchSpecies mocks base method.
func (m *MockInfoClient) FetchSpecies(pokemon *model.Pokemon, client *resty.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSpecies", pokemon, client)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchSpecies indicates an expected call of FetchSpecies.
func (mr *MockInfoClientMockRecorder) FetchSpecies(pokemon, client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSpecies", reflect.TypeOf((*MockInfoClient)(nil).FetchSpecies), pokemon, client)
}
