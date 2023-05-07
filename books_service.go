package bitkub

import (
	"context"
	"fmt"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type GetOpenBooksTx struct {
	c      *Client
	symbol string
	limit  int
}

func (s *GetOpenBooksTx) Symbol(symbol string) *GetOpenBooksTx {
	s.symbol = symbol
	return s
}

func (s *GetOpenBooksTx) Limit(limit int) *GetOpenBooksTx {
	s.limit = limit
	return s
}

func (s *GetOpenBooksTx) Do(ctx context.Context) (res types.OpenBooksResponse, err error) {
	err = s.validate()
	if err != nil {
		return res, err
	}

	endpoint := types.MarketBooksEndpoint.String() + "?sym=" + s.symbol + "&lmt=" + fmt.Sprint(s.limit)

	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.NewEndPoint(endpoint),
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

	return res, nil

}

func (s *GetOpenBooksTx) validate() error {
	if s.limit <= 0 {
		s.limit = 10
	}
	if s.symbol == "" {
		return fmt.Errorf("symbol is mandatory")
	}

	return nil
}
