package bitkub

import (
	"strconv"

	"github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type GetServerTimeTx struct {
	c *Client
}

func (s *GetServerTimeTx) Do() (*int, error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.ServertimeEndpoint,
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(r)
	if err != nil {
		return nil, err
	}
	resp, _ := strconv.Atoi(string(data))

	res := &resp

	return res, nil
}
