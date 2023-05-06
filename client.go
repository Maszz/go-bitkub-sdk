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

	hc "github.com/Maszz/go-bitkub-sdk/utils/http_client"

	BitkubTs "github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

const (
	x_btk_apikey   = "X-BTK-APIKEY"
	baseAPIMainURL = "https://api.bitkub.com"
)

const (
	/*
		Non-secure endpoints
	*/
	status_endpoint              = "/api/status"
	servertime_endpoint          = "/api/servertime"
	market_symbols_endpoint      = "/api/market/symbols"
	market_ticker_endpoint       = "/api/market/ticker"
	market_trades_endpoint       = "/api/market/trades"
	market_bids_endpoint         = "/api/market/bids"
	market_asks_endpoint         = "/api/market/asks"
	market_books_endpoint        = "/api/market/books"
	market_depth_endpoint        = "/api/market/depth"
	tradingview_history_endpoint = "/tradingview/history"

	/*
		Secure endpoints
	*/
	market_wallet_endpoint            = "/api/market/wallet"
	market_balances_endpoint          = "/api/market/balances"
	market_place_bid_endpoint         = "/api/market/place-bid"
	market_place_ask_endpoint         = "/api/market/place-ask"
	market_place_ask_by_fiat_endpoint = "/api/market/place-ask-by-fiat"
	market_cancel_order_endpoint      = "/api/market/cancel-order"
	market_my_open_orders_endpoint    = "/api/market/my-open-orders"
	market_my_order_history_endpoint  = "/api/market/my-order-history"
	market_order_info_endpoint        = "/api/market/order-info"
	crypto_addresses_endpoint         = "/api/crypto/addresses"
	crypto_withdraw_endpoint          = "/api/crypto/withdraw"
	crypto_internal_widraw_endpoint   = "/api/crypto/internal-withdraw"
	crypto_deposit_history_endpoint   = "/api/crypto/deposit-history"
	crypto_withdraw_history_endpoint  = "/api/crypto/withdraw-history"
	crypto_generatre_address          = "/api/crypto/generate-address"
	fiat_accounts_endpoint            = "/api/fiat/accounts"
	fiat_withdraw_endpoint            = "/api/fiat/withdraw"
	fiat_deposit_history_endpoint     = "/api/fiat/deposit-history"
	fiat_withdraw_history_endpoint    = "/api/fiat/withdraw-history"

	market_wstoken_endpoint = "/api/market/wstoken"
	user_limits_endpoint    = "/api/user/limits"
	user_trading_credits    = "/api/user/trading-credits"

	market_v2_place_bid_endpoint    = "/api/market/v2/place-bid"
	market_v2_place_ask_endpoint    = "/api/market/v2/place-ask"
	market_v2_cancel_order_endpoint = "/api/market/v2/cancel-order"

	place_bid_test_endpoint = "/api/market/place-bid/test"
	place_ask_test_endpoint = "/api/market/place-ask/test"
)

type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	HTTPClient *hc.HttpClient
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
}

/*
Not tested how Big Struct will affect performance.
*/
func init() {
	fmt.Println("warmup")
	var v BitkubTs.TickerResponse
	sonic.Pretouch(reflect.TypeOf(v))
}

func NewClient(apiKey, secretKey string) *Client {
	// warmup()

	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    baseAPIMainURL,
		UserAgent:  "Bitkub-sdk/golang",
		HTTPClient: hc.NewHttpClient(),
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
		headers.Set(x_btk_apikey, c.APIKey)
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

	return req, err
}

func (c *Client) NewTestService() *TestService {
	return &TestService{c: c}
}

func (c *Client) NewTestSignedService() *TestSignedService {
	return &TestSignedService{c: c}
}

func (c *Client) NewGetBalancesService() *GetBalancesService {
	return &GetBalancesService{c: c}
}

func (c *Client) NewGetStatusService() *GetStatusService {
	return &GetStatusService{c: c}
}

func (c *Client) NewGetServerTimeService() *GetServerTimeService {
	return &GetServerTimeService{c: c}
}

func (c *Client) NewGetTickerService() *GetTickerService {
	return &GetTickerService{c: c}
}

func (c *Client) NewGetTradesService() *GetTradesService {
	return &GetTradesService{c: c}
}

func (c *Client) NewGetBidsService() *GetBidsService {
	return &GetBidsService{c: c}
}

func (c *Client) NewGetAsksService() *GetAsksService {
	return &GetAsksService{c: c}
}
