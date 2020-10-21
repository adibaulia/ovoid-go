package ovoid

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type (
	//Login struct hold needed login field to interact with ovo API
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
	//Login2FA holds response when call Login2FA
	Login2FA struct {
		RefID string `json:"refId,omitempty"`
	}

	//AccessToken stores response from Login2FAVerify()
	AccessToken struct {
		Mobile            string `json:"mobile,omitempty"`
		Email             string `json:"email,omitempty"`
		FullName          string `json:"fullName,omitempty"`
		IsEmailVerified   bool   `json:"isEmailVerified,omitempty"`
		IsSecurityCodeSet bool   `json:"isSecurityCodeSet,omitempty"`
		UpdateAccessToken string `json:"updateAccessToken,omitempty"`
	}

	//Auth stores response from LoginSecurityCode()
	Auth struct {
		Token              string      `json:"token,omitempty"`
		TokenSeed          string      `json:"tokenSeed,omitempty"`
		TimeStamp          int         `json:"timeStamp,omitempty"`
		TokenSeedExpiredAt int         `json:"tokenSeedExpiredAt,omitempty"`
		DisplayMessage     interface{} `json:"displayMessage,omitempty"`
		Email              string      `json:"email,omitempty"`
		FullName           string      `json:"fullName,omitempty"`
		IsEmailVerified    bool        `json:"isEmailVerified,omitempty"`
		IsSecurityCodeSet  bool        `json:"isSecurityCodeSet,omitempty"`
		UpdateAccessToken  string      `json:"updateAccessToken,omitempty"`
	}
)

//NewOvoLogin initialize phone number and get login to ovo API
func NewOvoLogin(phone string) (*Login, error) {
	if phone == "" {
		return nil, fmt.Errorf("phone required")
	}
	return &Login{
		DeviceID: uuid.New().String(),
		Mobile:   phone,
	}, nil
}

//Login2FA login using phone number to OVO
func (l *Login) Login2FA() (*Login2FA, error) {
	req := &request{
		Method: "POST",
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
func (l *Login) Login2FAVerify(refID, verificationCode string) (*AccessToken, error) {

	if refID == "" {
		return nil, fmt.Errorf("refID required")
	}

	if verificationCode == "" {
		return nil, fmt.Errorf("verivicationCode required")
	}

	l = &Login{
		Mobile:             l.Mobile,
		RefID:              refID,
		VerificationCode:   verificationCode,
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
		log.Println(err)
		return nil, err
	}
	return result, nil
}

//LoginSecurityCode verify login from security pin code and updated access Token
func (l *Login) LoginSecurityCode(updateAccessToken string) (*Auth, error) {
	l = &Login{
		DeviceUnixtime:    time.Now().Unix(),
		SecurityCode:      l.SecurityCode,
		UpdateAccessToken: updateAccessToken,
		Message:           "",
	}

	req := &request{
		Method: "POST",
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
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}
