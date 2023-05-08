package bitkub

import (
	"context"

	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type GetStatusTx struct {
	c *Client
}

func (s *GetStatusTx) Do(ctx context.Context) (res *types.ServerStatusArray, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.StatusEndpoint,
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(types.ServerStatusArray)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
