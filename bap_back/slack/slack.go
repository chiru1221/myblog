/*
	Reference:
	https://github.com/GoogleCloudPlatform/golang-samples/blob/master/functions/slack/search.go
*/

package slack

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"example.com/bap/util/data"
	"example.com/bap/util/secrets"
)

type Slack interface {
	Construct()
	DialogApi(body io.Reader) error
	SendResponse(responseURL string) error
	TriggerId(r *http.Request) (string, error)
	Payload(r *http.Request) (string, error)
	MergePayload(pre *data.Blog, payload *data.Payload) *data.Blog
	VerifyWebHook(r *http.Request) (bool, error)
}

type SlackImpl struct {
	SlackSigningSecret string
	SlackToken         string
	DialogURL          string
}

func NewSlack() Slack {
	return &SlackImpl{}
}

type OldTimeStampError struct {
	s string
}

func (e *OldTimeStampError) Error() string {
	return e.s
}

const (
	version                     = "v0"
	slackRequestTimestampHeader = "X-Slack-Request-Timestamp"
	slackSignatureHeader        = "X-Slack-Signature"
	dialogURL                   = "https://slack.com/api/dialog.open"
)

func (slack *SlackImpl) Construct() {
	slack.SlackSigningSecret = secrets.SecretsFile(os.Getenv("SLACK_SS_FILE"))
	slack.SlackToken = secrets.SecretsFile(os.Getenv("SLACK_TOKEN_FILE"))
	slack.DialogURL = dialogURL
}

func (slack *SlackImpl) DialogApi(body io.Reader) error {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = "Bearer " + slack.SlackToken

	// create request
	req, err := http.NewRequest("POST", slack.DialogURL, body)
	if err != nil {
		return err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}

	// send request
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (slack *SlackImpl) SendResponse(responseURL string) error {
	header := make(map[string]string)
	header["Content-Type"] = "text/plane"

	// create request
	body := bytes.NewBuffer([]byte("ok"))
	req, err := http.NewRequest("POST", responseURL, body)
	if err != nil {
		return err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}

	// send request
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (slack *SlackImpl) TriggerId(r *http.Request) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return r.PostFormValue("trigger_id"), nil
}

func (slack *SlackImpl) Payload(r *http.Request) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return r.PostFormValue("payload"), nil
}

func (slack *SlackImpl) MergePayload(pre *data.Blog, payload *data.Payload) *data.Blog {
	body := payload.Body
	if body.Article != "" {
		pre.Article = body.Article
	}
	if body.Title != "" {
		pre.Title = body.Title
	}
	if body.Tag != "" {
		pre.Tag = strings.Split(body.Tag, ",")
	}
	if body.Open != "" {
		isOpen, err := strconv.ParseBool(body.Open)
		if err != nil {
			isOpen = false
		}
		pre.Open = isOpen
	}
	if body.Date != "" {
		pre.Date = body.Date
	}
	return pre
}

func (slack *SlackImpl) VerifyWebHook(r *http.Request) (bool, error) {
	timeStamp := r.Header.Get(slackRequestTimestampHeader)
	slackSignature := r.Header.Get(slackSignatureHeader)

	t, err := strconv.ParseInt(timeStamp, 10, 64)
	if err != nil {
		return false, fmt.Errorf("strconv.ParseInt(%s): %v", timeStamp, err)
	}

	if ageOk, age := checkTimestamp(t); !ageOk {
		return false, &OldTimeStampError{fmt.Sprintf("checkTimestamp(%v): %v %v", t, ageOk, age)}
	}

	if timeStamp == "" || slackSignature == "" {
		return false, fmt.Errorf("either timeStamp or signature headers were blank")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return false, fmt.Errorf("ioutil.ReadAll(%v): %v", r.Body, err)
	}

	// Reset the body so other calls won't fail.
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	baseString := fmt.Sprintf("%s:%s:%s", version, timeStamp, body)

	signature := GetSignature([]byte(baseString), []byte(slack.SlackSigningSecret))

	trimmed := strings.TrimPrefix(slackSignature, fmt.Sprintf("%s=", version))
	signatureInHeader, err := hex.DecodeString(trimmed)

	if err != nil {
		return false, fmt.Errorf("hex.DecodeString(%v): %v", trimmed, err)
	}

	return hmac.Equal(signature, signatureInHeader), nil
}

func GetSignature(base []byte, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write(base)

	return h.Sum(nil)
}

func checkTimestamp(timeStamp int64) (bool, time.Duration) {
	t := time.Since(time.Unix(timeStamp, 0))

	return t.Minutes() <= 5, t
}
