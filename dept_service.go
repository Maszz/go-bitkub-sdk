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
	symbol string
	limit  int
}

func (s *GetMarketDepthTx) Symbol(symbol string) *GetMarketDepthTx {
	s.symbol = symbol
	return s
}

func (s *GetMarketDepthTx) Limit(limit int) *GetMarketDepthTx {
	s.limit = limit
	return s
}

func (s *GetMarketDepthTx) Do(ctx context.Context) (res types.MarketDepthResponse, err error) {

	err = s.validate()
	if err != nil {
		return res, err
	}

	endpoint := types.MarketDepthEndpoint.String() + "?sym=" + s.symbol + "&lmt=" + fmt.Sprint(s.limit)

	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return res, err
	}
	err = sonic.Unmarshal(data, &res)
	if err != nil {
		return res, err
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
