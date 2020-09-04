package traefik_plugin_bridge

import (
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
