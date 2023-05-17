package bitkub

import (
	"fmt"
	"testing"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/suite"
)

type GetBalanceServiceTestSuite struct {
	baseTestSuite
	mockData          []byte
	unmarshalMockData []byte
	apiErrorMockData  []byte
}

func TestGetBalanceServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GetBalanceServiceTestSuite))
}

func (s *GetBalanceServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.mockData = []byte(`{
		"error": 0,
		"result": {
		  "THB": { "available": 0, "reserved": 0 },
		  "BTC": { "available": 0, "reserved": 0 },
		  "ETH": { "available": 0, "reserved": 0 }
		}
	  }
	  `)
	s.unmarshalMockData = []byte(`{
		"error": 0,
		"result": {
		  "THB": { "available": "0", "reserved": "0" },
		  "BTC": { "available": "0", "reserved": "0" },
		  "ETH": { "available": "0", "reserved": "0" }
		}
	  }
	  `)
	s.apiErrorMockData = []byte(`{
		"error": 3,
		"result": {}
	  }`)
}

func (s *GetBalanceServiceTestSuite) TestGetBalance() {
	s.mockDo(s.mockData, nil)
	mockDataStuct := new(types.BalancesResponse)
	err := sonic.Unmarshal(s.mockData, mockDataStuct)
	s.r().NoError(err)

	balances, err := s.client.NewGetBalancesTx().Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, balances)
}

func (s *GetBalanceServiceTestSuite) TestGetBalanceAny() {
	s.mockDo(s.mockData, nil)
	mockDataStuct := new(types.BalancesResponseAny)
	err := sonic.Unmarshal(s.mockData, mockDataStuct)
	s.r().NoError(err)

	balances, err := s.client.NewGetBalancesTx().DoAny()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, balances)
}

func (s *GetBalanceServiceTestSuite) TestGetBalanceHTTPError() {
	s.mockDo(nil, fmt.Errorf("http error"))
	_, err := s.client.NewGetBalancesTx().Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "http error")
}

func (s *GetBalanceServiceTestSuite) TestGetBalanceAnyHTTPError() {
	s.mockDo(nil, fmt.Errorf("http error"))
	_, err := s.client.NewGetBalancesTx().DoAny()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "http error")
}

func (s *GetBalanceServiceTestSuite) TestGetBalanceAPIError() {
	s.mockDo(s.apiErrorMockData, nil)

	data, err := s.client.NewGetBalancesTx().Do()
	defer s.assertDo()

	s.r().Nil(data)
	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")
}
func (s *GetBalanceServiceTestSuite) TestGetBalanceAnyAPIError() {
	s.mockDo(s.apiErrorMockData, nil)

	data, err := s.client.NewGetBalancesTx().DoAny()
	defer s.assertDo()

	s.r().Nil(data)
	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")
}

func (s *GetBalanceServiceTestSuite) TestGetBalanceUnMarshalError() {
	s.mockDo(s.unmarshalMockData, nil)

	data, err := s.client.NewGetBalancesTx().Do()
	defer s.assertDo()

	s.r().Nil(data)
	s.r().Error(err)
	s.r().EqualError(err, "json: cannot unmarshal string into Go struct field BalancesProps.result.THB.available of type float64")
}
