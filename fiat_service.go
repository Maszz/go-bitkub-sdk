package bitkub

import (
	"fmt"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/Maszz/go-bitkub-sdk/utils"

	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

type GetFiatAccountsTx struct {
	c     *Client
	page  int
	limit int
}

func (s *GetFiatAccountsTx) Page(page int) *GetFiatAccountsTx {
	s.page = page
	return s
}

func (s *GetFiatAccountsTx) Limit(limit int) *GetFiatAccountsTx {
	s.limit = limit
	return s
}

func (s *GetFiatAccountsTx) Do() (*types.FiatAccountsResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	endpoint := s.urlBuilder()

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeSigned,
	}

	payload := types.FiatAccountPayload{
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
	res := new(types.FiatAccountsResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetFiatAccountsTx) validate() error {
	if s.page < 0 {
		return types.ErrPageMustBePositive
	}
	if s.limit < 0 {
		return types.ErrLimitMustBePositive
	}
	return nil
}

func (s *GetFiatAccountsTx) urlBuilder() string {
	hasPage := s.page != 0
	hasLimit := s.limit != 0
	url := types.FiatAccountsEndpoint.String() + "?"
	if hasPage {
		url += "p=" + fmt.Sprint(s.page)
	}
	if hasLimit {
		url += "&lmt=" + fmt.Sprint(s.limit)
	}
	if !(hasPage || hasLimit) {
		return types.FiatAccountsEndpoint.String()
	}

	return url
}

type FiatWithdrawTx struct {
	c         *Client
	fiatAccID string
	amount    float64
}

func (s *FiatWithdrawTx) ID(id string) *FiatWithdrawTx {
	s.fiatAccID = id
	return s
}

func (s *FiatWithdrawTx) Amount(amount float64) *FiatWithdrawTx {
	s.amount = amount
	return s
}

func (s *FiatWithdrawTx) Do() (*types.FiatWithdrawResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.FiatWithdrawEndpoint,
		signed:   secTypeSigned,
	}

	payload := types.FiatWithdrawPayload{
		TS:        utils.CurrentTimestamp(),
		FiatAccID: s.fiatAccID,
		Amount:    s.amount,
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
	res := new(types.FiatWithdrawResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FiatWithdrawTx) validate() error {
	if s.fiatAccID == "" {
		return types.ErrFiatAccIDMandatory
	}
	if s.amount <= 0 {
		return types.ErrAmountMustBePositive
	}
	return nil
}

type GetFiatDepositsTx struct {
	c     *Client
	page  int
	limit int
}

func (s *GetFiatDepositsTx) Page(page int) *GetFiatDepositsTx {
	s.page = page
	return s
}

func (s *GetFiatDepositsTx) Limit(limit int) *GetFiatDepositsTx {
	s.limit = limit
	return s
}

func (s *GetFiatDepositsTx) Do() (*types.GetFiatDepositsResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	endpoint := s.urlBuilder()

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeSigned,
	}

	payload := types.GetFiatDepositsPayload{
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
	res := new(types.GetFiatDepositsResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetFiatDepositsTx) validate() error {
	if s.page < 0 {
		return types.ErrPageMustBePositive
	}
	if s.limit < 0 {
		return types.ErrLimitMustBePositive
	}
	return nil
}

func (s *GetFiatDepositsTx) urlBuilder() string {
	hasPage := s.page != 0
	hasLimit := s.limit != 0
	url := types.FiatDepositHistoryEndpoint.String() + "?"
	if hasPage {
		url += "p=" + fmt.Sprint(s.page)
	}
	if hasLimit {
		url += "&lmt=" + fmt.Sprint(s.limit)
	}
	if !(hasPage || hasLimit) {
		return types.FiatDepositHistoryEndpoint.String()
	}

	return url
}

type GetFiatWithdrawsTx struct {
	c     *Client
	page  int
	limit int
}

func (s *GetFiatWithdrawsTx) Page(page int) *GetFiatWithdrawsTx {
	s.page = page
	return s
}

func (s *GetFiatWithdrawsTx) Limit(limit int) *GetFiatWithdrawsTx {
	s.limit = limit
	return s
}

func (s *GetFiatWithdrawsTx) Do() (*types.GetFiatWithdrawsResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	endpoint := s.urlBuilder()

	r := &request{
		method:   fasthttp.MethodPost,
		endpoint: types.NewEndPoint(endpoint),
		signed:   secTypeSigned,
	}

	payload := types.GetFiatWithdrawsPayload{
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
	res := new(types.GetFiatWithdrawsResponse)
	err = sonic.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *GetFiatWithdrawsTx) validate() error {
	if s.page < 0 {
		return types.ErrPageMustBePositive
	}
	if s.limit < 0 {
		return types.ErrLimitMustBePositive
	}
	return nil
}

func (s *GetFiatWithdrawsTx) urlBuilder() string {
	hasPage := s.page != 0
	hasLimit := s.limit != 0
	url := types.FiatWithdrawHistoryEndpoint.String() + "?"
	if hasPage {
		url += "p=" + fmt.Sprint(s.page)
	}
	if hasLimit {
		url += "&lmt=" + fmt.Sprint(s.limit)
	}
	if !(hasPage || hasLimit) {
		return types.FiatWithdrawHistoryEndpoint.String()
	}

	return url
}
