package bitkub

import (
	"fmt"

	"github.com/Maszz/go-bitkub-sdk/types"
	"github.com/valyala/fasthttp"
)

type params map[string]interface{}
type secType int

const (
	secTypeNone   secType = iota
	secTypeSigned         // if the 'timestamp' parameter is required
)

type request struct {
	method   string
	endpoint types.EndPointType
	body     []byte
	query    fasthttp.Args
	headers  *fasthttp.RequestHeader
	fullUrl  string
	signed   secType
}

func NewRequest() *request {
	return &request{}
}

func (r *request) setParam(key string, value interface{}) *request {
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setFormParams set params with key/values to request form body
