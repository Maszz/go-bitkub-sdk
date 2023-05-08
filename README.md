# go-bitkub-sdk
A Golang sdk for bitkub api

## This Module is in initial development it not support majority of a api

some modules behave differently from documentation. This makes it difficult to Test.
Need Manualy call on every endpoint to test json responses it works correctly.
Make sure you have read Bitkub API document before continuing.

## Installation

```shell
go get github.com/Maszz/go-bitkub-sdk
```
updating a package 

```shell
go get -u github.com/Maszz/go-bitkub-sdk
```

### Importing

```golang
import (
    "github.com/Maszz/go-bitkub-sdk"
)
```

## API List 

Name  | Status
------------ | ------------ | 
[GET /api/status](#get-endpoints-status)| Implemented(Tested)
[GET /api/servertime](#get-server-time) | Implemented(Tested)
[GET /api/market/symbols](#list-all-symbols) | Implemented(Tested)
[GET /api/market/ticker](#get-ticker-information) | Implemented(Tested)
[GET /api/market/trades](#list-recent-trades) | Implemented(Tested)
[GET /api/market/bids](#list-open-buy-orders) | Implemented(Tested)
[GET /api/market/asks](#list-open-sell-orders) | Implemented(Tested)
[GET /api/market/books](#list-all-open-orders) | Implemented(Tested)
[GET /api/market/depth](#get-depth-information) | Implemented(Tested)
[GET /tradingview/history](#get-historical-data-from-tradingview) | Implemented(Tested)
POST /api/market/wallet | Implemented(Tested)
POST /api/market/balances | Implemented(Tested)
POST /api/market/place-bid | Not implemented(Depecated)
POST /api/market/place-ask | Not implemented(Depecated)
POST /api/market/place-ask-by-fiat | Not implemented(Depecated)
POST /api/market/cancel-order | Not implemented(Depecated)
POST /api/market/my-open-orders | Implemented(Tested)
POST /api/market/my-orders-history | Implemented(Tested)
POST /api/market/order-info | Implemented(Tested)
POST /api/crypto/addresses | Implemented(Tested)
POST /api/crypto/withdraw | Implemented(Tested)
POST /api/crypto/deposit-history | Implemented(Tested)
POST /api/crypto/withdraw-history | Implemented(Tested)
POST /api/crypto/generate-address | Can't be implemented(Lack of document).
POST /api/fiat/accounts | Implemented(Tested)
POST /api/fiat/withdraw | Implemented(Tested)
POST /api/fiat/deposit-history | Implemented(Tested)
POST /api/fiat/withdraw-history | Implemented(Tested)
POST /api/market/wstoken | Implemented(Tested)
POST /api/user/limits | Implemented(Tested)
POST /api/user/trading-credits | Implemented (Tested)
[POST /api/market/v2/place-bid](#create-asks-orders) | Implemented (Tested)
POST /api/market/v2/place-ask | Implemented (Tested)
POST /api/market/v2/cancel-order | Implemented (Tested)

## Documentation

### Get started

#### Setup

Init client for API services. 

```golang
client := bitkub.NewClient("api_key", "api_secret")
```

Simply call API in chain style. Call Do() in the end to send HTTP request.
All responses return in go struct.

### API Documentation

#### Get Endpoints Status
```golang
res, err := client.NewGetStatusTx().Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

#### Get Server time
```golang
res, err := client.NewGetServerTimeTx().Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

#### List all symbols
```golang
res, err := client.NewGetSymbolsTx().Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

#### Get ticker information
```golang
res, err := client.NewGetTickerTx().Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

#### List recent trades
```golang
res, err := client.NewGetTradesTx().Symbol(symbols.THB_BTC).Limit(20).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

#### List open buy orders
```golang
res, err := client.NewGetBidsTx().Symbol(symbols.THB_BTC).Limit(20).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

#### List open sell orders
```golang
res, err := client.NewGetAsksTx().Symbol(symbols.THB_BTC).Limit(20).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```

#### List all open orders
```golang
res, err := client.NewGetBooksTx().Symbol(symbols.THB_BTC).Limit(20).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```
#### Get depth information
```golang
res, err := client.NewGetMarketDepthTx().Symbol(symbols.THB_BTC).Limit(20).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```


#### Get historical data from tradingView
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

#### Create Asks Orders
```golang
res, err := client.NewPlaceAskTx().Symbol(symbols.THB_BTC).Amount(0.001).OrderType(types.OrderTypeMarket).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```




