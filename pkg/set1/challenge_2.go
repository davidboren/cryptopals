package set1

import (
	"encoding/hex"
	"errors"
)

// BytesToHex Converts a byteArray to a Hex string
func BytesToHex(byteArray []byte) string {
	return hex.EncodeToString(byteArray)
}

// XorBytes Converts a Hex string to a byteArray
func XorBytes(b1 []byte, b2 []byte) []byte {
	if len(b1) != len(b2) {
		panic(errors.New("Bytes must be the same length"))
	}
	newByteArray := make([]byte, 0, len(b1))
	for i, b := range b1 {
		newByteArray = append(newByteArray, b^b2[i])
	}
	return newByteArray
}

// XorHex Xors two hex strings and returns in hex
func XorHex(s1 string, s2 string) string {
	return BytesToHex(XorBytes(HexToBytes(s1), HexToBytes(s2)))
}

// XorSingleByte Xors a string against a single character
func XorSingleByte(b1 []byte, b2 byte) []byte {
	repeated := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		repeated[i] = b2
	}
	return XorBytes(b1, repeated)
}
