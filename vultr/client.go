// Vultr API in golang.
// You can use this SDK to create a Vultr API Client and then use the convenient
// methods that encapsulates Vultr API. There are also some useful methods which
// are made up of serveral methods and they can make your Vultr life better.
package vultr

import (
	"net"
	"net/http"
	"time"
)

const (
	DefaultDialTimeout           = 3
	DefaultResponseHeaderTimeout = 15
	DefaultKeepAliveTimeout      = 30

	HeaderAPIKey      = "API-Key"
	HeaderContentType = "Content-Type"
)

// SDK Config
type Config struct {
	// required
	APIKey string

	// optional
	DialTimeout           int
	KeepAliveTimeout      int
	ResponseHeaderTimeout int
}

// SDK Client
type Client struct {
	*Config

	client *http.Client
}

// Initialize a new client
func New(cfg *Config) (*Client, error) {
	if cfg == nil || cfg.APIKey == "" {
		return nil, ErrNoAPIKey
	}
	if cfg.DialTimeout <= 0 {
		cfg.DialTimeout = DefaultDialTimeout
	}
	if cfg.ResponseHeaderTimeout <= 0 {
		cfg.ResponseHeaderTimeout = DefaultResponseHeaderTimeout
	}
	if cfg.KeepAliveTimeout <= 0 {
		cfg.KeepAliveTimeout = DefaultKeepAliveTimeout
	}

	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   time.Duration(cfg.DialTimeout) * time.Second,
			KeepAlive: time.Duration(cfg.KeepAliveTimeout) * time.Second,
		}).Dial,
		ResponseHeaderTimeout: time.Duration(cfg.ResponseHeaderTimeout) * time.Second,
	}

	vc := &Client{
		Config: cfg,
		client: &http.Client{
			Transport: tr,
		},
	}

	return vc, nil
}
