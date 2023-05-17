package bitkub

import (
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/suite"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/types/symbols"
)

type DeptServiceTestSuite struct {
	baseTestSuite
	mockData          []byte
	unmarshalMockData []byte
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(DeptServiceTestSuite))
}

func (s *DeptServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.mockData = []byte(`{
		"asks": [
		  [916500, 0.20672591],
		  [916519.66, 0.00196395]
		],
		"bids": [
		  [916485.55, 0.02499999],
		  [916485.04, 1.79799999]
		]
	  }`) //

	s.unmarshalMockData = []byte(`{
		"asks": [
		  ["916500", "0.20672591"],
		  ["916519.66", "0.00196395"]
		],
		"bids": [
		  ["916485.55", "0.02499999"],
		  ["916485.04", "1.79799999"]
		]
	  }`) //
}

func (s *DeptServiceTestSuite) TestGetMarketDept() {
	s.mockDo(s.mockData, nil)
	mockDataStuct := new(types.MarketDepthResponse)
	err := sonic.Unmarshal(s.mockData, mockDataStuct)
	s.r().NoError(err)

	dept, err := s.client.NewGetMarketDepthTx().Symbol(symbols.THB_BTC).Limit(2).Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, dept)
}

func (s *DeptServiceTestSuite) TestSetSymbol() {
	tx := s.client.NewGetMarketDepthTx()
	symbol := types.Symbol("THB_BTC")
	tx.Symbol(symbol)
	s.r().Equal(symbol, tx.symbol)
}

func (s *DeptServiceTestSuite) TestSetLimit() {
	tx := s.client.NewGetMarketDepthTx()
	limit := 10
	tx.Limit(limit)
	s.r().Equal(limit, tx.limit)
}

func (s *DeptServiceTestSuite) TestValidateSymbol() {
	tx := s.client.NewGetMarketDepthTx()
	tx.Limit(10)
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("")
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("THB_BTC")
	s.r().NoError(tx.validate())
}

func (s *DeptServiceTestSuite) TestValidateLimit() {
	tx := s.client.NewGetMarketDepthTx()
	tx.Symbol(types.Symbol("THB_BTC"))
	s.r().ErrorIs(tx.validate(), types.ErrLimitMandatory)
	tx.Limit(-2)
	s.r().ErrorIs(tx.validate(), types.ErrLimitMustBePositive)
	tx.Limit(5)
	s.r().NoError(tx.validate())
}

func (s *DeptServiceTestSuite) TestDoValidate() {
	_, err := s.client.NewGetMarketDepthTx().Do()
	s.r().Error(err)
	s.r().ErrorIs(err, types.ErrSymbolMandatory)
}

func (s *DeptServiceTestSuite) TestGetMarketDeptHttpError() {
	s.mockDo(nil, fmt.Errorf("http error"))

	_, err := s.client.NewGetMarketDepthTx().Symbol(symbols.THB_BTC).Limit(2).Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "http error")
}

func (s *DeptServiceTestSuite) TestGetMarketDeptUnmarshalError() {
	s.mockDo(s.unmarshalMockData, nil)

	data, err := s.client.NewGetMarketDepthTx().Symbol(symbols.THB_BTC).Limit(2).Do()
	defer s.assertDo()

	s.r().Nil(data)
	s.r().Error(err)
	s.r().
		EqualError(err, "json: cannot unmarshal string into Go struct field MarketDepthResponse.asks of type float64")
}
