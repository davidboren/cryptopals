package cryptopals

import (
	b64 "encoding/base64"
)

type challenge19Encryptor struct {
	numBytes int
}

func (c challenge19Encryptor) GetEncryptedNonceBytes() []byte {
	secret := getSecret()
	return EncryptAes128Ecb(GetNonceBytes(uint64(0)), secret.consistentKey, c.numBytes)
}

func loadChallenge19() [][]byte {
	arr := [][]byte{}
	for _, s := range FileToArray("/Users/dboren/dev/go/src/github.com/davidboren/cryptopals/data/set3/challenge_19.txt") {
		b, _ := b64.StdEncoding.DecodeString(s)
		arr = append(arr, b)
	}
	return arr
}

func challenge19AsStrings() []string {
	s := []string{}
	fullArr := loadChallenge19()
	for _, arr := range fullArr {
		s = append(s, string(arr))
	}
	return s
}

func loadEncryptedChallenge19() [][]byte {
	c19 := loadChallenge19()
	fullArr := [][]byte{}
	for _, arr := range c19 {
		fullArr = append(
			fullArr,
			encryptCTR(arr, challenge19Encryptor{16}),
		)
	}
	return fullArr
}
