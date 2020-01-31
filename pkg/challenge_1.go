package cryptopals

import (
	b64 "encoding/base64"
	"encoding/hex"
	"errors"
)

// HexToBytes Converts a Hex string to a byteArray
func HexToBytes(hexString string) []byte {
	byteArray, err := hex.DecodeString(hexString)
	if err != nil {
		panic(errors.New("Bad hex string"))
	}
	return byteArray
}

// HexToB64 Converts a Hex string to a base64 encoded string
func HexToB64(hexString string) string {
	byteArray := HexToBytes(hexString)
	bitarrayEncoded := b64.StdEncoding.EncodeToString(byteArray)
	return bitarrayEncoded
}
