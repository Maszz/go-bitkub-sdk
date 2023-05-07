package bitkub

import (
	"context"

	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/utils"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/valyala/fasthttp"
)

type GetBalancesTx struct {
	c *Client
}

func (s *GetBalancesTx) Do(ctx context.Context) (res types.BalancesResponse, err error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketBalancesEndpoint,
		signed:   secTypeSigned,
	}
	/*
		// do hmac and sign payload + cal payload stuff.
	*/
	payload := types.BalancesPayload{
		Ts: utils.CurrentTimestamp(),
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return res, err
	}

	r.body = byteBody
	data, err := s.c.callAPI(ctx, r)

	if err != nil {
		return res, err
	}
	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return res, respErr
	}
	err = sonic.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	// setparmas stuff.

	return res, nil
}

func (s *GetBalancesTx) DoAny(ctx context.Context) (res types.BalancesResponseAny, err error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketBalancesEndpoint,
		signed:   secTypeSigned,
	}
	/*
		// do hmac and sign payload + cal payload stuff.
	*/
	payload := types.BalancesPayload{
		Ts: utils.CurrentTimestamp(),
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return res, err
	}

	r.body = byteBody
	data, err := s.c.callAPI(ctx, r)

	if err != nil {
		return types.BalancesResponseAny{}, err
	}
	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return res, respErr
	}
	// resp := BitkubTs.BalancesResponseAny{}
	err = sonic.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	// setparmas stuff.

	return res, nil
}
