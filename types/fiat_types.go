package types

type FiatAccountPayload struct {
	Ts  Timestamp `json:"ts,omitempty"`
	Sig Signature `json:"sig,omitempty"`
}

type FiatAccountsResponse struct {
	Error  int `json:"error"`
	Result []struct {
		ID   string `json:"id"`
		Bank string `json:"bank"`
		Name string `json:"name"`
		Time int    `json:"time"`
	} `json:"result"`
	Pagination struct {
		Page int `json:"page,omitempty"`
		Last int `json:"last"`
		Next int `json:"next"`
		Prev int `json:"prev"`
	} `json:"pagination"`
}

type FiatWithdrawResponse struct {
	Error  int `json:"error"`
	Result struct {
		Txn string  `json:"txn"`
		Acc string  `json:"acc"`
		Cur string  `json:"cur"`
		Amt float64 `json:"amt"`
		Fee float64 `json:"fee"`
		Rec float64 `json:"rec"`
		Ts  int     `json:"ts"`
	} `json:"result"`
}

type FiatWithdrawPayload struct {
	Ts        Timestamp `json:"ts,omitempty"`
	Sig       Signature `json:"sig,omitempty"`
	FiatAccId string    `json:"id,omitempty"`
	Amount    float64   `json:"amt,omitempty"`
}

type GetFiatDepositsPayload struct {
	Ts  Timestamp `json:"ts,omitempty"`
	Sig Signature `json:"sig,omitempty"`
}
type GetFiatDepositsResponse struct {
	Error  int `json:"error"`
	Result []struct {
		TxnID    string  `json:"txn_id"`
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
		Status   string  `json:"status"`
		Time     int     `json:"time"`
	} `json:"result"`
	Pagination struct {
		Page int `json:"page,omitempty"`
		Last int `json:"last"`
		Next int `json:"next"`
		Prev int `json:"prev"`
	} `json:"pagination"`
}

type GetFiatWithdrawsPayload struct {
	Ts  Timestamp `json:"ts,omitempty"`
	Sig Signature `json:"sig,omitempty"`
}
type GetFiatWithdrawsResponse struct {
	Error  int `json:"error"`
	Result []struct {
		TxnID    string `json:"txn_id"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
		Fee      int    `json:"fee"`
		Status   string `json:"status"`
		Time     int    `json:"time"`
	} `json:"result"`
	Pagination struct {
		Page int `json:"page"`
		Last int `json:"last"`
	} `json:"pagination"`
}
