package rpc

import (
	"github.com/stretchr/testify/assert"
	"go-sdk/response"
	"go-sdk/tests"
	"math/big"
	"testing"
)

func BenchmarkEth_GetNonce(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	buf := response.HexToAddress(tests.Testaddr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().GetTransactionCount(buf, response.Latest)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_ChainID(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().ChainID()
		assert.NoError(b, err)
	}
}

func BenchmarkEth_GetTransactionReceipt(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	buf := response.HexToHash(tests.Testtxnhash)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().GetTransactionReceipt(buf)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_GasPrice(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().GasPrice()
		assert.NoError(b, err)
	}
}

func BenchmarkEth_GetBalance(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	buf := response.HexToAddress(tests.Testaddr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().GetBalance(buf, response.Latest)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_GetStorageAt(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	buf := response.HexToAddress(tests.Testaddr)
	hx := response.HexToHash("0x0")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().GetStorageAt(buf, hx, response.Latest)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_GetTransactionCount(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	buf := response.HexToAddress(tests.Testaddr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().GetTransactionCount(buf, response.Latest)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_GetBlockTransactionCountByNumber(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().GetBlockTransactionCountByNumber(tests.Testblocknumber)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_GetCode(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	buf := response.HexToAddress(tests.Testaddr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().GetCode(buf, response.Latest)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_Call(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	From := response.HexToAddress(tests.Testaddr)
	To := response.HexToAddress(tests.Testaddr)
	data, _ := ParseHexBytes("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675")
	buf := &response.CallMsg{
		From:     From,
		To:       &To,
		Data:     data,
		GasPrice: 1000000000000,
		Gas:      big.NewInt(210000),
		Value:    big.NewInt(10000000000),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().Call(buf, response.Latest)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_EstimateGas(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	From := response.HexToAddress(tests.Testaddr)
	To := response.HexToAddress(tests.Testaddr)
	buf := &response.CallMsg{
		From:     From,
		To:       &To,
		GasPrice: 100000000000,
		Gas:      big.NewInt(210000),
		Value:    big.NewInt(10000000000),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().EstimateGas(buf)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_GetBlockByHash(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	data := response.HexToHash(tests.Testtxnhash)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().GetBlockByHash(data, true)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_GetBlockByNumber(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().GetBlockByNumber(response.BlockNumber(tests.Testblocknumber), true)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_BlockNumber(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().BlockNumber()
		assert.NoError(b, err)
	}
}

func BenchmarkEth_GetTransactionByHash(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	hash := response.HexToHash("0x3ca1d331567da2433a0690aaa632dfbdea236e5f27369669c8fe2461a6ccb813")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().GetTransactionByHash(hash)
		assert.NoError(b, err)
	}
}

func BenchmarkEth_NewFilter(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	buf := &response.LogFilter{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tie.Eth().NewFilter(buf)
		assert.NoError(b, err)
	}
}
