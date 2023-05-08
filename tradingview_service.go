package bitkub

import (
	"context"
	"fmt"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"
	"github.com/bytedance/sonic"
)

type GetTradingViewHistoryTx struct {
	c          *Client
	symbol     string
	from       int64
	to         int64
	resolution types.TimeResolution
}

func (s *GetTradingViewHistoryTx) Symbol(symbol string) *GetTradingViewHistoryTx {
	s.symbol = symbol
	return s
}

func (s *GetTradingViewHistoryTx) FromTimestamp(from int64) *GetTradingViewHistoryTx {
	s.from = from
	return s
}

func (s *GetTradingViewHistoryTx) ToTimestamp(to int64) *GetTradingViewHistoryTx {
	s.to = to
	return s
}

func (s *GetTradingViewHistoryTx) ToCurrent(to int64) *GetTradingViewHistoryTx {
	s.to = utils.RawCurrentTimestamp()
	return s
}
func (s *GetTradingViewHistoryTx) Resolution(resolution types.TimeResolution) *GetTradingViewHistoryTx {

	switch resolution {
	case types.Time_1m:
		s.resolution = types.Time_1m
	case types.Time_5m:
		s.resolution = types.Time_5m
	case types.Time_15m:
		s.resolution = types.Time_15m
	case types.Time_1h:
		s.resolution = types.Time_1h
	case types.Time_240m:
		s.resolution = types.Time_240m
	case types.Time_1d:
		s.resolution = types.Time_1d
	default:
		panic("Invalid resolution")
	}
	s.resolution = resolution
	return s
}

func (s *GetTradingViewHistoryTx) Do(ctx context.Context) (res *types.TradingViewHistoryResponse, err error) {

	if err = s.validate(); err != nil {
		return nil, err
	}

	endpoint := types.TradingviewHistoryEndpoint.String() + "?symbol=" + s.symbol + "&resolution=" + s.resolution.String() + "&from=" + fmt.Sprint(s.from) + "&to=" + fmt.Sprint(s.to)

	r := &request{
		method:   "GET",
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeNone,
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(types.TradingViewHistoryResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *GetTradingViewHistoryTx) validate() error {
	if s.from <= 0 {
		s.from = utils.RawCurrentTimestamp() - 86400
	}
	if s.to <= 0 {
		s.to = utils.RawCurrentTimestamp()
	}
	if s.resolution == "" {
		s.resolution = types.Time_1h
	}
	if s.symbol == "" {
		return fmt.Errorf("symbol is mandatory")
	}
	return nil
}
