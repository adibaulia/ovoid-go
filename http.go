package goovo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	request struct {
		Host          string
		Path          string
		Authorization string
		Method        string
		Body          interface{}
	}
)

func post(req *request) (*http.Response, error) {
	body, err := json.Marshal(req.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	request, err := http.NewRequest(req.Method, fmt.Sprintf("https://%s%s", req.Host, req.Path), bytes.NewBuffer(body))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("app-id", appID)
	request.Header.Set("App-Version", appVersion)
	request.Header.Set("OS", osName)
	if req.Authorization != "" {
		request.Header.Set("Authorization", req.Authorization)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp, err := client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(err)
		var responseError = new(ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		return nil, responseError
	}

	return resp, nil
}
