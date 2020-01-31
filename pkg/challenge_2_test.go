package cryptopals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXorHex(t *testing.T) {
	hexRes := XorHex(
		"1c0111001f010100061a024b53535009181c",
		"686974207468652062756c6c277320657965",
	)

	assert.Equal(t, hexRes, "746865206b696420646f6e277420706c6179")
	assert.Equal(t, string(HexToBytes(hexRes)), "the kid don't play")
}
