package cryptopals

import (
	"testing"
)

func TestChallenge20(t *testing.T) {
	challenge20 := challenge20AsStrings()
	encrypted20 := loadEncryptedChallenge20()
	key, decrypted := BreakRepeatingXorArrays(encrypted20, 30)
	t.Logf("Key: %v", string(key))
	for i, d := range decrypted {
		if challenge20[i] == string(d) {
			t.Logf("\nMatching: %v", challenge20[i])
		} else {
			t.Logf("\nReal: %v\nDecrypted: %v", challenge20[i], string(d))
		}
		// assert.Equal(t, challenge19[i], string(d))
	}
	// t.Fail()
}
