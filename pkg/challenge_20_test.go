package cryptopals

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChallenge20(t *testing.T) {
	challenge20 := challenge20AsStrings()
	encrypted20 := loadEncryptedChallenge20()
	key, decrypted := BreakRepeatingXorArrays(encrypted20, 32)
	t.Logf("Key: %v", string(key))
	t.Logf("Key: %v", []byte(string(key)))
	t.Logf("KeyLength: %v", len(key))
	allMatching := true
	for i, d := range decrypted {
		real := strings.ToUpper(challenge20[i])
		decrypted := strings.ToUpper(string(d))
		if real == decrypted {
			t.Logf("\nMatching: %v", challenge20[i])
		} else {
			t.Logf("\nReal: %v\nDecrypted: %v", challenge20[i], string(d))
			allMatching = false
		}
	}
	t.Logf("\nAllMatching: %v", allMatching)
	assert.True(t, allMatching)
}
