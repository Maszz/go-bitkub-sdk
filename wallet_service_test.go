package bitkub

import (
	"fmt"
	"testing"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/suite"
)

type WalletServiceTestSuite struct {
	baseTestSuite
	mockData          []byte
	unmarshalMockData []byte
	apiErrorMockData  []byte
}

func TestWalletService(t *testing.T) {
	suite.Run(t, new(WalletServiceTestSuite))
}

func (s *WalletServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.mockData = []byte(`{
		"error": 0,
		"result": {
		  "THB": 0,
		  "BTC": 0,
		  "ETH": 0,
		  "WAN": 0,
		  "ADA": 0
		}
	  }`)
	s.unmarshalMockData = []byte(`{
		"error": 0,
		"result": {
		  "THB": "0",
		  "BTC": "0",
		  "ETH": "0",
		  "WAN": "0",
		  "ADA": "0"
		}
	  }`)

	s.apiErrorMockData = []byte(`{
		"error": 3,
		"result": {}
	  }`)
}
func (s *WalletServiceTestSuite) TestGetWallet() {
	s.mockDo(s.mockData, nil)
	mockDataStuct := new(types.WalletResponse)
	err := sonic.Unmarshal(s.mockData, mockDataStuct)
	s.r().NoError(err)

	wallets, err := s.client.NewGetWalletsTx().Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, wallets)
}

func (s *WalletServiceTestSuite) TestGetWalletAny() {
	s.mockDo(s.mockData, nil)
	mockDataStuct := new(types.WalletResponseAny)
	err := sonic.Unmarshal(s.mockData, mockDataStuct)
	s.r().NoError(err)

	wallets, err := s.client.NewGetWalletsTx().DoAny()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, wallets)
}

func (s *WalletServiceTestSuite) TestGetWalletHTTPError() {
	s.mockDo(nil, fmt.Errorf("HTTP Error"))

	_, err := s.client.NewGetWalletsTx().Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "HTTP Error")
}

func (s *WalletServiceTestSuite) TestGetWalletAnyHTTPError() {
	s.mockDo(nil, fmt.Errorf("HTTP Error"))

	_, err := s.client.NewGetWalletsTx().DoAny()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "HTTP Error")
}

// func (s *WalletServiceTestSuite) TestGetWalletUnmarshalError() {
// 	s.mockDo(s.unmarshalMockData, nil)

// 	data, err := s.client.NewGetWalletsTx().Do()
// 	defer s.assertDo()

// 	s.r().Nil(data)
// 	s.r().Error(err)
// 	s.r().EqualError(err, "json: cannot unmarshal string into Go struct field WalletResult.result.THB of type float64")
// }

// func (s *WalletServiceTestSuite) TestGetWalletAnyUnmarshalError() {
// 	s.mockDo(s.unmarshalMockData, nil)

// 	data, err := s.client.NewGetWalletsTx().DoAny()
// 	defer s.assertDo()

// 	s.r().Nil(data)
// 	s.r().Error(err)
// 	s.r().EqualError(err, "json: cannot unmarshal string into Go struct field WalletResponseAny.result of type float64")
// }

func (s *WalletServiceTestSuite) TestGetWalletAPIError() {
	s.mockDo(s.apiErrorMockData, nil)
	_, err := s.client.NewGetWalletsTx().Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")
}

func (s *WalletServiceTestSuite) TestGetWalletAnyAPIError() {
	s.mockDo(s.apiErrorMockData, nil)
	_, err := s.client.NewGetWalletsTx().DoAny()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")
}
