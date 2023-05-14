package bitkub

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type serverServiceTestSuite struct {
	baseTestSuite
	mockData []byte
}

func TestServerService(t *testing.T) {
	suite.Run(t, new(serverServiceTestSuite))
}

func (s *serverServiceTestSuite) BeforeTest(suiteName, testName string) {
	s.mockData = []byte(`1499827319559`)
}

func (s *serverServiceTestSuite) TestGetServertime() {

	s.mockDo(s.mockData, nil)
	defer s.assertDo()
	serverTime, err := s.client.NewGetServerTimeTx().Do()
	s.r().NoError(err)
	mockInt, _ := strconv.Atoi(string(s.mockData))
	s.r().EqualValues(mockInt, *serverTime)
}

func (s *serverServiceTestSuite) TestGetServertimeHttpError() {

	s.mockDo(nil, fmt.Errorf("dummy error"))
	defer s.assertDo()
	_, err := s.client.NewGetServerTimeTx().Do()
	s.r().Error(err)
	s.r().EqualError(err, "dummy error")
}

func (s *serverServiceTestSuite) TestGetServertimeEmptyResponse() {
	data := []byte(``)

	s.mockDo(data, nil)
	defer s.assertDo()
	serverTime, err := s.client.NewGetServerTimeTx().Do()
	// fmt.Print(err)
	s.r().NoError(err)
	s.r().EqualValues(0, *serverTime)
}
