package types

type GetUserLimitsPayload struct {
	Ts  Timestamp `json:"ts,omitempty"`
	Sig Signature `json:"sig,omitempty"`
}

type GetUserLimitsResponse struct {
	Error  int `json:"error"`
	Result struct {
		Limits struct {
			Crypto struct {
				Deposit  float64 `json:"deposit"`
				Withdraw float64 `json:"withdraw"`
			} `json:"crypto"`
			Fiat struct {
				Deposit  float64 `json:"deposit"`
				Withdraw float64 `json:"withdraw"`
			} `json:"fiat"`
		} `json:"limits"`
		Usage struct {
			Crypto struct {
				Deposit               float64 `json:"deposit"`
				Withdraw              float64 `json:"withdraw"`
				DepositPercentage     float64 `json:"deposit_percentage"`
				WithdrawPercentage    float64 `json:"withdraw_percentage"`
				DepositThbEquivalent  float64 `json:"deposit_thb_equivalent"`
				WithdrawThbEquivalent float64 `json:"withdraw_thb_equivalent"`
			} `json:"crypto"`
			Fiat struct {
				Deposit            float64 `json:"deposit"`
				Withdraw           float64 `json:"withdraw"`
				DepositPercentage  float64 `json:"deposit_percentage"`
				WithdrawPercentage float64 `json:"withdraw_percentage"`
			} `json:"fiat"`
		} `json:"usage"`
		Rate float64 `json:"rate"`
	} `json:"result"`
}

type GetTradingCreditsPayload struct {
	Ts  Timestamp `json:"ts,omitempty"`
	Sig Signature `json:"sig,omitempty"`
}

type GetTradingCreditsResponse struct {
	Error  int     `json:"error"`
	Result float64 `json:"result"`
}
