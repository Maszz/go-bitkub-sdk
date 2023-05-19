package types

type BidsAsksResponse struct {
	Error  int             `json:"error"`
	Result [][]interface{} `json:"result"`
}

type BaseRequestPayload struct {
	TS  int64  `json:"ts,omitempty"`
	Sig string `json:"sig,omitempty"`
}

type PlaceBidAskResponse struct {
	Error  int `json:"error"`
	Result struct {
		ID   int     `json:"id"`
		Hash string  `json:"hash"`
		Typ  string  `json:"typ"`
		Amt  float64 `json:"amt"`
		Rat  int     `json:"rat"`
		Fee  float64 `json:"fee"`
		Cre  float64 `json:"cre"`
		Rec  float64 `json:"rec"`
		TS   int     `json:"ts"`
		Ci   string  `json:"ci"`
	} `json:"result"`
}

type PlaceBidAskPayload struct {
	TS       Timestamp `json:"ts,omitempty"`
	Sig      Signature `json:"sig,omitempty"`
	Symbol   Symbol    `json:"sym,omitempty"`
	Amount   float64   `json:"amt,omitempty"`
	Rate     float64   `json:"rat"`
	Type     OrderType `json:"typ,omitempty"`
	ClientID string    `json:"client_id,omitempty"`
}
