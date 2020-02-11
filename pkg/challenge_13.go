package cryptopals

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type challenge13Encryptor struct{}
type challenge13Generator struct{}

func validateCookieString(s string) {
	if strings.Contains(s, "&") || strings.Contains(s, "=") {
		panic(errors.New("String contains forbidden characters '&' or '='"))
	}
}

func mapToCookie(m map[string]string) string {
	sarr := []string{}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		validateCookieString(k)
		validateCookieString(m[k])
		sarr = append(sarr, k+"="+m[k])
	}
	return strings.Join(sarr, "&")
}

func parseCookie(s string) map[string]string {
	res1 := strings.SplitN(s, "&", -1)
	m := map[string]string{}
	for _, pair := range res1 {
		vals := strings.SplitN(pair, "=", -1)
		m[vals[0]] = vals[1]
	}
	return m
}

func cookieFor(email string) string {
	userData := map[string]string{
		"email": email,
		"uid":   fmt.Sprintf("%v", 1000),
		"role":  "user",
	}
	sarr := []string{}
	for _, k := range []string{"email", "uid", "role"} {
		validateCookieString(k)
		validateCookieString(userData[k])
		sarr = append(sarr, k+"="+userData[k])
	}
	return strings.Join(sarr, "&")
}

func encryptProfile(email string) []byte {
	return challenge13Encryptor{}.Encrypt([]byte(email))
}

func (c challenge13Encryptor) Encrypt(data []byte) []byte {
	return encryptConsistent([]byte(cookieFor(string(data))))
}

func (c challenge13Generator) Generate(i int, j int) ([]byte, int) {
	if i < c.MinGenerationSize() {
		panic(errors.New("Email generation requires a length of at least 7 characters"))
	}
	emailSuffix := "@a.com"
	dataLoc := i - len(emailSuffix) - 1
	data := []byte(strings.Repeat("A", i-len(emailSuffix)) + emailSuffix)
	data[dataLoc] = byte(j)
	return data, dataLoc
}

func (c challenge13Generator) MinGenerationSize() int {
	return 7
}

func getAdminProfile(c ConsistentEncryptor, d DataGenerator) []byte {
	blockSize := getBlockSize(c, d)
	prefixLength := getPrefixLength(c, d, blockSize)
	suffixLength := getSuffixLength(c, d, blockSize)

	// i + prefixLength + suffixLength - 4 == y*blockSize
	// i == y*blockSize - prefixLength - suffixLength + 4

	dataSize := blockSize - prefixLength - suffixLength + 4
	for dataSize < d.MinGenerationSize() {
		dataSize += blockSize
	}
	data, _ := d.Generate(dataSize, 0)
	encrypted := c.Encrypt(data)
	// fullDataLength := len(data) + prefixLength + suffixLength
	// fmt.Printf("\nDataSize: %v, lenData: %v, prefixLength: %v, suffixLength: %v", dataSize, len(data), prefixLength, suffixLength)
	trimmedEncrypted := encrypted[:len(encrypted)-blockSize]
	// fmt.Printf("\ndataSize: %v, fullDataLength: %v, encryptedSize: %v", dataSize, fullDataLength, len(encrypted))

	// fmt.Printf("\ntrimmedDecrypted: %v", string(decryptConsistent(trimmedEncrypted)))

	dataSize2 := blockSize - prefixLength
	for dataSize2 < d.MinGenerationSize() {
		dataSize2 += blockSize
	}
	data2, _ := d.Generate(dataSize2, 0)
	// fmt.Printf("\nDataSize2: %v, lenData2: %v, prefixLength: %v, suffixLength: %v", dataSize2, len(data2), prefixLength, suffixLength)
	adminData := append(data2, []byte("admin")...)
	encryptedAdmin := c.Encrypt(adminData)
	adminStart := len(data2) + prefixLength
	finalAdminBlock := encryptedAdmin[adminStart : adminStart+blockSize]
	// fmt.Printf("\nAdminDecrypted: %v", string(decryptConsistent(finalAdminBlock)))

	// fmt.Printf("\nAdminlen: %v\n", len(encryptedAdmin))
	// fmt.Printf("\ntrimmedEncryptedLen: %v\n", len(trimmedEncrypted))
	return append(trimmedEncrypted, finalAdminBlock...)

}
