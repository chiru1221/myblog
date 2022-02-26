package service_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"example.com/bap/drive"
	"example.com/bap/mongodb"
	"example.com/bap/service"
	"example.com/bap/util/data"

	"github.com/golang/mock/gomock"
)

func TestBlogOpenFilter(t *testing.T) {
	type testCase struct {
		name       string
		input      *data.Blog
		wantResult *data.Blog
	}

	tests := []testCase{
		{
			name:       "Open is ture",
			input:      &data.Blog{Title: "test", Open: true},
			wantResult: &data.Blog{Title: "test", Open: true},
		},
		{
			name:       "Open is false",
			input:      &data.Blog{Title: "test", Open: false},
			wantResult: nil,
		},
	}

	var srv service.ServiceImpl
	for _, tc := range tests {
		result := srv.BlogOpenFilter(tc.input)
		if tc.wantResult == nil {
			if result != nil {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult, result)
			}
		} else {
			if result.Title != tc.wantResult.Title {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult, result)
			}
		}
	}
}

func TestBlogsOpenFilter(t *testing.T) {
	type testCase struct {
		name       string
		input      []data.Blog
		wantResult []data.Blog
	}

	tests := []testCase{
		{
			name: "All open are ture",
			input: []data.Blog{
				{Title: "test1", Open: true},
				{Title: "test2", Open: true},
			},
			wantResult: []data.Blog{
				{Title: "test1", Open: true},
				{Title: "test2", Open: true},
			},
		},
		{
			name: "All open are false",
			input: []data.Blog{
				{Title: "test1", Open: false},
				{Title: "test2", Open: false},
			},
			wantResult: []data.Blog{},
		},
	}
	var srv service.ServiceImpl
	for _, tc := range tests {
		result := srv.BlogsOpenFilter(tc.input)
		if len(result) != len(tc.wantResult) {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult, result)
		} else {
			for i := 0; i < len(result); i++ {
				if result[i].Title != result[i].Title {
					t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult[i], result[i])
				}
			}
		}
	}
}

func TestBlogsArticleMask(t *testing.T) {
	type testCase struct {
		name       string
		input      []data.Blog
		wantResult []data.Blog
	}

	tests := []testCase{
		{
			name:       "Single slice",
			input:      []data.Blog{{Article: "test"}},
			wantResult: []data.Blog{{Article: ""}},
		},
		{
			name:       "Multiple slice",
			input:      []data.Blog{{Article: "test"}, {Article: "hoge"}},
			wantResult: []data.Blog{{Article: ""}, {Article: ""}},
		},
	}
	var srv service.ServiceImpl
	for _, tc := range tests {
		result := srv.BlogsArticleMask(tc.input)
		if len(result) != len(tc.wantResult) {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult, result)
		} else {
			for i := 0; i < len(result); i++ {
				if result[i].Article != tc.wantResult[i].Article {
					t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult[i].Article, result[i].Article)
				}
			}
		}
	}
}

func TestProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name            string
		dbMockFn        func(*mongodb.MockDB)
		driveMockFn     func(*drive.MockDrive)
		srvMockFn       func(*service.MockService)
		wantValueResult *data.Profile
		wantErrorResult error
	}

	var (
		testProfile = &data.Profile{Content: ""}
		testError   = errors.New("Error")
		testResp    = http.Response{Body: io.NopCloser(bytes.NewReader([]byte("")))}
	)
	tests := []testCase{
		{
			name: "All ok",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadProfile().Return(testProfile, nil)
			},
			driveMockFn: func(m *drive.MockDrive) {
				m.EXPECT().Export(gomock.Any(), gomock.Any(), gomock.Any()).Return(&testResp, nil)
			},
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().ZipToHtml(gomock.Any()).Return("", nil)
			},
			wantValueResult: testProfile,
			wantErrorResult: nil,
		},
		{
			name: "DB error",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadProfile().Return(nil, testError)
			},
			driveMockFn:     func(m *drive.MockDrive) {},
			srvMockFn:       func(m *service.MockService) {},
			wantValueResult: nil,
			wantErrorResult: testError,
		},
		{
			name: "Drive error",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadProfile().Return(testProfile, nil)
			},
			driveMockFn: func(m *drive.MockDrive) {
				m.EXPECT().Export(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, testError)
			},
			srvMockFn:       func(m *service.MockService) {},
			wantValueResult: nil,
			wantErrorResult: testError,
		},
		{
			name: "Zip error",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadProfile().Return(testProfile, nil)
			},
			driveMockFn: func(m *drive.MockDrive) {
				m.EXPECT().Export(gomock.Any(), gomock.Any(), gomock.Any()).Return(&testResp, nil)
			},
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().ZipToHtml(gomock.Any()).Return("", testError)
			},
			wantValueResult: nil,
			wantErrorResult: testError,
		},
	}

	mockDB := mongodb.NewMockDB(ctrl)
	mockDrive := drive.NewMockDrive(ctrl)
	mockSrv := service.NewMockService(ctrl)
	srv := service.ServiceImpl{nil, mockDB, mockDrive, mockSrv}
	for _, tc := range tests {
		tc.dbMockFn(mockDB)
		tc.driveMockFn(mockDrive)
		tc.srvMockFn(mockSrv)
		result, err := srv.Profile()
		if result != tc.wantValueResult {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantValueResult, result)
		}
		if err != tc.wantErrorResult {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantErrorResult, err)
		}
	}
}

func TestBlog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name            string
		dbMockFn        func(*mongodb.MockDB)
		driveMockFn     func(*drive.MockDrive)
		srvMockFn       func(*service.MockService)
		wantValueResult *data.Blog
		wantErrorResult error
	}

	var (
		testBlog  = &data.Blog{Article: ""}
		testError = errors.New("Error")
		testResp  = http.Response{Body: io.NopCloser(bytes.NewReader([]byte("")))}
	)
	tests := []testCase{
		{
			name: "All ok",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadBlog(gomock.Any()).Return(testBlog, nil)
			},
			driveMockFn: func(m *drive.MockDrive) {
				m.EXPECT().Export(gomock.Any(), gomock.Any(), gomock.Any()).Return(&testResp, nil)
			},
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().ZipToHtml(gomock.Any()).Return("", nil)
			},
			wantValueResult: testBlog,
			wantErrorResult: nil,
		},
		{
			name: "DB error",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadBlog(gomock.Any()).Return(nil, testError)
			},
			driveMockFn:     func(m *drive.MockDrive) {},
			srvMockFn:       func(m *service.MockService) {},
			wantValueResult: nil,
			wantErrorResult: testError,
		},
		{
			name: "Drive error",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadBlog(gomock.Any()).Return(testBlog, nil)
			},
			driveMockFn: func(m *drive.MockDrive) {
				m.EXPECT().Export(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, testError)
			},
			srvMockFn:       func(m *service.MockService) {},
			wantValueResult: nil,
			wantErrorResult: testError,
		},
		{
			name: "Zip error",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadBlog(gomock.Any()).Return(testBlog, nil)
			},
			driveMockFn: func(m *drive.MockDrive) {
				m.EXPECT().Export(gomock.Any(), gomock.Any(), gomock.Any()).Return(&testResp, nil)
			},
			srvMockFn: func(m *service.MockService) {
				m.EXPECT().ZipToHtml(gomock.Any()).Return("", testError)
			},
			wantValueResult: nil,
			wantErrorResult: testError,
		},
	}

	mockDB := mongodb.NewMockDB(ctrl)
	mockDrive := drive.NewMockDrive(ctrl)
	mockSrv := service.NewMockService(ctrl)
	srv := service.ServiceImpl{mockDB, nil, mockDrive, mockSrv}
	for _, tc := range tests {
		tc.dbMockFn(mockDB)
		tc.driveMockFn(mockDrive)
		tc.srvMockFn(mockSrv)
		result, err := srv.Blog("")
		if result != tc.wantValueResult {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantValueResult, result)
		}
		if err != tc.wantErrorResult {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantErrorResult, err)
		}
	}
}

func TestBlogMeta(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name            string
		dbMockFn        func(*mongodb.MockDB)
		wantValueResult *data.Blog
		wantErrorResult error
	}

	var (
		testBlog  = &data.Blog{}
		testError = errors.New("Error")
	)
	tests := []testCase{
		{
			name: "All ok",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadBlog(gomock.Any()).Return(testBlog, nil)
			},
			wantValueResult: testBlog,
			wantErrorResult: nil,
		},
		{
			name: "DB error",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadBlog(gomock.Any()).Return(nil, testError)
			},
			wantValueResult: nil,
			wantErrorResult: testError,
		},
	}

	mockDB := mongodb.NewMockDB(ctrl)
	srv := service.ServiceImpl{mockDB, nil, nil, nil}
	for _, tc := range tests {
		tc.dbMockFn(mockDB)
		result, err := srv.BlogMeta("test")
		if result != tc.wantValueResult {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantValueResult, result)
		}
		if err != tc.wantErrorResult {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantErrorResult, err)
		}
	}
}

func TestBlogs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name            string
		dbMockFn        func(*mongodb.MockDB)
		wantValueResult []data.Blog
		wantErrorResult error
	}

	var (
		testBlog  = []data.Blog{{Article: ""}}
		testError = errors.New("Error")
	)
	tests := []testCase{
		{
			name: "All ok",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadBlogs().Return(testBlog, nil)
			},
			wantValueResult: testBlog,
			wantErrorResult: nil,
		},
		{
			name: "DB error",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().ReadBlogs().Return(nil, testError)
			},
			wantValueResult: nil,
			wantErrorResult: testError,
		},
	}

	mockDB := mongodb.NewMockDB(ctrl)
	srv := service.ServiceImpl{mockDB, nil, nil, nil}
	for _, tc := range tests {
		tc.dbMockFn(mockDB)
		result, err := srv.Blogs()
		for i := range tc.wantValueResult {
			if result[i].Article != tc.wantValueResult[i].Article {
				t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantValueResult[i], result[i])
			}
		}
		if err != tc.wantErrorResult {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantErrorResult, err)
		}
	}
}

func TestNewBlog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name       string
		dbMockFn   func(*mongodb.MockDB)
		wantResult error
	}
	var testError = errors.New("Error")

	tests := []testCase{
		{
			name: "All ok",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().InsertBlog(gomock.Any()).Return(nil)
			},
			wantResult: nil,
		},
		{
			name: "DB error",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().InsertBlog(gomock.Any()).Return(testError)
			},
			wantResult: testError,
		},
	}
	mockDB := mongodb.NewMockDB(ctrl)
	srv := service.ServiceImpl{mockDB, nil, nil, nil}

	for _, tc := range tests {
		tc.dbMockFn(mockDB)

		err := srv.NewBlog(&data.Blog{})
		if err != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult, err)
		}
	}
}

func TestUpdateBlog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testCase struct {
		name       string
		dbMockFn   func(*mongodb.MockDB)
		wantResult error
	}

	var testError = errors.New("Error")

	tests := []testCase{
		{
			name: "All ok",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().UpdateBlog(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantResult: nil,
		},
		{
			name: "DB error",
			dbMockFn: func(m *mongodb.MockDB) {
				m.EXPECT().UpdateBlog(gomock.Any(), gomock.Any()).Return(testError)
			},
			wantResult: testError,
		},
	}
	mockDB := mongodb.NewMockDB(ctrl)
	srv := service.ServiceImpl{mockDB, nil, nil, nil}

	for _, tc := range tests {
		tc.dbMockFn(mockDB)
		err := srv.UpdateBlog(&data.Blog{}, "")
		if err != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult, err)
		}
	}
}
