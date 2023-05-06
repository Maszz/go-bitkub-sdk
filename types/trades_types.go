package types

/*
Result[0] -> [timestamp, rate, amount,side]
*/
type TradesResponse struct {
	Error  int             `json:"error"`
	Result [][]interface{} `json:"result"`
}
