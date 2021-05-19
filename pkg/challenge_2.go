package cryptopals

import (
	"encoding/hex"
	"errors"
)

// BytesToHex Converts a byteArray to a Hex string
func BytesToHex(byteArray []byte) string {
	return hex.EncodeToString(byteArray)
}

// MaxInt returns the max of a set of integers
func MaxInt(vals ...int) int {
	if len(vals) == 0 {
		panic(errors.New("Cannot take max of empty set of values"))
	}
	var max_val int
	for i, v := range vals {
		if i == 0 || v > max_val {
			max_val = v
		}
	}
	return max_val
}

// MinInt returns the min of a set of integers
func MinInt(vals ...int) int {
	if len(vals) == 0 {
		panic(errors.New("Cannot take max of empty set of values"))
	}
	var min_val int
	for i, v := range vals {
		if i == 0 || v < min_val {
			min_val = v
		}
	}
	return min_val
}

// XorBytes Xors two bytearrays sequentially
func XorBytes(b1 []byte, b2 []byte) []byte {
	newByteArray := make([]byte, 0, len(b1))
	max_len := MaxInt(len(b1), len(b2))
	for i := 0; i < max_len; i++ {
		newByteArray = append(newByteArray, b1[i%len(b1)]^b2[i%len(b2)])
	}
	return newByteArray
}

// XorHex Xors two hex strings and returns in hex
func XorHex(s1 string, s2 string) string {
	return BytesToHex(XorBytes(HexToBytes(s1), HexToBytes(s2)))
}
