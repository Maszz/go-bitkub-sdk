package httpclient

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type HTTPClient struct {
	client *fasthttp.Client
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &fasthttp.Client{},
	}
}

type APIError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

func (c *HTTPClient) DoRequest(url, method string, body []byte, header *fasthttp.RequestHeader) ([]byte, error) {
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

	err := c.client.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)

		return nil, err
	}
	if resp.StatusCode() >= fasthttp.StatusBadRequest {
		apiErr := new(APIError)
		e := json.Unmarshal(resp.Body(), apiErr)
		if e != nil {

			return nil, err
		}

		return nil, err
	}
	bodyBytes := resp.Body()

	return bodyBytes, nil
}

func (c *HTTPClient) Post(url string, body []byte) []byte {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.SetBody(body)

	err := c.client.Do(req, resp)
	if err != nil {
		fmt.Printf("Client post failed: %s\n", err)

		return nil
	}
	bodyBytes := resp.Body()

	return bodyBytes
}

func (c *HTTPClient) Get(url string) []byte {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(url)

	err := c.client.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)

		return nil
	}
	bodyBytes := resp.Body()

	return bodyBytes
}

func (c *HTTPClient) GetWithHeader(url string, header map[string]string) []byte {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(url)
	for k, v := range header {
		req.Header.Set(k, v)
	}

	err := c.client.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)

		return nil
	}

	bodyBytes := resp.Body()
	fmt.Printf("Client get success: %s\n", (bodyBytes))

	return bodyBytes
}

func (c *HTTPClient) PostWithHeader(url string, body []byte, header map[string]string) []byte {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release
	req.SetRequestURI(url)

	req.SetBody(body)
	for k, v := range header {
		req.Header.Set(k, v)
	}

	err := c.client.Do(req, resp)
	if err != nil {
		fmt.Printf("Client post failed: %s\n", err)

		return nil
	}
	bodyBytes := resp.Body()

	return bodyBytes
}
