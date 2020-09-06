package traefik_plugin_bridge

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequest_FillFromHTTPRequest(t *testing.T) {
	httpReq := &http.Request{
		Method:     "baz",
		RequestURI: "foo",
		Header: map[string][]string{
			"h1": []string{"v1"},
		},
		RemoteAddr: "bar",
	}

	req := &Request{}

	req.FillFromHTTPRequest(httpReq, &ConfigRequest{Headers: false, RemoteAddress: false, Method: false, URI: false})
	if len(req.Headers) != 0 {
		t.Fatalf("unexpected value")
	}
	if req.RemoteAddr != "" {
		t.Fatalf("unexpected value")
	}
	if req.Method != "" {
		t.Fatalf("unexpected value")
	}
	if req.RequestURI != "" {
		t.Fatalf("unexpected value")
	}

	req.FillFromHTTPRequest(httpReq, &ConfigRequest{Headers: true, RemoteAddress: true, Method: true, URI: true})
	if len(req.Headers) != 1 {
		t.Fatalf("unexpected value")
	}
	if req.RemoteAddr != "bar" {
		t.Fatalf("unexpected value")
	}
	if req.Method != "baz" {
		t.Fatalf("unexpected value")
	}
	if req.RequestURI != "foo" {
		t.Fatalf("unexpected value")
	}
}

func TestRequest_reset(t *testing.T) {
	req := &Request{
		Headers:    map[string][]string{"foo": []string{"bar"}},
		RemoteAddr: "a",
		Method:     "b",
		RequestURI: "c",
	}

	req.reset()

	if len(req.Headers) != 0 {
		t.Fatalf("unexpected value")
	}
	if req.RemoteAddr != "" {
		t.Fatalf("unexpected value")
	}
	if req.Method != "" {
		t.Fatalf("unexpected value")
	}
	if req.RequestURI != "" {
		t.Fatalf("unexpected value")
	}
}

func TestRequest_marshalJSON(t *testing.T) {
	req := &Request{
		Headers:    map[string][]string{"foo": []string{"bar"}},
		RemoteAddr: "a",
		Method:     "b",
		RequestURI: "c",
	}

	buf, err := req.ToJSON()
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if string(buf) != `{"1":{"foo":["bar"]},"2":"a","3":"b","4":"c"}` {
		t.Fatalf("unexpected text %s", buf)
	}
}

func TestRequest_unmarshalJSON(t *testing.T) {
	req := &Request{}

	err := req.FromJSON([]byte(`{"1":{"foo":["bar"]},"2":"a","3":"b","4":"c"}`))
	if err != nil {
		t.Fatalf("unexpected error")
	}
	if len(req.Headers) != 1 {
		t.Fatalf("unexpected value")
	}
	if req.RemoteAddr != "a" {
		t.Fatalf("unexpected value")
	}
	if req.Method != "b" {
		t.Fatalf("unexpected value")
	}
	if req.RequestURI != "c" {
		t.Fatalf("unexpected value")
	}
}

func TestRequestPool(t *testing.T) {
	req1 := AcquireRequest()
	p1 := fmt.Sprintf("%p", req1)
	ReleaseRequest(req1)
	req2 := AcquireRequest()
	p2 := fmt.Sprintf("%p", req2)
	if p1 != p2 {
		t.Fatalf("should be equal")
	}
}
