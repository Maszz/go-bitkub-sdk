package bitkub

import (
	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/utils"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/valyala/fasthttp"
)

type GetBalancesTx struct {
	c *Client
}

func (s *GetBalancesTx) Do() (*types.BalancesResponse, error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketBalancesEndpoint,
		signed:   secTypeSigned,
	}
	/*
		// do hmac and sign payload + cal payload stuff.
	*/
	payload := types.BalancesPayload{
		TS: utils.CurrentTimestamp(),
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return nil, err
	}

	r.body = byteBody
	data, err := s.c.callAPI(r)

	if err != nil {
		return nil, err
	}
	respErr := s.c.catchAPIError(data)
	if respErr != nil {
		return nil, respErr
	}
	res := new(types.BalancesResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetBalancesTx) DoAny() (*types.BalancesResponseAny, error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketBalancesEndpoint,
		signed:   secTypeSigned,
	}
	/*
		// do hmac and sign payload + cal payload stuff.
	*/
	payload := types.BalancesPayload{
		TS: utils.CurrentTimestamp(),
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return nil, err
	}

	r.body = byteBody
	data, err := s.c.callAPI(r)

	if err != nil {
		return nil, err
	}
	respErr := s.c.catchAPIError(data)
	if respErr != nil {
		return nil, respErr
	}
	res := new(types.BalancesResponseAny)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
