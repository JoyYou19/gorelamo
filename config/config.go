package config

import (
	"net/http"
	"time"
)

type Config struct {
	HTTPClient *http.Client
	Username   string
	Password   string
}

func DefaultConfig() *Config {
	return &Config{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func WithBasicAuth(user, pass string) Option {
	return func(cfg *Config) {
		cfg.Username = user
		cfg.Password = pass
	}
}
