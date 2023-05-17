package bitkub

import (
	"fmt"
	"testing"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	baseTestSuite
	getUserLimitMockData               []byte
	getTradingCreditsMockData          []byte
	getUserLimitUnmarshalMockData      []byte
	getTradingCreditsUnmarshalMockData []byte
	apiErrorMockData                   []byte
}

func TestUserServiceTestSuit(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))

}

func (s *UserServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.getUserLimitMockData = []byte(`{
		"error": 0,
		"result": {
		  "limits": {
			"crypto": { "deposit": 2.18140808, "withdraw": 2.18140808 },
			"fiat": { "deposit": 2000000, "withdraw": 2000000 }
		  },
		  "usage": {
			"crypto": {
			  "deposit": 0,
			  "withdraw": 0,
			  "deposit_percentage": 0,
			  "withdraw_percentage": 0,
			  "deposit_thb_equivalent": 0,
			  "withdraw_thb_equivalent": 0
			},
			"fiat": {
			  "deposit": 0,
			  "withdraw": 0,
			  "deposit_percentage": 0,
			  "withdraw_percentage": 0
			}
		  },
		  "rate": 916839
		}
	  }`)
	s.getUserLimitUnmarshalMockData = []byte(`{
		"error": 0,
		"result": {
		  "limits": {
			"crypto": { "deposit": 2.18140808, "withdraw": 2.18140808 },
			"fiat": { "deposit": "2000000", "withdraw": "2000000" }
		  },
		  "usage": {
			"crypto": {
			  "deposit": 0,
			  "withdraw": 0,
			  "deposit_percentage": 0,
			  "withdraw_percentage": 0,
			  "deposit_thb_equivalent": 0,
			  "withdraw_thb_equivalent": 0
			},
			"fiat": {
			  "deposit": 0,
			  "withdraw": 0,
			  "deposit_percentage": 0,
			  "withdraw_percentage": 0
			}
		  },
		  "rate": 916839
		}
	  }`)

	s.apiErrorMockData = []byte(`{
		"error": 3,
		"result": {}
	  }`)

	s.getTradingCreditsMockData = []byte(`{
		"error": 0,
		"result": 0
	  }`)

	s.getTradingCreditsUnmarshalMockData = []byte(`{
		"error": 0,
		"result": "0"
	  }`)

}
func (s *UserServiceTestSuite) TestGetUserLimit() {
	s.mockDo(s.getUserLimitMockData, nil)
	mockDataStuct := new(types.GetUserLimitsResponse)
	err := sonic.Unmarshal(s.getUserLimitMockData, mockDataStuct)
	s.r().NoError(err)

	userLimit, err := s.client.NewGetUserLimitsTx().Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, userLimit)
}

func (s *UserServiceTestSuite) TestGetUserLimitHTTPError() {
	s.mockDo(nil, fmt.Errorf("http error"))
	_, err := s.client.NewGetUserLimitsTx().Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "http error")
}

func (s *UserServiceTestSuite) TestGetUserLimitUnmarshalError() {
	s.mockDo(s.getUserLimitUnmarshalMockData, nil)
	data, err := s.client.NewGetUserLimitsTx().Do()
	defer s.assertDo()

	s.r().Nil(data)
	s.r().Error(err)
	s.r().EqualError(err, "json: cannot unmarshal string into Go struct field .result.limits.fiat.deposit of type float64")
}

func (s *UserServiceTestSuite) TestGetUserLimitAPIError() {
	s.mockDo(s.apiErrorMockData, nil)
	_, err := s.client.NewGetUserLimitsTx().Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")
}

func (s *UserServiceTestSuite) TestGetTradingCredit() {
	s.mockDo(s.getTradingCreditsMockData, nil)
	mockDataStuct := new(types.GetTradingCreditsResponse)
	err := sonic.Unmarshal(s.getTradingCreditsMockData, mockDataStuct)
	s.r().NoError(err)

	tradingCredits, err := s.client.NewGetTradingCreditsTx().Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, tradingCredits)
}

func (s *UserServiceTestSuite) TestGetTradingCreditHTTPError() {
	s.mockDo(nil, fmt.Errorf("http error"))
	_, err := s.client.NewGetTradingCreditsTx().Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "http error")
}

func (s *UserServiceTestSuite) TestGetTradingCreditUnmarshalError() {
	s.mockDo(s.getTradingCreditsUnmarshalMockData, nil)
	data, err := s.client.NewGetTradingCreditsTx().Do()
	defer s.assertDo()

	s.r().Nil(data)
	s.r().Error(err)
	s.r().EqualError(err, "json: cannot unmarshal string into Go struct field GetTradingCreditsResponse.result of type float64")
}

func (s *UserServiceTestSuite) TestGetTradingCreditAPIError() {
	s.mockDo(s.apiErrorMockData, nil)
	_, err := s.client.NewGetTradingCreditsTx().Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")
}
