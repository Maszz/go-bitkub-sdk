package bitkub

import (
	"context"

	"github.com/bytedance/sonic"

	BitkubTs "github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type GetStatusService struct {
	c *Client
}

func (s *GetStatusService) Do(ctx context.Context) (res BitkubTs.ServerStatusArray, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: status_endpoint,
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
