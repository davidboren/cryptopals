package cryptopals

import (
	"testing"
)

func TestChallenge19(t *testing.T) {
	challenge19 := challenge19AsStrings()
	encrypted19 := loadEncryptedChallenge19()
	key, decrypted := BreakRepeatingXorArrays(encrypted19, 16)
	t.Logf("Key: %v", string(key))
	for i, d := range decrypted {
		if challenge19[i] == string(d) {
			t.Logf("\nMatching: %v", challenge19[i])
		} else {
			t.Logf("\nReal: %v\nDecrypted: %v", challenge19[i], string(d))
		}
		// assert.Equal(t, challenge19[i], string(d))
	}
	// t.Fail()
}
