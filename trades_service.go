package bitkub

import (
	"fmt"

	"github.com/Maszz/go-bitkub-sdk/types"

	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type GetTradesTx struct {
	c      *Client
	symbol types.Symbol
	limit  int
}

func (s *GetTradesTx) Symbol(symbol types.Symbol) *GetTradesTx {
	s.symbol = symbol
	return s
}

func (s *GetTradesTx) Limit(limit int) *GetTradesTx {
	s.limit = limit
	return s
}

func (s *GetTradesTx) Do() (*types.TradesResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	endpoint := types.MarketTradesEndpoint.String() + "?sym=" + s.symbol.String() + "&lmt=" + fmt.Sprint(s.limit)
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
	res := new(types.TradesResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetTradesTx) validate() error {
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
