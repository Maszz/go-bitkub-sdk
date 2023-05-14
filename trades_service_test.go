package bitkub

import (
	"fmt"
	"testing"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/suite"
)

type tradesServiceTestSuite struct {
	baseTestSuite
	mockData          []byte
	unmarshalMockData []byte
	apiErrorMockData  []byte
}

func TestTradesService(t *testing.T) {
	suite.Run(t, new(tradesServiceTestSuite))
}

func (s *tradesServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.mockData = []byte(`{
		"error": 0,
		"result": [
		  [1684080070, 912654.8, 0.0000111, "BUY"],
		  [1684080040, 912700, 0.00001092, "BUY"],
		  [1684080018, 912700, 0.01000027, "BUY"],
		  [1684080017, 912700, 0.00546455, "BUY"],
		  [1684079981, 912600, 0.00323654, "SELL"],
		  [1684079970, 912795.1, 0.00109279, "BUY"],
		  [1684079907, 912617.3, 0.0000111, "SELL"],
		  [1684079891, 912617.28, 0.00708301, "SELL"],
		  [1684079891, 912756.67, 0.0033963, "SELL"],
		  [1684079889, 912795.1, 0.00001092, "BUY"]
		]
	  }
	  `)

	s.unmarshalMockData = []byte(`{
		"error": 0,
		"result":{}
	  }
	  `)
	s.apiErrorMockData = []byte(`{
		"error": 3,
		"result": []
	} `)
}

func (s *tradesServiceTestSuite) TestSetSymbol() {
	tx := s.client.NewGetTradesTx()
	symbol := types.Symbol("THB_BTC")
	tx.Symbol(symbol)
	s.r().Equal(symbol, tx.symbol)
}

func (s *tradesServiceTestSuite) TestSetLimit() {
	tx := s.client.NewGetTradesTx()
	limit := 10
	tx.Limit(limit)
	s.r().Equal(limit, tx.limit)
}

func (s *tradesServiceTestSuite) TestGetTrades() {
	s.mockDo(s.mockData, nil)
	mockDataStuct := new(types.TradesResponse)
	err := sonic.Unmarshal(s.mockData, mockDataStuct)
	s.r().NoError(err)

	tx, err := s.client.NewGetTradesTx().Symbol("THB_BTC").Limit(10).Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, tx)
}

func (s *tradesServiceTestSuite) TestGetTradesValidateSymbol() {
	tx := s.client.NewGetTradesTx()
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("")
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
}

func (s *tradesServiceTestSuite) TestGetTradesValidateLimit() {
	tx := s.client.NewGetTradesTx()
	tx.Symbol(types.Symbol("THB_BTC"))
	limit := -2
	tx.Limit(limit)
	s.r().ErrorIs(tx.validate(), types.ErrLimitMustBePositive)

}

func (s *tradesServiceTestSuite) TestGetTradesDoValidate() {

	_, err := s.client.NewGetTradesTx().Do()
	s.r().Error(err)
	s.r().ErrorIs(err, types.ErrSymbolMandatory)
}

func (s *tradesServiceTestSuite) TestGetTradesAPIError() {
	s.mockDo(s.apiErrorMockData, nil)
	mockDataStuct := new(types.TradesResponse)
	err := sonic.Unmarshal(s.apiErrorMockData, mockDataStuct)
	s.r().NoError(err)

	_, err = s.client.NewGetTradesTx().Symbol("THB_BTC").Limit(10).Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")
}

func (s *tradesServiceTestSuite) TestGetTradesHTTPError() {
	s.mockDo(nil, fmt.Errorf("dummy error"))
	_, err := s.client.NewGetTradesTx().Symbol("THB_BTC").Limit(10).Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "dummy error")
}

func (s *tradesServiceTestSuite) TestGetTradesUnmarshalError() {
	s.mockDo(s.unmarshalMockData, nil)
	mockDataStuct := new(types.TradesResponse)
	err := sonic.Unmarshal(s.unmarshalMockData, mockDataStuct)
	s.r().Error(err)
	s.r().EqualError(err, "json: cannot unmarshal object into Go struct field TradesResponse.result of type [][]interface {}")

	_, err = s.client.NewGetTradesTx().Symbol("THB_BTC").Limit(10).Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "json: cannot unmarshal object into Go struct field TradesResponse.result of type [][]interface {}")

}
