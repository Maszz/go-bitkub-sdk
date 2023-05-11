package types

type GetWsTokenPayload struct {
	TS  Timestamp `json:"ts,omitempty"`
	Sig Signature `json:"sig,omitempty"`
}

type GetWsTokenResponse struct {
	Error  int    `json:"error"`
	Result string `json:"result"`
}
