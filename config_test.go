package traefik_plugin_bridge

import "testing"

func TestFailStrategy_Validate(t *testing.T) {
	var f FailStrategy

	f = "PASS"
	if err := f.Validate(); err != nil {
		t.Fatalf("%s", err)
	}

	f = "INTERRUPT"
	if err := f.Validate(); err != nil {
		t.Fatalf("%s", err)
	}

	f = "WRONG"
	err := f.Validate()
	if err == nil {
		t.Fatalf("should be error")
	}
	if err.Error() != "wrong FailStrategy, expected values: PASS, INTERRUPT" {
		t.Fatalf("unexpected error text")
	}
}

func TestProtocol_Validate(t *testing.T) {
	var p Protocol

	p = "HTTP"
	if err := p.Validate(); err != nil {
		t.Fatalf("%s", err)
	}

	p = "BINARY"
	if err := p.Validate(); err != nil {
		t.Fatalf("%s", err)
	}

	p = "WRONG"
	err := p.Validate()
	if err == nil {
		t.Fatalf("should be error")
	}
	if err.Error() != "wrong Protocol, expected values: HTTP, BINARY" {
		t.Fatalf("unexpected error text")
	}
}

func TestConfig_Validate_EmptyAddress(t *testing.T) {
	c := Config{
		Protocol:     "HTTP",
		Address:      "",
		Timeout:      0,
		Request:      ConfigRequest{},
		FailStrategy: "PASS",
	}
	err := c.Validate()
	if err == nil {
		t.Fatalf("should be error")
	}
	if err.Error() != "address must be not empty" {
		t.Fatalf("unexpected error text")
	}
}

func TestConfig_Validate_WrongProtocol(t *testing.T) {
	c := Config{
		Protocol:     "WRONG",
		Address:      "localhost",
		Timeout:      0,
		Request:      ConfigRequest{},
		FailStrategy: "PASS",
	}
	err := c.Validate()
	if err == nil {
		t.Fatalf("should be error")
	}
	if err.Error() != "wrong Protocol, expected values: HTTP, BINARY" {
		t.Fatalf("unexpected error text")
	}
}

func TestConfig_Validate_WrongFailStrategy(t *testing.T) {
	c := Config{
		Protocol:     "HTTP",
		Address:      "localhost",
		Timeout:      0,
		Request:      ConfigRequest{},
		FailStrategy: "WRONG",
	}
	err := c.Validate()
	if err == nil {
		t.Fatalf("should be error")
	}
	if err.Error() != "wrong FailStrategy, expected values: PASS, INTERRUPT" {
		t.Fatalf("unexpected error text")
	}
}

func TestConfig_Validate(t *testing.T) {
	c := Config{
		Protocol:     "HTTP",
		Address:      "localhost",
		Timeout:      0,
		Request:      ConfigRequest{},
		FailStrategy: "PASS",
	}
	err := c.Validate()
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCreateConfig(t *testing.T) {
	c := CreateConfig()
	if c.Protocol != defaultProtocol {
		t.Fatalf("wrong default value")
	}
	if c.Address != defaultAddress {
		t.Fatalf("wrong default value")
	}
	if c.Timeout != defaultTimeout {
		t.Fatalf("wrong default value")
	}
	if c.FailStrategy != defaultFailStrategy {
		t.Fatalf("wrong default value")
	}
	if c.Request.Headers != true {
		t.Fatalf("wrong default value")
	}
	if c.Request.Headers != true {
		t.Fatalf("wrong default value")
	}
	if c.Request.RemoteAddress != true {
		t.Fatalf("wrong default value")
	}
	if c.Request.Method != true {
		t.Fatalf("wrong default value")
	}
	if c.Request.URI != true {
		t.Fatalf("wrong default value")
	}
}
