package types

type TradingViewHistoryResponse struct {
	C []float64 `json:"c"`
	H []float64 `json:"h"`
	L []float64 `json:"l"`
	O []float64 `json:"o"`
	S string    `json:"s"`
	T []int     `json:"t"`
	V []float64 `json:"v"`
}
