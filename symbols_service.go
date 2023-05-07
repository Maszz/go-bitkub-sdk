package bitkub

import (
	"context"

	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type GetSymbolsTx struct {
	c *Client
}

func (s *GetSymbolsTx) Do(ctx context.Context) (res types.SymbolsResponse, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.MarketSymbolsEndpoint,
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return res, err
	}
	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return res, respErr
	}
	err = sonic.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	// setparmas stuff.

	return res, nil
}
