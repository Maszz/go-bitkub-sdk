package bitkub

import (
	"encoding/json"

	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"

	"github.com/valyala/fasthttp"
)

type GetWalletsTx struct {
	c *Client
}

func (s *GetWalletsTx) Do() (*types.WalletResponse, error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketWalletEndpoint,
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
	res := new(types.WalletResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetWalletsTx) DoAny() (*types.WalletResponseAny, error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketWalletEndpoint,
		signed:   secTypeSigned,
	}

	payload := types.BalancesPayload{
		TS: utils.CurrentTimestamp(),
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := json.Marshal(payload)
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
	res := new(types.WalletResponseAny)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
