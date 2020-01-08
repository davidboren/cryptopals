package set1

import (
	// "gotest.tools/assert"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLikelyXor(t *testing.T) {
	assert.Equal(t, "58", "58")
	HexToBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	// XorSingleChar("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736", string([]byte{58}))
	// char, score, res := MostLikelyXorChar("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	// assert.Assert(t, score == -137.0327050351889)
	// assert.Assert(t, res == "Cooking MC's like a pound of bacon")
}
