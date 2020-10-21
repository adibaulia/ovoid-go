package ovoid

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

	//RespBalance holds response from GetAllBalance
	RespBalance struct {
		Balance Balance `json:"balance"`
	}

	//Balance holds balances response
	Balance struct {
		OvoPoint OvoPoint `json:"600"`
		OvoMain  OvoMain  `json:"000"`
		OvoCash  OvoCash  `json:"001"`
	}

	//OvoPoint holds balance from ovopoint
	OvoPoint struct {
		CardBalance   float64 `json:"card_balance"`
		CardNo        string  `json:"card_no"`
		PaymentMethod string  `json:"payment_method"`
	}

	//OvoMain holds balance from ovomain
	OvoMain struct {
		CardBalance   int    `json:"card_balance"`
		CardNo        string `json:"card_no"`
		PaymentMethod string `json:"payment_method"`
	}

	//OvoCash holds balance from ovocash
	OvoCash struct {
		CardBalance   int    `json:"card_balance"`
		CardNo        string `json:"card_no"`
		PaymentMethod string `json:"payment_method"`
	}

	//RespBudget holds response from GetAllBudgets()
	RespBudget struct {
		Budget struct {
			Amount   int `json:"amount"`
			Spending int `json:"spending"`
		} `json:"budget"`
		TotalSpending int `json:"totalSpending"`
		CycleDate     int `json:"cycleDate"`
		Summary       []struct {
			Amount     int `json:"amount"`
			CategoryID int `json:"categoryId"`
			Spending   int `json:"spending"`
		} `json:"summary"`
	}

	//RespNotifications holds response from GetAllNotifications()
	RespNotifications struct {
		Notifications []struct {
			ID          string      `json:"id"`
			ChannelType string      `json:"channelType"`
			MessageType string      `json:"messageType"`
			Subject     interface{} `json:"subject"`
			Message     string      `json:"message"`
			DateCreated string      `json:"dateCreated"`
			Status      string      `json:"status"`
			Receiver    struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"receiver"`
		} `json:"notifications"`
	}

	//RefBank holds response GetRefBank()
	RefBank struct {
		BankTypes []struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Isdefault int    `json:"isdefault"`
			Value     string `json:"value"`
		} `json:"bankTypes"`
		BankTypeDefault interface{} `json:"bankTypeDefault"`
	}
)

//Error for implement error method from error pkg
func (e *ErrorResp) Error() string {
	return ""
}
