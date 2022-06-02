package rpc

import (
	"encoding/hex"
	"fmt"
	"go-sdk/response"
	"math/big"
	"strconv"
	"strings"
	"unicode/utf8"
)

func EncodeUintToHex(i uint64) string {
	return fmt.Sprintf("0x%x", i)
}

func ParseBigInt(str string) *big.Int {
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
	}
	num := new(big.Int)
	num.SetString(str, 16)
	return num
}

func ParseUint64orHex(str string) (uint64, error) {
	base := 10
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
		base = 16
	}
	return strconv.ParseUint(str, base, 64)
}

func EncodeToHex(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

func ParseHexBytes(str string) ([]byte, error) {
	if !strings.HasPrefix(str, "0x") {
		return nil, fmt.Errorf("it does not have 0x prefix")
	}
	buf, err := hex.DecodeString(str[2:])
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func parseUint64orHex(str string) (uint64, error) {
	base := 10
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
		base = 16
	}
	return strconv.ParseUint(str, base, 64)
}

func SubStr(str string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(str); i++ {
		_, size = utf8.DecodeRuneInString(str[n:])
		n += size
	}
	return str[:n]
}

func Inputer(fun string, parm []string) ([]byte, error) {
	bs := []byte(fun)
	b := response.Keccak256(bs)
	res := EncodeToHex(b)
	r := SubStr(res, 10)
	var build strings.Builder
	build.WriteString(r)
	for i := 0; i < len(parm); i++ {
		hash := response.HexToHash(parm[i])
		s := hash.String()
		s2 := s[2:]
		build.WriteString(s2)
	}
	fmt.Println(build.String())
	return ParseHexBytes(build.String())
}
