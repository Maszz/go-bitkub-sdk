package bitkub

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/utils/httpclient"

	"github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	HTTPClient *httpclient.HTTPClient
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
}

/*
Not tested how Big Struct will affect performance.
*/
func init() {
	var v types.TickerResponse
	err := sonic.Pretouch(reflect.TypeOf(v))
	if err != nil {
		panic(err)
	}
}

func NewClient(apiKey, secretKey string) *Client {
	// warmup()

	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    types.BaseAPIMainURL,
		UserAgent:  "Bitkub-sdk/golang",
		HTTPClient: httpclient.NewHTTPClient(),
		Logger:     log.New(os.Stderr, "Bitkub-golang ", log.LstdFlags),
	}
}

func (c *Client) signPayload(payload interface{}) string {
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return ""
	}
	bodyString := string(byteBody)
	key := []byte(c.SecretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(bodyString))
	hmacSigned := h.Sum(nil)
	hmacSignedStr := hex.EncodeToString(hmacSigned)
	return hmacSignedStr
}

func (c *Client) parseRequest(r *request) {
	// do hmac and sign payload + cal payload stuff.
	urlWithBase := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	headers := &fasthttp.RequestHeader{}

	if r.signed == secTypeSigned {
		headers.Set("Accept", "application/json")
		headers.Set(types.X_btk_apikey, c.APIKey)
	}
	if len(r.body) > 0 {
		headers.Set("Content-Type", "application/json")
	}

	headers.Set("User-Agent", c.UserAgent)
	r.fullURL = urlWithBase
	fmt.Println("full url", r.fullURL)
	r.headers = headers
}

func (c *Client) callAPI(r *request) ([]byte, error) {
	c.parseRequest(r)

	// transform request object to fasthttp request object
	// fmt.Println("calling api", r.query.String())
	req, err := c.HTTPClient.DoRequest(r.fullURL, r.method, r.body, r.headers)
	// parse only error response
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) parseError(data []byte) types.APIError {
	var errResp types.APIResponseError
	err := sonic.Unmarshal(data, &errResp)
	if err != nil {
		return types.APIError{ErrorID: -1, ErrorDesc: "Unmarshal error"}
	}

	if errResp.Error == 0 {
		return types.APIError{ErrorID: 0, ErrorDesc: types.BitkubAPIErrors[types.APINoError]}
	}

	errMessage := types.BitkubAPIErrors[errResp.Error]

	return types.APIError{ErrorID: errResp.Error, ErrorDesc: errMessage}
}

func (c *Client) catchAPIError(data []byte) error {
	err := c.parseError(data)
	if err.ErrorID != 0 {
		errMsg := fmt.Sprintf("error id: %d, error message: %s", err.ErrorID, err.ErrorDesc)
		return errors.New(errMsg)
	}
	return nil
}

func (c *Client) NewGetStatusTx() *GetStatusTx {
	return &GetStatusTx{c: c}
}

func (c *Client) NewGetServerTimeTx() *GetServerTimeTx {
	return &GetServerTimeTx{c: c}
}

func (c *Client) NewGetSymbolsTx() *GetSymbolsTx {
	return &GetSymbolsTx{c: c}
}

func (c *Client) NewGetTickerTx() *GetTickerTx {
	return &GetTickerTx{c: c}
}

func (c *Client) NewGetTradesTx() *GetTradesTx {
	return &GetTradesTx{c: c}
}

// @param symbol string
// @param limit int
func (c *Client) NewGetBidsTx() *GetBidsTx {
	return &GetBidsTx{c: c}
}

func (c *Client) NewGetAsksTx() *GetAsksTx {
	return &GetAsksTx{c: c}
}

func (c *Client) NewGetBooksTx() *GetOpenBooksTx {
	return &GetOpenBooksTx{c: c}
}

func (c *Client) NewGetMarketDepthTx() *GetMarketDepthTx {
	return &GetMarketDepthTx{c: c}
}

func (c *Client) NewGetTradingviewHistoryTx() *GetTradingViewHistoryTx {
	return &GetTradingViewHistoryTx{c: c}
}

func (c *Client) NewGetWalletsTx() *GetWalletsTx {
	return &GetWalletsTx{c: c}
}

func (c *Client) NewGetBalancesTx() *GetBalancesTx {
	return &GetBalancesTx{c: c}
}

func (c *Client) NewPlaceBidTx() *PlaceBidTx {
	return &PlaceBidTx{c: c}
}

func (c *Client) NewPlaceAskTx() *PlaceAskTx {
	return &PlaceAskTx{c: c}
}

func (c *Client) NewCancelOrderTx() *CancelOrderTx {
	return &CancelOrderTx{c: c}
}

/*
Function to Get Open Orders

Parameters Description:
  - Symbol(types.Symbol) : Symbol of order

Parameters Should be set before call Do() function:

	func (*GetOpenOrdersTx) Symbol(symbol types.Symbol)

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetOpenOrdersTx().Symbol(symbols.THB_BTC).Do()
*/
func (c *Client) NewGetOpenOrdersTx() *GetOpenOrdersTx {
	return &GetOpenOrdersTx{c: c}
}

/*
Function to Get Order History

Parameters Description:
  - Symbol(types.Symbol) : Symbol of order
  - Page(Optional[int]) : Page of result
  - Limit(Optional[int]) : Limit of result
  - Start(Optional[types.Timestamp]) : Start Time
  - End(Optional[types.Timestamp]) : End time

Parameters Should be set before call Do() function:

	func (*GetOrderHistoryTx) Symbol(symbol types.Symbol)
	func (*GetOrderHistoryTx) Page(page int)
	func (*GetOrderHistoryTx) Limit(limit int)
	func (*GetOrderHistoryTx) Start(start types.Timestamp)
	func (*GetOrderHistoryTx) End(end types.Timestamp)

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetOrderHistoryTx().Symbol(symbols.THB_BTC).Do()
*/
func (c *Client) NewGetOrderHistoryTx() *GetOrderHistoryTx {
	return &GetOrderHistoryTx{c: c}
}

/*
Function to Get Order Information

Parameters Description:
  - Symbol(types.Symbol) : Page of result
  - OrderID(types.OrderID) : Order ID
  - OrderSide(types.OrderSide) : Order side
  - OrderHash(Optional[types.OrderHash]) : Order hash

Parameters Should be set before call Do() function:

	func (*GetOrderInfoTx) Symbol(symbol types.Symbol)
	func (*GetOrderInfoTx) OrderID(orderID types.OrderID)
	func (*GetOrderInfoTx) OrderSide(orderSide types.OrderSide)
	func (*GetOrderInfoTx) OrderHash(orderHash types.OrderHash)

When OrderHash is set, Symbol, OrderID and OrderSide will be ignored.

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetOrderInfoTx().OrderHash("OrderHash").Do()
*/
func (c *Client) NewGetOrderInfoTx() *GetOrderInfoTx {
	return &GetOrderInfoTx{c: c}
}

/*
Function to Get Crypto Wallet Address

Parameters Description:
  - Page(Optional[int]) : Page of result
  - Limit(Optional[int]) : Limit of result

Parameters Should be set before call Do() function:

	func (*GetCryptoAddressesTx) Page(page int)
	func (*GetCryptoAddressesTx) Limit(limit int)

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetCryptoAddressesTx().Page(1).Limit(5).Do()
*/
func (c *Client) NewGetCryptoAddressesTx() *GetCryptoAddressesTx {
	return &GetCryptoAddressesTx{c: c}
}

/*
Function to Withdraw Crypto to specific address

Parameters Description:
  - Address(string) : Address of receiver
  - Amount(float64) : Amount of crypto to withdraw
  - Currency(string) : Currency of crypto to withdraw
  - Memo(Optional[string]) : Memo of crypto to withdraw
  - Network(types.BlockChainNetwork) : Network of crypto to withdraw

Parameters Should be set before call Do() function:

	func (*CryptoWithdrawTx) Address(address string)
	func (*CryptoWithdrawTx) Amount(amount float64)
	func (*CryptoWithdrawTx) Currency(cur string)
	func (*CryptoWithdrawTx) Memo(memo string)
	func (*CryptoWithdrawTx) Network(network types.BlockChainNetwork)

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewCryptoWithdrawTx().Network(chains.BTC)
	.Address("address").Amount(0.01).Currency("BTC").Do()
*/
func (c *Client) NewCryptoWithdrawTx() *CryptoWithdrawTx {
	return &CryptoWithdrawTx{c: c}
}

/*
Function to Get History Crypto Deposit Transaction

Parameters Description:
  - Page(Optional[int]) : Page of result
  - Limit(Optional[int]) : Limit of result

Parameters Should be set before call Do() function:

	func (*GetCryptoDepositTx) Page(page int)
	func (*GetCryptoDepositTx) Limit(limit int)

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetCryptoDepositTx().Page(1).Limit(5).Do()
*/
func (c *Client) NewGetCryptoDepositTx() *GetCryptoDepositTx {
	return &GetCryptoDepositTx{c: c}
}

/*
Function to Get History Crypto Withdraw Transaction

Parameters Description:
  - Page(Optional[int]) : Page of result
  - Limit(Optional[int]) : Limit of result

Parameters Should be set before call Do() function:

	func (*GetCryptoWithdrawTx) Page(page int)
	func (*GetCryptoWithdrawTx) Limit(limit int)

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetCryptoWithdrawTx().Page(1).Limit(5).Do()
*/
func (c *Client) NewGetCryptoWithdrawTx() *GetCryptoWithdrawTx {
	return &GetCryptoWithdrawTx{c: c}
}

/*
Function to Get Fiat Accounts Transaction

Parameters Description:
  - Page(Optional[int]) : Page of result
  - Limit(Optional[int]) : Limit of result

Parameters Should be set before call Do() function:

	func (*GetFiatAccountsTx) Page(page int)
	func (*GetFiatAccountsTx) Limit(limit int)

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetFiatAccountsTx().Page(1).Limit(5).Do()
*/
func (c *Client) NewGetFiatAccountsTx() *GetFiatAccountsTx {
	return &GetFiatAccountsTx{c: c}
}

/*
Function to Get Fiat Withdraw Transaction

Parameters Description:
  - ID(string) : Fiat Account ID
  - Amount(float64) : Amount of fiat to withdraw

Parameters Should be set before call Do() function:

	func (*FiatWithdrawTx) Amount(amount float64)
	func (*FiatWithdrawTx) ID(id string) *FiatWithdrawTx

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewFiatWithdrawTx().Amount(1000).ID("Fiat Account ID").Do()
*/
func (c *Client) NewFiatWithdrawTx() *FiatWithdrawTx {
	return &FiatWithdrawTx{c: c}
}

/*
Function to Get Fiat Deposit History Transaction

Parameters Description:
  - Page(Optional[int]) : Page of result
  - Limit(Optional[int]) : Limit of result

Parameters Should be set before call Do() function:

	func (*GetFiatDepositsTx) Page(page int)
	func (*GetFiatDepositsTx) Limit(limit int)

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetFiatDepositsTx().Page(1).Limit(5).Do()
*/
func (c *Client) NewGetFiatDepositsTx() *GetFiatDepositsTx {
	return &GetFiatDepositsTx{c: c}
}

/*
Function to Get Fiat Withdraw History Transaction

Parameters Description:
  - Page(Optional[int]) : Page of result
  - Limit(Optional[int]) : Limit of result

Parameters Should be set before call Do() function:

	func (*GetFiatWithdrawsTx) Page(page int)
	func (*GetFiatWithdrawsTx) Limit(limit int)

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetFiatWithdralsTx().Page(1).Limit(5).Do()
*/
func (c *Client) NewGetFiatWithdralsTx() *GetFiatWithdrawsTx {
	return &GetFiatWithdrawsTx{c: c}
}

/*
Function to Get User limit Transaction

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetUserLimitsTx().Do()
*/
func (c *Client) NewGetUserLimitsTx() *GetUserLimitsTx {
	return &GetUserLimitsTx{c: c}
}

/*
Function to Get Trading credits Transaction

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetTradingCreditsTx().Do()
*/
func (c *Client) NewGetTradingCreditsTx() *GetTradingCreditsTx {
	return &GetTradingCreditsTx{c: c}
}

/*
Function to Get Websocket Token Transaction

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetWsTokenTx().Do()
*/
func (c *Client) NewGetWsTokenTx() *GetWsTokenTx {
	return &GetWsTokenTx{c: c}
}

// Not usable due to lack of documentation
// func (c *Client) NewCryptoGenerateAddressTx() *CryptoGenerateAddressTx {
// 	return &CryptoGenerateAddressTx{c: c}
// }
