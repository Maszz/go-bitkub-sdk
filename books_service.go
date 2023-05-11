package bitkub

import (
	"fmt"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type GetOpenBooksTx struct {
	c      *Client
	symbol types.Symbol
	limit  int
}

func (s *GetOpenBooksTx) Symbol(symbol types.Symbol) *GetOpenBooksTx {
	s.symbol = symbol
	return s
}

func (s *GetOpenBooksTx) Limit(limit int) *GetOpenBooksTx {
	s.limit = limit
	return s
}

func (s *GetOpenBooksTx) Do() (*types.OpenBooksResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	endpoint := types.MarketBooksEndpoint.String() + "?sym=" + s.symbol.String() + "&lmt=" + fmt.Sprint(s.limit)

	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(r)
	if err != nil {
		return nil, err
	}

	respErr := s.c.catchAPIError(data)
	if respErr != nil {
		return nil, respErr
	}
	res := new(types.OpenBooksResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetOpenBooksTx) validate() error {
	if s.limit <= 0 {
		s.limit = 10
	}
	if s.symbol == "" {
		return types.ErrSymbolMandatory
	}

	return nil
}
