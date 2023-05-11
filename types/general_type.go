package types

import "errors"

type EndPointType string
type OrderType string
type OrderSide string
type BlockChainNetwork string
type OrderID string // After April 19th, 2023 at 18:00PM(GMT+7) Response Field id, first, parent, last change type from Integer to String.
type OrderHash string
type Symbol string
type Timestamp int64
type Signature string
type BitKubAPIError int //

func (s Symbol) String() string {
	return string(s)
}

func NewEndPoint(endpoint string) EndPointType {
	return EndPointType(endpoint)
}

func (e EndPointType) String() string {
	return string(e)
}

const (
	X_btk_apikey   = "X-BTK-APIKEY"
	BaseAPIMainURL = "https://api.bitkub.com"

	OrderTypeLimit  OrderType = "LIMIT"
	OrderTypeMarket OrderType = "MARKET"

	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

const (
	/*
		Non-secure endpoints
	*/
	StatusEndpoint             EndPointType = "/api/status"
	ServertimeEndpoint         EndPointType = "/api/servertime"
	MarketSymbolsEndpoint      EndPointType = "/api/market/symbols"
	MarketTickerEndpoint       EndPointType = "/api/market/ticker"
	MarketTradesEndpoint       EndPointType = "/api/market/trades"
	MarketBidsEndpoint         EndPointType = "/api/market/bids"
	MarketAsksEndpoint         EndPointType = "/api/market/asks"
	MarketBooksEndpoint        EndPointType = "/api/market/books"
	MarketDepthEndpoint        EndPointType = "/api/market/depth"
	TradingviewHistoryEndpoint EndPointType = "/tradingview/history"

	/*
		Secure endpoints
	*/
	MarketWalletEndpoint          EndPointType = "/api/market/wallet"
	MarketBalancesEndpoint        EndPointType = "/api/market/balances"
	MarketPlaceBidEndpoint        EndPointType = "/api/market/place-bid"
	MarketPlaceAskEndpoint        EndPointType = "/api/market/place-ask"
	MarketPlaceAskByFiatEndpoint  EndPointType = "/api/market/place-ask-by-fiat"
	MarketCancelOrderEndpoint     EndPointType = "/api/market/cancel-order"
	MarketMyOpenOrdersEndpoint    EndPointType = "/api/market/my-open-orders"
	MarketMyOrderHistoryEndpoint  EndPointType = "/api/market/my-order-history"
	MarketOrderInfoEndpoint       EndPointType = "/api/market/order-info"
	CryptoAddressesEndpoint       EndPointType = "/api/crypto/addresses"
	CryptoWithdrawEndpoint        EndPointType = "/api/crypto/withdraw"
	CryptoInternalWidrawEndpoint  EndPointType = "/api/crypto/internal-withdraw"
	CryptoDepositHistoryEndpoint  EndPointType = "/api/crypto/deposit-history"
	CryptoWithdrawHistoryEndpoint EndPointType = "/api/crypto/withdraw-history"
	CryptoGeneratreAddress        EndPointType = "/api/crypto/generate-address"
	FiatAccountsEndpoint          EndPointType = "/api/fiat/accounts"
	FiatWithdrawEndpoint          EndPointType = "/api/fiat/withdraw"
	FiatDepositHistoryEndpoint    EndPointType = "/api/fiat/deposit-history"
	FiatWithdrawHistoryEndpoint   EndPointType = "/api/fiat/withdraw-history"

	MarketWstokenEndpoint EndPointType = "/api/market/wstoken"
	UserLimitsEndpoint    EndPointType = "/api/user/limits"
	UserTradingCredits    EndPointType = "/api/user/trading-credits"

	MarketPlaceBidEndpointV2     EndPointType = "/api/market/v2/place-bid"
	MarketPlaceAskEndpointV2     EndPointType = "/api/market/v2/place-ask"
	MarketPCancelOrderEndpointV2 EndPointType = "/api/market/v2/cancel-order"

	PlaceBidTestEndpoint EndPointType = "/api/market/place-bid/test"
	PlaceAskTestEndpoint EndPointType = "/api/market/place-ask/test"
)

type TimeResolution string

type APIResponseError struct {
	Error BitKubAPIError `json:"error"`
}

type APIError struct {
	ErrorID   BitKubAPIError `json:"error"`
	ErrorDesc string         `json:"error_description"`
}

func (t TimeResolution) String() string {
	return string(t)
}

const (
	Time1m   TimeResolution = "1"
	Time5m   TimeResolution = "5"
	Time15m  TimeResolution = "15"
	Time1h   TimeResolution = "60"
	Time240m TimeResolution = "240"
	Time1d   TimeResolution = "D"
)

var (
	ErrSymbolMandatory      = errors.New("symbol is mandatory")
	ErrAmountMandatory      = errors.New("amount is mandatory")
	ErrAmountMustBePositive = errors.New("amount must be positive")
	ErrRateMandatory        = errors.New("rate is mandatory")
	ErrOrderTypeMandatory   = errors.New("order type is mandatory")
	ErrPageMustBePositive   = errors.New("page must be positive")
	ErrLimitMustBePositive  = errors.New("limit must be positive")
	ErrCurrencyMandatory    = errors.New("currency is mandatory")
	ErrAddressMandatory     = errors.New("address is mandatory")
	ErrNetworkMandatory     = errors.New("network is mandatory")
	ErrFiatAccIDMandatory   = errors.New("fiat account id is mandatory")
	ErrInvalidOrderSide     = errors.New("invalid order side")
	ErrOrderIDMandatory     = errors.New("order id is mandatory")
	ErrInvalidTimeStamp     = errors.New("invalid timestamp")
)

const (
	APINoError BitKubAPIError = iota
	APIInvalidJSON
	APIMissingAPIKey
	APIInvalidAPIKey
	APIPendingForActivation
	APIIpNotAllowed
	APIInvalidSignature
	APIMissingTimeStamp
	APIInvalidTimeStamp
	APIInvalidUser
	APIInvalidParamiter
	APIInvalidSymbol
	APIInvalidAmount
	APIInvalidPrice
	APIImproperRate
	APIAmountTooLow
	APIFailedToGetBalance
	APIEmptyWallet
	APIInsufficientBalance
	APIInsertOrderFailed
	APIDeductBalanceFailed
	APIInvalidOrderForCancellation
	APIInvalidSide
	APIUpdateOrderStatusFailed
	APIInvalidOrderForLookup
	APIKYCRequired
	APILimitExceeded         BitKubAPIError = 30
	APIPendingWidrawalExists BitKubAPIError = iota + 13
	APIInvalidCurrencyForWithdrawal
	APIAddressIsNotWhitelisted
	APIDeductCryptoFailed
	APICreateWidthdrawalRecordFailed
	APINonceHasToBeNumeric
	APIInvalidNonce
	APIWithdrawalLimitExceeded
	APIInvalidBankAccount
	APIBankLimitExceeded
	APIPendingWitdrawalExists
	APIWitdrawalUnderMaintenance
	APIInvalidPermission
	APIInvalidInternalAddress
	APIAddressDepreciated
	APICancelOnlyMode
	APIUserSuspendedFromPurchasing
	APIUserSuspendedFromSelling
	APIServerError = 90
)

/* Work In progress Alot of miss information from APi document*/

var BitkubAPIErrors = map[BitKubAPIError]string{
	APINoError:                       "No Error",
	APIInvalidJSON:                   "Invalid JSON payload",
	APIMissingAPIKey:                 "Missing X-BTK-APIKEY",
	APIInvalidAPIKey:                 "Invalid API key",
	APIPendingForActivation:          "API pending for activation",
	APIIpNotAllowed:                  "IP not allowed",
	APIInvalidSignature:              "Missing / Invalid signature",
	APIMissingTimeStamp:              "Missing timestamp",
	APIInvalidTimeStamp:              "Invalid timestamp",
	APIInvalidUser:                   "Invalid user",
	APIInvalidParamiter:              "Invalid parameter",
	APIInvalidSymbol:                 "Invalid symbol",
	APIInvalidAmount:                 "Invalid amount",
	APIInvalidPrice:                  "Invalid price",
	APIImproperRate:                  "Improper rate",
	APIAmountTooLow:                  "Amount too low",
	APIFailedToGetBalance:            "Failed to get balance",
	APIEmptyWallet:                   "Wallet is empty",
	APIInsufficientBalance:           "Insufficient balance",
	APIInsertOrderFailed:             "Failed to insert order into db",
	APIDeductBalanceFailed:           "Failed to deduct balance",
	APIInvalidOrderForCancellation:   "Invalid order for cancellation",
	APIInvalidSide:                   "Invalid side",
	APIUpdateOrderStatusFailed:       "Failed to update order status",
	APIInvalidOrderForLookup:         "Invalid order for lookup",
	APIKYCRequired:                   "KYC level 1 is required to proceed",
	APILimitExceeded:                 "Limit exceeded",
	APIPendingWidrawalExists:         "Pending withdrawal exists",
	APIInvalidCurrencyForWithdrawal:  "Invalid currency for withdrawal",
	APIAddressIsNotWhitelisted:       "Address is not whitelisted",
	APIDeductCryptoFailed:            "Failed to deduct crypto",
	APICreateWidthdrawalRecordFailed: "Failed to create withdrawal record",
	APINonceHasToBeNumeric:           "Nonce has to be numeric",
	APIInvalidNonce:                  "Invalid nonce",
	APIWithdrawalLimitExceeded:       "Withdrawal limit exceeded",
	APIInvalidBankAccount:            "Invalid bank account",
	APIBankLimitExceeded:             "Bank limit exceeded",
	APIPendingWitdrawalExists:        "Pending withdrawal exists",
	APIWitdrawalUnderMaintenance:     "Withdrawal is under maintenance",
	APIInvalidPermission:             "Invalid permission",
	APIInvalidInternalAddress:        "Invalid internal address",
	APIAddressDepreciated:            "Address has been depreciated",
	APICancelOnlyMode:                "Cancel only mode",
	APIUserSuspendedFromPurchasing:   "User has been suspended from purchasing",
	APIUserSuspendedFromSelling:      "User has been suspended from selling",
	APIServerError:                   "Server error",
}
