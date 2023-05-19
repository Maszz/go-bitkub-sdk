package bitkub

import (
	"fmt"
	"testing"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/suite"
)

type statusServiceTestSuite struct {
	baseTestSuite
	mockData          []byte
	unmarshalMockData []byte
}

func TestStatusService(t *testing.T) {
	suite.Run(t, new(statusServiceTestSuite))
}

func (s *statusServiceTestSuite) BeforeTest(suiteName, testName string) {

	s.mockData = []byte(`[
		{ "name": "Non-secure endpoints", "status": "ok", "message": "" },
		{ "name": "Secure endpoints", "status": "ok", "message": "" }
	  ]`)

	s.unmarshalMockData = []byte(`[
		{ "name": "Non-secure endpoints", "status": 0, "message": "" },
		{ "name": "Secure endpoints", "status": 0, "message": "" }
	  ]`)
}

func (s *statusServiceTestSuite) TestGetStatus() {

	s.mockDo(s.mockData, nil)

	mockDataStuct := make(types.ServerStatusArray, 0)

	err := sonic.Unmarshal(s.mockData, &mockDataStuct)
	s.r().NoError(err)

	status, err2 := s.client.NewGetStatusTx().Do()
	defer s.assertDo()

	s.r().NoError(err2)
	s.r().EqualValues(mockDataStuct, *status)

}

func (s *statusServiceTestSuite) TestGetStatusHttpError() {
	s.mockDo(nil, fmt.Errorf("dummy error"))

	_, err := s.client.NewGetStatusTx().Do()
	defer s.assertDo()

	s.r().Error(err)
	s.r().EqualError(err, "dummy error")

}

// func (s *statusServiceTestSuite) TestGetStatusUnmarshalError() {

// 	s.mockDo(s.unmarshalMockData, nil)
// 	data, err := s.client.NewGetStatusTx().Do()
// 	defer s.assertDo()

// 	s.r().Nil(data)
// 	s.r().Error(err)
// 	s.r().EqualError(err, "json: cannot unmarshal number into Go struct field ServerStatus.status of type string")
// }
