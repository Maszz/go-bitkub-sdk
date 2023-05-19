package bitkub

import (
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

func (s *GetBidsTx) Do() (*types.BidsAsksResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	endpoint := types.MarketBidsEndpoint.String() + "?sym=" + s.symbol.String() + "&lmt=" + fmt.Sprint(s.limit)

	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(r)
	if err != nil {
		return nil, err
	}
	respErr := s.c.catchAPIError(data)
	if respErr != nil {
		return nil, respErr
	}
	res := new(types.BidsAsksResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetBidsTx) validate() error {
	if s.symbol == "" {
		return types.ErrSymbolMandatory
	}
	if s.limit == 0 {
		return types.ErrLimitMandatory
	}
	if s.limit < 0 {
		return types.ErrLimitMustBePositive
	}

	return nil
}

type GetAsksTx struct {
	c      *Client
	symbol types.Symbol
	limit  int
}

func (s *GetAsksTx) Symbol(symbol types.Symbol) *GetAsksTx {
	s.symbol = symbol
	return s
}

func (s *GetAsksTx) Limit(limit int) *GetAsksTx {
	s.limit = limit
	return s
}

func (s *GetAsksTx) Do() (*types.BidsAsksResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	endpoint := types.MarketAsksEndpoint.String() + "?sym=" + s.symbol.String() + "&lmt=" + fmt.Sprint(s.limit)

	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(r)
	if err != nil {
		return nil, err
	}
	respErr := s.c.catchAPIError(data)
	if respErr != nil {
		return nil, respErr
	}
	res := new(types.BidsAsksResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetAsksTx) validate() error {
	if s.symbol == "" {
		return types.ErrSymbolMandatory
	}
	if s.limit == 0 {
		return types.ErrLimitMandatory
	}
	if s.limit < 0 {
		return types.ErrLimitMustBePositive
	}
	return nil
}

type PlaceBidTx struct {
	c         *Client
	symbol    types.Symbol
	amount    float64
	rate      float64
	orderType types.OrderType
	clientID  string
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

func (s *PlaceBidTx) OrderType(orderType types.OrderType) *PlaceBidTx {
	s.orderType = orderType
	return s
}

func (s *PlaceBidTx) ClientID(clientID string) *PlaceBidTx {
	s.clientID = clientID
	return s
}

func (s *PlaceBidTx) Do() (*types.PlaceBidAskResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketPlaceBidEndpointV2,
		signed:   secTypeSigned,
	}

	payload := types.PlaceBidAskPayload{

		TS:       utils.CurrentTimestamp(),
		Symbol:   s.symbol,
		Amount:   s.amount,
		Rate:     s.rate,
		Type:     s.orderType,
		ClientID: s.clientID,
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
	res := new(types.PlaceBidAskResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PlaceBidTx) validate() error {
	if s.symbol == "" {
		return types.ErrSymbolMandatory
	}
	if s.amount == 0 {
		return types.ErrAmountMandatory
	}
	if s.rate == 0 && s.orderType != types.OrderTypeMarket {
		return types.ErrRateMandatory
	}
	if s.orderType == "" {
		return types.ErrOrderTypeMandatory
	}
	return nil
}

type PlaceAskTx struct {
	c         *Client
	symbol    types.Symbol
	amount    float64
	rate      float64
	orderType types.OrderType
	clientID  string
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

func (s *PlaceAskTx) OrderType(orderType types.OrderType) *PlaceAskTx {
	s.orderType = orderType
	return s
}

func (s *PlaceAskTx) ClientID(clientID string) *PlaceAskTx {
	s.clientID = clientID
	return s
}

func (s *PlaceAskTx) Do() (*types.PlaceBidAskResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketPlaceAskEndpointV2,
		signed:   secTypeSigned,
	}

	payload := types.PlaceBidAskPayload{
		TS:       utils.CurrentTimestamp(),
		Symbol:   s.symbol,
		Amount:   s.amount,
		Rate:     s.rate,
		Type:     s.orderType,
		ClientID: s.clientID,
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
	res := new(types.PlaceBidAskResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PlaceAskTx) validate() error {
	if s.symbol == "" {
		return types.ErrSymbolMandatory
	}
	if s.amount <= 0 {
		return types.ErrAmountMandatory
	}
	if s.rate == 0 && s.orderType != types.OrderTypeMarket {
		return types.ErrRateMandatory
	}
	if s.orderType == "" {
		return types.ErrOrderTypeMandatory
	}
	return nil
}
