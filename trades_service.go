package bitkub

import (
	"context"
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

func (s *GetTradesTx) Do(ctx context.Context) (res *types.TradesResponse, err error) {
	if err = s.validate(); err != nil {
		return nil, err
	}

	endpoint := types.MarketTradesEndpoint.String() + "?sym=" + s.symbol.String() + "&lmt=" + fmt.Sprint(s.limit)
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
	res = new(types.TradesResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *GetTradesTx) validate() error {
	if s.limit <= 0 {
		s.limit = 10
	}
	if s.symbol == "" {
		return fmt.Errorf("symbol is mandatory")
	}

	return nil
}
