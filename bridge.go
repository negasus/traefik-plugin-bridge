package traefik_plugin_bridge

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type bridgeMode int

const (
	bridgeModeHTTP   = 1
	bridgeModeBINARY = 2
)

// Bridge plugin.
type Bridge struct {
	mode          bridgeMode
	clientHTTP    *ClientHTTP
	clientBINARY  *ClientBINARY
	next          http.Handler
	name          string
	requestConfig ConfigRequest
}

// New created a new Bridge plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation error, %w", err)
	}

	b := &Bridge{
		next:          next,
		name:          name,
		requestConfig: config.Request,
	}

	var err error

	switch config.Protocol {
	case ProtocolBinary:
		b.mode = bridgeModeBINARY
		b.clientBINARY, err = newClientBinary(config.Address, time.Millisecond*time.Duration(config.Timeout))
		if err != nil {
			return nil, fmt.Errorf("error create the binary client, %w", err)
		}
	default:
		b.mode = bridgeModeHTTP
		b.clientHTTP, err = newClientHTTP(config.Address, time.Millisecond*time.Duration(config.Timeout))
		if err != nil {
			return nil, fmt.Errorf("error create the http client, %w", err)
		}
	}

	return b, nil
}

func (bridge *Bridge) log(format string, v ...interface{}) {
	log.Printf("[BRIDGE MIDDLEWARE: %s] %s", bridge.name, fmt.Sprintf(format, v...))
}
