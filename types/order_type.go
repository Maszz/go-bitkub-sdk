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
		Rate     int     `json:"rate"`
		Fee      float64 `json:"fee"`
		Credit   float64 `json:"credit"`
		Amount   float64 `json:"amount"`
		Receive  int     `json:"receive                                                                                                                                                                                                "`
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
		TxnID           string  `json:"txn_id"`
		OrderID         string  `json:"order_id"`
		Hash            string  `json:"hash"`
		ParentOrderID   string  `json:"parent_order_id"`
		ParentOrderHash string  `json:"parent_order_hash"`
		SuperOrderID    string  `json:"super_order_id"`
		SuperOrderHash  string  `json:"super_order_hash"`
		TakenByMe       bool    `json:"taken_by_me"`
		IsMaker         bool    `json:"is_maker"`
		Side            string  `json:"side"`
		Type            string  `json:"type"`
		Rate            string  `json:"rate"`
		Fee             string  `json:"fee"`
		Credit          string  `json:"credit"`
		Amount          float64 `json:"amount"`
		Ts              int     `json:"ts"`
		Date            string  `json:"date"`
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
	OrderHash OrderHash `json:"hash"`
}

type GetOrdersInfoResponse struct {
	Error  int `json:"error"`
	Result struct {
		ID            string  `json:"id"`
		First         string  `json:"first"`
		Parent        string  `json:"parent"`
		Last          string  `json:"last"`
		Amount        int     `json:"amount"`
		Rate          int     `json:"rate"`
		Fee           int     `json:"fee"`
		Credit        int     `json:"credit"`
		Filled        float64 `json:"filled"`
		Total         int     `json:"total"`
		Status        string  `json:"status"`
		PartialFilled bool    `json:"partial_filled"`
		Remaining     int     `json:"remaining"`
		History       []struct {
			Amount    float64 `json:"amount"`
			Credit    float64 `json:"credit"`
			Fee       float64 `json:"fee"`
			ID        string  `json:"id"`
			Rate      int     `json:"rate"`
			Timestamp int     `json:"timestamp"`
		} `json:"history"`
	} `json:"result"`
}
