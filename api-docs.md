
# Go Bitkub SDK Documentation
This documentation is in development may lack of information.

# API List 

- [GET /api/status](#get-endpoints-status)
- [GET /api/servertime](#get-server-time) 
- [GET /api/market/symbols](#list-all-symbols) 
- [GET /api/market/ticker](#get-ticker-information) 
- [GET /api/market/trades](#list-recent-trades) 
- [GET /api/market/bids](#list-open-buy-orders) 
- [GET /api/market/asks](#list-open-sell-orders) 
- [GET /api/market/books](#list-all-open-orders) 
- [GET /api/market/depth](#get-depth-information) 
- [GET /tradingview/history](#get-historical-data-from-tradingview) 
- [POST /api/market/wallet](#get-user-available-balances) 
- [POST /api/market/balances](#get-balances-info) 
- [POST /api/market/my-open-orders](#list-all-open-orders) 
- [POST /api/market/my-orders-history](#list-all-matched-orders) 
- [POST /api/market/order-info](#get-order-information) 
- [POST /api/crypto/addresses](#get-addresses)
- [POST /api/crypto/withdraw](#crypto-withdraw) 
- [POST /api/crypto/deposit-history](#get-deposit-history) 
- [POST /api/crypto/withdraw-history](#get-withdraw-history) 
- [POST /api/fiat/accounts](#get-fiat-accounts) 
- [POST /api/fiat/withdraw](#fiat-withdraw) 
- [POST /api/fiat/deposit-history](#get-deposit-history) 
- [POST /api/fiat/withdraw-history](#get-withdraw-history) 
- [POST /api/market/wstoken](#get-websocket-token) 
- [POST /api/user/limits](#get-user-limits) 
- [POST /api/user/trading-credits](#get-trading-credits)  
- [POST /api/market/v2/place-bid](#create-ask-orders) 
- [POST /api/market/v2/place-ask](#create-bid-orders)   
- POST /api/market/v2/cancel-order  

# API documentation
Example usage for each endpoint

### Get Endpoints Status
```golang
res, err := client.NewGetStatusTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```
# API Documentation

### Get Server time
```golang
res, err := client.NewGetServerTimeTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### List all symbols
```golang
res, err := client.NewGetSymbolsTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get ticker information
```golang
res, err := client.NewGetTickerTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### List recent trades
```golang
res, err := client.NewGetTradesTx().Symbol(symbols.THB_BTC).Limit(20).Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### List open buy orders
```golang
res, err := client.NewGetBidsTx().Symbol(symbols.THB_BTC).Limit(20).Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### List open sell orders
```golang
res, err := client.NewGetAsksTx().Symbol(symbols.THB_BTC).Limit(20).Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### List all open orders
```golang
res, err := client.NewGetBooksTx().Symbol(symbols.THB_BTC).Limit(20).Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```
### Get depth information
```golang
res, err := client.NewGetMarketDepthTx().Symbol(symbols.THB_BTC).Limit(20).Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```


### Get historical data from tradingView
```golang
// ToCurrent is replicsent of time.Now(), But you can specify the time by using ToTimestamp() instead.
res, err := client.NewGetTradingviewHistoryTx().Symbol(symbols.THB_BTC).FromTimestamp(1633424427).ToCurrent().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get user available balances
```golang
res, err := client.NewGetWalletsTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get balances info
```golang
res, err := client.NewGetBalancesTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### List all open orders
```golang
res, err := client.NewGetOpenOrdersTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```
### List all matched orders
```golang
res, err := client.NewGetOrderHistoryTx().Symbol(symbols.THB_BTC).Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```
### Get Order Information
```golang
res, err := client.OrderHash("Order hash").Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get Addresses
```golang
res, err := client.NewGetCryptoAddressesTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Crypto Withdraw
```golang
res, err := bitkubClient.NewCryptoWithdrawTx().Network(chains.BTC).Address("Address").Amount(0.01).Currency("BTC").Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```
> Address should be a whitelisted address.

### Get Deposit History
```golang
res, err := client.NewGetCryptoDepositTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get Withdraw History
```golang
res, err := client.NewGetCryptoWithdrawTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get Fiat Accounts
```golang
res, err := client.NewGetCryptoWithdrawTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Fiat Withdraw
```golang
res, err := client.NewFiatWithdrawTx().ID("Fiat account ID").Amount(1000).Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get Deposit History
```golang
res, err := client.NewGetFiatDepositsTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get Withdraw History
```golang
res, err := client.NewGetCryptoWithdrawTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get Websocket Token
```golang
res, err := client.NewGetWsTokenTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get User Limits
```golang
res, err := client.NewGetUserLimitsTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get Trading Credits
```golang
res, err := client.NewGetTradingCreditsTx().Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Create Ask Orders
```golang
res, err := client.NewPlaceAskTx().Symbol(symbols.THB_BTC).Amount(0.001).OrderType(types.OrderTypeMarket).Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Create Bid Orders
```golang
res, err := client.NewPlaceBidTx().Symbol(symbols.THB_BTC).Amount(1000).OrderType(types.OrderTypeMarket).Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Cancel Order
```golang
res, err := client.NewCancelOrderTx().OrderHash("OrderHash").Do()
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```



