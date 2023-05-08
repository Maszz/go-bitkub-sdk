package bitkub

import (
	"context"
	"encoding/binary"

	"github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type GetServerTimeTx struct {
	c *Client
}

func (s *GetServerTimeTx) Do(ctx context.Context) (res *uint64, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.ServertimeEndpoint,
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := binary.BigEndian.Uint64(data)
	res = &resp

	return res, nil
}
