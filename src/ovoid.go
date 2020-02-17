package main

import (
	"encoding/json"
	"fmt"
	"log"
	"ovoid/src/config"
	ovo "ovoid/src/helpers/http"
	"ovoid/src/models/response"
	"time"

	"github.com/google/uuid"
)

type (
	Login struct {
		DeviceID           string `json:"deviceId,omitempty"`
		Mobile             string `json:"mobile,omitempty"`
		AppVersion         string `json:"appVersion,omitempty"`
		MacAddress         string `json:"macAddress,omitempty"`
		OsName             string `json:"osName,omitempty"`
		OsVersion          string `json:"osVersion,omitempty"`
		PushNotificationID string `json:"pushNotificationId,omitempty"`
		RefID              string `json:"refId,omitempty"`
		VerificationCode   string `json:"verificationCode,omitempty"`
		DeviceUnixtime     int64  `json:"deviceUnixtime"`
		SecurityCode       string `json:"securityCode"`
		UpdateAccessToken  string `json:"updateAccessToken"`
		Message            string `json:"message"`
	}

	OVOID struct {
		AuthToken string `json:"token,omitempty"`
	}
)

func main() {
	ovoid := OVOID{
		AuthToken: `eyJhbGciOiJSUzI1NiJ9.eyJleHBpcnlJbk1pbGxpU2Vjb25kcyI6NjA0ODAwMDAwLCJjcmVhdGVUaW1lIjoxNTgxOTA3NzczNzkzLCJzZWNyZXQiOiJVaUhBN0ExSnA2czhLblBvNWY1OStMUWVnTTE1QTRiNTFWVVY1WE1tSW5rQ0xwTWtkc0tzTUFYTmtWTnhXN0VxWmRJR0RVYmJQMHJva3dScHFyWDRGYW4yV1FKV0Q5OVJRdGRBQUhDNGptRTdBNzI0WVRIdVJsQnRjeHlVdFpkODlPU2hVaGFSMS9ZdjQrdjZJdXU5VENXMVR1Nnc2MTN3NUI2d0hYaFQ0cXYyU2d3bDdxMHJFcVVpRWVpejYxZGFTUG9nOTFrVVJsWG5zZmZQYmdhRmxEQktKdU1PaXpDNStQcFh4OGFYalFGUnFvcndYRkZrQis0VHBBRmtHeG9kempuMGt6VlF6ekJrMDFDRkNHeVh3ZUxGVnFPSkxLVXZuRkRLeTNNVHkyMDJ2MDBWWmFYV2owZkdEdzZhYWRhMFBMU3ptVG5jeTM3NlBwK0RTUWROTmZLVVRmd3hGV0lFbk9BeFIzMVEwY3M9In0.Vwd1fnZL6meqlfRuTSY_m0koE_8mdC5QGBjLViVOKrS0wJy4gtkD1m0AsZAME7ndOu320K1rUQQcFAvHJqd0gR42SwgTD4h5xr51Ckq8L0YIrEXOtuLll6dvnMireGQIThW38sCNxmTIdVMuvE4lIAVOv3dTrcj813xnmNDVrri-Kst7OYxDUZVq6HFi_z-B6uTiOUcZeYlhV2p272WZ8CnXQ_xxWa1vOyjr4eLpN2Y5QORmW9DViM72LAU-Auo2a3DrEZ00-2y2kYfM8Y2THfubW7owL1fZfRPr4syiPR8qPtv0cOjbXKGjpG0GEernwwkLMhtoN2HG8LmS8ep-Mg`,
	}
	resp, respErr, err := ovoid.GetRefBank()
	if err != nil {
		panic(err)
	}
	log.Print(respErr)
	log.Print(resp)
	jsonResp, err := json.Marshal(resp)
	log.Print(string(jsonResp))

	//login := Login{Mobile: "081217179281"}
	// refId, _, _ := login.Login2FA()
	// log.Print(refId.RefID)
	//login.RefID = "0b1b7d9d66b747ebba734df5bfc7c837"
	//login.VerificationCode = "5653"
	//res, _, _ := login.Login2FAVerify()
	//log.Print(res)
	// login.UpdateAccessToken = "2b7c6a0187a14b6aa8a5398c56628b8a"
	// login.SecurityCode = "191613"
	// res, _, _ := login.LoginSecurityCode()
	// log.Print(res)

}

func (l *Login) Login2FA() (*response.Login2FA, *response.ErrorResp, error) {
	l.DeviceID = uuid.New().String()
	req := &ovo.Request{
		Method: "POST",
		Host:   config.BASE_ENDPOINT,
		Path:   "v2.0/api/auth/customer/login2FA",
		Body:   l,
	}
	resp, err := ovo.Post(req)
	if err != nil || resp.StatusCode != 200 {
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	var result = new(response.Login2FA)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	defer resp.Body.Close()
	return result, nil, nil

}

func (l *Login) Login2FAVerify() (*response.AccessToken, *response.ErrorResp, error) {

	l = &Login{
		Mobile:             l.Mobile,
		RefID:              l.RefID,
		VerificationCode:   l.VerificationCode,
		AppVersion:         config.APP_VERSION,
		DeviceID:           uuid.New().String(),
		MacAddress:         config.MAC_ADDRESS,
		OsName:             config.OS_NAME,
		OsVersion:          config.OS_VERSION,
		PushNotificationID: `FCM|f4OXYs_ZhuM:APA91bGde-ie2YBhmbALKPq94WjYex8gQDU2NMwJn_w9jYZx0emAFRGKHD2NojY6yh8ykpkcciPQpS0CBma-MxTEjaet-5I3T8u_YFWiKgyWoH7pHk7MXChBCBRwGRjMKIPdi3h0p2z7`,
	}

	log.Print(l.Mobile)
	req := &ovo.Request{
		Method: "POST",
		Host:   config.BASE_ENDPOINT,
		Path:   "v2.0/api/auth/customer/login2FA/verify",
		Body:   l,
	}
	resp, err := ovo.Post(req)
	if err != nil || resp.StatusCode != 200 {
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	var result = new(response.AccessToken)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	defer resp.Body.Close()
	return result, nil, nil
}

func (l *Login) LoginSecurityCode() (*response.Auth, *response.ErrorResp, error) {
	l = &Login{
		DeviceUnixtime:    time.Now().Unix(),
		SecurityCode:      l.SecurityCode,
		UpdateAccessToken: l.UpdateAccessToken,
		Message:           "",
	}

	req := &ovo.Request{
		Method: "POST",
		Host:   config.BASE_ENDPOINT,
		Path:   "v2.0/api/auth/customer/loginSecurityCode/verify",
		Body:   l,
	}
	resp, err := ovo.Post(req)
	if err != nil || resp.StatusCode != 200 {
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	var result = new(response.Auth)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	defer resp.Body.Close()
	return result, nil, nil
}

func (o *OVOID) GetAllBalance() (*response.RespBalance, *response.ErrorResp, error) {
	req := &ovo.Request{
		Method:        "GET",
		Host:          config.BASE_ENDPOINT,
		Path:          "v1.0/api/front/",
		Authorization: o.AuthToken,
	}
	resp, err := ovo.Post(req)
	if err != nil || resp.StatusCode != 200 {
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	var result = new(response.RespBalance)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	defer resp.Body.Close()
	return result, nil, nil
}

func (o *OVOID) GetBudget() (*response.RespBudget, *response.ErrorResp, error) {
	req := &ovo.Request{
		Method:        "GET",
		Host:          config.BASE_ENDPOINT,
		Path:          "v1.0/budget/detail",
		Authorization: o.AuthToken,
	}
	resp, err := ovo.Post(req)
	if err != nil || resp.StatusCode != 200 {
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	var result = new(response.RespBudget)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	defer resp.Body.Close()
	return result, nil, nil
}

func (o *OVOID) GetUnreadHistory() (map[string]int, *response.ErrorResp, error) {
	req := &ovo.Request{
		Method:        "GET",
		Host:          config.BASE_ENDPOINT,
		Path:          "v1.0/notification/status/count/UNREAD",
		Authorization: o.AuthToken,
	}
	resp, err := ovo.Post(req)
	if err != nil || resp.StatusCode != 200 {
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	var result = map[string]int{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	defer resp.Body.Close()
	return result, nil, nil
}

func (o *OVOID) GetAllNotification() (*response.RespNotifications, *response.ErrorResp, error) {
	req := &ovo.Request{
		Method:        "GET",
		Host:          config.BASE_ENDPOINT,
		Path:          "v1.0/notification/status/all",
		Authorization: o.AuthToken,
	}
	resp, err := ovo.Post(req)
	if err != nil || resp.StatusCode != 200 {
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	var result = new(response.RespNotifications)
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	defer resp.Body.Close()
	return result, nil, nil
}

func (o *OVOID) GetRefBank() (*response.RefBank, *response.ErrorResp, error) {
	req := &ovo.Request{
		Method:        "GET",
		Host:          config.BASE_ENDPOINT,
		Path:          "v1.0/reference/master/ref_bank",
		Authorization: o.AuthToken,
	}
	resp, err := ovo.Post(req)
	if err != nil || resp.StatusCode != 200 {
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	var result = new(response.RefBank)
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		var responseError = new(response.ErrorResp)
		json.NewDecoder(resp.Body).Decode(responseError)
		defer resp.Body.Close()
		return nil, responseError, err
	}
	defer resp.Body.Close()
	return result, nil, nil
}
