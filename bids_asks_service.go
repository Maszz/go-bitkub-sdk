package bitkub

import (
	"context"
	"fmt"

	"github.com/bytedance/sonic"

	BitkubTs "github.com/Maszz/go-bitkub-sdk/types"

	"github.com/valyala/fasthttp"
)

type GetBidsService struct {
	c      *Client
	symbol string
	limit  int
}

func (s *GetBidsService) Symbol(symbol string) *GetBidsService {
	s.symbol = symbol
	return s
}

func (s *GetBidsService) Limit(limit int) *GetBidsService {
	s.limit = limit
	return s
}

func (s *GetBidsService) Do(ctx context.Context) (res BitkubTs.BidsAsksResponse, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: status_endpoint + "?sym=" + s.symbol + "&limit=" + fmt.Sprint(s.limit),
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

type GetAsksService struct {
	c      *Client
	symbol string
	limit  int
}

func (s *GetAsksService) Symbol(symbol string) *GetAsksService {
	s.symbol = symbol
	return s
}

func (s *GetAsksService) Limit(limit int) *GetAsksService {
	s.limit = limit
	return s
}

func (s *GetAsksService) Do(ctx context.Context) (res BitkubTs.BidsAsksResponse, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: status_endpoint + "?sym=" + s.symbol + "&limit=" + fmt.Sprint(s.limit),
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
