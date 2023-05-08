package bitkub

import (
	"context"
	"fmt"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type GetMarketDepthTx struct {
	c      *Client
	symbol types.Symbol
	limit  int
}

func (s *GetMarketDepthTx) Symbol(symbol types.Symbol) *GetMarketDepthTx {
	s.symbol = symbol
	return s
}

func (s *GetMarketDepthTx) Limit(limit int) *GetMarketDepthTx {
	s.limit = limit
	return s
}

func (s *GetMarketDepthTx) Do(ctx context.Context) (res *types.MarketDepthResponse, err error) {

	if err = s.validate(); err != nil {
		return nil, err
	}

	endpoint := types.MarketDepthEndpoint.String() + "?sym=" + s.symbol.String() + "&lmt=" + fmt.Sprint(s.limit)

	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(types.MarketDepthResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetMarketDepthTx) validate() error {
	if s.limit <= 0 {
		s.limit = 10
	}
	if s.symbol == "" {
		return fmt.Errorf("symbol is mandatory")
	}

	return nil
}
