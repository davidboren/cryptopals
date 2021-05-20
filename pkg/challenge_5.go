package cryptopals

//RepeatingKeyXor xors byte arrays using a key bytearray in a loop
func RepeatingKeyXor(b1 []byte, key []byte) []byte {
	xord := make([]byte, len(b1))
	for i, b := range b1 {
		xord[i] = b ^ key[i%len(key)]
	}
	return xord
}

//RepeatingKeyXorStrings xors string using a key string in a loop
func RepeatingKeyXorStrings(toEncrpt string, key string) string {
	b1 := []byte(toEncrpt)
	b2 := []byte(key)
	return BytesToHex(RepeatingKeyXor(b1, b2))
}
