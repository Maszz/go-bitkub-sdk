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
	mockData []byte
}

func TestStatusService(t *testing.T) {
	suite.Run(t, new(statusServiceTestSuite))
}

func (s *statusServiceTestSuite) BeforeTest(suiteName, testName string) {

	s.mockData = []byte(`[
		{ "name": "Non-secure endpoints", "status": "ok", "message": "" },
		{ "name": "Secure endpoints", "status": "ok", "message": "" }
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
