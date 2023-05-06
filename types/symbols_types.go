package types

type SymbolsResponse struct {
	Error  int `json:"error"`
	Result []struct {
		ID     int    `json:"id"`
		Info   string `json:"info"`
		Symbol string `json:"symbol"`
	} `json:"result"`
}
