package bitkub

import (
	"context"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"
	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type GetWsTokenTx struct {
	c *Client
}

func (s *GetWsTokenTx) Do(ctx context.Context) (res *types.GetWsTokenResponse, err error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.MarketWstokenEndpoint,
		signed:   secTypeSigned,
	}
	/*
		// do hmac and sign payload + cal payload stuff.
	*/
	payload := types.GetWsTokenPayload{
		Ts: utils.CurrentTimestamp(),
	}
	payload.Sig = types.Signature(s.c.signPayload(payload))
	byteBody, err := sonic.Marshal(payload)
	if err != nil {
		return nil, err
	}

	r.body = byteBody
	data, err := s.c.callAPI(ctx, r)

	if err != nil {
		return nil, err
	}
	respErr := s.c.catchApiError(data)
	if respErr != nil {
		return nil, respErr
	}
	res = new(types.GetWsTokenResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
