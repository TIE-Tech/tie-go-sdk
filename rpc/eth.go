package rpc

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-sdk/response"
	"go-sdk/wallet"
	"math/big"

	crypto2 "go-sdk/crypto"
)

type Eth struct {
	c *Client
}

func (c *Client) Eth() *Eth {
	return c.endpoints.e
}

// Code for the specified address
func (e *Eth) GetCode(addr response.Address, block response.BlockNumberOrHash) (string, error) {
	var res string
	if err := e.c.Call("eth_getCode", &res, addr, block.Location()); err != nil {
		return "", err
	}
	return res, nil
}

// Specifies the value of the address storage location
func (e *Eth) GetStorageAt(addr response.Address, slot response.Hash, block response.BlockNumberOrHash) (response.Hash, error) {
	var hash response.Hash
	err := e.c.Call("eth_getStorageAt", &hash, addr, slot, block.Location())
	return hash, err
}

// Number of the latest block
func (e *Eth) BlockNumber() (uint64, error) {
	var out string
	if err := e.c.Call("eth_blockNumber", &out); err != nil {
		return 0, err
	}
	return ParseUint64orHex(out)
}

// Block information for the specified height
func (e *Eth) GetBlockByNumber(i response.BlockNumber, full bool) (*response.Block, error) {
	var b *response.Block
	if err := e.c.Call("eth_getBlockByNumber", &b, i.String(), full); err != nil {
		return nil, err
	}
	return b, nil
}

// Specifies the block information of the hash
func (e *Eth) GetBlockByHash(hash response.Hash, full bool) (*response.Block, error) {
	var b *response.Block
	if err := e.c.Call("eth_getBlockByHash", &b, hash, full); err != nil {
		return nil, err
	}
	return b, nil
}

// Polls the specified filter and returns an array of newly generated logs since the last poll
func (e *Eth) GetFilterChanges(id string) ([]*response.Log, error) {
	var raw string
	err := e.c.Call("eth_getFilterChanges", &raw, id)
	if err != nil {
		return nil, err
	}
	var res []*response.Log
	if err := json.Unmarshal([]byte(raw), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// Specify the transaction corresponding to the hash
func (e *Eth) GetTransactionByHash(hash response.Hash) (*response.Transaction, error) {
	var txn *response.Transaction
	err := e.c.Call("eth_getTransactionByHash", &txn, hash)
	return txn, err
}

// Create filter
func (e *Eth) NewFilter(filter *response.LogFilter) (string, error) {
	var id string
	err := e.c.Call("eth_newFilter", &id, filter)
	return id, err
}

// Create a filter in the node (to notify when a new block is generated)
func (e *Eth) NewBlockFilter() (string, error) {
	var id string
	err := e.c.Call("eth_newBlockFilter", &id, nil)
	return id, err
}

// Unloads the filter with the specified ID
func (e *Eth) UninstallFilter(id string) (bool, error) {
	var res bool
	err := e.c.Call("eth_uninstallFilter", &res, id)
	return res, err
}

// Returns the receipt for the specified transaction
func (e *Eth) GetTransactionReceipt(hash response.Hash) (*response.Receipt, error) {
	var receipt *response.Receipt
	err := e.c.Call("eth_getTransactionReceipt", &receipt, hash)
	return receipt, err
}

// Number of transactions occurring at the specified address
func (e *Eth) GetTransactionCount(addr response.Address, blockNumber response.BlockNumberOrHash) (uint64, error) {
	var nonce string
	if err := e.c.Call("eth_getTransactionCount", &nonce, addr, blockNumber.Location()); err != nil {
		return 0, err
	}
	return ParseUint64orHex(nonce)
}

// Number of transactions within the specified block (height lookup)
func (e *Eth) GetBlockTransactionCountByNumber(blockNumber response.BlockNumber) (uint64, error) {
	var nonce uint64
	if err := e.c.Call("eth_getBlockTransactionCountByNumber", &nonce, blockNumber.String()); err != nil {
		return 0, err
	}
	return nonce, nil
}

// Balance of account at specified address
func (e *Eth) GetBalance(addr response.Address, blockNumber response.BlockNumberOrHash) (*big.Int, error) {
	var out string
	if err := e.c.Call("eth_getBalance", &out, addr, blockNumber.Location()); err != nil {
		return nil, err
	}
	b, ok := new(big.Int).SetString(out[2:], 16)
	if !ok {
		return nil, fmt.Errorf("failed to convert to big.int")
	}
	return b, nil
}

// Current gas price
func (e *Eth) GasPrice() (uint64, error) {
	var out string
	if err := e.c.Call("eth_gasPrice", &out); err != nil {
		return 0, err
	}
	return ParseUint64orHex(out)
}

// Create message call transaction or contract creation
func (e *Eth) SendTransaction(txn *response.Transaction) (response.Hash, error) {
	var hash response.Hash
	err := e.c.Call("eth_sendTransaction", &hash, txn)
	return hash, err
}

// Create a new message for the signed transaction and invoke transaction or contract creation
func (e *Eth) SendRawTransaction(data []byte) (response.Hash, error) {
	var hash response.Hash
	hexData := "0x" + hex.EncodeToString(data)

	err := e.c.Call("eth_sendRawTransaction", &hash, hexData)
	return hash, err
}

// Create a new message for the signed transaction and invoke transaction or contract creation
func (e *Eth) SendRawTransactionStr(data []byte) (response.Hash, error) {
	var hash response.Hash
	hexData := "0x" + string(data)

	err := e.c.Call("eth_sendRawTransaction", &hash, hexData)
	return hash, err
}

// Execute a new message call immediately without creating a transaction on the blockchain
func (e *Eth) Call(msg *response.CallMsg, block response.BlockNumber) (string, error) {
	var out string
	if err := e.c.Call("eth_call", &out, msg, block.String()); err != nil {
		return "", err
	}
	return out, nil
}

// Estimated transaction gas fee
func (e *Eth) EstimateGas(msg *response.CallMsg) (uint64, error) {
	var out string
	if err := e.c.Call("eth_estimateGas", &out, msg); err != nil {
		return 0, err
	}
	return ParseUint64orHex(out)
}

// Specify all logs in the filter
func (e *Eth) GetLogs(filter *response.LogFilter) ([]*response.Log, error) {
	var out []*response.Log
	if err := e.c.Call("eth_getLogs", &out, filter); err != nil {
		return nil, err
	}
	return out, nil
}

// Chain ID
func (e *Eth) ChainID() (*big.Int, error) {
	var out string
	if err := e.c.Call("eth_chainId", &out); err != nil {
		return nil, err
	}
	return ParseBigInt(out), nil
}

// Get nonce
func (e *Eth) GetNonce(addr response.Address, blockNumber response.BlockNumberOrHash) (uint64, error) {
	var nonce string
	if err := e.c.Call("eth_getTransactionCount", &nonce, addr, blockNumber.Location()); err != nil {
		return 0, err
	}
	return parseUint64orHex(nonce)
}

// Generate transaction signature
func (e *Eth) Sign(hexprv string, txn *response.Transaction, chainId uint64) ([]byte, error) {

	// Set nonce
	nonce, _ := e.GetNonce(txn.From, response.Latest)
	txn.Nonce = nonce

	// start
	signer1 := wallet.NewEIP155Signer(chainId)

	// Private key processing
	eckey, err := crypto2.HexToECDSA(hexprv) // Resolve private key

	if err != nil {
		return nil, err
	}

	// Private key structure
	ecdsaKey := new(ecdsa.PrivateKey)
	ecdsaKey.PublicKey = eckey.PublicKey
	ecdsaKey.D = eckey.D

	// Create a new private key
	key := wallet.NewKey(ecdsaKey)

	// Create signature
	signtxn, err := signer1.SignTx(txn, key)
	if err != nil {
		return nil, err
	}

	// serialize
	b, err := signtxn.MarshalRLPTo(nil)
	if err != nil {
		return nil, err
	}

	// tranStr := "0x" + hex.EncodeToString(b)
	tranStr := hex.EncodeToString(b)

	return []byte(tranStr), nil
}
