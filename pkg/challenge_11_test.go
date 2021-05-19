package cryptopals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectECB(t *testing.T) {
	blockSize := 16
	data := CBCDecrypt(loadChallenge10(), []byte("YELLOW SUBMARINE"), 16, []byte("\x00"))
	min_dups := 2
	// count_actual := 0
	// count_pred := 0
	for i := 1; i < 100; i++ {
		encrypted, isAES := RandEncrypt([]byte(data), blockSize)
		predictedAES := DetectECB(encrypted, blockSize, min_dups)
		// if predictedAES {
		// 	count_pred++
		// 	t.Logf("Actual: %v\nPredicted: %v", isAES, true)
		// }
		// if isAES {
		// 	count_actual++
		// }
		assert.Equal(t, isAES, predictedAES)
	}
	// t.Logf("Actual Count: %v\nPredicted Count: %v", count_actual, count_pred)
	// t.Fail()
}
