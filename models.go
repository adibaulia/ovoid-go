package ovoid

type (

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

	//Notifications holds response from GetAllNotifications()
	Notifications struct {
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
