package http_client

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type HttpClient struct {
	cliient *fasthttp.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		cliient: &fasthttp.Client{},
	}
}

type APIError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

func (c *HttpClient) DoRequest(url string, method string, body []byte, header *fasthttp.RequestHeader) []byte {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release
	header.CopyTo(&req.Header)

	req.SetRequestURI(url)

	req.Header.SetMethod(method)

	if body != nil {
		req.SetBody(body)
	}

	// query.CopyTo(req.URI().QueryArgs())

	err := c.cliient.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
		return nil

	}
	if resp.StatusCode() >= fasthttp.StatusBadRequest {
		apiErr := new(APIError)
		e := json.Unmarshal(resp.Body(), apiErr)
		// fmt.Println("err on request", string(resp.Body()))
		if e != nil {
			fmt.Printf("failed to unmarshal json: %s", e)
			return nil
		}
		// return nil, apiErr

		// fmt.Println("err on request", apiErr)
		return nil

	}

	bodyBytes := resp.Body()
	// fmt.Printf("Client get success: %s\n", *(*string)(unsafe.Pointer(&bodyBytes)))

	return bodyBytes
}

func (c *HttpClient) Post(url string, body []byte) []byte {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.SetBody(body)

	err := c.cliient.Do(req, resp)
	if err != nil {
		fmt.Printf("Client post failed: %s\n", err)
		return nil
	}

	bodyBytes := resp.Body()
	return bodyBytes
}

func (c *HttpClient) Get(url string) []byte {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(url)

	err := c.cliient.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
		return nil
	}

	bodyBytes := resp.Body()
	return bodyBytes
}

func (c *HttpClient) GetWithHeader(url string, header map[string]string) []byte {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(url)
	for k, v := range header {
		req.Header.Set(k, v)
	}

	err := c.cliient.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
		return nil
	}

	bodyBytes := resp.Body()
	fmt.Printf("Client get success: %s\n", (bodyBytes))

	return bodyBytes
}

func (c *HttpClient) PostWithHeader(url string, body []byte, header map[string]string) []byte {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release
	req.SetRequestURI(url)

	req.SetBody(body)
	for k, v := range header {
		req.Header.Set(k, v)
	}

	err := c.cliient.Do(req, resp)
	if err != nil {
		fmt.Printf("Client post failed: %s\n", err)
		return nil
	}

	bodyBytes := resp.Body()
	return bodyBytes
}
