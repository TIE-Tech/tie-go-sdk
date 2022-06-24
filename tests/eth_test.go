package tests

import (
	"github.com/TIE-Tech/tie-go-sdk/response"
	"github.com/TIE-Tech/tie-go-sdk/rpc"
	"math/big"
	"testing"
)

func TestEth_GetNonce(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	buf := response.HexToAddress(Testaddr)
	_, err := tie.Eth().GetTransactionCount(buf, response.Latest)
	if err != nil {
		t.Error("GetNonce", err)
	}
	_, err = tie.Eth().GetTransactionCount(buf, response.BlockNumber(80481))
	if err != nil {
		t.Error("GetNonce", err)
	}
}

func TestEth_ChainID(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	_, err := tie.Eth().ChainID()
	if err != nil {
		t.Error("ChainID", err)
	}
}

func TestEth_GetTransactionReceipt(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	buf := response.HexToHash(Testtxnhash)
	_, err := tie.Eth().GetTransactionReceipt(buf)
	if err != nil {
		t.Error("TransactionReceipt", err)
	}
}

func TestEth_GasPrice(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	_, err := tie.Eth().GasPrice()
	if err != nil {
		t.Error("GasPrice", err)
	}
}

func TestEth_GetBalance(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	buf := response.HexToAddress(Testaddr)
	_, err := tie.Eth().GetBalance(buf, response.Latest)
	if err != nil {
		t.Error("GetBalance", err)
	}
	_, err = tie.Eth().GetBalance(buf, response.BlockNumber(80481))
	if err != nil {
		t.Error("GetBalance", err)
	}
}

func TestEth_GetStorageAt(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	buf := response.HexToAddress(Testaddr)
	hx := response.HexToHash("0x0")
	_, err := tie.Eth().GetStorageAt(buf, hx, response.Latest)
	if err != nil {
		t.Error("GetStorageAt", err)
	}
	hx = response.HexToHash(Testtxnhash)
	_, err = tie.Eth().GetStorageAt(buf, hx, response.Latest)
	if err != nil {
		t.Error("GetStorageAt", err)
	}
}

func TestEth_GetTransactionCount(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	buf := response.HexToAddress(Testaddr)
	_, err := tie.Eth().GetTransactionCount(buf, response.Latest)
	if err != nil {
		t.Error("GetTransactionCount", err)
	}
	_, err = tie.Eth().GetTransactionCount(buf, response.BlockNumber(80481))
	if err != nil {
		t.Error("GetTransactionCount", err)
	}
}

func TestEth_GetBlockTransactionCountByNumber(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	_, err := tie.Eth().GetBlockTransactionCountByNumber(response.Latest)
	if err != nil {
		t.Error("GetBlockTransactionCountByNumber", err)
	}
}

func TestEth_GetCode(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	buf := response.HexToAddress(Testaddr)
	_, err := tie.Eth().GetCode(buf, response.Latest)
	if err != nil {
		t.Error("GetCode", err)
	}
}

func TestEth_Call(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	From := response.HexToAddress(Testaddr)
	To := response.HexToAddress(Testaddr)
	data, _ := rpc.ParseHexBytes("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675")
	buf := &response.CallMsg{
		From:     From,
		To:       &To,
		Data:     data,
		GasPrice: 1000000000000,
		Gas:      big.NewInt(210000),
		Value:    big.NewInt(10000000000),
	}
	_, err := tie.Eth().Call(buf, response.Latest)
	if err != nil {
		t.Error("Call", err)
	}
}

func TestEth_EstimateGas(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	From := response.HexToAddress(Testaddr)
	To := response.HexToAddress(Testaddr)
	buf := &response.CallMsg{
		From:     From,
		To:       &To,
		GasPrice: 100000000000,
		Gas:      big.NewInt(210000),
		Value:    big.NewInt(10000000000),
	}
	_, err := tie.Eth().EstimateGas(buf)
	if err != nil {
		t.Error("EstimateGas", err)
	}
}

func TestEth_GetBlockByHash(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	data := response.HexToHash(Testtxnhash)
	_, err := tie.Eth().GetBlockByHash(data, true)
	if err != nil {
		t.Error("GetBlockByHash", err)
	}
}

func TestEth_GetBlockByNumber(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	_, err := tie.Eth().GetBlockByNumber(response.BlockNumber(18745349), true)
	if err != nil {
		t.Error("GetBlockByNumber", err)
	}
}

func TestEth_BlockNumber(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	_, err := tie.Eth().BlockNumber()
	if err != nil {
		t.Error("BlockNumber", err)
	}
}

func TestEth_GetTransactionByHash(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	hash := response.HexToHash(Testtxnhash)
	_, err := tie.Eth().GetTransactionByHash(hash)
	if err != nil {
		t.Error("GetTransactionByHash", err)
	}
}

func TestEth_NewFilter(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	buf := &response.LogFilter{}
	_, err := tie.Eth().NewFilter(buf)
	if err != nil {
		t.Error("NewFilter", err)
	}
}
