package goovo

type (
	Login2FA struct {
		RefID string `json:"refId,omitempty"`
	}

	AccessToken struct {
		Mobile            string `json:"mobile,omitempty"`
		Email             string `json:"email,omitempty"`
		FullName          string `json:"fullName,omitempty"`
		IsEmailVerified   bool   `json:"isEmailVerified,omitempty"`
		IsSecurityCodeSet bool   `json:"isSecurityCodeSet,omitempty"`
		UpdateAccessToken string `json:"updateAccessToken,omitempty"`
	}

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
	RespBalance struct {
		Balance struct {
			OvoPoint struct {
				CardBalance   float64 `json:"card_balance"`
				CardNo        string  `json:"card_no"`
				PaymentMethod string  `json:"payment_method"`
			} `json:"600"`
			OvoMain struct {
				CardBalance   int    `json:"card_balance"`
				CardNo        string `json:"card_no"`
				PaymentMethod string `json:"payment_method"`
			} `json:"000"`
			OvoCash struct {
				CardBalance   int    `json:"card_balance"`
				CardNo        string `json:"card_no"`
				PaymentMethod string `json:"payment_method"`
			} `json:"001"`
		} `json:"balance"`
	}

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

func (e *ErrorResp) Error() string {
	return ""
}
