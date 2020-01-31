package cryptopals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectECB(t *testing.T) {
	blockSize := 16
	data := CBCDecrypt(loadChallenge10(), []byte("YELLOW SUBMARINE"), 16, []byte("\x00"))
	for i := 1; i < 100; i++ {
		encrypted, isAES := RandEncrypt([]byte(data), blockSize)
		assert.Equal(t, isAES, DetectECB(encrypted, blockSize, 1))
	}
}
