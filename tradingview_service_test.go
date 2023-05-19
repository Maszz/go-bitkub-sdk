package bitkub

import (
	"fmt"
	"testing"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/suite"
)

type tradingViewServiceTestSuite struct {
	baseTestSuite
	mockData          []byte
	unmarshalMockData []byte
}

func TestTradingViewService(t *testing.T) {
	suite.Run(t, new(tradingViewServiceTestSuite))
}

func (s *tradingViewServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.mockData = []byte(`{
		"c": [1691200, 1683819.64],
		"h": [1692222.22, 1693000],
		"l": [1671000, 1680000],
		"o": [1682500, 1692222.17],
		"s": "ok",
		"t": [1633424400, 1633428000],
		"v": [22.001944110000036, 13.138984619999986]
	  }
	  `)

	s.unmarshalMockData = []byte(`{
		"c": [1691200, 1683819.64],
		"h": [1692222.22, 1693000],
		"l": [1671000, 1680000],
		"o": [1682500, 1692222.17],
		"s": 0,
		"t": [1633424400, 1633428000],
		"v": [22.001944110000036, 13.138984619999986]
	  }
	  `)
}

func (s *tradingViewServiceTestSuite) TestSetSymbol() {
	tx := s.client.NewGetTradingviewHistoryTx()
	symbol := types.Symbol("THB_BTC")
	tx.Symbol(symbol)
	s.r().Equal(symbol, tx.symbol)
}

func (s *tradingViewServiceTestSuite) TestSetFromTimestamp() {
	tx := s.client.NewGetTradingviewHistoryTx()
	fromTimestamp := int64(1000)
	tx.FromTimestamp(fromTimestamp)
	s.r().Equal(fromTimestamp, tx.from)
}

func (s *tradingViewServiceTestSuite) TestSetToTimestamp() {
	tx := s.client.NewGetTradingviewHistoryTx()
	toTimestamp := int64(1000)
	tx.ToTimestamp(toTimestamp)
	s.r().Equal(toTimestamp, tx.to)
}

func (s *tradingViewServiceTestSuite) TestSetToCurrentTimestamp() {
	tx := s.client.NewGetTradingviewHistoryTx()
	tx.ToCurrent()
	s.r().NotEqual(int64(0), tx.to)
}

func (s *tradingViewServiceTestSuite) TestSetResolution() {
	tx := s.client.NewGetTradingviewHistoryTx()
	tx.Resolution(types.Time1m)
	s.r().Equal(types.Time1m, tx.resolution)

}

func (s *tradingViewServiceTestSuite) TestGetTradingViewHistoryValidateSymbol() {
	tx := s.client.NewGetTradingviewHistoryTx()
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("")
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)

}

func (s *tradingViewServiceTestSuite) TestGetTradingViewHistoryValidateFromTimestamp() {
	tx := s.client.NewGetTradingviewHistoryTx()
	tx.Symbol("THB_BTC")
	s.r().ErrorIs(tx.validate(), types.ErrInvalidTimeStamp)
	tx.FromTimestamp(0)
	s.r().ErrorIs(tx.validate(), types.ErrInvalidTimeStamp)
}

func (s *tradingViewServiceTestSuite) TestGetTradingViewHistoryValidateToTimestamp() {
	tx := s.client.NewGetTradingviewHistoryTx()
	tx.Symbol("THB_BTC")
	tx.FromTimestamp(1000)
	s.r().ErrorIs(tx.validate(), types.ErrInvalidTimeStamp)
	tx.ToTimestamp(0)
	s.r().ErrorIs(tx.validate(), types.ErrInvalidTimeStamp)
}

func (s *tradingViewServiceTestSuite) TestGetTradingViewHistoryValidateResolution() {
	tx := s.client.NewGetTradingviewHistoryTx()
	tx.Symbol("THB_BTC")
	tx.FromTimestamp(1000)
	tx.ToTimestamp(2000)
	s.r().ErrorIs(tx.validate(), types.ErrInvalidTimeResolution)
	tx.Resolution("")
	s.r().ErrorIs(tx.validate(), types.ErrInvalidTimeResolution)
	tx.Resolution("1y")
	s.r().ErrorIs(tx.validate(), types.ErrInvalidTimeResolution)

	tx.Resolution(types.Time1m)
	s.r().NoError(tx.validate())
}

func (s *tradingViewServiceTestSuite) TestGetTradingViewHistoryDoValidate() {
	_, err := s.client.NewGetTradingviewHistoryTx().Do()
	s.r().Error(err)
	s.r().ErrorIs(err, types.ErrSymbolMandatory)

}

func (s *tradingViewServiceTestSuite) TestGetTradingViewHistory() {
	s.mockDo(s.mockData, nil)
	mockDataStuct := new(types.TradingViewHistoryResponse)
	err := sonic.Unmarshal(s.mockData, mockDataStuct)
	s.r().NoError(err)

	tx, err := s.client.NewGetTradingviewHistoryTx().FromTimestamp(1633424400).ToTimestamp(1633428000).Resolution(types.Time1m).Symbol("THB_BTC").Do()
	defer s.assertDo()
	s.r().NoError(err)
	s.r().Equal(mockDataStuct, tx)
}

func (s *tradingViewServiceTestSuite) TestGetTradingViewHistoryHTTPError() {
	s.mockDo(nil, fmt.Errorf("dummy error"))
	_, err := s.client.NewGetTradingviewHistoryTx().FromTimestamp(1633424400).ToTimestamp(1633428000).Resolution(types.Time1m).Symbol("THB_BTC").Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "dummy error")
}

// func (s *tradingViewServiceTestSuite) TestGetTradingViewUnmarshalError() {
// 	s.mockDo(s.unmarshalMockData, nil)
// 	data, err := s.client.NewGetTradingviewHistoryTx().FromTimestamp(1633424400).ToTimestamp(1633428000).Resolution(types.Time1m).Symbol("THB_BTC").Do()
// 	defer s.assertDo()

// 	s.r().Nil(data)
// 	s.r().Error(err)
// 	s.r().EqualError(err, "json: cannot unmarshal number into Go struct field TradingViewHistoryResponse.s of type string")
// }
