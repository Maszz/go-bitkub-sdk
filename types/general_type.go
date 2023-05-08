package types

type EndPointType string
type OrderType string
type OrderSide string
type BlockChainNetwork string
type OrderId string // After April 19th, 2023 at 18:00PM(GMT+7) Response Field id, first, parent, last change type from Integer to String.
type OrderHash string
type Symbol string
type Timestamp int64
type Signature string
type BitKubApiError int //

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

	ERC20 BlockChainNetwork = "ETH"
	BEP20 BlockChainNetwork = "BSC"
	KAP20 BlockChainNetwork = "BKC"
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

type ApiResponseError struct {
	Error BitKubApiError `json:"error"`
}

type ApiError struct {
	ErrorId   BitKubApiError `json:"error"`
	ErrorDesc string         `json:"error_description"`
}

func (t TimeResolution) String() string {
	return string(t)
}

const (
	Time_1m   TimeResolution = "1"
	Time_5m   TimeResolution = "5"
	Time_15m  TimeResolution = "15"
	Time_1h   TimeResolution = "60"
	Time_240m TimeResolution = "240"
	Time_1d   TimeResolution = "D"
)

const (
	ApiNoError BitKubApiError = iota
	ApiInvalidJson
	ApiMissingApiKey
	ApiInvalidApiKey
	ApiPendingForActivation
	ApiIpNotAllowed
	ApiInvalidSignature
	ApiMissingTimeStamp
	ApiInvalidTimeStamp
	ApiInvalidUser
	ApiInvalidParamiter
	ApiInvalidSymbol
	ApiInvalidAmount
	ApiInvalidPrice
	ApiImproperRate
	ApiAmountTooLow
	ApiFailedToGetBalance
	ApiEmptyWallet
	ApiInsufficientBalance
	ApiInsertOrderFailed
	ApiDeductBalanceFailed
	ApiInvalidOrderForCancellation
	ApiInvalidSide
	ApiUpdateOrderStatusFailed
	ApiInvalidOrderForLookup
	ApiKYCRequired
	ApiLimitExceeded         BitKubApiError = 30
	ApiPendingWidrawalExists BitKubApiError = iota + 13
	ApiInvalidCurrencyForWithdrawal
	ApiAddressIsNotWhitelisted
	ApiDeductCryptoFailed
	ApiCreateWidthdrawalRecordFailed
	ApiNonceHasToBeNumeric
	ApiInvalidNonce
	ApiWithdrawalLimitExceeded
	ApiInvalidBankAccount
	ApiBankLimitExceeded
	ApiPendingWitdrawalExists
	ApiWitdrawalUnderMaintenance
	ApiInvalidPermission
	ApiInvalidInternalAddress
	ApiAddressDepreciated
	ApiCancelOnlyMode
	ApiUserSuspendedFromPurchasing
	ApiUserSuspendedFromSelling
	ApiServerError = 90
)

/* Work In progress Alot of miss information from APi document*/

var BitkubApiErrors = map[BitKubApiError]string{
	ApiNoError:                       "No Error",
	ApiInvalidJson:                   "Invalid JSON payload",
	ApiMissingApiKey:                 "Missing X-BTK-APIKEY",
	ApiInvalidApiKey:                 "Invalid API key",
	ApiPendingForActivation:          "API pending for activation",
	ApiIpNotAllowed:                  "IP not allowed",
	ApiInvalidSignature:              "Missing / Invalid signature",
	ApiMissingTimeStamp:              "Missing timestamp",
	ApiInvalidTimeStamp:              "Invalid timestamp",
	ApiInvalidUser:                   "Invalid user",
	ApiInvalidParamiter:              "Invalid parameter",
	ApiInvalidSymbol:                 "Invalid symbol",
	ApiInvalidAmount:                 "Invalid amount",
	ApiInvalidPrice:                  "Invalid price",
	ApiImproperRate:                  "Improper rate",
	ApiAmountTooLow:                  "Amount too low",
	ApiFailedToGetBalance:            "Failed to get balance",
	ApiEmptyWallet:                   "Wallet is empty",
	ApiInsufficientBalance:           "Insufficient balance",
	ApiInsertOrderFailed:             "Failed to insert order into db",
	ApiDeductBalanceFailed:           "Failed to deduct balance",
	ApiInvalidOrderForCancellation:   "Invalid order for cancellation",
	ApiInvalidSide:                   "Invalid side",
	ApiUpdateOrderStatusFailed:       "Failed to update order status",
	ApiInvalidOrderForLookup:         "Invalid order for lookup",
	ApiKYCRequired:                   "KYC level 1 is required to proceed",
	ApiLimitExceeded:                 "Limit exceeded",
	ApiPendingWidrawalExists:         "Pending withdrawal exists",
	ApiInvalidCurrencyForWithdrawal:  "Invalid currency for withdrawal",
	ApiAddressIsNotWhitelisted:       "Address is not whitelisted",
	ApiDeductCryptoFailed:            "Failed to deduct crypto",
	ApiCreateWidthdrawalRecordFailed: "Failed to create withdrawal record",
	ApiNonceHasToBeNumeric:           "Nonce has to be numeric",
	ApiInvalidNonce:                  "Invalid nonce",
	ApiWithdrawalLimitExceeded:       "Withdrawal limit exceeded",
	ApiInvalidBankAccount:            "Invalid bank account",
	ApiBankLimitExceeded:             "Bank limit exceeded",
	ApiPendingWitdrawalExists:        "Pending withdrawal exists",
	ApiWitdrawalUnderMaintenance:     "Withdrawal is under maintenance",
	ApiInvalidPermission:             "Invalid permission",
	ApiInvalidInternalAddress:        "Invalid internal address",
	ApiAddressDepreciated:            "Address has been depreciated",
	ApiCancelOnlyMode:                "Cancel only mode",
	ApiUserSuspendedFromPurchasing:   "User has been suspended from purchasing",
	ApiUserSuspendedFromSelling:      "User has been suspended from selling",
	ApiServerError:                   "Server error",
}
