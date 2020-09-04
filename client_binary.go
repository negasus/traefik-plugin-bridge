package traefik_plugin_bridge

import (
	"fmt"
	"time"
)

type ClientBINARY struct {
}

func newClientBinary(address string, timeout time.Duration) (*ClientBINARY, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *ClientBINARY) Call(req *Request) (*Response, error) {
	return nil, fmt.Errorf("not implemented")
}
