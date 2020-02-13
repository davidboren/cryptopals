package cryptopals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChallenge18(t *testing.T) {
	expected := "Yo, VIP Let's kick it Ice, Ice, baby Ice, Ice, baby "
	assert.Equal(t, expected, string(encryptCTR(loadChallenge18(), challenge18Encryptor{"YELLOW SUBMARINE"})))
}
