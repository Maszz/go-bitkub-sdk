package types

import "github.com/bytedance/sonic"

type BalancesPayload struct {
	Ts  Timestamp `json:"ts,omitempty"`
	Sig Signature `json:"sig,omitempty"`
}
type BalancesProps struct {
	Available float64 `json:"available"`
	Reserved  float64 `json:"reserved"`
}

type BalancesResponseAny struct {
	Error  int                      `json:"error"`
	Result map[string]BalancesProps `json:"result"`
}

func (b *BalancesResponseAny) Compile() (res BalancesResponse) {
	resultByte, err := sonic.Marshal(b.Result)
	if err != nil {
		return BalancesResponse{}
	}
	err = sonic.Unmarshal(resultByte, &res.Result)
	if err != nil {
		return BalancesResponse{}
	}
	res.Error = b.Error
	return res

}

type BalancesResult struct {
	THB   BalancesProps `json:"THB"`
	BTC   BalancesProps `json:"BTC"`
	ETH   BalancesProps `json:"ETH"`
	WAN   BalancesProps `json:"WAN"`
	ADA   BalancesProps `json:"ADA"`
	OMG   BalancesProps `json:"OMG"`
	BCH   BalancesProps `json:"BCH"`
	USDT  BalancesProps `json:"USDT"`
	XRP   BalancesProps `json:"XRP"`
	ZIL   BalancesProps `json:"ZIL"`
	SNT   BalancesProps `json:"SNT"`
	CVC   BalancesProps `json:"CVC"`
	LINK  BalancesProps `json:"LINK"`
	IOST  BalancesProps `json:"IOST"`
	ZRX   BalancesProps `json:"ZRX"`
	KNC   BalancesProps `json:"KNC"`
	ABT   BalancesProps `json:"ABT"`
	MANA  BalancesProps `json:"MANA"`
	CTXC  BalancesProps `json:"CTXC"`
	XLM   BalancesProps `json:"XLM"`
	SIX   BalancesProps `json:"SIX"`
	JFIN  BalancesProps `json:"JFIN"`
	BNB   BalancesProps `json:"BNB"`
	POW   BalancesProps `json:"POW"`
	DOGE  BalancesProps `json:"DOGE"`
	DAI   BalancesProps `json:"DAI"`
	BAND  BalancesProps `json:"BAND"`
	KSM   BalancesProps `json:"KSM"`
	DOT   BalancesProps `json:"DOT"`
	USDC  BalancesProps `json:"USDC"`
	BAT   BalancesProps `json:"BAT"`
	NEAR  BalancesProps `json:"NEAR"`
	SCRT  BalancesProps `json:"SCRT"`
	GLM   BalancesProps `json:"GLM"`
	YFI   BalancesProps `json:"YFI"`
	UNI   BalancesProps `json:"UNI"`
	COMP  BalancesProps `json:"COMP"`
	MKR   BalancesProps `json:"MKR"`
	DON   BalancesProps `json:"DON"`
	AAVE  BalancesProps `json:"AAVE"`
	KUB   BalancesProps `json:"KUB"`
	ENJ   BalancesProps `json:"ENJ"`
	ALPHA BalancesProps `json:"ALPHA"`
	BAL   BalancesProps `json:"BAL"`
	CRV   BalancesProps `json:"CRV"`
	AXS   BalancesProps `json:"AXS"`
	SAND  BalancesProps `json:"SAND"`
	SUSHI BalancesProps `json:"SUSHI"`
}

type BalancesResponse struct {
	Error  int            `json:"error"`
	Result BalancesResult `json:"result"`
}
