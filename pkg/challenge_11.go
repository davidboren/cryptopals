package cryptopals

import (
	"math/rand"
	"time"
)

func randBytes(numBytes int) []byte {
	rand.Seed(time.Now().UnixNano())
	output := make([]byte, numBytes)
	rand.Read(output)
	return output
}

// RandEncrypt chooses a random encryption for a block of data
func RandEncrypt(data []byte, blockSize int) ([]byte, bool) {
	prefixCount := 5 + rand.Int()%5
	suffixCount := 5 + rand.Int()%5
	data = AddBytes(randBytes(prefixCount), data)
	data = AddBytes(data, randBytes(suffixCount))
	data = PadBlocks(data, blockSize, byte('\x04'))
	key := randBytes(16)
	iv := randBytes(16)
	mode := rand.Int() % 2
	var encrypted []byte
	if mode == 0 {
		encrypted = CBCEncrypt(data, key, blockSize, iv)
	} else {
		encrypted = EncryptAes128Ecb(data, key, blockSize)
	}
	return encrypted, mode == 1

}

// DetectECB detects AES
func DetectECB(data []byte, blockSize int, countThreshold int) bool {
	counts := getDupCount(data, blockSize)
	if counts >= countThreshold {
		return true
	}
	return false
}
