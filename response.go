package traefik_plugin_bridge

import (
	"encoding/json"
	"sync"
)

var (
	responsePool = sync.Pool{
		New: func() interface{} {
			return &Response{}
		},
	}
)

func acquireResponse() *Response {
	return responsePool.Get().(*Response)
}

func releaseResponse(r *Response) {
	r.reset()
	responsePool.Put(r)
}

type ResponseInterruptRequest struct {
	ResponseCode int    `json:"1,omitempty"`
	Body         string `json:"2,omitempty"`
}

type Response struct {
	AddRequestHeaders    map[string]string         `json:"1,omitempty"`
	AddResponseHeaders   map[string]string         `json:"2,omitempty"`
	DeleteRequestHeaders []string                  `json:"3,omitempty"`
	InterruptRequest     *ResponseInterruptRequest `json:"4,omitempty"`
}

func (r *Response) reset() {
	for key := range r.AddRequestHeaders {
		delete(r.AddRequestHeaders, key)
	}
	for key := range r.AddResponseHeaders {
		delete(r.AddResponseHeaders, key)
	}
	r.DeleteRequestHeaders = r.DeleteRequestHeaders[:0]
	r.InterruptRequest = nil
}

func (r *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Response) UnmarshalJSON(buf []byte) error {
	return json.Unmarshal(buf, r)
}
