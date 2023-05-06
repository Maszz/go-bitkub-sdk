package types

import "github.com/bytedance/sonic"

type WalletResponseAny map[string]float64

func (w *WalletResponseAny) Compile() (res WalletResponse) {
	resultByte, err := sonic.Marshal(w)
	if err != nil {
		return WalletResponse{}
	}
	err = sonic.Unmarshal(resultByte, &res.Result)
	if err != nil {
		return WalletResponse{}
	}
	return res
}

type WalletResult struct {
	THB   float64 `json:"THB"`
	BTC   float64 `json:"BTC"`
	ETH   float64 `json:"ETH"`
	WAN   float64 `json:"WAN"`
	ADA   float64 `json:"ADA"`
	OMG   float64 `json:"OMG"`
	BCH   float64 `json:"BCH"`
	USDT  float64 `json:"USDT"`
	XRP   float64 `json:"XRP"`
	ZIL   float64 `json:"ZIL"`
	SNT   float64 `json:"SNT"`
	CVC   float64 `json:"CVC"`
	LINK  float64 `json:"LINK"`
	IOST  float64 `json:"IOST"`
	ZRX   float64 `json:"ZRX"`
	KNC   float64 `json:"KNC"`
	ABT   float64 `json:"ABT"`
	MANA  float64 `json:"MANA"`
	CTXC  float64 `json:"CTXC"`
	XLM   float64 `json:"XLM"`
	SIX   float64 `json:"SIX"`
	JFIN  float64 `json:"JFIN"`
	BNB   float64 `json:"BNB"`
	POW   float64 `json:"POW"`
	DOGE  float64 `json:"DOGE"`
	DAI   float64 `json:"DAI"`
	BAND  float64 `json:"BAND"`
	KSM   float64 `json:"KSM"`
	DOT   float64 `json:"DOT"`
	USDC  float64 `json:"USDC"`
	BAT   float64 `json:"BAT"`
	NEAR  float64 `json:"NEAR"`
	SCRT  float64 `json:"SCRT"`
	GLM   float64 `json:"GLM"`
	YFI   float64 `json:"YFI"`
	UNI   float64 `json:"UNI"`
	COMP  float64 `json:"COMP"`
	MKR   float64 `json:"MKR"`
	DON   float64 `json:"DON"`
	AAVE  float64 `json:"AAVE"`
	KUB   float64 `json:"KUB"`
	ENJ   float64 `json:"ENJ"`
	ALPHA float64 `json:"ALPHA"`
	BAL   float64 `json:"BAL"`
	CRV   float64 `json:"CRV"`
	AXS   float64 `json:"AXS"`
	SAND  float64 `json:"SAND"`
	SUSHI float64 `json:"SUSHI"`
}

type WalletResponse struct {
	Error  float64      `json:"error"`
	Result WalletResult `json:"result"`
}
