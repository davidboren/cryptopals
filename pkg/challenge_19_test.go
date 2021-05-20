package cryptopals

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChallenge19(t *testing.T) {
	challenge19 := challenge19AsStrings()
	encrypted19 := loadEncryptedChallenge19()
	key, decrypted := BreakRepeatingXorArrays(encrypted19, 32)
	t.Logf("Key: %v", string(key))
	t.Logf("KeyLength: %v", len(key))
	allMatching := true
	for i, d := range decrypted {
		real := strings.ToUpper(challenge19[i])
		decrypted := strings.ToUpper(string(d))
		decrypted2 := RepeatingKeyXor(encrypted19[i], key)
		if real == decrypted {
			t.Logf("\nMatching: %v", challenge19[i])
		} else {
			t.Logf("\nReal: %v\nDecrypted: %v\nDecrypted2: %v", challenge19[i], string(d), string(decrypted2))
			allMatching = false
		}
	}
	t.Logf("\nAllMatching: %v", allMatching)
	assert.True(t, allMatching)
}
