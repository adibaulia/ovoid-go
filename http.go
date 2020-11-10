package ovoid

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type (
	//ErrorResp stores responses error from ovo API
	ErrorResp struct {
		Message       string `json:"message,omitempty"`
		Code          int    `json:"code,omitempty"`
		URL           string `json:"url,omitempty"`
		Method        string `json:"method,omitempty"`
		RemoteAddress string `json:"remoteAddress,omitempty"`
		Unixtime      int    `json:"unixtime,omitempty"`
		Retry         int    `json:"retry,omitempty"`
		Timestamp     int    `json:"timestamp,omitempty"`
		Content       struct {
			Present bool `json:"present,omitempty"`
		} `json:"content,omitempty"`
	}
	request struct {
		Host          string
		Path          string
		Authorization string
		Method        string
		Timeout       time.Duration
		Body          interface{}
	}
)

//Error for implement error method from error pkg
func (e *ErrorResp) Error() string {
	return ""
}

func post(req *request) (*http.Response, error) {
	body, err := json.Marshal(req.Body)
	if err != nil {
		return nil, fmt.Errorf("error when marshalling body:%v", err)
	}

	client := &http.Client{
		Timeout: req.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	req.Host = baseEndpoint

	request, err := http.NewRequest(req.Method, fmt.Sprintf("https://%s%s", req.Host, req.Path), bytes.NewBuffer(body))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("app-id", appID)
	request.Header.Set("App-Version", appVersion)
	request.Header.Set("OS", osName)
	if req.Authorization != "" {
		request.Header.Set("Authorization", req.Authorization)
	}
	if err != nil {
		return nil, fmt.Errorf("error when creating new request:%v", err)
	}

	resp, err := client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		var responseError = new(ErrorResp)
		err := json.NewDecoder(resp.Body).Decode(responseError)
		if err != nil {
			return nil, fmt.Errorf("error when decoding error response ovo:%v", err)
		}
		return nil, responseError
	}

	return resp, nil
}
