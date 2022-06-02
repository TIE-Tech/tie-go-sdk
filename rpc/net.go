package rpc

type Net struct {
	c *Client
}

func (c *Client) Net() *Net {
	return c.endpoints.n
}

// Returns the ID of the currently connected network
func (n *Net) Version() (uint64, error) {
	var out string
	if err := n.c.Call("net_version", &out); err != nil {
		return 0, err
	}
	return ParseUint64orHex(out)
}

// Monitor the status of network connection
func (n *Net) Listening() (bool, error) {
	var out bool
	err := n.c.Call("net_listening", &out)
	return out, err
}

// Number of peer nodes connected by the client
func (n *Net) PeerCount() (uint64, error) {
	var out string
	if err := n.c.Call("net_peerCount", &out); err != nil {
		return 0, err
	}
	return ParseUint64orHex(out)
}
