package cryptopals

import (
	"crypto/aes"
	"errors"
)

func loadChallenge7() []byte {
	return B64ArrayToBytes(FileToArray("/Users/dboren/dev/go/src/github.com/davidboren/cryptopals/data/set1/challenge_7.txt"))
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// DecryptAes128Ecb decrypts data with a key using Aes 128Ecb
func DecryptAes128Ecb(data, key []byte, blockSize int) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	decrypted := make([]byte, len(data))

	for bs := 0; bs < len(data); bs += blockSize {
		cipher.Decrypt(decrypted[bs:bs+blockSize], data[bs:bs+blockSize])
	}

	return decrypted
}

// EncryptAes128Ecb decrypts data with a key using Aes 128Ecb
func EncryptAes128Ecb(data, key []byte, blockSize int) []byte {
	if len(data)%blockSize != 0 {
		panic(errors.New("Your data should be padded prior to encrypting"))
	}
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted := make([]byte, len(data))

	for bs := 0; bs < len(data); bs += blockSize {
		cipher.Encrypt(encrypted[bs:bs+blockSize], data[bs:bs+blockSize])
	}

	return encrypted
}
