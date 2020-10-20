package goovo

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

const (
	appID        = "C7UMRSMFRZ46D9GW9IK7"
	appVersion   = "3.2.0"
	osName       = "Android"
	osVersion    = "8.1.0"
	macAddress   = "d8:8e:35:4d:bd:88"
	baseEndpoint = "api.ovo.id/"
	aws          = "apigw01.aws.ovo.id/"
	transferOvo  = "trf_ovo"
	transferBank = "trf_other_bank"
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

//Login2FA login using phone number to OVO
func (l *Login) Login2FA() (*Login2FA, error) {
	l.DeviceID = uuid.New().String()
	req := &request{
		Method: "POST",
		Host:   baseEndpoint,
		Path:   "v2.0/api/auth/customer/login2FA",
		Body:   l,
	}
	resp, err := post(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var result = new(Login2FA)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

//Login2FAVerify is verify OTP sent to phonenumber
func (l *Login) Login2FAVerify() (*AccessToken, error) {
	l = &Login{
		Mobile:             l.Mobile,
		RefID:              l.RefID,
		VerificationCode:   l.VerificationCode,
		AppVersion:         appVersion,
		DeviceID:           uuid.New().String(),
		MacAddress:         macAddress,
		OsName:             osName,
		OsVersion:          osVersion,
		PushNotificationID: `FCM|f4OXYs_ZhuM:APA91bGde-ie2YBhmbALKPq94WjYex8gQDU2NMwJn_w9jYZx0emAFRGKHD2NojY6yh8ykpkcciPQpS0CBma-MxTEjaet-5I3T8u_YFWiKgyWoH7pHk7MXChBCBRwGRjMKIPdi3h0p2z7`,
	}

	log.Print(l.Mobile)
	req := &request{
		Method: "POST",
		Host:   baseEndpoint,
		Path:   "v2.0/api/auth/customer/login2FA/verify",
		Body:   l,
	}
	resp, err := post(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var result = new(AccessToken)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

//LoginSecurityCode verify login from security pin code and updated access Token
func (l *Login) LoginSecurityCode() (*Auth, error) {
	l = &Login{
		DeviceUnixtime:    time.Now().Unix(),
		SecurityCode:      l.SecurityCode,
		UpdateAccessToken: l.UpdateAccessToken,
		Message:           "",
	}

	req := &request{
		Method: "POST",
		Host:   baseEndpoint,
		Path:   "v2.0/api/auth/customer/loginSecurityCode/verify",
		Body:   l,
	}
	resp, err := post(req)
	if err != nil {
		return nil, err
	}
	var result = new(Auth)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}

//GetAllBalance gets all balance in account
func (o *OVOID) GetAllBalance() (*RespBalance, error) {
	req := &request{
		Method:        "GET",
		Host:          baseEndpoint,
		Path:          "v1.0/api/front/",
		Authorization: o.AuthToken,
	}
	resp, err := post(req)
	if err != nil {
		return nil, err
	}
	var result = new(RespBalance)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}

//GetBudget gets all budget in account
func (o *OVOID) GetBudget() (*RespBudget, error) {
	req := &request{
		Method:        "GET",
		Host:          baseEndpoint,
		Path:          "v1.0/budget/detail",
		Authorization: o.AuthToken,
	}
	resp, err := post(req)
	if err != nil {
		return nil, err
	}
	var result = new(RespBudget)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}

//GetUnreadHistory get notification that unread
func (o *OVOID) GetUnreadHistory() (map[string]int, error) {
	req := &request{
		Method:        "GET",
		Host:          baseEndpoint,
		Path:          "v1.0/notification/status/count/UNREAD",
		Authorization: o.AuthToken,
	}
	resp, err := post(req)
	if err != nil {
		return nil, err
	}
	var result = map[string]int{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}

//GetAllNotification gets all notification
func (o *OVOID) GetAllNotification() (*RespNotifications, error) {
	req := &request{
		Method:        "GET",
		Host:          baseEndpoint,
		Path:          "v1.0/notification/status/all",
		Authorization: o.AuthToken,
	}
	resp, err := post(req)
	if err != nil {
		return nil, err
	}
	var result = new(RespNotifications)
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}

//GetRefBank Get bank reference
func (o *OVOID) GetRefBank() (*RefBank, error) {
	req := &request{
		Method:        "GET",
		Host:          baseEndpoint,
		Path:          "v1.0/reference/master/ref_bank",
		Authorization: o.AuthToken,
	}
	resp, err := post(req)
	if err != nil {
		return nil, err
	}
	var result = new(RefBank)
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}
