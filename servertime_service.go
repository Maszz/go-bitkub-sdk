package bitkub

import (
	"context"
	"encoding/binary"

	"github.com/bytedance/sonic"

	"github.com/valyala/fasthttp"
)

type GetServerTimeService struct {
	c *Client
}

func (s *GetServerTimeService) Do(ctx context.Context) (res uint64, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: servertime_endpoint,
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return 0, err
	}
	resp := binary.BigEndian.Uint64(data)
	err = sonic.Unmarshal(data, &resp)
	if err != nil {
		return 0, err
	}

	// setparmas stuff.

	return resp, nil
}
