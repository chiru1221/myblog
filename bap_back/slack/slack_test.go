/*
	Reference:
	https://github.com/GoogleCloudPlatform/golang-samples/blob/master/functions/slack/signing_test.go
*/
package slack

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"example.com/bap/util/data"
)

func TestDialogApi(t *testing.T) {
	type testCase struct {
		name       string
		url        string
		wantResult error
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	tests := []testCase{
		{
			name:       "All ok",
			url:        ts.URL,
			wantResult: nil,
		},
	}

	for _, tc := range tests {
		si := SlackImpl{DialogURL: tc.url}
		result := si.DialogApi(nil)
		if result != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %v, but got: %v,", tc.name, tc.wantResult, result)
		}
	}
}

func TestSendResponse(t *testing.T) {
	type testCase struct {
		name       string
		url        string
		wantResult error
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	tests := []testCase{
		{
			name:       "All ok",
			url:        ts.URL,
			wantResult: nil,
		},
	}

	si := SlackImpl{}
	for _, tc := range tests {
		result := si.SendResponse(tc.url)
		if result != tc.wantResult {
			t.Errorf("Test: %s - Wanted: %v, but got: %v,", tc.name, tc.wantResult, result)
		}
	}
}

func TestMergePayload(t *testing.T) {
	type testCase struct {
		name         string
		inputPre     *data.Blog
		inputPayload *data.Payload
		wantResult   *data.Blog
	}

	tests := []testCase{
		{
			name:     "Full",
			inputPre: &data.Blog{},
			inputPayload: &data.Payload{
				Body: data.SubmitBody{
					Id:      "0000",
					Title:   "test",
					Tag:     "test,hoge",
					Article: "abc",
					Open:    "true",
					Date:    "2022/01/01",
				},
			},
			wantResult: &data.Blog{
				Title:   "test",
				Tag:     []string{"test", "hoge"},
				Article: "abc",
				Open:    true,
				Date:    "2022/01/01",
			},
		},
		{
			name:     "Part",
			inputPre: &data.Blog{},
			inputPayload: &data.Payload{
				Body: data.SubmitBody{
					Id:   "0000",
					Tag:  "test,hoge",
					Open: "true",
					Date: "2022/01/01",
				},
			},
			wantResult: &data.Blog{
				Tag:  []string{"test", "hoge"},
				Open: true,
				Date: "2022/01/01",
			},
		},
	}

	si := SlackImpl{}
	for _, tc := range tests {
		result := si.MergePayload(tc.inputPre, tc.inputPayload)
		if result.Title != tc.wantResult.Title {
			t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, tc.wantResult.Title, result.Title)
		}
		if len(result.Tag) != len(tc.wantResult.Tag) {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Tag, result.Tag)
		} else {
			for i := 0; i < len(result.Tag); i++ {
				if result.Tag[i] != tc.wantResult.Tag[i] {
					t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Tag, result.Tag)
				}
			}
		}
		if result.Article != tc.wantResult.Article {
			t.Errorf("Test: %s - Wanted: %s, but got: %s", tc.name, tc.wantResult.Article, result.Article)
		}
		if result.Open != tc.wantResult.Open {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Open, result.Open)
		}
		if result.Date != tc.wantResult.Date {
			t.Errorf("Test: %s - Wanted: %v, but got: %v", tc.name, tc.wantResult.Date, result.Date)
		}
	}
}

func TestSigningWithSecret(t *testing.T) {
	type testCase struct {
		name       string
		signature  string
		timeStamp  string
		wantResult bool
	}

	secret := "slack_ss"
	version := "v0"
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	body := "somebody"
	base := fmt.Sprintf("%s:%s:%s", version, ts, body)
	correctSHA2Signature := fmt.Sprintf("%s=%s", version, hex.EncodeToString(GetSignature([]byte(base), []byte(secret))))

	tests := []testCase{
		{name: "Good request", signature: correctSHA2Signature, timeStamp: ts, wantResult: true},
		{name: "Bad signature", signature: "v0=146abde6763faeba19adc4d9fe4961668f4be11f7405a1c05b636f29312eac2e", timeStamp: ts, wantResult: false},
		{name: "Old timestamp", signature: correctSHA2Signature, timeStamp: "12345", wantResult: false},
	}

	var slackAuth = SlackImpl{SlackSigningSecret: secret}

	for _, tc := range tests {
		rq := httptest.NewRequest("POST", "https://someurl.com", strings.NewReader("somebody"))
		rq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		rq.Header.Add("X-Slack-Request-Timestamp", tc.timeStamp)
		rq.Header.Add("X-Slack-Signature", tc.signature)

		got, err := slackAuth.VerifyWebHook(rq)
		if err != nil {
			// Any error other then the expected one is a failed test.
			if _, ok := err.(*OldTimeStampError); !ok {
				t.Errorf("verifyWebHook: %v", err)
			}
		}
		if tc.wantResult != got {
			t.Errorf("Test: %v - Wanted: %v but got: %v", tc.name, tc.wantResult, got)
		}
	}
}
