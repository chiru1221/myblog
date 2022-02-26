package main

import (
	"context"
	"errors"
	"testing"

	pb "example.com/bap/blogprofile"
	"example.com/bap/service"
	"example.com/bap/util/data"

	"github.com/golang/mock/gomock"
)

func TestBlogsGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name       string
		srvMockFn  func(*service.MockService)
		wantResult error
	}

	var (
		testBlog  []data.Blog
		testError = errors.New("Error")
	)

	tests := []testCase{
		{
			name: "All ok",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().Blogs().Return(testBlog, nil)
				m.EXPECT().BlogsOpenFilter(gomock.Any()).Return(testBlog)
				m.EXPECT().BlogsArticleMask(gomock.Any()).Return(testBlog)
				m.EXPECT().BlogsToBlogList(gomock.Any())
			},
			wantResult: nil,
		},
		{
			name: "Service error",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().Blogs().Return(nil, testError)
			},
			wantResult: testError,
		},
	}

	mockSrv := service.NewMockService(ctrl)
	server := Server{Srv: mockSrv}
	for _, tc := range tests {
		tc.srvMockFn(mockSrv)
		_, err := server.Blogs(context.Background(), &pb.NoId{})
		if err != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %d, but got: %d", tc.name, tc.wantResult, err)
		}
	}
}

func TestBlogGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name       string
		srvMockFn  func(*service.MockService)
		wantResult error
	}

	var (
		testBlog  *data.Blog
		testError = errors.New("Error")
	)

	tests := []testCase{
		{
			name: "All ok",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().Blog(gomock.Any()).Return(testBlog, nil)
				m.EXPECT().BlogOpenFilter(gomock.Any()).Return(testBlog)
				m.EXPECT().BlogToBlogDetail(gomock.Any())
			},
			wantResult: nil,
		},
		{
			name: "Service error",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().Blog(gomock.Any()).Return(nil, testError)
			},
			wantResult: testError,
		},
	}
	mockSrv := service.NewMockService(ctrl)
	server := Server{Srv: mockSrv}
	for _, tc := range tests {
		tc.srvMockFn(mockSrv)
		_, err := server.Blog(context.Background(), &pb.BlogId{Id: "0000"})
		if err != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %d, but got: %d", tc.name, tc.wantResult, err)
		}
	}
}

func TestProfileGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name       string
		srvMockFn  func(*service.MockService)
		wantResult error
	}

	var (
		testProfile *data.Profile
		testError   = errors.New("Error")
	)

	tests := []testCase{
		{
			name: "All ok",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().Profile().Return(testProfile, nil)
				m.EXPECT().ProfileToProfileDetail(gomock.Any())
			},
			wantResult: nil,
		},
		{
			name: "Service error",
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().Profile().Return(nil, testError)
			},
			wantResult: testError,
		},
	}

	mockSrv := service.NewMockService(ctrl)
	server := Server{Srv: mockSrv}
	for _, tc := range tests {
		tc.srvMockFn(mockSrv)
		_, err := server.Profile(context.Background(), &pb.NoId{})
		if err != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %d, but got: %d", tc.name, tc.wantResult, err)
		}
	}
}
