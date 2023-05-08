package bitkub

import (
	"context"
	"fmt"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"
	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type CancelOrderTx struct {
	c          *Client
	symbol     types.Symbol
	order_id   types.OrderId
	order_side types.OrderSide
	order_hash types.OrderHash
}

func (s *CancelOrderTx) Symbol(symbol types.Symbol) *CancelOrderTx {
	s.symbol = symbol
	return s
}

func (s *CancelOrderTx) OrderID(order_id types.OrderId) *CancelOrderTx {
	s.order_id = order_id
	return s
}

func (s *CancelOrderTx) OrderSide(order_side types.OrderSide) *CancelOrderTx {
	switch order_side {
	case types.OrderSideBuy:
		s.order_side = types.OrderSideBuy
	case types.OrderSideSell:
		s.order_side = types.OrderSideSell
	default:
		panic("Invalid order_side")
	}
	return s
}

func (s *CancelOrderTx) OrderHash(order_hash types.OrderHash) *CancelOrderTx {
	s.order_hash = order_hash
	return s
}

func (s *CancelOrderTx) Do(ctx context.Context) (res *types.CancelOrderResponse, err error) {

	if err = s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketCancelOrderEndpoint,
		signed:   secTypeSigned,
	}

	payload := types.CancelOrderPayload{

		Ts:        utils.CurrentTimestamp(),
		Symbol:    s.symbol,
		OrderID:   s.order_id,
		OrderSide: s.order_side,
		OrderHash: s.order_hash,
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return nil, err
	}
	r.body = byteBody
	data, err := s.c.callAPI(ctx, r)

	if err != nil {
		return nil, err
	}
	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return nil, respErr
	}
	res = new(types.CancelOrderResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *CancelOrderTx) validate() (err error) {
	if s.order_hash != "" {
		return nil
	}
	if s.symbol == "" {
		return fmt.Errorf("invalid symbol")
	}
	if s.order_id == "" {
		return fmt.Errorf("invalid order_id")
	}
	if s.order_side == "" {
		return fmt.Errorf("invalid order_side")
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

func (s *GetOpenOrdersTx) Do(ctx context.Context) (res *types.GetOpenOrdersResponse, err error) {
	if err = s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketMyOpenOrdersEndpoint,
		signed:   secTypeSigned,
	}

	payload := types.GetOpenOrdersPayload{
		Ts:     utils.CurrentTimestamp(),
		Symbol: s.symbol,
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return nil, err
	}
	r.body = byteBody
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return nil, respErr
	}
	res = new(types.GetOpenOrdersResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *GetOpenOrdersTx) validate() (err error) {
	if s.symbol == "" {
		return fmt.Errorf("invalid symbol")
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

func (s *GetOrderHistoryTx) Do(ctx context.Context) (res *types.GetOrderHistoryResponse, err error) {
	if err = s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketMyOrderHistoryEndpoint,
		signed:   secTypeSigned,
	}

	payload := types.GetOrderHistoryPayload{
		Ts:     utils.CurrentTimestamp(),
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
	data, err := s.c.callAPI(ctx, r)

	if err != nil {
		return nil, err
	}
	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return nil, respErr
	}
	res = new(types.GetOrderHistoryResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetOrderHistoryTx) validate() (err error) {
	if s.symbol == "" {
		return fmt.Errorf("require symbol")
	}
	if s.page < 0 {
		return fmt.Errorf("invalid page")
	}
	if s.limit < 0 {
		return fmt.Errorf("invalid limit")
	}
	if s.start < 0 {
		return fmt.Errorf("invalid start")
	}
	if s.end < 0 {
		return fmt.Errorf("invalid end")
	}
	return nil
}

type GetOrderInfoTx struct {
	c          *Client
	symbol     types.Symbol
	order_id   types.OrderId
	order_side types.OrderSide
	order_hash types.OrderHash
}

func (s *GetOrderInfoTx) Symbol(symbol types.Symbol) *GetOrderInfoTx {
	s.symbol = symbol
	return s
}

func (s *GetOrderInfoTx) OrderID(order_id types.OrderId) *GetOrderInfoTx {
	s.order_id = order_id
	return s
}

func (s *GetOrderInfoTx) OrderSide(order_side types.OrderSide) *GetOrderInfoTx {
	switch order_side {
	case types.OrderSideBuy:
		s.order_side = types.OrderSideBuy
	case types.OrderSideSell:
		s.order_side = types.OrderSideSell
	default:
		panic("Invalid order_side")
	}
	return s
}

func (s *GetOrderInfoTx) OrderHash(order_hash types.OrderHash) *GetOrderInfoTx {
	s.order_hash = order_hash
	return s
}

func (s *GetOrderInfoTx) Do(ctx context.Context) (res *types.GetOrdersInfoResponse, err error) {
	if err = s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketOrderInfoEndpoint,
		signed:   secTypeSigned,
	}

	payload := types.GetOrdersInfoPayload{

		Ts:        utils.CurrentTimestamp(),
		Symbol:    s.symbol,
		OrderID:   s.order_id,
		OrderSide: s.order_side,
		OrderHash: s.order_hash,
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return nil, err
	}
	r.body = byteBody
	data, err := s.c.callAPI(ctx, r)

	if err != nil {
		return nil, err
	}
	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return nil, respErr
	}
	res = new(types.GetOrdersInfoResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetOrderInfoTx) validate() (err error) {
	if s.order_hash != "" {
		return nil
	}
	if s.symbol == "" {
		return fmt.Errorf("invalid symbol")
	}
	if s.order_id == "" {
		return fmt.Errorf("invalid order_id")
	}
	if s.order_side == "" {
		return fmt.Errorf("invalid order_side")
	}
	return nil

}
