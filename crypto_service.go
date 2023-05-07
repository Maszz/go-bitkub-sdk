package bitkub

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"
	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type GetCryptoAddressesTx struct {
	c           *Client
	page        int
	limit       int
	hasAllQuery bool
}

func (s *GetCryptoAddressesTx) Page(page int) *GetCryptoAddressesTx {
	s.page = page
	return s
}

func (s *GetCryptoAddressesTx) Limit(limit int) *GetCryptoAddressesTx {
	s.limit = limit
	return s
}

func (s *GetCryptoAddressesTx) Do(ctx context.Context) (res types.CryptoAddressesResponse, err error) {
	err = s.validate()
	if err != nil {
		return res, err
	}

	endpoint := s.urlBuilder()

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeSigned,
	}

	payload := types.CryptoAddressesPayload{
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

	if s.hasAllQuery {
		err = s.tranformHasAllQueryResponse(data, &res)
		if err != nil {
			return res, err
		}
		return res, nil
	}

	err = sonic.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil

}

func (s *GetCryptoAddressesTx) validate() (err error) {
	if s.page < 0 {
		return fmt.Errorf("page must be positive number")
	}
	if s.limit < 0 {
		return fmt.Errorf("limit must be positive number")
	}
	return nil
}

func (s *GetCryptoAddressesTx) urlBuilder() string {
	hasPage := s.page != 0
	hasLimit := s.limit != 0
	url := types.CryptoAddressesEndpoint.String() + "?"
	if hasPage {
		url += "p=" + fmt.Sprint(s.page)
	}
	if hasLimit {
		url += "&lmt=" + fmt.Sprint(s.limit)
	}
	if hasPage && hasLimit {
		s.hasAllQuery = true
	}
	if !(hasPage || hasLimit) {
		return types.CryptoAddressesEndpoint.String()
	}

	return url
}

func (s *GetCryptoAddressesTx) tranformHasAllQueryResponse(data []byte, res *types.CryptoAddressesResponse) (err error) {
	resWithAllQuery := types.CryptoAddressesResponseWithAllQuery{}
	err = sonic.Unmarshal(data, &resWithAllQuery)
	if err != nil {
		return err
	}
	page, err := strconv.Atoi(resWithAllQuery.Pagination.Page)
	if err != nil {
		return err
	}
	resWithAllQuery.Pagination.Page = ""
	resWithAllQueryStr, err := sonic.Marshal(resWithAllQuery)
	if err != nil {
		return err
	}
	err = sonic.Unmarshal(resWithAllQueryStr, res)
	if err != nil {
		return err
	}
	res.Pagination.Page = page
	return nil

}

type CryptoWithdrawTx struct {
	c       *Client
	cur     string
	amount  float64
	address string
	memo    string
	network types.BlockChainNetwork
}

func (s *CryptoWithdrawTx) Currency(cur string) *CryptoWithdrawTx {
	s.cur = cur
	return s
}

func (s *CryptoWithdrawTx) Amount(amount float64) *CryptoWithdrawTx {
	s.amount = amount
	return s
}

func (s *CryptoWithdrawTx) Address(address string) *CryptoWithdrawTx {
	s.address = address
	return s
}

func (s *CryptoWithdrawTx) Memo(memo string) *CryptoWithdrawTx {
	s.memo = memo
	return s
}

func (s *CryptoWithdrawTx) Network(network types.BlockChainNetwork) *CryptoWithdrawTx {
	s.network = network
	return s
}

func (s *CryptoWithdrawTx) Do(ctx context.Context) (res types.CryptoWithdrawResponse, err error) {

	err = s.validate()
	if err != nil {
		return res, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.CryptoWithdrawEndpoint,
		signed:   secTypeSigned,
	}

	payload := types.CryptoWithdrawPayload{
		Ts:      utils.CurrentTimestamp(),
		Cur:     s.cur,
		Amount:  s.amount,
		Address: s.address,
		Memo:    s.memo,
		Network: s.network,
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

	return res, nil
}

func (s *CryptoWithdrawTx) validate() (err error) {
	if s.cur == "" {
		return fmt.Errorf("currency is required")
	}
	if s.amount <= 0 {
		return fmt.Errorf("amount must be positive number")
	}
	if s.address == "" {
		return fmt.Errorf("address is required")
	}
	if s.network == "" {
		return fmt.Errorf("network is required")
	}
	return nil
}

type GetCryptoDepositTx struct {
	c           *Client
	page        int
	limit       int
	hasAllQuery bool
}

func (s *GetCryptoDepositTx) Page(page int) *GetCryptoDepositTx {
	s.page = page
	return s
}

func (s *GetCryptoDepositTx) Limit(limit int) *GetCryptoDepositTx {
	s.limit = limit
	return s
}

func (s *GetCryptoDepositTx) Do(ctx context.Context) (res types.CryptoDepositResponse, err error) {
	err = s.validate()
	if err != nil {
		return res, err
	}

	endpoint := s.urlBuilder()

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeSigned,
	}

	payload := types.CryptoDepositPayload{
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

	return res, nil
}

func (s *GetCryptoDepositTx) validate() (err error) {
	if s.page < 0 {
		return fmt.Errorf("page must be positive number")
	}
	if s.limit < 0 {
		return fmt.Errorf("limit must be positive number")
	}
	return nil
}

func (s *GetCryptoDepositTx) urlBuilder() string {
	hasPage := s.page != 0
	hasLimit := s.limit != 0
	url := types.CryptoDepositHistoryEndpoint.String() + "?"
	if hasPage {
		url += "p=" + fmt.Sprint(s.page)
	}
	if hasLimit {
		url += "&lmt=" + fmt.Sprint(s.limit)
	}
	if hasPage && hasLimit {
		s.hasAllQuery = true
	}
	if !(hasPage || hasLimit) {
		return types.CryptoDepositHistoryEndpoint.String()
	}

	return url

}
