// Code generated by MockGen. DO NOT EDIT.
// Source: slack.go

// Package slack is a generated GoMock package.
package slack

import (
	data "example.com/bap/util/data"
	io "io"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSlack is a mock of Slack interface.
type MockSlack struct {
	ctrl     *gomock.Controller
	recorder *MockSlackMockRecorder
}

// MockSlackMockRecorder is the mock recorder for MockSlack.
type MockSlackMockRecorder struct {
	mock *MockSlack
}

// NewMockSlack creates a new mock instance.
func NewMockSlack(ctrl *gomock.Controller) *MockSlack {
	mock := &MockSlack{ctrl: ctrl}
	mock.recorder = &MockSlackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlack) EXPECT() *MockSlackMockRecorder {
	return m.recorder
}

// Construct mocks base method.
func (m *MockSlack) Construct() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Construct")
}

// Construct indicates an expected call of Construct.
func (mr *MockSlackMockRecorder) Construct() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Construct", reflect.TypeOf((*MockSlack)(nil).Construct))
}

// DialogApi mocks base method.
func (m *MockSlack) DialogApi(body io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DialogApi", body)
	ret0, _ := ret[0].(error)
	return ret0
}

// DialogApi indicates an expected call of DialogApi.
func (mr *MockSlackMockRecorder) DialogApi(body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DialogApi", reflect.TypeOf((*MockSlack)(nil).DialogApi), body)
}

// MergePayload mocks base method.
func (m *MockSlack) MergePayload(pre *data.Blog, payload *data.Payload) *data.Blog {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergePayload", pre, payload)
	ret0, _ := ret[0].(*data.Blog)
	return ret0
}

// MergePayload indicates an expected call of MergePayload.
func (mr *MockSlackMockRecorder) MergePayload(pre, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergePayload", reflect.TypeOf((*MockSlack)(nil).MergePayload), pre, payload)
}

// Payload mocks base method.
func (m *MockSlack) Payload(r *http.Request) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Payload", r)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Payload indicates an expected call of Payload.
func (mr *MockSlackMockRecorder) Payload(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Payload", reflect.TypeOf((*MockSlack)(nil).Payload), r)
}

// SendResponse mocks base method.
func (m *MockSlack) SendResponse(responseURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendResponse", responseURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendResponse indicates an expected call of SendResponse.
func (mr *MockSlackMockRecorder) SendResponse(responseURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendResponse", reflect.TypeOf((*MockSlack)(nil).SendResponse), responseURL)
}

// TriggerId mocks base method.
func (m *MockSlack) TriggerId(r *http.Request) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerId", r)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerId indicates an expected call of TriggerId.
func (mr *MockSlackMockRecorder) TriggerId(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerId", reflect.TypeOf((*MockSlack)(nil).TriggerId), r)
}

// VerifyWebHook mocks base method.
func (m *MockSlack) VerifyWebHook(r *http.Request) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyWebHook", r)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyWebHook indicates an expected call of VerifyWebHook.
func (mr *MockSlackMockRecorder) VerifyWebHook(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyWebHook", reflect.TypeOf((*MockSlack)(nil).VerifyWebHook), r)
}
