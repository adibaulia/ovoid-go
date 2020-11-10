package ovoid

import (
	"encoding/json"
	"fmt"
	"time"
)

type (

	//Ovo holds struct token and hold methods
	Ovo struct {
		AuthToken string `json:"token,omitempty"`
		Timeout   time.Duration
	}
)

//NewClient creates new instance ovo struct inside for authToken
func NewClient(authToken string, timeout time.Duration) (*Ovo, error) {
	if authToken == "" {
		return nil, fmt.Errorf("authToken required")
	}

	return &Ovo{AuthToken: authToken, Timeout: timeout}, nil
}

//GetAllBalances gets all balances in account
func (o *Ovo) GetAllBalances() (*RespBalance, error) {
	req := &request{
		Method:        "GET",
		Path:          "v1.0/api/front/",
		Authorization: o.AuthToken,
		Timeout:       o.Timeout,
	}
	resp, err := post(req)
	if err != nil {
		return nil, err
	}
	var result = new(RespBalance)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}

//GetBudgets gets all budget in account
func (o *Ovo) GetBudgets() (*RespBudget, error) {
	req := &request{
		Method:        "GET",
		Path:          "v1.0/budget/detail",
		Authorization: o.AuthToken,
		Timeout:       o.Timeout,
	}
	resp, err := post(req)
	if err != nil {
		return nil, err
	}
	var result = new(RespBudget)
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}

//GetCountUnreadNotifications get count number of notification that unread
//will return pointer of int because the number can be zero
func (o *Ovo) GetCountUnreadNotifications() (*int, error) {
	req := &request{
		Method:        "GET",
		Path:          "v1.0/notification/status/count/UNREAD",
		Authorization: o.AuthToken,
		Timeout:       o.Timeout,
	}
	resp, err := post(req)
	if err != nil {
		return nil, err
	}
	var result = map[string]int{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res := result["Total"]
	return &res, nil
}

//GetAllNotifications gets all notifications
func (o *Ovo) GetAllNotifications() ([]Notifications, error) {
	req := &request{
		Method:        "GET",
		Path:          "v1.0/notification/status/all",
		Authorization: o.AuthToken,
		Timeout:       o.Timeout,
	}
	resp, err := post(req)
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	var response []Notifications
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := json.Marshal(res["notifications"])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//GetRefBank Get bank referenceID
func (o *Ovo) GetRefBank() (*RefBank, error) {
	req := &request{
		Method:        "GET",
		Path:          "v1.0/reference/master/ref_bank",
		Authorization: o.AuthToken,
		Timeout:       o.Timeout,
	}
	resp, err := post(req)
	if err != nil {
		return nil, err
	}
	var result = new(RefBank)
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return result, nil
}
