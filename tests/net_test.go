package tests

import (
	"go-sdk/rpc"
	"testing"
)

func TestNet_Listening(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	_, err := tie.Net().Listening()
	if err != nil {
		t.Error("Listening", err)
	}
}

func TestNet_Version(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	_, err := tie.Net().Version()
	if err != nil {
		t.Error("Version", err)
	}
}

func TestNet_PeerCount(t *testing.T) {
	tie, _ := rpc.NewClient(Testurl)
	_, err := tie.Net().PeerCount()
	if err != nil {
		t.Error("PeerCount", err)
	}
}
