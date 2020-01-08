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

// XorSingleChar Xors a string against a single character
func XorSingleChar(s1 string, c string) string {
	b1 := HexToBytes(s1)
	b2 := HexToBytes(c)
	if len(b2) != 1 {
		panic(errors.New("Character string must be one byte in length"))
	}
	newByteArray := make([]byte, 0, len(b1))
	for i := range b1 {
		newByteArray[i] = b2[0]
	}
	return string(XorBytes(b1, newByteArray))
}
