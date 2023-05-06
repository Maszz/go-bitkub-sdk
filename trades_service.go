package bitkub

import (
	"context"
	"fmt"

	BitkubTs "github.com/Maszz/go-bitkub-sdk/types"

	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type GetTradesService struct {
	c      *Client
	symbol string
	limit  int
}

func (s *GetTradesService) Symbol(symbol string) *GetTradesService {
	s.symbol = symbol
	return s
}

func (s *GetTradesService) Limit(limit int) *GetTradesService {
	s.limit = limit
	return s
}

func (s *GetTradesService) Do(ctx context.Context) (res BitkubTs.TradesResponse, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: market_trades_endpoint + "?sym=" + s.symbol + "&limit=" + fmt.Sprint(s.limit),
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

	return res, nil

}
