package bitkub

import (
	"context"

	"github.com/bytedance/sonic"

	BitkubTs "github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type GetTickerService struct {
	c *Client
}

func (s *GetTickerService) Do(ctx context.Context) (res BitkubTs.TickerResponse, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: market_ticker_endpoint,
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

	// setparmas stuff.

	return res, nil
}

func (s *GetTickerService) DoAny(ctx context.Context) (res BitkubTs.TickerResponseAny, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: market_ticker_endpoint,
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

	// setparmas stuff.

	return res, nil
}
