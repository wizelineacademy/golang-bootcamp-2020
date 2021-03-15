// Code generated by MockGen. DO NOT EDIT.
// Source: service/http/http.go

// Package mock is a generated GoMock package.
package mock

import (
	model "pokeapi/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNewHttpService is a mock of NewHttpService interface.
type MockNewHttpService struct {
	ctrl     *gomock.Controller
	recorder *MockNewHttpServiceMockRecorder
}

// MockNewHttpServiceMockRecorder is the mock recorder for MockNewHttpService.
type MockNewHttpServiceMockRecorder struct {
	mock *MockNewHttpService
}

// NewMockNewHttpService creates a new mock instance.
func NewMockNewHttpService(ctrl *gomock.Controller) *MockNewHttpService {
	mock := &MockNewHttpService{ctrl: ctrl}
	mock.recorder = &MockNewHttpServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNewHttpService) EXPECT() *MockNewHttpServiceMockRecorder {
	return m.recorder
}

// GetPokemons mocks base method.
func (m *MockNewHttpService) GetPokemons() ([]model.SinglePokeExternal, *model.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPokemons")
	ret0, _ := ret[0].([]model.SinglePokeExternal)
	ret1, _ := ret[1].(*model.Error)
	return ret0, ret1
}

// GetPokemons indicates an expected call of GetPokemons.
func (mr *MockNewHttpServiceMockRecorder) GetPokemons() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPokemons", reflect.TypeOf((*MockNewHttpService)(nil).GetPokemons))
}
