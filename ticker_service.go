package bitkub

import (
	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type GetTickerTx struct {
	c *Client
}

func (s *GetTickerTx) Do() (*types.TickerResponse, error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.MarketTickerEndpoint,
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(r)
	if err != nil {
		return nil, err
	}

	res := new(types.TickerResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetTickerTx) DoAny() (*types.TickerResponseAny, error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.MarketTickerEndpoint,
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(r)
	if err != nil {
		return nil, err
	}

	res := make(types.TickerResponseAny, 0)
	err = sonic.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
