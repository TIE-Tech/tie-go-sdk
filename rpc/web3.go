package rpc

type Web3 struct {
	c *Client
}

func (c *Client) Web3() *Web3 {
	return c.endpoints.w
}

// Returns the current client version
func (w *Web3) ClientVersion() (string, error) {
	var out string
	err := w.c.Call("web3_clientVersion", &out)
	return out, err
}

// Returns the keccak-256 hash value of the specified data
func (w *Web3) Sha3(val []byte) ([]byte, error) {
	var out string
	if err := w.c.Call("web3_sha3", &out, EncodeToHex(val)); err != nil {
		return nil, err
	}
	return ParseHexBytes(out)
}
