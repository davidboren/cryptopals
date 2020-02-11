package cryptopals

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseCookie(t *testing.T) {
	cookie := "foo=bar&baz=qux&zap=zazzle"
	cMap := map[string]string{
		"foo": "bar",
		"baz": "qux",
		"zap": "zazzle",
	}
	assert.Equal(t, cMap, parseCookie(cookie))
	assert.Equal(t, cMap, parseCookie(mapToCookie(cMap)))
}

func TestCookieFor(t *testing.T) {
	cMap := map[string]string{
		"email": "bar",
		"uid":   "1000",
		"role":  "user",
	}
	cookie := fmt.Sprintf("email=%v&uid=%v&role=%v", cMap["email"], cMap["uid"], cMap["role"])
	assert.Equal(t, cookie, cookieFor(cMap["email"]))
}

func TestGeneration(t *testing.T) {
	gen := challenge13Generator{}
	for i := gen.MinGenerationSize(); i < gen.MinGenerationSize()+16*2; i++ {
		data, _ := gen.Generate(i, 0)
		assert.Equal(t, i, len(data))

	}
}

func TestPrefixLength13(t *testing.T) {
	assert.Equal(t, len("email="), getPrefixLength(challenge13Encryptor{}, challenge13Generator{}, 16))
}

func TestSuffixLength13(t *testing.T) {
	// gen, _ := challenge13Generator{}.Generate(10, 0)
	// t.Logf("Stuff: %v", string(decryptConsistent(challenge13Encryptor{}.Encrypt(gen))))
	assert.Equal(t, len("&uid=1000&role=user"), getSuffixLength(challenge13Encryptor{}, challenge13Generator{}, 16))
}

func TestAdminProfile(t *testing.T) {
	adminEncrypted := getAdminProfile(challenge13Encryptor{}, challenge13Generator{})
	// t.Logf("\nlen: %v\n", len(adminEncrypted))
	// t.Logf("\nDecrypted: %v", string(decryptConsistent(adminEncrypted)))
	assert.True(
		t,
		strings.Contains(string(decryptConsistent(adminEncrypted)), ".com&uid=1000&role=admin"),
	)
}
