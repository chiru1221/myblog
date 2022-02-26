package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/bap/service"
	"example.com/bap/slack"
	"example.com/bap/util/data"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name           string
		srvMockFn      func(*service.MockService)
		slackSrvMockFn func(*slack.MockSlack)
		wantResult     int
	}

	var (
		testBlog []data.Blog
		w        *httptest.ResponseRecorder
		r        *http.Request
	)

	tests := []testCase{
		{
			name: "Auth ok",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().Blogs().Return(testBlog, nil)
			},
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().VerifyWebHook(gomock.Any()).Return(true, nil)
			},
			wantResult: 200,
		},
		{
			name:      "Auth Ng",
			srvMockFn: func(m *service.MockService) {},
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().VerifyWebHook(gomock.Any()).Return(false, nil)
			},
			wantResult: 403,
		},
	}

	mockSrv := service.NewMockService(ctrl)
	mockSlackSrv := slack.NewMockSlack(ctrl)
	handler := Handler{Srv: mockSrv, SlackSrv: mockSlackSrv}
	router := mux.NewRouter()
	router.HandleFunc("/view", handler.viewBlog).Methods(http.MethodPost)
	router.Use(handler.authMiddleware)
	r = httptest.NewRequest("POST", "/view", nil)
	for _, tc := range tests {
		tc.srvMockFn(mockSrv)
		tc.slackSrvMockFn(mockSlackSrv)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, r)
		if w.Code != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %d, but got: %d", tc.name, tc.wantResult, w.Code)
		}
	}
}

func TestView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name       string
		srvMockFn  func(*service.MockService)
		wantResult int
	}

	var (
		testBlog  []data.Blog
		testError = errors.New("Error")
		w         *httptest.ResponseRecorder
		r         *http.Request
	)

	tests := []testCase{
		{
			name: "All ok",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().Blogs().Return(testBlog, nil)
			},
			wantResult: 200,
		},
		{
			name: "Service error",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().Blogs().Return(nil, testError)
			},
			wantResult: 500,
		},
	}

	mockSrv := service.NewMockService(ctrl)
	handler := Handler{Srv: mockSrv, SlackSrv: nil}
	r = httptest.NewRequest("POST", "/view", nil)
	for _, tc := range tests {
		tc.srvMockFn(mockSrv)
		w = httptest.NewRecorder()
		http.HandlerFunc(handler.viewBlog).ServeHTTP(w, r)
		if w.Code != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %d, but got: %d", tc.name, tc.wantResult, w.Code)
		}
	}
}

func TestNewTrigger(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name           string
		slackSrvMockFn func(*slack.MockSlack)
		wantResult     int
	}

	var (
		testError = errors.New("Error")
		w         *httptest.ResponseRecorder
		r         *http.Request
	)

	tests := []testCase{
		{
			name: "All ok",
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().TriggerId(gomock.Any()).Return("", nil)
				m.EXPECT().DialogApi(gomock.Any()).Return(nil)
			},
			wantResult: 200,
		},
		{
			name: "Trigger error",
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().TriggerId(gomock.Any()).Return("", testError)
			},
			wantResult: 500,
		},
		{
			name: "DialogApi error",
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().TriggerId(gomock.Any()).Return("", nil)
				m.EXPECT().DialogApi(gomock.Any()).Return(testError)
			},
			wantResult: 500,
		},
	}

	mockSlackSrv := slack.NewMockSlack(ctrl)
	handler := Handler{Srv: nil, SlackSrv: mockSlackSrv}
	r = httptest.NewRequest("POST", "/new", nil)
	for _, tc := range tests {
		tc.slackSrvMockFn(mockSlackSrv)
		w = httptest.NewRecorder()
		http.HandlerFunc(handler.newTrigger).ServeHTTP(w, r)
		if w.Code != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %d, but got: %d", tc.name, tc.wantResult, w.Code)
		}
	}
}

func TestUpdateTrigger(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name           string
		slackSrvMockFn func(*slack.MockSlack)
		wantResult     int
	}

	var (
		testError = errors.New("Error")
		w         *httptest.ResponseRecorder
		r         *http.Request
	)

	tests := []testCase{
		{
			name: "All ok",
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().TriggerId(gomock.Any()).Return("", nil)
				m.EXPECT().DialogApi(gomock.Any()).Return(nil)
			},
			wantResult: 200,
		},
		{
			name: "Trigger error",
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().TriggerId(gomock.Any()).Return("", testError)
			},
			wantResult: 500,
		},
		{
			name: "DialogApi error",
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().TriggerId(gomock.Any()).Return("", nil)
				m.EXPECT().DialogApi(gomock.Any()).Return(testError)
			},
			wantResult: 500,
		},
	}

	mockSlackSrv := slack.NewMockSlack(ctrl)
	handler := Handler{Srv: nil, SlackSrv: mockSlackSrv}
	r = httptest.NewRequest("POST", "/update", nil)
	for _, tc := range tests {
		tc.slackSrvMockFn(mockSlackSrv)
		w = httptest.NewRecorder()
		http.HandlerFunc(handler.updateTrigger).ServeHTTP(w, r)
		if w.Code != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %d, but got: %d", tc.name, tc.wantResult, w.Code)
		}
	}
}

func TestSubmit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name           string
		srvMockFn      func(*service.MockService)
		slackSrvMockFn func(*slack.MockSlack)
		wantResult     int
	}

	var (
		testError = errors.New("Error")
		w         *httptest.ResponseRecorder
		r         *http.Request
	)

	createPayload := func(callbackId string) (payload string) {
		payload = fmt.Sprintf(`{
			"callback_id":"%s",
			"type":"dialog_submission",
			"response_url":"http://",
			"submission":{
				"id":"0000",
				"title":"test",
				"article":"https",
				"tag":"test,test",
				"open":"false",
				"date":"2022/01/01"
			}
		}`, callbackId)
		payload = strings.ReplaceAll(payload, "\t", "")
		payload = strings.ReplaceAll(payload, "\n", "")
		payload = strings.ReplaceAll(payload, "\r", "")
		payload = strings.ReplaceAll(payload, " ", "")
		return payload
	}

	tests := []testCase{
		{
			name: "All ok / new",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().NewBlog(gomock.Any()).Return(nil)
			},
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().Payload(gomock.Any()).Return(
					createPayload("new-blog"), nil,
				)
				m.EXPECT().MergePayload(gomock.Any(), gomock.Any())
				m.EXPECT().SendResponse(gomock.Any()).Return(nil)
			},
			wantResult: 200,
		},
		{
			name: "All ok / update",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().BlogMeta(gomock.Any()).Return(nil, nil)
				m.EXPECT().UpdateBlog(gomock.Any(), gomock.Any()).Return(nil)
			},
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().Payload(gomock.Any()).Return(
					createPayload("update-blog"), nil,
				)
				m.EXPECT().MergePayload(gomock.Any(), gomock.Any())
				m.EXPECT().SendResponse(gomock.Any()).Return(nil)
			},
			wantResult: 200,
		},
		{
			name:      "Payload error",
			srvMockFn: func(m *service.MockService) {},
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().Payload(gomock.Any()).Return("", testError)
			},
			wantResult: 500,
		},
		{
			name: "New blog error",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().NewBlog(gomock.Any()).Return(testError)
			},
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().Payload(gomock.Any()).Return(
					createPayload("new-blog"), nil,
				)
				m.EXPECT().MergePayload(gomock.Any(), gomock.Any())
			},
			wantResult: 500,
		},
		{
			name: "Read blog error",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().BlogMeta(gomock.Any()).Return(nil, testError)
			},
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().Payload(gomock.Any()).Return(
					createPayload("update-blog"), nil,
				)
			},
			wantResult: 500,
		},
		{
			name: "Update blog error",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().BlogMeta(gomock.Any()).Return(nil, nil)
				m.EXPECT().UpdateBlog(gomock.Any(), gomock.Any()).Return(testError)
			},
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().Payload(gomock.Any()).Return(
					createPayload("update-blog"), nil,
				)
				m.EXPECT().MergePayload(gomock.Any(), gomock.Any())
			},
			wantResult: 500,
		},
		{
			name:      "Callback error",
			srvMockFn: func(m *service.MockService) {},
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().Payload(gomock.Any()).Return("", nil)
			},
			wantResult: 500,
		},
		{
			name: "Send res error",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().NewBlog(gomock.Any()).Return(nil)
			},
			slackSrvMockFn: func(m *slack.MockSlack) {
				m.EXPECT().Payload(gomock.Any()).Return(
					createPayload("new-blog"), nil,
				)
				m.EXPECT().MergePayload(gomock.Any(), gomock.Any())
				m.EXPECT().SendResponse(gomock.Any()).Return(testError)
			},
			wantResult: 500,
		},
	}

	mockSrv := service.NewMockService(ctrl)
	mockSlackSrv := slack.NewMockSlack(ctrl)
	handler := Handler{Srv: mockSrv, SlackSrv: mockSlackSrv}
	r = httptest.NewRequest("POST", "/update", nil)
	for _, tc := range tests {
		tc.srvMockFn(mockSrv)
		tc.slackSrvMockFn(mockSlackSrv)
		w = httptest.NewRecorder()
		http.HandlerFunc(handler.submit).ServeHTTP(w, r)
		if w.Code != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %d, but got: %d", tc.name, tc.wantResult, w.Code)
		}
	}
}
