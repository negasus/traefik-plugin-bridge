package traefik_plugin_bridge

import "testing"

func TestClientBinary_new_not_implemented(t *testing.T) {
	_, err := newClientBinary("", 0)
	if err == nil {
		t.Fatalf("expect an error")
	}
	if err.Error() != "not implemented" {
		t.Fatalf("unexpected error text")
	}
}

func TestClientBinary_call_not_implemented(t *testing.T) {
	c := &ClientBINARY{}
	_, err := c.Call(nil)
	if err == nil {
		t.Fatalf("expect an error")
	}
	if err.Error() != "not implemented" {
		t.Fatalf("unexpected error text")
	}
}
