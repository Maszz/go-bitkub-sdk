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
GET /api/status  | Implemented(Tested)
Get /api/servertime | Implemented(Tested)
Get /api/market/symbols | Implemented(Tested)
Get /api/market/ticker | Implemented(Tested)
Get /api/market/trades | Implemented(Tested)
Get /api/market/bids | Implemented(Tested)
Get /api/market/asks | Implemented(Tested)
Get /api/market/books | Implemented(Tested)
Get /api/market/depth | Implemented(Tested)
Get /tradingview/history | Implemented(Tested)
Post /api/market/wallet | Implemented(Tested)
Post /api/market/balances | Implemented(Tested)
Post /api/market/place-bid | Not implemented(Depecated)
Post /api/market/place-ask | Not implemented(Depecated)
Post /api/market/place-ask-by-fiat | Not implemented(Depecated)
Post /api/market/cancel-order | Not implemented(Depecated)
Post /api/market/my-open-orders | Implemented(Tested)
Post /api/market/my-orders-history | Implemented(Tested)
Post /api/market/order-info | Implemented(Tested)
Post /api/crypto/addresses | Implemented(Tested)
Post /api/crypto/withdraw | Implemented(Tested)
Post /api/crypto/deposit-history | Implemented(Tested)
Post /api/crypto/withdraw-history | Implemented(Tested)
Post /api/crypto/generate-address | Can't be implemented(Lack of document).
Post /api/fiat/accounts | Implemented(Tested)
Post /api/fiat/withdraw | Implemented(Tested)
Post /api/fiat/deposit-history | Implemented(Tested)
Post /api/fiat/withdraw-history | Implemented(Tested)
Post /api/market/wstoken | Implemented(Tested)
Post /api/user/limits | Implemented(Tested)
Post /api/user/trading-credits | Implemented (Tested)
Post /api/market/v2/place-bid | Implemented (Tested)
Post /api/market/v2/place-ask | Implemented (Tested)
Post /api/market/v2/cancel-order | Implemented (Tested)

## Documentation

#### Setup

Init client for API services. 

```golang
bitkubClient := bitkub.NewClient("api_key", "api_secret")
```

Simply call API in chain style. Call Do() in the end to send HTTP request.
All responses return in go struct.

#### Create Asks Orders
```golang
res, err := bitkubClient.NewPlaceAskTx().Symbol(symbols.THB_BTC).Amount(0.001).OrderType(types.OrderTypeMarket).Do(context.Background())
if err != nil {
		fmt.Println(err)
		return
	}
jsonEnc, _ := json.Marshal(res)
fmt.Println(string(jsonEnc))
```





