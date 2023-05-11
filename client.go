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

func (c *Client) NewGetOpenOrdersTx() *GetOpenOrdersTx {
	return &GetOpenOrdersTx{c: c}
}

func (c *Client) NewGetOrderHistoryTx() *GetOrderHistoryTx {
	return &GetOrderHistoryTx{c: c}
}

func (c *Client) NewGetOrderInfoTx() *GetOrderInfoTx {
	return &GetOrderInfoTx{c: c}
}

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
	.Address("address").Amount(0.01).Currency("BTC").Do(context.Background())
*/
func (c *Client) NewCryptoWithdrawTx() *CryptoWithdrawTx {
	return &CryptoWithdrawTx{c: c}
}

/*
Function to Get History Crypto Deposit Transaction

Parameters Description:
  - Page : Page of result
  - Limit : Limit of result

Parameters Should be set before call Do() function:

	func (*GetCryptoDepositTx) Page(page int)
	func (*GetCryptoDepositTx) Limit(limit int)

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetCryptoDepositTx().Page(1).Limit(5).Do(context.Background())
*/
func (c *Client) NewGetCryptoDepositTx() *GetCryptoDepositTx {
	return &GetCryptoDepositTx{c: c}
}

/*
Function to Get History Crypto Withdraw Transaction

Parameters Description:
  - Page(int) : Page of result
  - Limit(int) : Limit of result

Parameters Should be set before call Do() function:

	func (*GetCryptoWithdrawTx) Page(page int)
	func (*GetCryptoWithdrawTx) Limit(limit int)

Example of usage:

	client := bitkub.NewClient()
	res, err := client.NewGetCryptoWithdrawTx().Page(1).Limit(5).Do(context.Background())
*/
func (c *Client) NewGetCryptoWithdrawTx() *GetCryptoWithdrawTx {
	return &GetCryptoWithdrawTx{c: c}
}

func (c *Client) NewGetFiatAccountsTx() *GetFiatAccountsTx {
	return &GetFiatAccountsTx{c: c}
}

func (c *Client) NewFiatWithdrawTx() *FiatWithdrawTx {
	return &FiatWithdrawTx{c: c}
}

func (c *Client) NewGetFiatDepositsTx() *GetFiatDepositsTx {
	return &GetFiatDepositsTx{c: c}
}

func (c *Client) NewGetFiatWithdralsTx() *GetFiatWithdrawsTx {
	return &GetFiatWithdrawsTx{c: c}
}

func (c *Client) NewGetUserLimitsTx() *GetUserLimitsTx {
	return &GetUserLimitsTx{c: c}
}

func (c *Client) NewGetTradingCreditsTx() *GetTradingCreditsTx {
	return &GetTradingCreditsTx{c: c}
}

func (c *Client) NewGetWsTokenTx() *GetWsTokenTx {
	return &GetWsTokenTx{c: c}
}

// Not usable due to lack of documentation
// func (c *Client) NewCryptoGenerateAddressTx() *CryptoGenerateAddressTx {
// 	return &CryptoGenerateAddressTx{c: c}
// }
