package gorelamo

import (
	"github.com/JoyYou19/gorelamo/config"
	"github.com/JoyYou19/gorelamo/database"
	"github.com/JoyYou19/gorelamo/transport"
)

type Client struct {
	baseURL   string
	transport *transport.HTTP
}

func NewClient(baseURL string, opts ...config.Option) *Client {
	cfg := config.DefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	return &Client{
		baseURL:   baseURL,
		transport: transport.NewHTTP(cfg),
	}
}

func (c *Client) Database(name string) *database.Database {
	return database.New(c.transport, c.baseURL, name)
}
