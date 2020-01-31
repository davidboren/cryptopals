package cryptopals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPadding(t *testing.T) {
	assert.Equal(t, "YELLOW SUBMARINE\x04\x04\x04\x04", string(Pad([]byte("YELLOW SUBMARINE"), 20, byte('\x04'))))
}
