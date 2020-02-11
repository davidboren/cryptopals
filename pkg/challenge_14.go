package cryptopals

import (
	b64 "encoding/base64"
	"math/rand"
	"time"
)

type encryptionPrefix struct {
	consistentPrefix []byte
}

var challenge14Prefix *encryptionPrefix

type challenge14Encryptor struct{}

func (c challenge14Encryptor) GetPrefix() []byte {
	if challenge14Prefix == nil {
		rand.Seed(time.Now().UnixNano())
		challenge14Prefix = &encryptionPrefix{
			consistentPrefix: randBytes(rand.Int() % 100),
		}
	}
	return challenge14Prefix.consistentPrefix
}

func (c challenge14Encryptor) GetSuffix() []byte {
	decoded, _ := b64.StdEncoding.DecodeString(appendText)
	return decoded

}

func (c challenge14Encryptor) Encrypt(data []byte) []byte {
	return encryptConsistent(append(c.GetPrefix(), append(data, c.GetSuffix()...)...))
}
