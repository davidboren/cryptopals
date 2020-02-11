package cryptopals

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	// "math"
	"math/rand"
	"strings"
	"time"
)

type encryptionSecret struct {
	consistentKey       []byte
	consistentBlockSize int
	consistentIV        []byte
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

// ConsistentEncryptor provides a simple encryption interface using internal settings
type ConsistentEncryptor interface {
	Encrypt(data []byte) []byte
	// Decrypt(data []byte) []byte
}

// DataGenerator creates byte arrays with a length input (i) and a variation input (j) and returns the variation index
type DataGenerator interface {
	Generate(i int, j int) ([]byte, int)
	MinGenerationSize() int
}

func getSecret() *encryptionSecret {
	if secret == nil {
		rand.Seed(time.Now().UnixNano())
		// keySizes := []int{16, 24, 32}
		secret = &encryptionSecret{
			consistentKey: randomAESKey(),
			// Harcoded in cypher/aes
			consistentBlockSize: 16,
			consistentIV:        randBytes(16),
		}
	}
	return secret
}

func encryptConsistent(data []byte) []byte {
	secret := getSecret()
	data = PadBlocks(data, secret.consistentBlockSize, byte('\x04'))
	encrypted := EncryptAes128Ecb(data, secret.consistentKey, secret.consistentBlockSize)
	return encrypted
}

func decryptConsistent(data []byte) []byte {
	secret := getSecret()
	return DecryptAes128Ecb(data, secret.consistentKey, secret.consistentBlockSize)
}

func getBlockSize(c ConsistentEncryptor, d DataGenerator) int {
	blockSize := 0
	minSize := d.MinGenerationSize()
	lastEncryptedSize := 0
	for i := minSize; i <= 128; i++ {
		data, _ := d.Generate(i, 0)
		encrypted := c.Encrypt(data)
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

func getPrefixLength(c ConsistentEncryptor, d DataGenerator, blockSize int) int {
	minSize := d.MinGenerationSize()
	data, changeIndex := d.Generate(minSize, 0)
	data2, changeIndex2 := d.Generate(minSize, 1)
	if changeIndex != changeIndex2 {
		panic(fmt.Errorf("Bad data generation.  The change index should be identical for inputs %v, %v", 0, 1))
	}
	encryptedData := c.Encrypt(data)
	encryptedData2 := c.Encrypt(data2)
	for j := 0; j < len(encryptedData); j += blockSize {
		if !blocksEqual(encryptedData[j:j+blockSize], encryptedData2[j:j+blockSize]) {
			for k := minSize; k <= blockSize+minSize; k++ {
				data, changeIndex = d.Generate(k, 0)
				data2, changeIndex2 = d.Generate(k, 1)
				encryptedData = c.Encrypt(data)
				encryptedData2 = c.Encrypt(data2)
				// fmt.Printf("\n    ESize: %v, Size: %v, Data: %v, Data2: %v, J: %v, ChangeIndex: %v, LE1: %v, LE2: %v\n    EData:  %v\n    EData2: %v", len(data), k, data, data2, j, changeIndex, len(encryptedData), len(encryptedData2), encryptedData[j:j+blockSize], encryptedData2[j:j+blockSize])
				if blocksEqual(encryptedData[j:j+blockSize], encryptedData2[j:j+blockSize]) {
					// fmt.Printf("\nEntered")
					return j + blockSize - changeIndex
				}
			}
		}
	}
	panic(errors.New("Could not determine prefix length"))
}

func getSuffixLength(c ConsistentEncryptor, d DataGenerator, blockSize int) int {
	minSize := d.MinGenerationSize()
	prefixLength := getPrefixLength(c, d, blockSize)
	prevBlocks := 0
	for i := minSize; i < blockSize+minSize+1; i++ {
		data, _ := d.Generate(i, 0)
		arr := c.Encrypt(data)
		numBlocks := len(arr) / blockSize
		if prevBlocks != 0 && numBlocks == prevBlocks+1 {
			return len(arr) - blockSize - len(data) - prefixLength + 1

		}
		prevBlocks = numBlocks
	}
	return 0
}

func getData(c ConsistentEncryptor, d DataGenerator, blockSize int) (int, [][]byte, [][]byte) {
	minSize := d.MinGenerationSize()
	encryptedArrays := make([][]byte, blockSize)
	dataArrays := make([][]byte, blockSize)
	prefixLength := getPrefixLength(c, d, blockSize)
	for i := minSize; i < blockSize+minSize; i++ {
		data, _ := d.Generate(i, 0)
		finalLen := len(data) + prefixLength
		loc := blockSize - finalLen%blockSize - 1
		dataArrays[loc] = data
		encryptedArrays[loc] = c.Encrypt(dataArrays[loc])
	}
	return prefixLength, dataArrays, encryptedArrays
}

func decryptMultiCallECB(c ConsistentEncryptor, d DataGenerator) []byte {
	largeData, _ := d.Generate(100, 0)
	blockSize := getBlockSize(c, d)
	isECB := DetectECB(c.Encrypt(largeData), blockSize, 1)
	if !isECB {
		panic(errors.New("Cannot detect ECB"))
	}

	prefixLength, dataArrays, encryptedArrays := getData(c, d, blockSize)

	msgLength := len(encryptedArrays[blockSize-1]) - len(dataArrays[blockSize-1]) - prefixLength
	messageBytes := make([]byte, msgLength)

	for j := 0; j < int((msgLength)/blockSize); j++ {
		for i, arr := range encryptedArrays {
			for b := byte(0); b < 255; b++ {
				byteList := []byte{b}
				newMsg := append(dataArrays[i], messageBytes[:j*blockSize+i]...)
				newData := append(newMsg, byteList...)
				encrypted := c.Encrypt(newData)
				finalPrefixLength := len(newMsg) + prefixLength
				blockStart := finalPrefixLength - finalPrefixLength%blockSize
				blockEnd := blockStart + blockSize
				// fmt.Printf("Block Start: %v, End: %v\n", blockStart, blockEnd)
				if blocksEqual(encrypted[blockStart:blockEnd], arr[blockStart:blockEnd]) {
					messageBytes[j*blockSize+i] = b
				}
			}
		}

	}
	return messageBytes
}

type challenge12Encryptor struct{}

// StringGenerator is used to create a string with char "A" of length i
type StringGenerator struct{}

func (c challenge12Encryptor) GetPrefix() []byte {
	return []byte("AAAAAA1234")
}

func (c challenge12Encryptor) GetSuffix() []byte {
	decoded, _ := b64.StdEncoding.DecodeString(appendText)
	return decoded

}

func (c challenge12Encryptor) Encrypt(data []byte) []byte {
	return encryptConsistent(append(c.GetPrefix(), append(data, c.GetSuffix()...)...))
}

// Generate creates a string with char "A" of length i, and a variation of the last character according to j
func (s StringGenerator) Generate(i int, j int) ([]byte, int) {
	if i < s.MinGenerationSize() {
		panic(errors.New("String generation requires a length of at least 1 characters"))
	}
	changeLoc := i - s.MinGenerationSize()
	data := []byte(strings.Repeat("A", i))
	data[changeLoc] = byte(j)
	return data, changeLoc
}

// MinGenerationSize creates a string with char "A" of length i, and a variation of the last character according to j
func (s StringGenerator) MinGenerationSize() (i int) {
	return 7
}
