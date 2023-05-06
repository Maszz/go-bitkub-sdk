package bitkub

import (
	"context"

	"github.com/bytedance/sonic"

	"github.com/valyala/fasthttp"
)

type TestService struct {
	c      *Client
	symbol string
	limit  *int
	fromID *int64
}

type ServerStatus struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ServerStatusArray []ServerStatus

func (s *TestService) Do(ctx context.Context) (res ServerStatusArray, err error) {
	r := &request{
		method:   fasthttp.MethodGet,
		endpoint: "/api/status",
		signed:   secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return make(ServerStatusArray, 0), err
	}
	resp := make(ServerStatusArray, 0)
	err = sonic.Unmarshal(data, &resp)
	if err != nil {
		return make(ServerStatusArray, 0), err
	}

	// setparmas stuff.

	return resp, nil
}
