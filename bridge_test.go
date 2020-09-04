package traefik_plugin_bridge

import (
	"bytes"
	"context"
	"log"
	"strings"
	"testing"
)

func TestBridgeNew_WrongConfig(t *testing.T) {
	_, err := New(context.Background(), nil, &Config{}, "foo")
	if err == nil {
		t.Fatalf("should be not nil")
	}
}

func TestBridgeNew_UnimplementedBinary(t *testing.T) {
	_, err := New(context.Background(), nil, &Config{Address: "foo", Protocol: "BINARY", FailStrategy: defaultFailStrategy}, "foo")
	if err == nil {
		t.Fatalf("should be not nil")
	}
	if err.Error() != "error create the binary client, not implemented" {
		t.Fatalf("unexpected error text")
	}
}

func TestBridgeNew(t *testing.T) {
	_, err := New(context.Background(), nil, CreateConfig(), "foo")
	if err != nil {
		t.Fatalf("should be nil")
	}
}

func TestBridgeLog(t *testing.T) {
	out := &bytes.Buffer{}
	log.SetOutput(out)

	b := &Bridge{}

	b.log("foo %s", "bar")

	if !strings.HasSuffix(out.String(), "[BRIDGE MIDDLEWARE] foo bar\n") {
		t.Fatalf("wrong text")
	}
}
