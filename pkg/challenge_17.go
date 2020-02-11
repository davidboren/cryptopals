package cryptopals

import (
	"math/rand"
)

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
	return HasValidPadding(decrypted, byte(4))
}
