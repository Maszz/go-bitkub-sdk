package types

type OpenBooksResponse struct {
	Error  int `json:"error"`
	Result struct {
		Asks [][]interface{} `json:"asks"`
		Bids [][]interface{} `json:"bids"`
	} `json:"result"`
}
