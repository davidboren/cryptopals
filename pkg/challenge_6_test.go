package cryptopals

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHamming(t *testing.T) {
	dist, _ := hammingStrings("this is a test", "wokka wokka!!!")
	assert.Equal(t, 37, dist)
}

func TestFullDecode6(t *testing.T) {
	bestXorKey, decodedBytes := BreakRepeatingXor(loadChallenge6(), 32)
	assert.Equal(t, string(bestXorKey), "Terminator X: Bring the noise")
	assert.True(t, strings.HasPrefix(string(decodedBytes), "I'm back and I'm ringin' the bell"))
	assert.True(t, strings.HasSuffix(string(decodedBytes), "Play that funky music \n"))
}
