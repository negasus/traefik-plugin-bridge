package traefik_plugin_bridge

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ClientHTTP struct {
	address string
	client  *http.Client
}

func newClientHTTP(address string, timeout time.Duration) (*ClientHTTP, error) {
	c := &ClientHTTP{
		address: address,
		client: &http.Client{
			Timeout: timeout,
		},
	}

	return c, nil
}

func (c *ClientHTTP) Call(req *Request) (*Response, error) {
	buf, err := req.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("error marshaling, %w", err)
	}
	httpReq, err := http.NewRequest(http.MethodPost, c.address, bytes.NewReader(buf))
	if err != nil {
		return nil, fmt.Errorf("error create the request, %w", err)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error send the request, %w", err)
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("error read the response body, %w", err)
	}

	resp := acquireResponse()
	err = resp.FromJSON(body)
	if err != nil {
		releaseResponse(resp)
		return nil, fmt.Errorf("error unmarshaling, %w", err)
	}

	return resp, nil
}
