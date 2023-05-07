package bitkub

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/utils/http_client"

	"github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	HTTPClient *http_client.HttpClient
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
}

/*
Not tested how Big Struct will affect performance.
*/
func init() {
	fmt.Println("warmup")
	var v types.TickerResponse
	sonic.Pretouch(reflect.TypeOf(v))
}

func NewClient(apiKey, secretKey string) *Client {
	// warmup()

	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    types.BaseAPIMainURL,
		UserAgent:  "Bitkub-sdk/golang",
		HTTPClient: http_client.NewHttpClient(),
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

func (c *Client) parseRequest(r *request) (err error) {
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
	r.fullUrl = urlWithBase
	fmt.Println("full url", r.fullUrl)
	r.headers = headers

	return nil

}

func (c *Client) callAPI(ctx context.Context, r *request) (data []byte, err error) {
	err = c.parseRequest(r)
	if err != nil {
		return []byte{}, err
	}
	// tranform request object to fasthttp request object
	fmt.Println("calling api", r.query.String())
	req := c.HTTPClient.DoRequest(r.fullUrl, r.method, r.body, r.headers)
	// parse only error response

	return req, err
}

func (c *Client) parseError(data []byte) types.ApiError {
	var errResp types.ApiResponseError
	err := sonic.Unmarshal(data, &errResp)
	if err != nil {
		return types.ApiError{ErrorId: -1, ErrorDesc: "Unmarshal error"}
	}

	if errResp.Error == 0 {
		return types.ApiError{ErrorId: 0, ErrorDesc: types.BitkubApiErrors[types.ApiNoError]}
	}

	errMessage := types.BitkubApiErrors[errResp.Error]

	return types.ApiError{ErrorId: errResp.Error, ErrorDesc: errMessage}
}

func (c *Client) catchApiError(data []byte) error {
	err := c.parseError(data)
	if err.ErrorId != 0 {
		return fmt.Errorf("error id: %d, error message: %s", err.ErrorId, err.ErrorDesc)
	}
	return nil
}

// func (c *Client) NewTestService() *TestService {
// 	return &TestService{c: c}
// }
// func (c *Client) NewTestSignedService() *TestSignedService {
// 	return &TestSignedService{c: c}
// }

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

func (c *Client) NewCryptoWithdrawTx() *CryptoWithdrawTx {
	return &CryptoWithdrawTx{c: c}
}

// Internal withdraw is not supported yet due to lack of documentation. and can't tested cause of KYB.

func (c *Client) NewGetCryptoDepositTx() *GetCryptoDepositTx {
	return &GetCryptoDepositTx{c: c}
}
