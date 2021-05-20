package cryptopals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChallenge18(t *testing.T) {
	expected := "Yo, VIP Let's kick it Ice, Ice, baby Ice, Ice, baby "
	encryptor := challenge18Encryptor{"YELLOW SUBMARINE", uint64(0)}
	decrypted := encryptCTR(loadChallenge18(), &encryptor)
	assert.Equal(t, expected, string(decrypted))
}
