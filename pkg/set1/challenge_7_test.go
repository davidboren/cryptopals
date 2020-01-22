package set1

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAes(t *testing.T) {
	decodedBytes := DecryptAes128Ecb(loadChallenge7(), []byte("YELLOW SUBMARINE"), 16)
	assert.True(t, strings.HasPrefix(string(decodedBytes), "I'm back and I'm ringin' the bell"))
}
