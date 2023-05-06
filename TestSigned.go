package bitkub

import (
	"context"

	"github.com/bytedance/sonic"

	BitkubTs "github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"

	"github.com/valyala/fasthttp"
)

type TestSignedService struct {
	c      *Client
	symbol string
	limit  *int
	fromID *int64
}

type TestSignedPayload struct {
	Ts  int64  `json:"ts"`
	Sig string `json:"sig,omitempty"`
}

func (s *TestSignedService) Do(ctx context.Context) (res BitkubTs.BalancesResponse, err error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: "/api/market/balances",
		signed:   secTypeSigned,
	}
	/*
		// do hmac and sign payload + cal payload stuff.
	*/
	payload := TestSignedPayload{
		Ts: utils.CurrentTimestamp(),
	}
	payload.Sig = s.c.signPayload(payload)
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return BitkubTs.BalancesResponse{}, err
	}

	r.body = byteBody
	data, err := s.c.callAPI(ctx, r)

	if err != nil {
		return BitkubTs.BalancesResponse{}, err
	}
	resp := BitkubTs.BalancesResponse{}
	err = sonic.Unmarshal(data, &resp)
	if err != nil {
		return BitkubTs.BalancesResponse{}, err
	}

	// setparmas stuff.

	return resp, nil
}
