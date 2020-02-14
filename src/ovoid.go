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
	// ovoid := OVOID{
	// 	AuthToken: "eyJhbGciOiJSUzI1NiJ9.eyJleHBpcnlJbk1pbGxpU2Vjb25kcyI6NjA0ODAwMDAwLCJjcmVhdGVUaW1lIjoxNTgxNjUyNzEyMzg3LCJzZWNyZXQiOiJVaUhBN0ExSnA2czhLblBvNWY1OStMUWVnTTE1QTRiNTFWVVY1WE1tSW5rQ0xwTWtkc0tzTUFYTmtWTnhXN0VxWmRJR0RVYmJQMHJva3dScHFyWDRGYW4yV1FKV0Q5OVJRdGRBQUhDNGptRTdBNzI0WVRIdVJsQnRjeHlVdFpkODlPU2hVaGFSMS9ZdjQrdjZJdXU5VExVMHRaVjZobFRGZlhYNnlkdFh2NFNLaGdmNVdOSW5kQzdTN093U2pBSWFNSE13S21FL01oTml2TFl1TUZtNTY2NElybFJkZlphQmZjSzhVaFA1cHhMRHF3bXE2Q1o4cUpJYlRFYkljY2tqL3J1L1hmMnFLYnJxd0xOaVByVzZ3V0x2TGJvQjJraXFpVDhvdGVTVmxuZ3Nja0g3SEg3RnpsTm5adDA3dlFGUFNEZzFiMU5lZHpkN1p0cjA0SHM3ayt5Mk1lUUNISUNXa3V6TTRmdTJidFU9In0.DNW3DeGNxeQq3Q2MO0TimGE_Pg3GT7h64acP9aVBEoJ9GjKWhK65LbI25wjakYeM-t2vKRonchOxLcXJ14X7YENV6Xzc3pCN6xHX2BK-0Rsgy4dZiFXV0lUHsX4lgZonSL6J9pcoFfb-hH9_K1YjFXWadANjuyp7VBRrBFwRqAQEOATV91-5kWAXNhcZ2SL-JdY_EdA0gowF2aUF53QBy4zccXKK26feViROC_2jzYIFvxv348No8M9qqrmLasBCCtJT2sfG3Z_qnOH-qZd5kGyS-z8m2zqz8-QKL8X5Lp9DaWEt5-BoWWeVClQBsOsGVx3tFHDejCVCbJItjBuRmw",
	// }
	// resp, respErr, err := ovoid.GetAllBalance()
	// if err != nil {
	// 	panic(err)
	// }
	// log.Print(respErr)
	// log.Print(resp.Balance.OvoPoint.CardBalance)

	login := Login{Mobile: "081217179281"}
	login.Login2FA()
}

func (l *Login) Login2FA() (*response.Login2FA, *response.ErrorResp, error) {
	l.DeviceID = uuid.New().String()
	req := &ovo.Request{
		Method: "POST",
		Host:   config.BASE_ENDPOINT,
		Path:   "v2.0/api/auth/customer/login2FA",
		Body:   l,
	}
	log.Print(l.DeviceID)
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
		AppVersion:         config.APP_VERSION,
		DeviceID:           uuid.New().String(),
		MacAddress:         config.MAC_ADDRESS,
		OsName:             config.OS_NAME,
		OsVersion:          config.OS_VERSION,
		PushNotificationID: `FCM|f4OXYs_ZhuM:APA91bGde-ie2YBhmbALKPq94WjYex8gQDU2NMwJn_w9jYZx0emAFRGKHD2NojY6yh8ykpkcciPQpS0CBma-MxTEjaet-5I3T8u_YFWiKgyWoH7pHk7MXChBCBRwGRjMKIPdi3h0p2z7`,
	}

	log.Print(l)
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
	l.DeviceUnixtime = time.Now().Unix()
	l.Message = ""
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
