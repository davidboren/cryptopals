package cryptopals

import (
	"fmt"
)

var paddingBytes = []byte{0, 1, 2, 3, 4, 5}

func isPaddingByte(b byte) bool {
	for _, p := range paddingBytes {
		if p == b {
			return true
		}
	}
	return false

}

// HasValidPadding validates padding for a byteArray
func HasValidPadding(b []byte, paddedWith byte) bool {
	isFinal := true
	for i := len(b) - 1; i > 0; i-- {
		if isFinal {
			if isPaddingByte(b[i]) {
				if b[i] != paddedWith {
					return false
				}
			} else {
				isFinal = false
				if b[i] != paddedWith {
					return true
				}
			}
		}
	}
	return true
}

// StripPadding strips byteArray of a padding byte while validating
func StripPadding(b []byte, paddedWith byte) []byte {
	if !HasValidPadding(b, paddedWith) {
		panic(fmt.Errorf("Bytes are padded with %v, not %v", b[len(b)-1], paddedWith))
	}
	isFinal := true
	for i := len(b) - 1; i > 0; i-- {
		if isFinal {
			if !isPaddingByte(b[i]) {
				isFinal = false
				if b[i] != paddedWith {
					return b[:i+1]
				}
			}
		}
	}
	return b
}

// StripPKCS7 strips byteArray of padding byte \x04 while validating
func StripPKCS7(b []byte) []byte {
	paddedWith := byte(4)
	return StripPadding(b, paddedWith)
}
