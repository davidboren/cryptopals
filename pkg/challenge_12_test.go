package cryptopals

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetBlockSize(t *testing.T) {
	assert.Equal(t, getSecret().consistentBlockSize, getBlockSize(challenge12Encryptor{}, StringGenerator{}))
}

func TestPrefixLength12(t *testing.T) {
	assert.Equal(t, len(challenge12Encryptor{}.GetPrefix()), getPrefixLength(challenge12Encryptor{}, StringGenerator{}, 16))
}

func TestSuffixLength12(t *testing.T) {
	assert.Equal(t, len(challenge12Encryptor{}.GetSuffix()), getSuffixLength(challenge12Encryptor{}, StringGenerator{}, 16))
}

func TestChallenge12String(t *testing.T) {
	msg := `Rollin' in my 5.0
With my rag-top down so my hair can blow
The girlies on standby waving just to say hi
Did you stop? No, I just drove by
`
	decoded := string(decryptMultiCallECB(challenge12Encryptor{}, StringGenerator{}))
	assert.Equal(t, msg, strings.Trim(decoded, string([]byte{byte('\x04')})))
}
