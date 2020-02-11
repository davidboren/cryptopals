package cryptopals

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestChallenge16Encryption(t *testing.T) {
	e := challenge16Encryptor{}
	data := []byte("A verry verry veryy very very very longish string-a-ding")
	assert.True(t, strings.Contains(string(CBCDecryptConsistent(e.Encrypt(data))), string(data)))
}

func TestChallenge16Prefix(t *testing.T) {
	e := challenge16Encryptor{}
	g := StringGenerator{}
	blockSize := getBlockSize(e, g)
	assert.Equal(t, 16, blockSize)
	assert.Equal(t, len(e.GetPrefix()), getPrefixLength(e, g, blockSize))
}

func TestChallenge16Suffix(t *testing.T) {
	e := challenge16Encryptor{}
	g := StringGenerator{}
	blockSize := getBlockSize(e, g)
	assert.Equal(t, len(e.GetSuffix()), getSuffixLength(e, g, blockSize))
}

func TestChallenge16Stuff3(t *testing.T) {
	e := challenge16Encryptor{}
	encrypted := crackChallenge16(e, StringGenerator{})
	assert.True(t, isAdmin(string(encrypted)))
}
