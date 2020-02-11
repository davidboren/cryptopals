package cryptopals

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestChallenge14String(t *testing.T) {
	msg := `Rollin' in my 5.0
With my rag-top down so my hair can blow
The girlies on standby waving just to say hi
Did you stop? No, I just drove by
`
	decoded := string(decryptMultiCallECB(challenge14Encryptor{}, StringGenerator{}))
	assert.Equal(t, msg, strings.Trim(decoded, string([]byte{byte('\x04')})))
}
