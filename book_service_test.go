package bitkub

import (
	"fmt"
	"testing"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/types/symbols"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/suite"
)

type BookServiceTestSuite struct {
	baseTestSuite
	mockData          []byte
	unmarshalMockData []byte
	apiErrorMockData  []byte
}

func TestBookServiceTestSuit(t *testing.T) {
	suite.Run(t, new(BookServiceTestSuite))
}

func (s *BookServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.mockData = []byte(`{
		"error": 0,
		"result": {
		  "asks": [
			["185637710", 1684365394, 5917.5, 913000, 0.00648139],
			["185637778", 1684340421, 91.3, 913000, 0.0001]
		  ],
		  "bids": [
			["208234158", 1684340426, 137127.99, 912420, 0.15029043],
			["208234184", 1684340432, 933405.64, 912420, 1.02299998]
		  ]
		}
	  }`)
	s.unmarshalMockData = []byte(`{
		"error": 0,
		"result": []
	  }`)
	s.apiErrorMockData = []byte(`{
		"error": 3,
		"result": {}
	  }`)
}

func (s *BookServiceTestSuite) TestGetOpenBook() {
	s.mockDo(s.mockData, nil)
	mockDataStruct := new(types.OpenBooksResponse)
	err := sonic.Unmarshal(s.mockData, mockDataStruct)
	s.r().NoError(err)

	book, err := s.client.NewGetBooksTx().Symbol(symbols.THB_BTC).Limit(2).Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStruct, book)
}

func (s *BookServiceTestSuite) TestSetSymbol() {
	tx := s.client.NewGetBooksTx()
	symbol := types.Symbol("THB_BTC")
	tx.Symbol(symbol)
	s.r().Equal(symbol, tx.symbol)
}

func (s *BookServiceTestSuite) TestSetLimit() {
	tx := s.client.NewGetBooksTx()
	limit := 10
	tx.Limit(limit)
	s.r().Equal(limit, tx.limit)
}

func (s *BookServiceTestSuite) TestValidateSymbol() {
	tx := s.client.NewGetBooksTx()
	tx.Limit(10)
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("")
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("THB_BTC")
	s.r().NoError(tx.validate())
}

func (s *BookServiceTestSuite) TestValidateLimit() {
	tx := s.client.NewGetBooksTx()
	tx.Symbol(types.Symbol("THB_BTC"))
	s.r().ErrorIs(tx.validate(), types.ErrLimitMandatory)
	tx.Limit(-2)
	s.r().ErrorIs(tx.validate(), types.ErrLimitMustBePositive)
	tx.Limit(5)
	s.r().NoError(tx.validate())
}

func (s *BookServiceTestSuite) TestDoValidate() {
	_, err := s.client.NewGetBooksTx().Do()
	s.r().Error(err)
	s.r().ErrorIs(err, types.ErrSymbolMandatory)
}

func (s *BookServiceTestSuite) TestGetMarketDeptHttpError() {
	s.mockDo(nil, fmt.Errorf("http error"))

	_, err := s.client.NewGetBooksTx().Symbol(symbols.THB_BTC).Limit(2).Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "http error")
}

func (s *BookServiceTestSuite) TestGetOpenBookUnmarshalError() {
	s.mockDo(s.unmarshalMockData, nil)

	data, err := s.client.NewGetBooksTx().Symbol(symbols.THB_BTC).Limit(2).Do()
	defer s.assertDo()

	s.r().Nil(data)
	s.r().Error(err)
	s.r().EqualError(err, "json: cannot unmarshal array into Go struct field OpenBooksResponse.result of type struct { Asks [][]interface {} \"json:\\\"asks\\\"\"; Bids [][]interface {} \"json:\\\"bids\\\"\" }")
}

func (s *BookServiceTestSuite) TestGetOpenBookAPIError() {
	s.mockDo(s.apiErrorMockData, nil)

	data, err := s.client.NewGetBooksTx().Symbol(symbols.THB_BTC).Limit(2).Do()
	defer s.assertDo()

	s.r().Nil(data)
	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")

}
