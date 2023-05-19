package bitkub

import (
	"fmt"
	"testing"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/suite"
)

type tickerServiceTestSuite struct {
	baseTestSuite
	mockData []byte
}

func TestTickerService(t *testing.T) {
	suite.Run(t, new(tickerServiceTestSuite))
}

func (s *tickerServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.mockData = []byte(`{
	  "THB_1INCH": {
		"id": 1,
		"last": 13.79,
		"lowestAsk": 14.06,
		"highestBid": 13.84,
		"percentChange": -0.29,
		"baseVolume": 1260.47601297,
		"quoteVolume": 17477.39,
		"isFrozen": 0,
		"high24hr": 14.33,
		"low24hr": 13.69,
		"change": -0.04,
		"prevClose": 13.79,
		"prevOpen": 13.83
	  },
	  "THB_ABT": {
		"id": 22,
		"last": 3.61,
		"lowestAsk": 3.62,
		"highestBid": 3.61,
		"percentChange": 1.12,
		"baseVolume": 356708.32419128,
		"quoteVolume": 1295019.22,
		"isFrozen": 0,
		"high24hr": 3.82,
		"low24hr": 3.5,
		"change": 0.04,
		"prevClose": 3.61,
		"prevOpen": 3.57
	  }
	}`)

}

func (s *tickerServiceTestSuite) TestGetTicker() {

	s.mockDo(s.mockData, nil)
	mockDataStuctures := new(types.TickerResponse)
	err := sonic.Unmarshal(s.mockData, mockDataStuctures)
	s.r().NoError(err)

	ticker, err := s.client.NewGetTickerTx().Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().EqualValues(mockDataStuctures, ticker)
}

func (s *tickerServiceTestSuite) TestGetTickerAny() {
	s.mockDo(s.mockData, nil)
	mockDataStuctures := make(types.TickerResponseAny, 0)
	err := sonic.Unmarshal(s.mockData, &mockDataStuctures)
	s.r().NoError(err)

	ticker, err := s.client.NewGetTickerTx().DoAny()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().EqualValues(mockDataStuctures, *ticker)
}

func (s *tickerServiceTestSuite) TestGetTickerHTTPError() {
	s.mockDo(nil, fmt.Errorf("HTTP Error"))

	_, err := s.client.NewGetTickerTx().Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().Contains(err.Error(), "HTTP Error")
}

func (s *tickerServiceTestSuite) TestGetTickerAnyHTTPError() {
	s.mockDo(nil, fmt.Errorf("HTTP Error"))

	_, err := s.client.NewGetTickerTx().DoAny()
	defer s.assertDo()

	s.r().Error(err)
	s.r().Contains(err.Error(), "HTTP Error")
}
