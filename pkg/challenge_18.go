package cryptopals

import (
	b64 "encoding/base64"
	"encoding/binary"
)

var challenge18 string = `
L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ==
`

func loadChallenge18() []byte {
	b, err := b64.StdEncoding.DecodeString(challenge18)
	if err != nil {
		panic(err)
	}
	return b
}

func nonceToEncryptedBytes(v uint64, e ConsistentEncryptor) []byte {
	return e.Encrypt(nonceToBytes(v))

}

func nonceToBytes(v uint64) []byte {
	b := make([]byte, 16)
	binary.LittleEndian.PutUint64(b[8:], v)
	return b

}

func encryptCTR(b []byte, e ConsistentEncryptor) []byte {
	nonce := uint64(0)
	j := 0
	nonceBytes := nonceToEncryptedBytes(nonce, e)
	encrypted := make([]byte, len(b))
	for i, v := range b {
		encrypted[i] = v ^ nonceBytes[j]
		j++
		if j == len(nonceBytes) {
			nonce++
			nonceBytes = nonceToEncryptedBytes(nonce, e)
			j = 0
		}
	}
	return encrypted
}

type challenge18Encryptor struct {
	Key string
}

func (c challenge18Encryptor) Encrypt(data []byte) []byte {
	return EncryptAes128Ecb(data, []byte(c.Key), 16)
}
