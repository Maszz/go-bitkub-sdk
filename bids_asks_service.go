package bitkub

import (
	"context"
	"fmt"

	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"

	"github.com/valyala/fasthttp"
)

type GetBidsTx struct {
	c      *Client
	symbol types.Symbol
	limit  int
}

func (s *GetBidsTx) Symbol(symbol types.Symbol) *GetBidsTx {
	s.symbol = symbol
	return s
}

func (s *GetBidsTx) Limit(limit int) *GetBidsTx {
	s.limit = limit
	return s
}

func (s *GetBidsTx) Do(ctx context.Context) (res *types.BidsAsksResponse, err error) {

	if err = s.validate(); err != nil {
		return nil, err
	}

	endpoint := types.MarketBidsEndpoint.String() + "?sym=" + s.symbol.String() + "&lmt=" + fmt.Sprint(s.limit)

	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return nil, respErr
	}
	res = new(types.BidsAsksResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetBidsTx) validate() error {
	if s.limit == 0 {
		s.limit = 10
	}
	if s.symbol == "" {
		return fmt.Errorf("symbol is mandatory")
	}
	return nil
}

type GetAsksTx struct {
	c      *Client
	symbol string
	limit  int
}

func (s *GetAsksTx) Symbol(symbol string) *GetAsksTx {
	s.symbol = symbol
	return s
}

func (s *GetAsksTx) Limit(limit int) *GetAsksTx {
	s.limit = limit
	return s
}

func (s *GetAsksTx) Do(ctx context.Context) (res *types.BidsAsksResponse, err error) {
	if err = s.validate(); err != nil {
		return nil, err
	}

	endpoint := types.MarketAsksEndpoint.String() + "?sym=" + s.symbol + "&lmt=" + fmt.Sprint(s.limit)

	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return nil, respErr
	}
	res = new(types.BidsAsksResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *GetAsksTx) validate() error {
	if s.limit == 0 {
		s.limit = 10
	}
	if s.symbol == "" {
		return fmt.Errorf("symbol is mandatory")
	}

	return nil
}

type PlaceBidTx struct {
	c          *Client
	symbol     types.Symbol
	amount     float64
	rate       float64
	order_type types.OrderType
	client_id  string
}

func (s *PlaceBidTx) Symbol(symbol types.Symbol) *PlaceBidTx {
	s.symbol = symbol
	return s
}

func (s *PlaceBidTx) Amount(amount float64) *PlaceBidTx {
	s.amount = amount
	return s
}

func (s *PlaceBidTx) Rate(rate float64) *PlaceBidTx {
	s.rate = rate
	return s
}

func (s *PlaceBidTx) OrderType(order_type types.OrderType) *PlaceBidTx {
	s.order_type = order_type
	return s
}

func (s *PlaceBidTx) ClientID(client_id string) *PlaceBidTx {
	s.client_id = client_id
	return s
}

func (s *PlaceBidTx) Do(ctx context.Context) (res *types.PlaceBidAskResponse, err error) {

	if err = s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketPlaceBidEndpointV2,
		signed:   secTypeSigned,
	}

	payload := types.PlaceBidAskPayload{

		Ts:       utils.CurrentTimestamp(),
		Symbol:   s.symbol,
		Amount:   s.amount,
		Rate:     s.rate,
		Type:     s.order_type,
		ClientID: s.client_id,
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
	res = new(types.PlaceBidAskResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *PlaceBidTx) validate() error {
	if s.symbol == "" {
		return fmt.Errorf("symbol is mandatory")
	}
	if s.amount == 0 {
		return fmt.Errorf("amount is mandatory")
	}
	if s.rate == 0 && s.order_type != types.OrderTypeMarket {
		return fmt.Errorf("rate is mandatory")
	}
	if s.order_type == "" {
		return fmt.Errorf("order_type is mandatory")
	}
	return nil
}

type PlaceAskTx struct {
	c          *Client
	symbol     types.Symbol
	amount     float64
	rate       float64
	order_type types.OrderType
	client_id  string
}

func (s *PlaceAskTx) Symbol(symbol types.Symbol) *PlaceAskTx {
	s.symbol = symbol
	return s
}

func (s *PlaceAskTx) Amount(amount float64) *PlaceAskTx {
	s.amount = amount
	return s
}

func (s *PlaceAskTx) Rate(rate float64) *PlaceAskTx {
	s.rate = rate
	return s
}

func (s *PlaceAskTx) OrderType(order_type types.OrderType) *PlaceAskTx {
	s.order_type = order_type
	return s
}

func (s *PlaceAskTx) ClientID(client_id string) *PlaceAskTx {
	s.client_id = client_id
	return s
}

func (s *PlaceAskTx) Do(ctx context.Context) (res *types.PlaceBidAskResponse, err error) {
	if err = s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketPlaceAskEndpointV2,
		signed:   secTypeSigned,
	}

	payload := types.PlaceBidAskPayload{
		Ts:       utils.CurrentTimestamp(),
		Symbol:   s.symbol,
		Amount:   s.amount,
		Rate:     s.rate,
		Type:     s.order_type,
		ClientID: s.client_id,
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
	res = new(types.PlaceBidAskResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *PlaceAskTx) validate() error {
	if s.symbol == "" {
		return fmt.Errorf("symbol is mandatory")
	}
	if s.amount <= 0 {
		return fmt.Errorf("invalid amount")
	}
	if s.rate == 0 && s.order_type != types.OrderTypeMarket {
		return fmt.Errorf("rate is mandatory")
	}
	if s.order_type == "" {
		return fmt.Errorf("order_type is mandatory")
	}
	return nil
}
