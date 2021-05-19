package cryptopals

import (
	"fmt"
)

func isPaddingByte(b byte, blockSize int) bool {
	for i := 1; i <= blockSize; i++ {
		if byte(i) == b {
			return true
		}
	}
	return false

}

// HasValidPadding validates padding for a byteArray
func HasValidPadding(b []byte, blockSize int) bool {
	paddedWith := b[len(b)-1]
	if !isPaddingByte(paddedWith, blockSize) {
		return true
	}
	requiredPaddings := int(paddedWith)
	for i := 0; i < requiredPaddings; i++ {
		if b[len(b)-1-i] != paddedWith {
			return false
		}
	}
	return true
}

// StripPadding strips byteArray of a padding byte while validating
func StripPadding(b []byte, blockSize int) []byte {
	paddedWith := b[len(b)-1]
	if !HasValidPadding(b, blockSize) {
		panic(fmt.Errorf("Bytes are not padded correctly in block"))
	}
	if !isPaddingByte(paddedWith, blockSize) {
		return b
	}
	return b[:len(b)-int(paddedWith)]
}

// StripPKCS7 strips byteArray of padding byte \x04 while validating
func StripPKCS7(b []byte, blockSize int) []byte {
	return StripPadding(b, blockSize)
}
