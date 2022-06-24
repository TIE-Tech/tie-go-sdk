package rpc

import "github.com/TIE-Tech/tie-go-sdk/transport"

// rpc_client
type Client struct {
	transport transport.Transport
	endpoints endpoints
}

type endpoints struct {
	w *Web3
	e *Eth
	n *Net
}

type Config struct {
	headers map[string]string
}

type ConfigOption func(*Config)

func NewClient(addr string, opts ...ConfigOption) (*Client, error) {
	config := &Config{headers: map[string]string{}}
	for _, opt := range opts {
		opt(config)
	}

	c := &Client{}
	c.endpoints.w = &Web3{c}
	c.endpoints.e = &Eth{c}
	c.endpoints.n = &Net{c}

	t, err := transport.NewTransport(addr, config.headers)
	if err != nil {
		return nil, err
	}
	c.transport = t
	return c, nil
}

func (c *Client) Close() error {
	return c.transport.Close()
}

func (c *Client) Call(method string, out interface{}, params ...interface{}) error {
	return c.transport.Call(method, out, params...)
}
