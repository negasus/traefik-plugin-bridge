package traefik_plugin_bridge

import (
	"fmt"
	"testing"
)

func TestResponsePool(t *testing.T) {
	res1 := acquireResponse()
	p1 := fmt.Sprintf("%p", res1)
	releaseResponse(res1)
	res2 := acquireResponse()
	p2 := fmt.Sprintf("%p", res2)
	if p1 != p2 {
		t.Fatalf("should be equal")
	}
}

func TestResponse_reset(t *testing.T) {
	resp := &Response{
		AddRequestHeaders:    map[string]string{"1": "1"},
		AddResponseHeaders:   map[string]string{"2": "2"},
		DeleteRequestHeaders: []string{"3", "4"},
		InterruptRequest: &ResponseInterruptRequest{
			ResponseCode: 200,
			Body:         "5",
		},
	}
	resp.reset()
	if len(resp.AddRequestHeaders) != 0 {
		t.Fatalf("unexpected value")
	}
	if len(resp.AddResponseHeaders) != 0 {
		t.Fatalf("unexpected value")
	}
	if len(resp.DeleteRequestHeaders) != 0 {
		t.Fatalf("unexpected value")
	}
	if resp.InterruptRequest != nil {
		t.Fatalf("unexpected value")
	}
}

func TestResponse_marshalJSON(t *testing.T) {
	resp := &Response{
		AddRequestHeaders:    map[string]string{"a": "b"},
		AddResponseHeaders:   map[string]string{"c": "d"},
		DeleteRequestHeaders: []string{"e", "f"},
		InterruptRequest: &ResponseInterruptRequest{
			ResponseCode: 200,
			Body:         "g",
		},
	}

	buf, err := resp.ToJSON()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if string(buf) != `{"1":{"a":"b"},"2":{"c":"d"},"3":["e","f"],"4":{"1":200,"2":"g"}}` {
		t.Fatalf("unexpected value")
	}
}

func TestResponse_unmarshalJSON(t *testing.T) {
	resp := &Response{}

	err := resp.FromJSON([]byte(`{"1":{"a":"b"},"2":{"c":"d"},"3":["e","f"],"4":{"1":200,"2":"g"}}`))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if len(resp.AddRequestHeaders) != 1 {
		t.Fatalf("unexpected value")
	}
	if len(resp.AddResponseHeaders) != 1 {
		t.Fatalf("unexpected value")
	}
	if len(resp.DeleteRequestHeaders) != 2 {
		t.Fatalf("unexpected value")
	}
	if resp.InterruptRequest == nil {
		t.Fatalf("unexpected value")
	}
}
