
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
- POST /api/market/wallet 
- POST /api/market/balances 
- POST /api/market/place-bid 
- POST /api/market/place-ask 
- POST /api/market/place-ask-by-fiat 
- POST /api/market/cancel-order 
- POST /api/market/my-open-orders 
- POST /api/market/my-orders-history 
- POST /api/market/order-info 
- POST /api/crypto/addresses 
- POST /api/crypto/withdraw 
- POST /api/crypto/deposit-history 
- POST /api/crypto/withdraw-history 
- POST /api/crypto/generate-address  
- POST /api/fiat/accounts 
- POST /api/fiat/withdraw 
- POST /api/fiat/deposit-history 
- POST /api/fiat/withdraw-history 
- POST /api/market/wstoken 
- POST /api/user/limits 
- POST /api/user/trading-credits  
- [POST /api/market/v2/place-bid](#create-asks-orders) 
- POST /api/market/v2/place-ask  
- POST /api/market/v2/cancel-order  

# API documentation
Example usage fir each endpoint

### Get Endpoints Status
```golang
res, err := client.NewGetStatusTx().Do(context.Background())
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
res, err := client.NewGetServerTimeTx().Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### List all symbols
```golang
res, err := client.NewGetSymbolsTx().Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Get ticker information
```golang
res, err := client.NewGetTickerTx().Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### List recent trades
```golang
res, err := client.NewGetTradesTx().Symbol(symbols.THB_BTC).Limit(20).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### List open buy orders
```golang
res, err := client.NewGetBidsTx().Symbol(symbols.THB_BTC).Limit(20).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### List open sell orders
```golang
res, err := client.NewGetAsksTx().Symbol(symbols.THB_BTC).Limit(20).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### List all open orders
```golang
res, err := client.NewGetBooksTx().Symbol(symbols.THB_BTC).Limit(20).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```
### Get depth information
```golang
res, err := client.NewGetMarketDepthTx().Symbol(symbols.THB_BTC).Limit(20).Do(context.Background())
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
res, err := client.NewGetTradingviewHistoryTx().Symbol(symbols.THB_BTC).FromTimestamp(1633424427).ToCurrent().Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

### Create Asks Orders
```golang
res, err := client.NewPlaceAskTx().Symbol(symbols.THB_BTC).Amount(0.001).OrderType(types.OrderTypeMarket).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```





