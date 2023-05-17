package bitkub

import (
	"fmt"
	"testing"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/suite"
)

type symbolsServiceTestSuite struct {
	baseTestSuite
	mockData          []byte
	apiErrorMockData  []byte
	unmarshalMockData []byte
}

func TestSymbolsService(t *testing.T) {
	suite.Run(t, new(symbolsServiceTestSuite))
}

func (s *symbolsServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.mockData = []byte(`{
		"error": 0,
		"result": [
		  { "id": 1, "info": "Thai Baht to Bitcoin", "symbol": "THB_BTC" },
		  { "id": 2, "info": "Thai Baht to Ethereum", "symbol": "THB_ETH" },
		  { "id": 3, "info": "Thai Baht to Wancoin", "symbol": "THB_WAN" }
		]
	  }`)

	s.apiErrorMockData = []byte(`{
		"error": 3,
		"result": {}
	  }`)

	s.unmarshalMockData = []byte(`{
		"error": 0,
		"result": [
		  { "id": "1", "info": "Thai Baht to Bitcoin", "symbol": "THB_BTC" },
		  { "id": "2", "info": "Thai Baht to Ethereum", "symbol": "THB_ETH" },
		  { "id": "3", "info": "Thai Baht to Wancoin", "symbol": "THB_WAN" }
		]
	  }`)

}

func (s *symbolsServiceTestSuite) TestGetSymbols() {

	s.mockDo(s.mockData, nil)

	mockDataStuct := new(types.SymbolsResponse)
	err := sonic.Unmarshal(s.mockData, mockDataStuct)
	s.r().NoError(err)

	symbols, err2 := s.client.NewGetSymbolsTx().Do()
	defer s.assertDo()

	s.r().NoError(err2)
	s.r().EqualValues(mockDataStuct, symbols)

}

func (s *symbolsServiceTestSuite) TestGetSymbolsAPIError() {
	s.mockDo(s.apiErrorMockData, nil)

	_, err2 := s.client.NewGetSymbolsTx().Do()
	defer s.assertDo()

	s.r().Error(err2)
	s.r().Contains(err2.Error(), "Invalid API key")
}

func (s *symbolsServiceTestSuite) TestGetSymbolsUnmarshalError() {

	s.mockDo(s.unmarshalMockData, nil)
	data, err2 := s.client.NewGetSymbolsTx().Do()
	defer s.assertDo()

	s.r().Nil(data)
	s.r().Error(err2)
	s.r().EqualError(err2, `json: cannot unmarshal string into Go struct field .result.id of type int`)
}

func (s *symbolsServiceTestSuite) TestGetSymbolsHttpError() {
	s.mockDo(nil, fmt.Errorf("errFakeResponse"))

	_, err := s.client.NewGetSymbolsTx().Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "errFakeResponse")
}
