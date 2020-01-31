package cryptopals

import (
	b64 "encoding/base64"
	"errors"
	"math/rand"
	"strings"
	"time"
)

type encryptionSecret struct {
	consistentKey       []byte
	consistentBlockSize int
}

var secret *encryptionSecret

var appendText string = `
Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`

func randomAESKey() []byte {
	keySizes := []int{16, 24, 32}
	return randBytes(keySizes[rand.Int()%3])
}

func getSecret() *encryptionSecret {
	if secret == nil {
		rand.Seed(time.Now().UnixNano())
		// keySizes := []int{16, 24, 32}
		secret = &encryptionSecret{
			consistentKey: randomAESKey(),
			// Harcoded in cypher/aes
			consistentBlockSize: 16,
		}
	}
	return secret
}

func encryptConsistent(data []byte) []byte {
	secret := getSecret()
	decoded, _ := b64.StdEncoding.DecodeString(appendText)
	data = PadBlocks(AddBytes(data, decoded), secret.consistentBlockSize, byte('\x04'))
	encrypted := EncryptAes128Ecb(data, secret.consistentKey, secret.consistentBlockSize)
	return encrypted
}

func getBlockSize() int {
	blockSize := 0
	lastEncryptedSize := 0
	for i := 1; i <= 64; i++ {
		data := []byte(strings.Repeat("A", i))
		encrypted := encryptConsistent(data)
		if len(encrypted) > lastEncryptedSize+1 && lastEncryptedSize != 0 {
			blockSize = len(encrypted) - lastEncryptedSize
			break
		}
		lastEncryptedSize = len(encrypted)
	}
	if blockSize == 0 {
		panic(errors.New("found no non-zero blockSize for encrypted data"))
	}
	return blockSize
}

func decryptMultiCallECB(dataArrays [][]byte, encryptedArrays [][]byte, blockSize int, msgLength int) []byte {
	messageBytes := make([]byte, msgLength)
	for j := 0; j < int((msgLength)/blockSize); j++ {
		for i, arr := range encryptedArrays {
			for b := byte(0); b < 255; b++ {
				byteList := []byte{b}
				newMsg := AddBytes(dataArrays[i], messageBytes[:j*blockSize+i])
				newData := AddBytes(newMsg, byteList)
				encrypted := encryptConsistent(newData)
				if blocksEqual(encrypted[j*blockSize:(j+1)*blockSize], arr[j*blockSize:(j+1)*blockSize]) {
					messageBytes[j*blockSize+i] = b
				}
			}
		}

	}
	return messageBytes
}

func decodeChallenge12() []byte {
	blockSize := getBlockSize()
	encryptedArrays := make([][]byte, blockSize)
	dataArrays := make([][]byte, blockSize)
	isECB := DetectECB(encryptConsistent([]byte(strings.Repeat("A", 100))), blockSize, 1)
	if !isECB {
		panic(errors.New("Cannot detect ECB"))
	}

	for i := blockSize - 1; i >= 0; i-- {
		dataArrays[blockSize-i-1] = []byte(strings.Repeat("A", i))
		encryptedArrays[blockSize-i-1] = encryptConsistent(dataArrays[blockSize-i-1])
	}
	msgLength := len(encryptedArrays[blockSize-1])
	return decryptMultiCallECB(dataArrays, encryptedArrays, blockSize, msgLength)
}
