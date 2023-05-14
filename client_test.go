package bitkub

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"
)

type baseTestSuite struct {
	suite.Suite
	client    *mockedClient
	apiKey    string
	secretKey string
}
type assertReqFunc func(r *request)

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseTestSuite) SetupTest() {
	s.apiKey = "dummyAPIKey"
	s.secretKey = "dummySecretKey"
	s.client = newMockedClient(s.apiKey, s.secretKey)
}

type mockedClient struct {
	mock.Mock
	*Client
	assertReq assertReqFunc
}

func newMockedClient(apiKey, secretKey string) *mockedClient {
	m := new(mockedClient)
	m.Client = NewClient(apiKey, secretKey)
	return m
}

func (s *baseTestSuite) mockDo(data []byte, err error, statusCode ...int) {
	s.client.Client.do = s.client.do

	s.client.On("do", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(data, err)
}

func (m *mockedClient) do(url, method string, body []byte, header *fasthttp.RequestHeader) ([]byte, error) {

	args := m.Called(url, method, body, header)
	return args.Get(0).([]byte), args.Error(1)
}

func (s *baseTestSuite) assertDo() {
	s.client.AssertCalled(s.T(), "do", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

type clientTestSuite struct {
	baseTestSuite
}

func TestClient(t *testing.T) {
	suite.Run(t, new(clientTestSuite))
}

func (s *clientTestSuite) TestDoError() {

	s.mockDo(nil, fmt.Errorf("dummy error"))
	defer s.assertDo()
	_, err := s.client.callAPI(&request{})

	// s.r().Error(err)
	s.r().Contains(err.Error(), "dummy error")
}
