package set2

import (
	"github.com/davidboren/cryptopals/pkg/set1"
)

// CBCEncrypt encrypts irregularly sized messages
func CBCEncrypt() {

}

// CBCDecrypt decrypts irregularly sized messages
func CBCDecrypt() {

}

func loadChallenge9() [][]byte {
	loadedFile := set1.FileToArray("/Users/dboren/dev/go/src/github.com/davidboren/cryptopals/data/set2/challenge_10.txt")
	arr := make([][]byte, len(loadedFile))
	for i, str := range loadedFile {
		arr[i] = []byte(str)
	}
	return arr
}
