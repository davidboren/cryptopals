package cryptopals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLikelyXor(t *testing.T) {
	char, score, res := MostLikelyXorChar(HexToBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	assert.Equal(t, byte(88), char)
	// assert.InDelta(t, -123.40657307311709, score, 0.000000001)
	// assert.InDelta(t, -0.5751794248011923, score, 0.000000001)
	assert.InDelta(t, -110.35122500937223, score, 0.000000001)

	assert.Equal(t, "Cooking MC's like a pound of bacon", string(res))
}
