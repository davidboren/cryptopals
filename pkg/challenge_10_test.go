package cryptopals

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCBC(t *testing.T) {
	key := "YELLOW SUBMARINE"
	iv := []byte("\x00\x00\x00")
	blockSize := 16
	padWith := byte('\x04')
	data := "I am the very model of a modern major general.  I've information vegetable, animal, and mineral"
	padded := PadBlocks([]byte(data), blockSize, padWith)

	encrypted := CBCEncrypt(padded, []byte(key), blockSize, iv)
	assert.Equal(t, len(padded), len(encrypted))

	decrypted := CBCDecrypt(encrypted, []byte(key), blockSize, iv)
	assert.True(t, strings.HasPrefix(string(decrypted), data))
	assert.Equal(t, string(padded), string(decrypted))

}

func TestChallenge10(t *testing.T) {
	data := loadChallenge10()

	decrypted := CBCDecrypt(data, []byte("YELLOW SUBMARINE"), 16, []byte("\x00"))
	assert.True(t, strings.HasPrefix(string(decrypted), "I'm back and I'm ringin' the bell"))

}
