package cryptopals

import (
	b64 "encoding/base64"
)

func loadChallenge20() [][]byte {
	arr := [][]byte{}
	for _, s := range FileToArray("/Users/dboren/dev/go/src/github.com/davidboren/cryptopals/data/set3/challenge_20.txt") {
		b, _ := b64.StdEncoding.DecodeString(s)
		arr = append(arr, b)
	}
	return arr
}

func challenge20AsStrings() []string {
	s := []string{}
	fullArr := loadChallenge20()
	for _, arr := range fullArr {
		s = append(s, string(arr))
	}
	return s
}

func loadEncryptedChallenge20() [][]byte {
	c20 := loadChallenge20()
	fullArr := [][]byte{}
	for _, arr := range c20 {
		fullArr = append(
			fullArr,
			encryptCTR(arr, challenge19Encryptor{16}),
		)
	}
	return fullArr
}
