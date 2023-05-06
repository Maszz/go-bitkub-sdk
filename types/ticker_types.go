package types

import "github.com/bytedance/sonic"

type TickerResponseAny map[string]TickerProperty

func (t *TickerResponseAny) Compile() (res TickerResponse) {
	resultByte, err := sonic.Marshal(t)
	if err != nil {
		return TickerResponse{}
	}
	err = sonic.Unmarshal(resultByte, &res)
	if err != nil {
		return TickerResponse{}
	}
	return res
}

type TickerProperty struct {
	ID            int     `json:"id"`
	Last          float64 `json:"last"`
	LowestAsk     float64 `json:"lowestAsk"`
	HighestBid    float64 `json:"highestBid"`
	PercentChange float64 `json:"percentChange"`
	BaseVolume    float64 `json:"baseVolume"`
	QuoteVolume   float64 `json:"quoteVolume"`
	IsFrozen      int     `json:"isFrozen"`
	High24Hr      float64 `json:"high24hr"`
	Low24Hr       float64 `json:"low24hr"`
	Change        float64 `json:"change"`
	PrevClose     float64 `json:"prevClose"`
	PrevOpen      float64 `json:"prevOpen"`
}

type TickerResponse struct {
	THB1INCH TickerProperty `json:"THB_1INCH"`
	THBAAVE  TickerProperty `json:"THB_AAVE"`
	THBABT   TickerProperty `json:"THB_ABT"`
	THBADA   TickerProperty `json:"THB_ADA"`
	THBALGO  TickerProperty `json:"THB_ALGO"`
	THBALPHA TickerProperty `json:"THB_ALPHA"`
	THBAPE   TickerProperty `json:"THB_APE"`
	THBARB   TickerProperty `json:"THB_ARB"`
	THBATOM  TickerProperty `json:"THB_ATOM"`
	THBAVAX  TickerProperty `json:"THB_AVAX"`
	THBAXL   TickerProperty `json:"THB_AXL"`
	THBAXS   TickerProperty `json:"THB_AXS"`
	THBBAL   TickerProperty `json:"THB_BAL"`
	THBBAND  TickerProperty `json:"THB_BAND"`
	THBBAT   TickerProperty `json:"THB_BAT"`
	THBBCH   TickerProperty `json:"THB_BCH"`
	THBBNB   TickerProperty `json:"THB_BNB"`
	THBBTC   TickerProperty `json:"THB_BTC"`
	THBBUSD  TickerProperty `json:"THB_BUSD"`
	THBCELO  TickerProperty `json:"THB_CELO"`
	THBCHZ   TickerProperty `json:"THB_CHZ"`
	THBCOMP  TickerProperty `json:"THB_COMP"`
	THBCRV   TickerProperty `json:"THB_CRV"`
	THBCTXC  TickerProperty `json:"THB_CTXC"`
	THBCVC   TickerProperty `json:"THB_CVC"`
	THBDAI   TickerProperty `json:"THB_DAI"`
	THBDOGE  TickerProperty `json:"THB_DOGE"`
	THBDOT   TickerProperty `json:"THB_DOT"`
	THBDYDX  TickerProperty `json:"THB_DYDX"`
	THBENJ   TickerProperty `json:"THB_ENJ"`
	THBENS   TickerProperty `json:"THB_ENS"`
	THBETH   TickerProperty `json:"THB_ETH"`
	THBFLOW  TickerProperty `json:"THB_FLOW"`
	THBFTM   TickerProperty `json:"THB_FTM"`
	THBFXS   TickerProperty `json:"THB_FXS"`
	THBGAL   TickerProperty `json:"THB_GAL"`
	THBGALA  TickerProperty `json:"THB_GALA"`
	THBGF    TickerProperty `json:"THB_GF"`
	THBGLM   TickerProperty `json:"THB_GLM"`
	THBGRT   TickerProperty `json:"THB_GRT"`
	THBGT    TickerProperty `json:"THB_GT"`
	THBHBAR  TickerProperty `json:"THB_HBAR"`
	THBHFT   TickerProperty `json:"THB_HFT"`
	THBILV   TickerProperty `json:"THB_ILV"`
	THBIMX   TickerProperty `json:"THB_IMX"`
	THBIOST  TickerProperty `json:"THB_IOST"`
	THBJFIN  TickerProperty `json:"THB_JFIN"`
	THBKNC   TickerProperty `json:"THB_KNC"`
	THBKSM   TickerProperty `json:"THB_KSM"`
	THBKUB   TickerProperty `json:"THB_KUB"`
	THBLDO   TickerProperty `json:"THB_LDO"`
	THBLINK  TickerProperty `json:"THB_LINK"`
	THBLRC   TickerProperty `json:"THB_LRC"`
	THBLUNA  TickerProperty `json:"THB_LUNA"`
	THBLYXE  TickerProperty `json:"THB_LYXE"`
	THBMANA  TickerProperty `json:"THB_MANA"`
	THBMATIC TickerProperty `json:"THB_MATIC"`
	THBMKR   TickerProperty `json:"THB_MKR"`
	THBNEAR  TickerProperty `json:"THB_NEAR"`
	THBOCEAN TickerProperty `json:"THB_OCEAN"`
	THBOMG   TickerProperty `json:"THB_OMG"`
	THBOP    TickerProperty `json:"THB_OP"`
	THBPOW   TickerProperty `json:"THB_POW"`
	THBSAND  TickerProperty `json:"THB_SAND"`
	THBSCRT  TickerProperty `json:"THB_SCRT"`
	THBSIX   TickerProperty `json:"THB_SIX"`
	THBSNT   TickerProperty `json:"THB_SNT"`
	THBSNX   TickerProperty `json:"THB_SNX"`
	THBSOL   TickerProperty `json:"THB_SOL"`
	THBSTG   TickerProperty `json:"THB_STG"`
	THBSUSHI TickerProperty `json:"THB_SUSHI"`
	THBTRX   TickerProperty `json:"THB_TRX"`
	THBUNI   TickerProperty `json:"THB_UNI"`
	THBUSDC  TickerProperty `json:"THB_USDC"`
	THBUSDT  TickerProperty `json:"THB_USDT"`
	THBWAN   TickerProperty `json:"THB_WAN"`
	THBXLM   TickerProperty `json:"THB_XLM"`
	THBXRP   TickerProperty `json:"THB_XRP"`
	THBXTZ   TickerProperty `json:"THB_XTZ"`
	THBYFI   TickerProperty `json:"THB_YFI"`
	THBZIL   TickerProperty `json:"THB_ZIL"`
	THBZRX   TickerProperty `json:"THB_ZRX"`
}
