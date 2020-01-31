package cryptopals

import (
	"github.com/stretchr/testify/assert"
	// "strings"
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
