package cryptopals

import (
	"github.com/stretchr/testify/assert"
	// "strings"
	"testing"
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestChallenge15Strip(t *testing.T) {
	assert.Equal(t, StripPKCS7([]byte("ICE ICE BABY\x04\x04\x04\x04")), []byte("ICE ICE BABY"))
}

func TestChallenge15Panic(t *testing.T) {
	fn := func() {
		StripPKCS7([]byte("ICE ICE BABY\x01\x02\x03\x04"))
	}
	assertPanic(t, fn)
}

func TestChallenge15Panic2(t *testing.T) {
	fn := func() {
		StripPKCS7([]byte("ICE ICE BABY\x05\x05\x05\x05"))
	}
	assertPanic(t, fn)
}
