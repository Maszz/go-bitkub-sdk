package bitkub

import (
	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"
	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type CancelOrderTx struct {
	c         *Client
	symbol    types.Symbol
	orderID   types.OrderID
	orderSide types.OrderSide
	orderHash types.OrderHash
}

func (s *CancelOrderTx) Symbol(symbol types.Symbol) *CancelOrderTx {
	s.symbol = symbol
	return s
}

func (s *CancelOrderTx) OrderID(orderID types.OrderID) *CancelOrderTx {
	s.orderID = orderID
	return s
}

func (s *CancelOrderTx) OrderSide(orderSide types.OrderSide) *CancelOrderTx {
	switch orderSide {
	case types.OrderSideBuy:
		s.orderSide = types.OrderSideBuy
	case types.OrderSideSell:
		s.orderSide = types.OrderSideSell
	default:
		panic(types.ErrInvalidOrderSide)
	}
	return s
}

func (s *CancelOrderTx) OrderHash(orderHash types.OrderHash) *CancelOrderTx {
	s.orderHash = orderHash
	return s
}

func (s *CancelOrderTx) Do() (*types.CancelOrderResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketCancelOrderEndpoint,
		signed:   secTypeSigned,
	}

	payload := types.CancelOrderPayload{

		TS:        utils.CurrentTimestamp(),
		Symbol:    s.symbol,
		OrderID:   s.orderID,
		OrderSide: s.orderSide,
		OrderHash: s.orderHash,
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return nil, err
	}
	r.body = byteBody
	data, err := s.c.callAPI(r)

	if err != nil {
		return nil, err
	}
	respErr := s.c.catchAPIError(data)
	if respErr != nil {
		return nil, respErr
	}
	res := new(types.CancelOrderResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CancelOrderTx) validate() error {
	if s.orderHash != "" {
		return nil
	}
	if s.symbol == "" {
		return types.ErrSymbolMandatory
	}
	if s.orderID == "" {
		return types.ErrOrderIDMandatory
	}
	if s.orderSide == "" {
		return types.ErrInvalidOrderSide
	}
	return nil
}

type GetOpenOrdersTx struct {
	c      *Client
	symbol types.Symbol
}

func (s *GetOpenOrdersTx) Symbol(symbol types.Symbol) *GetOpenOrdersTx {
	s.symbol = symbol
	return s
}

func (s *GetOpenOrdersTx) Do() (*types.GetOpenOrdersResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketMyOpenOrdersEndpoint,
		signed:   secTypeSigned,
	}

	payload := types.GetOpenOrdersPayload{
		TS:     utils.CurrentTimestamp(),
		Symbol: s.symbol,
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return nil, err
	}
	r.body = byteBody
	data, err := s.c.callAPI(r)
	if err != nil {
		return nil, err
	}
	respErr := s.c.catchAPIError(data)
	if respErr != nil {
		return nil, respErr
	}
	res := new(types.GetOpenOrdersResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetOpenOrdersTx) validate() error {
	if s.symbol == "" {
		return types.ErrSymbolMandatory
	}
	return nil
}

type GetOrderHistoryTx struct {
	c      *Client
	symbol types.Symbol
	page   int
	limit  int
	start  types.Timestamp
	end    types.Timestamp
}

func (s *GetOrderHistoryTx) Symbol(symbol types.Symbol) *GetOrderHistoryTx {
	s.symbol = symbol
	return s
}

func (s *GetOrderHistoryTx) Page(page int) *GetOrderHistoryTx {
	s.page = page
	return s
}

func (s *GetOrderHistoryTx) Limit(limit int) *GetOrderHistoryTx {
	s.limit = limit
	return s
}

func (s *GetOrderHistoryTx) Start(start types.Timestamp) *GetOrderHistoryTx {
	s.start = start
	return s
}

func (s *GetOrderHistoryTx) End(end types.Timestamp) *GetOrderHistoryTx {
	s.end = end
	return s
}

func (s *GetOrderHistoryTx) Do() (*types.GetOrderHistoryResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketMyOrderHistoryEndpoint,
		signed:   secTypeSigned,
	}

	payload := types.GetOrderHistoryPayload{
		TS:     utils.CurrentTimestamp(),
		Symbol: s.symbol,
		Page:   s.page,
		Limit:  s.limit,
		Start:  s.start,
		End:    s.end,
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return nil, err
	}
	r.body = byteBody
	data, err := s.c.callAPI(r)

	if err != nil {
		return nil, err
	}
	respErr := s.c.catchAPIError(data)
	if respErr != nil {
		return nil, respErr
	}
	res := new(types.GetOrderHistoryResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetOrderHistoryTx) validate() error {
	if s.symbol == "" {
		return types.ErrSymbolMandatory
	}
	if s.page < 0 {
		return types.ErrPageMustBePositive
	}
	if s.limit < 0 {
		return types.ErrLimitMustBePositive
	}
	if s.start < 0 {
		return types.ErrInvalidTimeStamp
	}
	if s.end < 0 {
		return types.ErrInvalidTimeStamp
	}
	return nil
}

type GetOrderInfoTx struct {
	c         *Client
	symbol    types.Symbol
	orderID   types.OrderID
	orderSide types.OrderSide
	orderHash types.OrderHash
}

func (s *GetOrderInfoTx) Symbol(symbol types.Symbol) *GetOrderInfoTx {
	s.symbol = symbol
	return s
}

func (s *GetOrderInfoTx) OrderID(orderID types.OrderID) *GetOrderInfoTx {
	s.orderID = orderID
	return s
}

func (s *GetOrderInfoTx) OrderSide(orderSide types.OrderSide) *GetOrderInfoTx {
	switch orderSide {
	case types.OrderSideBuy:
		s.orderSide = types.OrderSideBuy
	case types.OrderSideSell:
		s.orderSide = types.OrderSideSell
	default:
		panic(types.ErrInvalidOrderSide)
	}
	return s
}

func (s *GetOrderInfoTx) OrderHash(orderHash types.OrderHash) *GetOrderInfoTx {
	s.orderHash = orderHash
	return s
}

func (s *GetOrderInfoTx) Do() (*types.GetOrdersInfoResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketOrderInfoEndpoint,
		signed:   secTypeSigned,
	}

	payload := types.GetOrdersInfoPayload{
		TS:        utils.CurrentTimestamp(),
		Symbol:    s.symbol,
		OrderID:   s.orderID,
		OrderSide: s.orderSide,
		OrderHash: s.orderHash,
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return nil, err
	}
	r.body = byteBody
	data, err := s.c.callAPI(r)

	if err != nil {
		return nil, err
	}
	respErr := s.c.catchAPIError(data)
	if respErr != nil {
		return nil, respErr
	}
	res := new(types.GetOrdersInfoResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetOrderInfoTx) validate() error {
	if s.orderHash != "" {
		return nil
	}
	if s.symbol == "" {
		return types.ErrSymbolMandatory
	}
	if s.orderID == "" {
		return types.ErrOrderIDMandatory
	}
	if s.orderSide == "" {
		return types.ErrInvalidOrderSide
	}
	return nil
}
