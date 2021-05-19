package cryptopals

import (
	"math/rand"
)

// PadBlocks rounds out an array of blocks with PKCS7 padding
func PadBlocksPKCS7(b []byte, blockSize int) []byte {
	return PadBlocks(b, blockSize, byte(len(b)%blockSize))
}

func loadChallenge17() []string {
	return FileToArray("/Users/dboren/dev/go/src/github.com/davidboren/cryptopals/data/set3/challenge_17.txt")
}

func encryptRandom17() []byte {
	c17 := loadChallenge17()
	data := []byte(c17[rand.Int()%len(c17)])
	data = PadBlocks(data, 16, byte(4))
	return CBCencryptConsistent([]byte(data))
}

func decryptRandom17(b []byte) bool {
	decrypted := CBCDecryptConsistent([]byte(b))
	return HasValidPadding(decrypted, 16)
}
