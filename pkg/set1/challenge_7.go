package set1

import (
	"crypto/aes"
)

func loadChallenge7() []byte {
	return B64ArrayToBytes(FileToArray("/Users/dboren/dev/go/src/github.com/davidboren/cryptopals/data/set1/challenge_7.txt"))
}

// DecryptAes128Ecb decrypts data with a key using Aes 128Ecb
func DecryptAes128Ecb(data, key []byte, blockSize int) []byte {
	cipher, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))

	for bs := 0; bs < len(data); bs += blockSize {
		cipher.Decrypt(decrypted[bs:bs+blockSize], data[bs:bs+blockSize])
	}

	return decrypted
}
