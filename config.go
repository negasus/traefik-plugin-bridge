package traefik_plugin_bridge

import (
	"fmt"
	"log"
)

const (
	defaultProtocol     = ProtocolHTTP
	defaultAddress      = "http://127.0.0.1:8000"
	defaultTimeout      = 1000
	defaultFailStrategy = FailStrategyPass
)

const (
	FailStrategyPass      string = "PASS"
	FailStrategyInterrupt string = "INTERRUPT"
	ProtocolHTTP          string = "HTTP"
	ProtocolBinary        string = "BINARY"
)

// ValidateFailStrategy FailStrategy value
func ValidateFailStrategy(f string) error {
	if f != FailStrategyInterrupt && f != FailStrategyPass {
		return fmt.Errorf("wrong FailStrategy, expected values: %s, %s. got = %s", FailStrategyPass, FailStrategyInterrupt, f)
	}
	return nil
}

// ValidateProtocol Protocol value
func ValidateProtocol(p string) error {
	log.Printf("%T (%v)", ProtocolBinary, ProtocolBinary)
	if p != ProtocolHTTP && p != ProtocolBinary {
		return fmt.Errorf("wrong Protocol, expected values: %s, %s", ProtocolHTTP, ProtocolBinary)
	}
	return nil
}

// Config the plugin configuration.
type Config struct {
	Protocol     string        `json:"protocol,omitempty"`
	Address      string        `json:"address,omitempty"`
	Timeout      int           `json:"timeout,omitempty"`
	Request      ConfigRequest `json:"request,omitempty"`
	FailStrategy string        `json:"failstrategy,omitempty"`
}

// ConfigRequest describe fields for pass to the middleware backend
type ConfigRequest struct {
	Headers       bool `json:"headers,omitempty"`
	RemoteAddress bool `json:"remote_address,omitempty"`
	Method        bool `json:"method,omitempty"`
	URI           bool `json:"uri,omitempty"`
}

// Validate config
func (c *Config) Validate() error {
	log.Printf("config: %+v", c)

	if c.Address == "" {
		return fmt.Errorf("address must be not empty")
	}
	if err := ValidateProtocol(c.Protocol); err != nil {
		return err
	}
	if err := ValidateFailStrategy(c.FailStrategy); err != nil {
		return err
	}
	return nil
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Protocol: defaultProtocol,
		Address:  defaultAddress,
		Timeout:  defaultTimeout,
		Request: ConfigRequest{
			Headers:       true,
			RemoteAddress: true,
			Method:        true,
			URI:           true,
		},
		FailStrategy: defaultFailStrategy,
	}
}
