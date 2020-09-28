package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"ovoid/src/config"
	"ovoid/src/models"
)

type (
	// Request request parameter
	Request struct {
		Host          string
		Path          string
		Authorization string
		Method        string
		Body          interface{}
	}
)

func Post(req *Request) (*http.Response, error) {
	body, err := json.Marshal(req.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	client := http.Client{}
	request, err := http.NewRequest(req.Method, fmt.Sprintf("https://%s%s", req.Host, req.Path), bytes.NewBuffer(body))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("app-id", config.APP_ID)
	request.Header.Set("App-Version", config.APP_VERSION)
	request.Header.Set("OS", config.OS_NAME)
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
		var responseError = new(models.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		return nil, responseError
	}

	return resp, nil
}
