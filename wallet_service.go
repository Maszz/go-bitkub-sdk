package bitkub

import (
	"context"
	"encoding/json"

	"github.com/bytedance/sonic"

	"github.com/Maszz/go-bitkub-sdk/types"
	BitkubTs "github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"

	"github.com/valyala/fasthttp"
)

type GetWalletsTx struct {
	c *Client
}

func (s *GetWalletsTx) Do(ctx context.Context) (res BitkubTs.WalletResponse, err error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketWalletEndpoint,
		signed:   secTypeSigned,
	}
	/*
		// do hmac and sign payload + cal payload stuff.
	*/
	payload := BitkubTs.BalancesPayload{
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

	err = sonic.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	// setparmas stuff.

	return res, nil
}

func (s *GetWalletsTx) DoAny(ctx context.Context) (res BitkubTs.WalletResponseAny, err error) {

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketWalletEndpoint,
		signed:   secTypeSigned,
	}
	/*
		// do hmac and sign payload + cal payload stuff.
	*/
	payload := BitkubTs.BalancesPayload{
		Ts: utils.CurrentTimestamp(),
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := json.Marshal(payload)
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

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	// setparmas stuff.

	return res, nil
}
