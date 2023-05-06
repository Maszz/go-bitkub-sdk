package bitkub

import (
	"context"

	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/utils"

	BitkubTs "github.com/Maszz/go-bitkub-sdk/types"
	"github.com/valyala/fasthttp"
)

type GetBalancesService struct {
	c *Client
}

func (s *GetBalancesService) Do(ctx context.Context) (res BitkubTs.BalancesResponse, err error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: market_balances_endpoint,
		signed:   secTypeSigned,
	}
	/*
		// do hmac and sign payload + cal payload stuff.
	*/
	payload := BitkubTs.BalancesPayload{
		Ts: utils.CurrentTimestamp(),
	}
	payload.Sig = s.c.signPayload(payload)
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return res, err
	}

	r.body = byteBody
	data, err := s.c.callAPI(ctx, r)

	if err != nil {
		return res, err
	}
	err = sonic.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	// setparmas stuff.

	return res, nil
}

func (s *GetBalancesService) DoAny(ctx context.Context) (res BitkubTs.BalancesResponseAny, err error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: market_balances_endpoint,
		signed:   secTypeSigned,
	}
	/*
		// do hmac and sign payload + cal payload stuff.
	*/
	payload := BitkubTs.BalancesPayload{
		Ts: utils.CurrentTimestamp(),
	}
	payload.Sig = s.c.signPayload(payload)
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return res, err
	}

	r.body = byteBody
	data, err := s.c.callAPI(ctx, r)

	if err != nil {
		return BitkubTs.BalancesResponseAny{}, err
	}
	// resp := BitkubTs.BalancesResponseAny{}
	err = sonic.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	// setparmas stuff.

	return res, nil
}
