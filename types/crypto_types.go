package types

/*
In some Response behavior of pagination, it's not the same as the documentation. and not consistent.
All Response that have pagination, Will have all Possible keys.
*/
type CryptoAddressesResponse struct {
	Error  int `json:"error"`
	Result []struct {
		Currency string `json:"currency"`
		Address  string `json:"address"`
		Tag      string `json:"tag"`
		Time     int    `json:"time"`
		Network  string `json:"network,omitempty"`
	} `json:"result"`
	Pagination struct {
		Page int `json:"page,omitempty"`
		Last int `json:"last"`
		Next int `json:"next"`
		Prev int `json:"prev"`
	} `json:"pagination"`
}

type CryptoAddressesResponseWithAllQuery struct {
	Error  int `json:"error"`
	Result []struct {
		Currency string `json:"currency"`
		Address  string `json:"address"`
		Tag      string `json:"tag"`
		Time     int    `json:"time"`
		Network  string `json:"network,omitempty"`
	} `json:"result"`
	Pagination struct {
		Page string `json:"page,omitempty"`
		Last int    `json:"last"`
		Next int    `json:"next"`
		Prev int    `json:"prev"`
	} `json:"pagination"`
}
type CryptoAddressesPayload struct {
	TS  Timestamp `json:"ts,omitempty"`
	Sig Signature `json:"sig,omitempty"`
}

type CryptoWithdrawPayload struct {
	TS      Timestamp         `json:"ts,omitempty"`
	Sig     Signature         `json:"sig,omitempty"`
	Cur     string            `json:"cur,omitempty"`
	Amount  float64           `json:"amt,omitempty"`
	Address string            `json:"adr,omitempty"`
	Memo    string            `json:"mem,omitempty"`
	Network BlockChainNetwork `json:"net,omitempty"`
}

type CryptoWithdrawResponse struct {
	Error  int `json:"error"`
	Result struct {
		Txn string  `json:"txn"`
		Adr string  `json:"adr"`
		Mem string  `json:"mem"`
		Cur string  `json:"cur"`
		Amt float64 `json:"amt"`
		Fee float64 `json:"fee"`
		TS  int     `json:"ts"`
	} `json:"result"`
}

type CryptoInternalWidthdrawPayload struct {
	TS      Timestamp `json:"ts,omitempty"`
	Sig     Signature `json:"sig,omitempty"`
	Cur     string    `json:"cur,omitempty"`
	Amount  float64   `json:"amt,omitempty"`
	Address string    `json:"adr,omitempty"`
	Memo    string    `json:"mem,omitempty"`
}

type GetCryptoDepositPayload struct {
	TS  Timestamp `json:"ts,omitempty"`
	Sig Signature `json:"sig,omitempty"`
}

type GetCryptoDepositResponse struct {
	Error  int `json:"error"`
	Result []struct {
		Hash          string      `json:"hash"`
		Currency      string      `json:"currency"`
		Amount        float64     `json:"amount"`
		Address       interface{} `json:"address"`
		Confirmations int         `json:"confirmations"`
		Status        string      `json:"status"`
		Note          interface{} `json:"note"`
		Time          int         `json:"time"`
	} `json:"result"`
	Pagination struct {
		Page int `json:"page,omitempty"`
		Last int `json:"last"`
		Next int `json:"next"`
		Prev int `json:"prev"`
	} `json:"pagination"`
}

type GetCryptoWithdrawPayload struct {
	TS  Timestamp `json:"ts,omitempty"`
	Sig Signature `json:"sig,omitempty"`
}

type GetCryptoWithdrawResponse struct {
	Error  int `json:"error"`
	Result []struct {
		TxnID    string  `json:"txn_id"`
		Hash     string  `json:"hash"`
		Currency string  `json:"currency"`
		Amount   string  `json:"amount"`
		Fee      float64 `json:"fee"`
		Address  string  `json:"address"`
		Memo     string  `json:"memo"`
		Status   string  `json:"status"`
		Note     string  `json:"note"`
		Time     int     `json:"time"`
	} `json:"result"`
	Pagination struct {
		Page int `json:"page,omitempty"`
		Last int `json:"last"`
		Next int `json:"next"`
		Prev int `json:"prev"`
	} `json:"pagination"`
}
type CryptoGenerateAddressPayload struct {
	TS     Timestamp `json:"ts,omitempty"`
	Sig    Signature `json:"sig,omitempty"`
	Symbol Symbol    `json:"sym,omitempty"`
}
type CryptoGenerateAddressResponse struct {
	Error  int `json:"error"`
	Result []struct {
		Currency string `json:"currency"`
		Address  string `json:"address"`
		Memo     string `json:"memo"`
	} `json:"result"`
}
