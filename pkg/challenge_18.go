package cryptopals

import (
	b64 "encoding/base64"
	"encoding/binary"
)

var challenge18 string = `
L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ==
`

// NonceEncryptor requires a method for generating encrypted nonce bytes given a uint64
type NonceEncryptor interface {
	GetEncryptedNonceBytes() []byte
}

type challenge18Encryptor struct {
	Key   string
	Nonce uint64
}

func (c *challenge18Encryptor) GetEncryptedNonceBytes() []byte {
	encrypted := EncryptAes128Ecb(GetNonceBytes(c.Nonce), []byte(c.Key), 16)
	c.Nonce++
	return encrypted
}

func GetNonceBytes(i uint64) []byte {
	b := make([]byte, 16)
	binary.LittleEndian.PutUint64(b[8:], i)
	return b
}

func encryptCTR(b []byte, e NonceEncryptor) []byte {
	nonceBytes := e.GetEncryptedNonceBytes()
	j := 0
	encrypted := make([]byte, len(b))
	for i, v := range b {
		if j == len(nonceBytes) {
			nonceBytes = e.GetEncryptedNonceBytes()
			j = 0
		}
		encrypted[i] = v ^ nonceBytes[j]
		j++
	}
	return encrypted
}

// func encryptCTR(b []byte, e NonceEncryptor) []byte {
// 	nonce := uint64(0)
// 	nonceBytes := make([]byte, 16)
// 	j := len(nonceBytes)
// 	encrypted := make([]byte, len(b))
// 	for i, v := range b {
// 		if j == len(nonceBytes) {
// 			if i != 0 {
// 				nonce++
// 			}
// 			nonceBytes = e.GetEncryptedNonceBytes(nonce)
// 			j = 0
// 		}
// 		encrypted[i] = v ^ nonceBytes[j]
// 		j++
// 	}
// 	return encrypted
// }

func loadChallenge18() []byte {
	b, err := b64.StdEncoding.DecodeString(challenge18)
	if err != nil {
		panic(err)
	}
	return b
}
