package bitkub

import (
	"fmt"
	"testing"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/suite"
)

type BidAskServiceTestSuite struct {
	baseTestSuite
	mockData                  []byte
	unmarshalError            []byte
	apiErrorMockData          []byte
	placeBidMockData          []byte
	placeBidUnmarshalMockData []byte
	placeAskMockData          []byte
	placeAskUnmarshalMockData []byte
}

func TestBidAskTestSuite(t *testing.T) {
	suite.Run(t, new(BidAskServiceTestSuite))
}

func (s *BidAskServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.mockData = []byte(`{
	"error": 0,
	"result": [
	  ["208248209", 1684344197, 2795948.56, 915284.12, 3.05473296],
	  ["208248215", 1684344199, 1362858.04, 915284.12, 1.48899999],
	  ["208248066", 1684344135, 1999.99, 915284.09, 0.00218511],
	  ["208248245", 1684344214, 14164.96, 915153.05, 0.01547825],
	  ["208248246", 1684344214, 402372.52, 914483, 0.44],
	  ["208248243", 1684344214, 20703.05, 914321.16, 0.02264309],
	  ["208248244", 1684344214, 34357.62, 914321.16, 0.0375772],
	  ["208248242", 1684344214, 981066.56, 914321.13, 1.07299999],
	  ["208247311", 1684343872, 128642.22, 914321.11, 0.14069699],
	  ["208248207", 1684344196, 4537.57, 914098.62, 0.00496399]
	]
  }
  `)
	s.unmarshalError = []byte(`{
	"error": 0,
	"result": {}
  }
  `)
	s.apiErrorMockData = []byte(`{
	"error": 3,
	"result": []
  }`)
	s.placeBidMockData = []byte(`{
		"error": 0,
		"result": {
		  "id": 1,
		  "hash": "fwQ6dnQWQPs4cbatF5Am2xCDP1J",
		  "amt": 1000,
		  "rat": 15000,
		  "fee": 2.5,
		  "cre": 2.5,
		  "rec": 0.06666666,
		  "ts": 1533834547,
		  "ci": "input_client_id"
		}
	  }`)
	s.placeBidUnmarshalMockData = []byte(`{
	"error": 0,
	"result": {
	  "id": 1,
	  "hash": "fwQ6dnQWQPs4cbatF5Am2xCDP1J",
	  "amt": "1000",
	  "rat": 15000,
	  "fee": 2.5,
	  "cre": 2.5,
	  "rec": 0.06666666,
	  "ts": 1533834547,
	  "ci": "input_client_id"
	}
  }
  `)

	s.placeAskMockData = []byte(`{
		"error": 0,
		"result": {
		  "id": 1,
		  "hash": "fwQ6dnQWQPs4cbatFGc9LPnpqyu",
		  "typ": "limit",
		  "amt":  1.00000000,
		  "rat": 15000,
		  "fee": 37.5,
		  "cre": 37.5,
		  "rec": 15000,
		  "ts": 1533834844,
		  "ci": "input_client_id"
		}
	  }`)

	s.placeAskUnmarshalMockData = []byte(`{
		"error": 0,
		"result": {
		  "id": 1,
		  "hash": "fwQ6dnQWQPs4cbatFGc9LPnpqyu",
		  "typ": "limit",
		  "amt":  "1.00000000",
		  "rat": 15000,
		  "fee": 37.5,
		  "cre": 37.5,
		  "rec": 15000,
		  "ts": 1533834844,
		  "ci": "input_client_id"
		}
	}`)
}

func (s *BidAskServiceTestSuite) TestGetBids() {
	s.mockDo(s.mockData, nil)
	mockDataStuct := new(types.BidsAsksResponse)
	err := sonic.Unmarshal(s.mockData, mockDataStuct)
	s.r().NoError(err)

	bids, err := s.client.NewGetBidsTx().Symbol("THB_BTC").Limit(5).Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, bids)

}

func (s *BidAskServiceTestSuite) TestGetBidsHTTPError() {
	s.mockDo(nil, fmt.Errorf("http error"))
	bids, err := s.client.NewGetBidsTx().Symbol("THB_BTC").Limit(5).Do()
	defer s.assertDo()

	s.r().Nil(bids)
	s.r().Error(err)
	s.r().EqualError(err, "http error")
}

func (s *BidAskServiceTestSuite) TestGetBidsUnmarshalError() {
	s.mockDo(s.unmarshalError, nil)
	bids, err := s.client.NewGetBidsTx().Symbol("THB_BTC").Limit(5).Do()
	defer s.assertDo()

	s.r().Nil(bids)
	s.r().Error(err)
	s.r().EqualError(err, "json: cannot unmarshal object into Go struct field BidsAsksResponse.result of type [][]interface {}")
}

func (s *BidAskServiceTestSuite) TestGetBidsAPIError() {
	s.mockDo(s.apiErrorMockData, nil)

	bids, err := s.client.NewGetBidsTx().Symbol("THB_BTC").Limit(5).Do()
	defer s.assertDo()

	s.r().Nil(bids)
	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")
}

func (s *BidAskServiceTestSuite) TestGetBidsSetSymbol() {
	tx := s.client.NewGetBidsTx()
	symbol := types.Symbol("THB_BTC")
	tx.Symbol(symbol)
	s.r().Equal(symbol, tx.symbol)
}

func (s *BidAskServiceTestSuite) TestGetBidsSetLimit() {
	tx := s.client.NewGetBidsTx()
	limit := 10
	tx.Limit(limit)
	s.r().Equal(limit, tx.limit)
}

func (s *BidAskServiceTestSuite) TestGetBidsValidateSymbol() {
	tx := s.client.NewGetBidsTx()
	tx.Limit(10)
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("")
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("THB_BTC")
	s.r().NoError(tx.validate())
}

func (s *BidAskServiceTestSuite) TestGetBidsValidateLimit() {
	tx := s.client.NewGetBidsTx()
	tx.Symbol(types.Symbol("THB_BTC"))
	s.r().ErrorIs(tx.validate(), types.ErrLimitMandatory)
	tx.Limit(-2)
	s.r().ErrorIs(tx.validate(), types.ErrLimitMustBePositive)
	tx.Limit(5)
	s.r().NoError(tx.validate())
}

func (s *BidAskServiceTestSuite) TestGetBidsDoValidate() {
	_, err := s.client.NewGetBidsTx().Do()
	s.r().Error(err)
	s.r().ErrorIs(err, types.ErrSymbolMandatory)
}

func (s *BidAskServiceTestSuite) TestGetAsks() {
	s.mockDo(s.mockData, nil)
	mockDataStuct := new(types.BidsAsksResponse)
	err := sonic.Unmarshal(s.mockData, mockDataStuct)
	s.r().NoError(err)

	bids, err := s.client.NewGetAsksTx().Symbol("THB_BTC").Limit(5).Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, bids)
}
func (s *BidAskServiceTestSuite) TestGetAsksSHTTPError() {
	s.mockDo(nil, fmt.Errorf("http error"))
	bids, err := s.client.NewGetAsksTx().Symbol("THB_BTC").Limit(5).Do()
	defer s.assertDo()

	s.r().Nil(bids)
	s.r().Error(err)
	s.r().EqualError(err, "http error")
}

func (s *BidAskServiceTestSuite) TestGetAsksUnmarshalError() {
	s.mockDo(s.unmarshalError, nil)
	bids, err := s.client.NewGetAsksTx().Symbol("THB_BTC").Limit(5).Do()
	defer s.assertDo()

	s.r().Nil(bids)
	s.r().Error(err)
	s.r().EqualError(err, "json: cannot unmarshal object into Go struct field BidsAsksResponse.result of type [][]interface {}")
}

func (s *BidAskServiceTestSuite) TestGetAsksAPIError() {
	s.mockDo(s.apiErrorMockData, nil)

	bids, err := s.client.NewGetAsksTx().Symbol("THB_BTC").Limit(5).Do()
	defer s.assertDo()

	s.r().Nil(bids)
	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")
}

func (s *BidAskServiceTestSuite) TestGetAsksSetSymbol() {
	tx := s.client.NewGetAsksTx()
	symbol := types.Symbol("THB_BTC")
	tx.Symbol(symbol)
	s.r().Equal(symbol, tx.symbol)
}

func (s *BidAskServiceTestSuite) TestGetAsksSetLimit() {
	tx := s.client.NewGetAsksTx()
	limit := 10
	tx.Limit(limit)
	s.r().Equal(limit, tx.limit)
}

func (s *BidAskServiceTestSuite) TestGetAsksValidateSymbol() {
	tx := s.client.NewGetAsksTx()
	tx.Limit(10)
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("")
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("THB_BTC")
	s.r().NoError(tx.validate())
}
func (s *BidAskServiceTestSuite) TestGetAsksValidateLimit() {
	tx := s.client.NewGetAsksTx()
	tx.Symbol(types.Symbol("THB_BTC"))
	s.r().ErrorIs(tx.validate(), types.ErrLimitMandatory)
	tx.Limit(-2)
	s.r().ErrorIs(tx.validate(), types.ErrLimitMustBePositive)
	tx.Limit(5)
	s.r().NoError(tx.validate())
}

func (s *BidAskServiceTestSuite) TestGetAsksDoValidate() {
	_, err := s.client.NewGetAsksTx().Do()
	s.r().Error(err)
	s.r().ErrorIs(err, types.ErrSymbolMandatory)
}

/*
Place Bid
*/
func (s *BidAskServiceTestSuite) TestPlaceBid() {
	s.mockDo(s.placeBidMockData, nil)
	mockDataStuct := new(types.PlaceBidAskResponse)
	err := sonic.Unmarshal(s.placeBidMockData, mockDataStuct)
	s.r().NoError(err)

	bids, err := s.client.NewPlaceBidTx().Symbol("THB_BTC").Amount(100).OrderType(types.OrderTypeMarket).Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, bids)
}

func (s *BidAskServiceTestSuite) TestPlaceBidHTTPError() {
	s.mockDo(nil, fmt.Errorf("http error"))
	bids, err := s.client.NewPlaceBidTx().Symbol("THB_BTC").Amount(100).OrderType(types.OrderTypeMarket).Do()
	defer s.assertDo()

	s.r().Nil(bids)
	s.r().Error(err)
	s.r().EqualError(err, "http error")
}

func (s *BidAskServiceTestSuite) TestPlaceBidUnmarshalError() {
	s.mockDo(s.placeBidUnmarshalMockData, nil)
	bids, err := s.client.NewPlaceBidTx().Symbol("THB_BTC").Amount(100).OrderType(types.OrderTypeMarket).Do()
	defer s.assertDo()

	s.r().Nil(bids)
	s.r().Error(err)
	s.r().EqualError(err, "json: cannot unmarshal string into Go struct field .result.amt of type float64")
}

func (s *BidAskServiceTestSuite) TestPlaceBidAPIError() {
	s.mockDo(s.apiErrorMockData, nil)

	bids, err := s.client.NewPlaceBidTx().Symbol("THB_BTC").Amount(100).OrderType(types.OrderTypeMarket).Do()
	defer s.assertDo()

	s.r().Nil(bids)
	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")
}

func (s *BidAskServiceTestSuite) TestPlaceBidSetSymbol() {
	tx := s.client.NewPlaceBidTx()
	symbol := types.Symbol("THB_BTC")
	tx.Symbol(symbol)
	s.r().Equal(symbol, tx.symbol)
}

func (s *BidAskServiceTestSuite) TestPlaceBidSetAmount() {
	tx := s.client.NewPlaceBidTx()
	amount := 10.0
	tx.Amount(amount)
	s.r().Equal(amount, tx.amount)
}

func (s *BidAskServiceTestSuite) TestPlaceBidSetOrderType() {
	tx := s.client.NewPlaceBidTx()
	orderType := types.OrderTypeMarket
	tx.OrderType(orderType)
	s.r().Equal(orderType, tx.orderType)
}

func (s *BidAskServiceTestSuite) TestPlaceBidSetRate() {
	tx := s.client.NewPlaceBidTx()
	rate := 10.0
	tx.Rate(rate)
	s.r().Equal(rate, tx.rate)
}

func (s *BidAskServiceTestSuite) TestPlaceBidSetClientID() {
	tx := s.client.NewPlaceBidTx()
	clientID := "clientID"
	tx.ClientID(clientID)
	s.r().Equal(clientID, tx.clientID)
}

func (s *BidAskServiceTestSuite) TestPlaceBidValidateSymbol() {
	tx := s.client.NewPlaceBidTx()
	tx.Amount(10)
	tx.OrderType(types.OrderTypeMarket)
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("")
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("THB_BTC")
	s.r().NoError(tx.validate())
}

func (s *BidAskServiceTestSuite) TestPlaceBidValidateAmount() {
	tx := s.client.NewPlaceBidTx()
	tx.Symbol(types.Symbol("THB_BTC"))
	tx.OrderType(types.OrderTypeMarket)
	s.r().ErrorIs(tx.validate(), types.ErrAmountMandatory)
	tx.Amount(-2)
	s.r().ErrorIs(tx.validate(), types.ErrAmountMustBePositive)
	tx.Amount(5)
	s.r().NoError(tx.validate())
}

func (s *BidAskServiceTestSuite) TestPlaceBidValidateOrderTypeAndRate() {
	tx := s.client.NewPlaceBidTx()
	tx.Symbol(types.Symbol("THB_BTC"))
	tx.Amount(10)
	s.r().ErrorIs(tx.validate(), types.ErrOrderTypeMandatory)
	tx.OrderType(types.OrderTypeMarket)
	s.r().NoError(tx.validate())
	tx.OrderType(types.OrderTypeLimit)
	s.r().ErrorIs(tx.validate(), types.ErrRateMandatory)
	tx.Rate(10)
	s.r().NoError(tx.validate())
}

func (s *BidAskServiceTestSuite) TestPlaceBidDoValidate() {
	_, err := s.client.NewPlaceBidTx().Do()
	s.r().Error(err)
	s.r().ErrorIs(err, types.ErrSymbolMandatory)
}

func (s *BidAskServiceTestSuite) TestPlaceAsk() {
	s.mockDo(s.placeAskMockData, nil)
	mockDataStuct := new(types.PlaceBidAskResponse)
	err := sonic.Unmarshal(s.placeAskMockData, mockDataStuct)
	s.r().NoError(err)

	asks, err := s.client.NewPlaceAskTx().Symbol("THB_BTC").Amount(100).OrderType(types.OrderTypeMarket).Do()
	defer s.assertDo()

	s.r().NoError(err)
	s.r().Equal(mockDataStuct, asks)
}

func (s *BidAskServiceTestSuite) TestPlaceAskHTTPError() {
	s.mockDo(nil, fmt.Errorf("http error"))
	asks, err := s.client.NewPlaceAskTx().Symbol("THB_BTC").Amount(100).OrderType(types.OrderTypeMarket).Do()
	defer s.assertDo()

	s.r().Nil(asks)
	s.r().Error(err)
	s.r().EqualError(err, "http error")
}

func (s *BidAskServiceTestSuite) TestPlaceAskUnmarshalError() {
	s.mockDo(s.placeAskUnmarshalMockData, nil)
	asks, err := s.client.NewPlaceAskTx().Symbol("THB_BTC").Amount(100).OrderType(types.OrderTypeMarket).Do()
	defer s.assertDo()

	s.r().Nil(asks)
	s.r().Error(err)
	s.r().EqualError(err, "json: cannot unmarshal string into Go struct field .result.amt of type float64")
}

func (s *BidAskServiceTestSuite) TestPlaceAskAPIError() {
	s.mockDo(s.apiErrorMockData, nil)

	asks, err := s.client.NewPlaceAskTx().Symbol("THB_BTC").Amount(100).OrderType(types.OrderTypeMarket).Do()
	defer s.assertDo()

	s.r().Nil(asks)
	s.r().Error(err)
	s.r().EqualError(err, "error id: 3, error message: Invalid API key")
}

func (s *BidAskServiceTestSuite) TestPlaceAskSetSymbol() {
	tx := s.client.NewPlaceAskTx()
	symbol := types.Symbol("THB_BTC")
	tx.Symbol(symbol)
	s.r().Equal(symbol, tx.symbol)
}

func (s *BidAskServiceTestSuite) TestPlaceAskSetAmount() {
	tx := s.client.NewPlaceAskTx()
	amount := 10.0
	tx.Amount(amount)
	s.r().Equal(amount, tx.amount)
}

func (s *BidAskServiceTestSuite) TestPlaceAskSetOrderType() {
	tx := s.client.NewPlaceAskTx()
	orderType := types.OrderTypeMarket
	tx.OrderType(orderType)
	s.r().Equal(orderType, tx.orderType)
}

func (s *BidAskServiceTestSuite) TestPlaceAskSetRate() {
	tx := s.client.NewPlaceAskTx()
	rate := 10.0
	tx.Rate(rate)
	s.r().Equal(rate, tx.rate)
}

func (s *BidAskServiceTestSuite) TestPlaceAskSetClientID() {
	tx := s.client.NewPlaceAskTx()
	clientID := "clientID"
	tx.ClientID(clientID)
	s.r().Equal(clientID, tx.clientID)
}

func (s *BidAskServiceTestSuite) TestPlaceAskValidateSymbol() {
	tx := s.client.NewPlaceAskTx()
	tx.Amount(10)
	tx.OrderType(types.OrderTypeMarket)
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("")
	s.r().ErrorIs(tx.validate(), types.ErrSymbolMandatory)
	tx.Symbol("THB_BTC")
	s.r().NoError(tx.validate())
}

func (s *BidAskServiceTestSuite) TestPlaceAskValidateAmount() {
	tx := s.client.NewPlaceAskTx()
	tx.Symbol(types.Symbol("THB_BTC"))
	tx.OrderType(types.OrderTypeMarket)
	s.r().ErrorIs(tx.validate(), types.ErrAmountMandatory)
	tx.Amount(-2)
	s.r().ErrorIs(tx.validate(), types.ErrAmountMustBePositive)
	tx.Amount(5)
	s.r().NoError(tx.validate())
}

func (s *BidAskServiceTestSuite) TestPlaceAskValidateOrderTypeAndRate() {
	tx := s.client.NewPlaceAskTx()
	tx.Symbol(types.Symbol("THB_BTC"))
	tx.Amount(10)
	s.r().ErrorIs(tx.validate(), types.ErrOrderTypeMandatory)
	tx.OrderType(types.OrderTypeMarket)
	s.r().NoError(tx.validate())
	tx.OrderType(types.OrderTypeLimit)
	s.r().ErrorIs(tx.validate(), types.ErrRateMandatory)
	tx.Rate(10)
	s.r().NoError(tx.validate())
}

func (s *BidAskServiceTestSuite) TestPlaceAskDoValidate() {
	_, err := s.client.NewPlaceAskTx().Do()
	s.r().Error(err)
	s.r().ErrorIs(err, types.ErrSymbolMandatory)
}
