package cryptopals

import (
	"fmt"
	"math/rand"
	"strings"
)

func sanitizeCookieString(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", ""), "=", "")
}

func mapToCookie(m map[string]string) string {
	sarr := []string{}
	for k, v := range m {
		sarr = append(sarr, sanitizeCookieString(k)+"="+sanitizeCookieString(v))
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

func profileFor(email string) string {
	return mapToCookie(map[string]string{
		"email": email,
		"uid":   fmt.Sprintf("%v", rand.Int()),
		"role":  "user",
	})
}

func decodeChallenge13() []byte {
	return []byte{}
}
