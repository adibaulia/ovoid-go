package ovoid

import (
	"encoding/json"
	"fmt"
)

type (
	ovo struct {
		AuthToken string `json:"token,omitempty"`
	}
)

//NewClient creates new instance ovo struct inside for authToken
func NewClient(authToken string) (*ovo, error) {
	if authToken == "" {
		return nil, fmt.Errorf("authToken required")
	}

	return &ovo{AuthToken: authToken}, nil
}

//GetAllBalance gets all balance in account
func (o *ovo) GetAllBalance() (*RespBalance, error) {
	req := &request{
		Method:        "GET",
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
func (o *ovo) GetBudget() (*RespBudget, error) {
	req := &request{
		Method:        "GET",
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
func (o *ovo) GetUnreadHistory() (map[string]int, error) {
	req := &request{
		Method:        "GET",
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
func (o *ovo) GetAllNotification() (*RespNotifications, error) {
	req := &request{
		Method:        "GET",
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
func (o *ovo) GetRefBank() (*RefBank, error) {
	req := &request{
		Method:        "GET",
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
