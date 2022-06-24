package tests

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	crypto2 "github.com/TIE-Tech/tie-go-sdk/crypto"
	"github.com/TIE-Tech/tie-go-sdk/response"
	"github.com/TIE-Tech/tie-go-sdk/wallet"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSigner_EIP1155(t *testing.T) {
	signer1 := wallet.NewEIP155Signer(Testblocknumber)

	addr0 := response.Address{0x1}
	key, err := wallet.GenerateKey()
	assert.NoError(t, err)

	txn := &response.Transaction{
		To:       &addr0,
		Value:    big.NewInt(10),
		GasPrice: 0,
	}
	txn, err = signer1.SignTx(txn, key)
	assert.NoError(t, err)

	from, err := signer1.RecoverSender(txn)
	assert.NoError(t, err)
	assert.Equal(t, from, key.Addr)

	/*
		// try to use a signer with another chain id
		signer2 := NewEIP155Signer(2)
		from2, err := signer2.RecoverSender(txn)
		assert.NoError(t, err)
		assert.NotEqual(t, from, from2)
	*/
}

func TestKey_Sign(t *testing.T) {
	signer1 := wallet.NewEIP155Signer(Testblocknumber)

	from := response.HexToAddress(Testaddr)
	to := response.HexToAddress(Testaddr)
	hexprv := "22a90d9711350a0b7c7c697ccb26dd1224ffbf16f6430220d28f0a30235fb01e"

	eckey, err := crypto2.HexToECDSA(hexprv)
	if err != nil {
		t.Fatal(err)
	}

	ecdsaKey := new(ecdsa.PrivateKey)
	ecdsaKey.PublicKey = eckey.PublicKey
	ecdsaKey.D = eckey.D

	key := wallet.NewKey(ecdsaKey)
	hexData := "40c10f19000000000000000000000000d615c42cf7856e0634404b7584ef8fcd6cc9b8960000000000000000000000000000000000000000000000000000000000000001"
	data, err := hex.DecodeString(hexData)
	if err != nil {
		t.Fatal(err)
	}
	txn := &response.Transaction{
		Nonce:    50,
		From:     from,
		To:       &to,
		Value:    big.NewInt(0),
		Gas:      234723,     // 394e3
		GasPrice: 5000000000, // 12a05f200
		Input:    data,
	}
	fmt.Println("input==>", hex.EncodeToString(txn.Input))
	signtxn, err := signer1.SignTx(txn, key)
	if err != nil {
		t.Fatal("signtx err", err)
	}

	b, _ := signtxn.MarshalRLPTo(nil)
	t.Log("0x" + hex.EncodeToString(b))

	signer2 := wallet.NewEIP155Signer(Testblocknumber)
	from2, err := signer2.RecoverSender(txn)
	if err != nil {
		t.Fatal("recover err", err)
	}
	t.Log("recover==>", from2.String())
}
