package traefik_plugin_bridge

import (
	"fmt"
)

const (
	defaultProtocol     = ProtocolHTTP
	defaultAddress      = "localhost:8000"
	defaultTimeout      = 500
	defaultFailStrategy = FailStrategyPass
)

const (
	FailStrategyPass      FailStrategy = "PASS"
	FailStrategyInterrupt FailStrategy = "INTERRUPT"
	ProtocolHTTP          Protocol     = "HTTP"
	ProtocolBinary        Protocol     = "BINARY"
)

// FailStrategy describe strategy on fail request to the backend
type FailStrategy string

// Validate FailStrategy value
func (f FailStrategy) Validate() error {
	if f != FailStrategyInterrupt && f != FailStrategyPass {
		return fmt.Errorf("wrong FailStrategy, expected values: %s, %s", FailStrategyPass, FailStrategyInterrupt)
	}
	return nil
}

// Protocol describe interaction protocol
type Protocol string

// Validate Protocol value
func (p Protocol) Validate() error {
	if p != ProtocolHTTP && p != ProtocolBinary {
		return fmt.Errorf("wrong Protocol, expected values: %s, %s", ProtocolHTTP, ProtocolBinary)
	}
	return nil
}

// Config the plugin configuration.
type Config struct {
	Protocol     Protocol      `json:"protocol,omitempty"`
	Address      string        `json:"address,omitempty"`
	Timeout      int           `json:"timeout,omitempty"`
	Request      ConfigRequest `json:"request,omitempty"`
	FailStrategy FailStrategy  `json:"fail_strategy,omitempty"`
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
	if c.Address == "" {
		return fmt.Errorf("address must be not empty")
	}
	if err := c.Protocol.Validate(); err != nil {
		return err
	}
	if err := c.FailStrategy.Validate(); err != nil {
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
