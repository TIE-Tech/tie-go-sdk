package keystore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestV3_EncodeDecode(t *testing.T) {
	data := []byte{0x1, 0x2}
	password := "abcd"

	encrypted, err := EncryptV3(data, password)
	assert.NoError(t, err)

	found, err := DecryptV3(encrypted, password)
	assert.NoError(t, err)

	assert.Equal(t, data, found)
}

func BenchmarkV3EncodeDecode(b *testing.B) {
	data := []byte{0x1, 0x2}
	password := "abcd"
	for i := 0; i < b.N; i++ {
		encrypted, err := EncryptV3(data, password)
		assert.NoError(b, err)
		found, err := DecryptV3(encrypted, password)
		assert.NoError(b, err)

		assert.Equal(b, data, found)
	}
}
