package tests

import (
	"github.com/TIE-Tech/tie-go-sdk/rpc"
	"testing"
)

func TestWeb3_ClientVersion(t *testing.T) {
	tie, _ := rpc.NewClient(rpc.Url)
	_, err := tie.Web3().ClientVersion()
	if err != nil {
		t.Error("ClientVersion", err)
	}
}

func TestWeb3_Sha3(t *testing.T) {
	tie, _ := rpc.NewClient(rpc.Url)
	_, err := tie.Web3().Sha3([]byte("0x68656c6c6f20776f726c64"))
	if err != nil {
		t.Error("Sha3", err)
	}
}
