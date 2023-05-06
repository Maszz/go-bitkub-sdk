package bitkub

import (
	"context"

	"github.com/bytedance/sonic"

	BitkubTs "github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type SymbolsService struct {
	c *Client
}

func (s *SymbolsService) Do(ctx context.Context) (res BitkubTs.SymbolsResponse, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: market_symbols_endpoint,
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return BitkubTs.SymbolsResponse{}, err
	}
	resp := BitkubTs.SymbolsResponse{}
	err = sonic.Unmarshal(data, &resp)
	if err != nil {
		return BitkubTs.SymbolsResponse{}, err
	}

	// setparmas stuff.

	return resp, nil
}
