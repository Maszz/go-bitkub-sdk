package bitkub

import (
	"context"
	"encoding/json"

	"github.com/bytedance/sonic"

	BitkubTs "github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"

	"github.com/valyala/fasthttp"
)

type WalletService struct {
	c   *Client
	any bool
}

func (s *WalletService) Any(any bool) *WalletService {
	s.any = any
	return s
}

func (s *WalletService) Do(ctx context.Context) (res BitkubTs.WalletResponse, err error) {
	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: market_wallet_endpoint,
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

func (s *WalletService) DoAny(ctx context.Context) (res BitkubTs.WalletResponseAny, err error) {

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: market_wallet_endpoint,
		signed:   secTypeSigned,
	}
	/*
		// do hmac and sign payload + cal payload stuff.
	*/
	payload := BitkubTs.BalancesPayload{
		Ts: utils.CurrentTimestamp(),
	}
	payload.Sig = s.c.signPayload(payload)
	byteBody, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	r.body = byteBody
	data, err := s.c.callAPI(ctx, r)

	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	// setparmas stuff.

	return res, nil
}
