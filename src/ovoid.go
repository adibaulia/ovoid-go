package main

import (
	"encoding/json"
	"fmt"
	"log"
	"ovoid/src/config"
	ovo "ovoid/src/helpers/http"
	"ovoid/src/models"
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

	ovoAccount := OVOID{}

	balance, _ := ovoAccount.GetAllBalance()
	fmt.Printf("%+v", balance)
}

//Login2FA login using phone number to OVO
func (l *Login) Login2FA() (*models.Login2FA, error) {
	l.DeviceID = uuid.New().String()
	req := &ovo.Request{
		Method: "POST",
		Host:   config.BASE_ENDPOINT,
		Path:   "v2.0/api/auth/customer/login2FA",
		Body:   l,
	}
	resp, err := ovo.Post(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var result = new(models.Login2FA)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

//Login2FAVerify is verify OTP sent to phonenumber
func (l *Login) Login2FAVerify() (*models.AccessToken, error) {
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
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var result = new(models.AccessToken)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

//LoginSecurityCode verify login from security pin code and updated access Token
func (l *Login) LoginSecurityCode() (*models.Auth, error) {
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
	if err != nil {
		return nil, err
	}
	var result = new(models.Auth)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}

//GetAllBalance gets all balance in account
func (o *OVOID) GetAllBalance() (*models.RespBalance, error) {
	req := &ovo.Request{
		Method:        "GET",
		Host:          config.BASE_ENDPOINT,
		Path:          "v1.0/api/front/",
		Authorization: o.AuthToken,
	}
	resp, err := ovo.Post(req)
	if err != nil {
		return nil, err
	}
	var result = new(models.RespBalance)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}

//GetBudget gets all budget in account
func (o *OVOID) GetBudget() (*models.RespBudget, error) {
	req := &ovo.Request{
		Method:        "GET",
		Host:          config.BASE_ENDPOINT,
		Path:          "v1.0/budget/detail",
		Authorization: o.AuthToken,
	}
	resp, err := ovo.Post(req)
	if err != nil {
		return nil, err
	}
	var result = new(models.RespBudget)
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
	req := &ovo.Request{
		Method:        "GET",
		Host:          config.BASE_ENDPOINT,
		Path:          "v1.0/notification/status/count/UNREAD",
		Authorization: o.AuthToken,
	}
	resp, err := ovo.Post(req)
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
func (o *OVOID) GetAllNotification() (*models.RespNotifications, error) {
	req := &ovo.Request{
		Method:        "GET",
		Host:          config.BASE_ENDPOINT,
		Path:          "v1.0/notification/status/all",
		Authorization: o.AuthToken,
	}
	resp, err := ovo.Post(req)
	if err != nil {
		return nil, err
	}
	var result = new(models.RespNotifications)
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}

//GetRefBank Get bank reference
func (o *OVOID) GetRefBank() (*models.RefBank, error) {
	req := &ovo.Request{
		Method:        "GET",
		Host:          config.BASE_ENDPOINT,
		Path:          "v1.0/reference/master/ref_bank",
		Authorization: o.AuthToken,
	}
	resp, err := ovo.Post(req)
	if err != nil {
		return nil, err
	}
	var result = new(models.RefBank)
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}
