package cryptopals

import (
	"strings"
)

type challenge16Encryptor struct{}

func (c challenge16Encryptor) GetPrefix() []byte {
	return []byte("comment1=cooking%20MCs;userdata=")
}

func (c challenge16Encryptor) GetSuffix() []byte {
	return []byte(";comment2=%20like%20a%20pound%20of%20bacon")
}

func (c challenge16Encryptor) prepareData(data []byte) []byte {
	dataString := string(data)
	dataString = strings.ReplaceAll(dataString, ";", "\";\"")
	dataString = strings.ReplaceAll(dataString, "=", "\"=\"")
	return append(c.GetPrefix(), append([]byte(dataString), c.GetSuffix()...)...)
}

func (c challenge16Encryptor) Encrypt(data []byte) []byte {
	return CBCencryptConsistent(c.prepareData(data))
}

// CBCencryptConsistent uses a consistent key and iv to encrypt
func CBCencryptConsistent(data []byte) []byte {
	secret := getSecret()
	data = PadBlocks(data, secret.consistentBlockSize, byte('\x04'))
	encrypted := CBCEncrypt(data, secret.consistentKey, secret.consistentBlockSize, secret.consistentIV)
	return encrypted
}

// CBCDecryptConsistent uses a consistent key and iv to decrypt
func CBCDecryptConsistent(data []byte) []byte {
	secret := getSecret()
	encrypted := CBCDecrypt(data, secret.consistentKey, secret.consistentBlockSize, secret.consistentIV)
	return encrypted
}

func isAdmin(s string) bool {
	return strings.Contains(string(CBCDecryptConsistent([]byte(s))), ";admin=true;")
}

func crackChallenge16(e challenge16Encryptor, g StringGenerator) []byte {
	blockSize := getBlockSize(e, g)
	prefixLength := getPrefixLength(e, g, blockSize)
	data := []byte("!admin?true")
	encrypted1 := e.Encrypt(data)
	semiColonXor := byte('!') ^ byte(';')
	equalsXor := byte('?') ^ byte('=')
	semiLoc := prefixLength
	equalsLoc := prefixLength + 6
	encrypted1[semiLoc-blockSize] = semiColonXor ^ encrypted1[semiLoc-blockSize]
	encrypted1[equalsLoc-blockSize] = equalsXor ^ encrypted1[equalsLoc-blockSize]
	return encrypted1
}
