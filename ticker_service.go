package bitkub

import (
	"context"

	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type GetTickerTx struct {
	c *Client
}

func (s *GetTickerTx) Do(ctx context.Context) (res types.TickerResponse, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.MarketTickerEndpoint,
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

	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return res, respErr
	}

	// setparmas stuff.

	return res, nil
}

func (s *GetTickerTx) DoAny(ctx context.Context) (res types.TickerResponseAny, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.MarketTickerEndpoint,
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

	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return res, respErr
	}

	// setparmas stuff.

	return res, nil
}
