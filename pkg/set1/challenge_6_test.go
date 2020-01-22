package set1

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHamming(t *testing.T) {
	dist, _ := hammingStrings("this is a test", "wokka wokka!!!")
	assert.Equal(t, 37, dist)
}

func TestFullDecode(t *testing.T) {
	bestXorKey, decodedBytes := DecodeRepeatingXor(loadChallenge6(), 10)
	assert.Equal(t, string(bestXorKey), "Terminator X: Bring the noise")
	assert.True(t, strings.HasPrefix(string(decodedBytes), "I'm back and I'm ringin' the bell"))
}
