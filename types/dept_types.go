package types

type MarketDepthResponse struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}
