package traefik_plugin_bridge

import (
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
	r, err := req.JSONReader()
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest(http.MethodPost, c.address, r)
	if err != nil {
		return nil, err
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	resp := acquireResponse()
	err = resp.UnmarshalJSON(body)
	if err != nil {
		releaseResponse(resp)
		return nil, err
	}

	return resp, nil
}
