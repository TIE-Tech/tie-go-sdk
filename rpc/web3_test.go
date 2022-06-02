package rpc

import (
	"go-sdk/tests"
	"testing"
)

func BenchmarkWeb3_ClientVersion(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := tie.Web3().ClientVersion()
		if err != nil {
			return
		}
	}
}

func BenchmarkWeb3_Sha3(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := tie.Web3().Sha3([]byte("0x68656c6c6f20776f726c64"))
		if err != nil {
			return
		}
	}
}
