package cryptopals

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAes(t *testing.T) {
	decodedBytes := DecryptAes128Ecb(loadChallenge7(), []byte("YELLOW SUBMARINE"), 16)
	assert.True(t, strings.HasPrefix(string(decodedBytes), "I'm back and I'm ringin' the bell"))
}

func TestAesIdentity(t *testing.T) {
	testString := "A secret string is a very very very long strings"
	testKey := "YELLOW SUBMARINE"
	padded := PadBlocks([]byte(testString), 16, byte('\x04'))
	encryptedBytes := EncryptAes128Ecb(padded, []byte(testKey), 16)
	decryptedBytes := DecryptAes128Ecb(encryptedBytes, []byte(testKey), 16)
	assert.Equal(t, string(padded), string(decryptedBytes))
}
