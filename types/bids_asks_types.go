package types

/*
Result[0] -> [order_id, timestamp, volumn,rate, amount]
*/
type BidsAsksResponse struct {
	Error  int             `json:"error"`
	Result [][]interface{} `json:"result"`
}
