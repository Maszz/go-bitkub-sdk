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
[GET /api/status] | Implemented(Tested)
[GET /api/servertime] | Implemented(Tested)
GET /api/market/symbols | Implemented(Tested)
GET /api/market/ticker | Implemented(Tested)
GET /api/market/trades | Implemented(Tested)
GET /api/market/bids | Implemented(Tested)
GET /api/market/asks| Implemented(Tested)
GET /api/market/books | Implemented(Tested)
GET /api/market/depth | Implemented(Tested)
GET /tradingview/history | Implemented(Tested)
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
POST /api/market/v2/place-bid | Implemented (Tested)
POST /api/market/v2/place-ask | Implemented (Tested)
POST /api/market/v2/cancel-order | Implemented (Tested)


# Get started

### Setup

Init client for API services. 

```golang
client := bitkub.NewClient("api_key", "api_secret")
```

Simply call API in chain style. Call Do() in the end to send HTTP request.
All responses return in go struct.

>For more information about This library read the [documentation](/api_doc.md)

### Create Sell Order
```golang
res, err := client.NewPlaceAskTx().Symbol(symbols.THB_BTC).Amount(0.001).OrderType(types.OrderTypeMarket).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
``` 

# Dependencies
Instead of using net/http, this library uses fasthttp as its HTTP client and utilizes sonic for JSON serialization and deserialization. However, if you prefer not to use these libraries or if your application is affected by any edge cases, future versions of this library may include an adapter to support external libraries for serializing, deserializing, and making HTTP requests.