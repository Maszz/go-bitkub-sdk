package bitkub

import (
	"fmt"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"
	"github.com/bytedance/sonic"
)

type GetTradingViewHistoryTx struct {
	c          *Client
	symbol     types.Symbol
	from       int64
	to         int64
	resolution types.TimeResolution
}

func (s *GetTradingViewHistoryTx) Symbol(symbol types.Symbol) *GetTradingViewHistoryTx {
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

func (s *GetTradingViewHistoryTx) ToCurrent() *GetTradingViewHistoryTx {
	s.to = utils.RawCurrentTimestamp()
	return s
}
func (s *GetTradingViewHistoryTx) Resolution(resolution types.TimeResolution) *GetTradingViewHistoryTx {
	s.resolution = resolution
	return s
}

func (s *GetTradingViewHistoryTx) Do() (*types.TradingViewHistoryResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	endpoint := types.TradingviewHistoryEndpoint.String() + "?symbol=" + s.symbol.String() + "&resolution=" + s.resolution.String() + "&from=" + fmt.Sprint(s.from) + "&to=" + fmt.Sprint(s.to)

	r := &request{
		method:   "GET",
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeNone,
	}

	data, err := s.c.callAPI(r)
	if err != nil {
		return nil, err
	}
	res := new(types.TradingViewHistoryResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *GetTradingViewHistoryTx) validate() error {
	if s.symbol == "" {
		return types.ErrSymbolMandatory
	}
	if s.from <= 0 {
		// s.from = utils.RawCurrentTimestamp() - 86400
		return types.ErrInvalidTimeStamp
	}
	if s.to <= 0 {
		// s.to = utils.RawCurrentTimestamp()
		return types.ErrInvalidTimeStamp
	}
	if err := s.validateResolution(); err != nil {
		return err
	}
	return nil
}

func (s *GetTradingViewHistoryTx) validateResolution() error {
	for _, v := range types.TimeResolutions {
		if s.resolution == v {
			return nil
		}
	}
	return types.ErrInvalidTimeResolution
}
