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

type Config struct {
	DialTimeout           int
	KeepAliveTimeout      int
	ResponseHeaderTimeout int

	APIKey string
}

type Client struct {
	*Config

	client *http.Client
}

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
