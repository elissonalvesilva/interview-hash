package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidateParam is a mock of ValidateParam interface.
type MockValidateParam struct {
	ctrl     *gomock.Controller
	recorder *MockValidateParamMockRecorder
}

// MockValidateParamMockRecorder is the mock recorder for MockValidateParam.
type MockValidateParamMockRecorder struct {
	mock *MockValidateParam
}

// NewMockValidateParam creates a new mock instance.
func NewMockValidateParam(ctrl *gomock.Controller) *MockValidateParam {
	mock := &MockValidateParam{ctrl: ctrl}
	mock.recorder = &MockValidateParamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidateParam) EXPECT() *MockValidateParamMockRecorder {
	return m.recorder
}

// ValidateRequestParams mocks base method.
func (m *MockValidateParam) ValidateRequestParams(params interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRequestParams", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateRequestParams indicates an expected call of ValidateRequestParams.
func (mr *MockValidateParamMockRecorder) ValidateRequestParams(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRequestParams", reflect.TypeOf((*MockValidateParam)(nil).ValidateRequestParams), params)
}
