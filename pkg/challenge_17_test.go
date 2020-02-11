package cryptopals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChallenge17Padding(t *testing.T) {
	e := encryptRandom17()
	assert.True(t, decryptRandom17(e))
}
