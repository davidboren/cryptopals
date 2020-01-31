package cryptopals

import (
	"fmt"
)

// AddBytes combines two bytearrays into one
func AddBytes(b1 []byte, b2 []byte) []byte {

	newBytes := make([]byte, len(b1)+len(b2))
	for i := 0; i < len(b1); i++ {
		newBytes[i] = b1[i]
	}
	for i := 0; i < len(b2); i++ {
		newBytes[len(b1)+i] = b2[i]
	}
	return newBytes
}

// Pad a block with a specific number of padding bytes
func Pad(block []byte, padding int, padWith byte) []byte {
	if len(block) == padding {
		return block
	}
	if len(block) > padding {
		panic(fmt.Errorf("Your block of size '%v' is of greater length than your padding length '%v'", len(block), padding))
	}
	newBlock := make([]byte, padding)
	for i := 0; i < len(block); i++ {
		newBlock[i] = block[i]

	}
	for i := len(block); i < padding; i++ {
		newBlock[i] = padWith

	}
	return newBlock
}

// PadBlocks rounds out an array of blocks with padding
func PadBlocks(b []byte, blockSize int, padWith byte) []byte {
	if len(b)%blockSize == 0 {
		return b
	}
	lastBlockStart := int(len(b)/blockSize) * blockSize
	return AddBytes(b[:lastBlockStart], Pad(b[lastBlockStart:], blockSize, padWith))
}
