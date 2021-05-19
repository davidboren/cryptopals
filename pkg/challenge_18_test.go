package cryptopals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChallenge18(t *testing.T) {
	expected := "Yo, VIP Let's kick it Ice, Ice, baby Ice, Ice, baby "
	decrypted := encryptCTR(loadChallenge18(), challenge18Encryptor{"YELLOW SUBMARINE"})
	assert.Equal(t, expected, string(decrypted))
}
