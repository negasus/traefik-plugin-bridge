package traefik_plugin_bridge

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

var (
	requestPool = sync.Pool{
		New: func() interface{} {
			return &Request{
				Headers: make(map[string][]string),
			}
		},
	}
)

func AcquireRequest() *Request {
	return requestPool.Get().(*Request)
}

func ReleaseRequest(r *Request) {
	r.reset()
	requestPool.Put(r)
}

type Request struct {
	Headers    map[string][]string `json:"1,omitempty"`
	RemoteAddr string              `json:"2,omitempty"`
	Method     string              `json:"3,omitempty"`
	RequestURI string              `json:"4,omitempty"`
}

func (r *Request) reset() {
	for key := range r.Headers {
		delete(r.Headers, key)
	}
	r.RemoteAddr = ""
	r.Method = ""
	r.RequestURI = ""
}

func (r *Request) FillFromHTTPRequest(httpReq *http.Request) {
	r.Headers = httpReq.Header
	r.RemoteAddr = httpReq.RemoteAddr
	r.Method = httpReq.Method
	r.RequestURI = httpReq.RequestURI
}

func (r *Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Request) UnmarshalJSON(buf []byte) error {
	return json.Unmarshal(buf, r)
}

func (r *Request) JSONReader() (io.Reader, error) {
	buf, err := r.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buf), nil
}
