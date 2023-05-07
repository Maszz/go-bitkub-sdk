# go-bitkub-sdk
A Golang sdk for bitkub api

## This Module is in initial development it not support majority of a api

some modules behave differently from documentation. This makes it difficult to Test.
Need Manualy call on every endpoint to test json responses it works correctly.
Make sure you have read Bitkub API document before continuing.


## API List 

Name  | Status
------------ | ------------ | 
GET /api/status  | Implemented
Get /api/servertime | Implemented
Get /api/market/symbols | Implemented
Get /api/market/ticker | Implemented
Get /api/market/trades | Implemented
Get /api/market/bids | Implemented
Get /api/market/asks | Implemented
Get /api/market/books | Implemented
Get /api/market/depth | Implemented
Get /tradingview/history | Implemented
Post /api/market/wallet | Implemented
Post /api/market/balances | Implemented
Post /api/market/place-bid | Not implemented(Depecated)
Post /api/market/place-ask | Not implemented(Depecated)
Post /api/market/place-ask-by-fiat | Not implemented(Depecated)
Post /api/market/cancel-order | Not implemented(Depecated)
Post /api/market/my-open-orders | Implemented
Post /api/market/my-orders-history | Implemented
Post /api/market/order-info | Implemented
Post /api/crypto/addresses | Implemented
Post /api/crypto/withdraw | Implemented
Post /api/crypto/deposit-history | Implemented
Post /api/crypto/withdraw-history | Implemented
Post /api/crypto/generate-address | Not Implemented
Post /api/fiat/accounts | Not Implemented
Post /api/fiat/withdraw | Not Implemented
Post /api/fiat/deposit-history | Not implemented
Post /api/fiat/withdraw-history | Not implemented
Post /api/market/wstoken | Not Implemented
Post /api/user/limits | Not Implemented
Post /api/user/trading-credits | Not Implemented
Post /api/market/v2/place-bid | Implemented
Post /api/market/v2/place-ask | Implemented
Post /api/market/v2/cancel-order | Implemented

##

