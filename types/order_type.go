package types

type CancelOrderPayload struct {
	Ts        Timestamp `json:"ts,omitempty"`
	Sig       Signature `json:"sig,omitempty"`
	Symbol    Symbol    `json:"sym"`
	OrderID   OrderId   `json:"id"` // OrderID is string in cancel order payload
	OrderSide OrderSide `json:"sd"`
	OrderHash OrderHash `json:"hash"`
}

type CancelOrderResponse struct {
	Error int `json:"error"`
}

type GetOpenOrdersPayload struct {
	Ts     Timestamp `json:"ts,omitempty"`
	Sig    Signature `json:"sig,omitempty"`
	Symbol Symbol    `json:"sym"`
}

type GetOpenOrdersResponse struct {
	Error  int `json:"error"`
	Result []struct {
		ID       string  `json:"id"`
		Hash     string  `json:"hash"`
		Side     string  `json:"side"`
		Type     string  `json:"type"`
		Rate     float64 `json:"rate"`
		Fee      float64 `json:"fee"`
		Credit   float64 `json:"credit"`
		Amount   float64 `json:"amount"`
		Receive  float64 `json:"receive                                                                                                                                                                                                "`
		ParentID string  `json:"parent_id"`
		SuperID  string  `json:"super_id"`
		ClientID string  `json:"client_id"`
		Ts       int     `json:"ts"`
	} `json:"result"`
}

type GetOrderHistoryPayload struct {
	Ts     Timestamp `json:"ts,omitempty"`
	Sig    Signature `json:"sig,omitempty"`
	Symbol Symbol    `json:"sym"`
	Page   int       `json:"page,omitempty"`
	Limit  int       `json:"limit,omitempty"`
	Start  Timestamp `json:"start,omitempty"`
	End    Timestamp `json:"end,omitempty"`
}

type GetOrderHistoryResponse struct {
	Error  int `json:"error"`
	Result []struct {
		TxnID           string `json:"txn_id"`
		OrderID         string `json:"order_id"`
		Hash            string `json:"hash"`
		ParentOrderID   string `json:"parent_order_id"`
		ParentOrderHash string `json:"parent_order_hash"`
		SuperOrderID    string `json:"super_order_id"`
		SuperOrderHash  string `json:"super_order_hash"`
		TakenByMe       bool   `json:"taken_by_me"`
		IsMaker         bool   `json:"is_maker"`
		Side            string `json:"side"`
		Type            string `json:"type"`
		Rate            string `json:"rate"`
		Fee             string `json:"fee"`
		Credit          string `json:"credit"`
		Amount          string `json:"amount"`
		Ts              int    `json:"ts"`
		Date            string `json:"date"`
	} `json:"result"`
	Pagination struct {
		Page int `json:"page,omitempty"`
		Last int `json:"last"`
		Next int `json:"next"`
		Prev int `json:"prev"`
	} `json:"pagination"`
}
type GetOrdersInfoPayload struct {
	Ts        Timestamp `json:"ts,omitempty"`
	Sig       Signature `json:"sig,omitempty"`
	Symbol    Symbol    `json:"sym"`
	OrderID   OrderId   `json:"id"` // OrderId is a string in this case **
	OrderSide OrderSide `json:"sd"`
	OrderHash OrderHash `json:"hash,omitempty"`
}

type GetOrdersInfoResponse struct {
	Error  int `json:"error"`
	Result struct {
		Amount   int     `json:"amount"`
		ClientID string  `json:"client_id"`
		Credit   int     `json:"credit"`
		Fee      float64 `json:"fee"`
		Filled   int     `json:"filled"`
		First    string  `json:"first"`
		History  []struct {
			Amount    int     `json:"amount"`
			Credit    int     `json:"credit"`
			Fee       float64 `json:"fee"`
			Hash      string  `json:"hash"`
			ID        string  `json:"id"`
			Rate      int     `json:"rate"`
			Timestamp int     `json:"timestamp"`
			TxnID     string  `json:"txn_id"`
		} `json:"history"`
		ID            string  `json:"id"`
		Last          string  `json:"last"`
		Parent        string  `json:"parent"`
		PartialFilled bool    `json:"partial_filled"`
		Rate          float64 `json:"rate"`
		Remaining     float64 `json:"remaining"`
		Side          string  `json:"side"`
		Status        string  `json:"status"`
		Total         int     `json:"total"`
	} `json:"result"`
}
