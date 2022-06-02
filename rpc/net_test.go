package rpc

import (
	"go-sdk/tests"
	"testing"
)

func BenchmarkNet_Version(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := tie.Net().Version()
		if err != nil {
			return
		}
	}
}

func BenchmarkNet_Listening(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := tie.Net().Listening()
		if err != nil {
			return
		}
	}
}

func BenchmarkNet_PeerCount(b *testing.B) {
	tie, _ := NewClient(tests.Testurl)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := tie.Net().PeerCount()
		if err != nil {
			return
		}
	}
}
