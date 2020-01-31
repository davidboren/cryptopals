package cryptopals

import (
	"errors"
	// "fmt"
)

// CBCEncrypt encrypts with chained Xor
func CBCEncrypt(data []byte, key []byte, blockSize int, iv []byte) []byte {
	if len(data)%blockSize != 0 {
		panic(errors.New("Your data should be padded prior to encrypting"))
	}
	cypherText := iv
	encrypted := make([]byte, len(data))
	for bs := 0; bs < len(data); bs += blockSize {
		cypherText = EncryptAes128Ecb(RepeatingKeyXor(data[bs:bs+blockSize], cypherText), key, blockSize)
		for i := 0; i < blockSize; i++ {
			encrypted[bs+i] = cypherText[i]
		}
	}
	return encrypted
}

// CBCDecrypt decrypts with chained Xor
func CBCDecrypt(data []byte, key []byte, blockSize int, iv []byte) []byte {
	xorWith := iv
	decrypted := make([]byte, len(data))
	for bs := 0; bs < len(data); bs += blockSize {
		for i, b := range RepeatingKeyXor(DecryptAes128Ecb(data[bs:bs+blockSize], key, blockSize), xorWith) {
			decrypted[bs+i] = b
		}
		xorWith = data[bs : bs+blockSize]
	}
	return decrypted
}

func loadChallenge10() []byte {
	return B64ArrayToBytes(FileToArray("/Users/dboren/dev/go/src/github.com/davidboren/cryptopals/data/set2/challenge_10.txt"))
}
