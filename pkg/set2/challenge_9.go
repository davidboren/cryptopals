package set2

import (
	"fmt"
)

// Pad a block with a specific number of padding bytes
func Pad(block []byte, padding int) []byte {
	if len(block) > padding {
		panic(fmt.Errorf("Your block of size '%v' is of greater length than your padding length '%v'", len(block), padding))
	}
	newBlock := make([]byte, padding)
	for i := 0; i < len(block); i++ {
		newBlock[i] = block[i]

	}
	for i := len(block); i < padding; i++ {
		newBlock[i] = byte('\x04')

	}
	return newBlock
}
