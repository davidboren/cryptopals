package set1

import (
	b64 "encoding/base64"
	"errors"
	"sort"
	"strings"
)

type scoredKeySize struct {
	KeySize int
	Score   float64
}

func byteArraysToStrings(byteArr [][]byte) []string {
	arr := make([]string, len(byteArr))
	for i, b := range byteArr {
		arr[i] = string(b)
	}
	return arr
}

func hamming(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Bytes are not the same length")
	}

	diff := 0
	for i, aByte := range a {
		bByte := b[i]
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if (aByte & mask) != (bByte & mask) {
				diff++
			}

		}

	}
	return diff, nil
}

func hammingStrings(s1, s2 string) (int, error) {
	return hamming([]byte(s1), []byte(s2))
}

func findBestKeySizes(a []byte, numBest int) []int {
	scores, keySizes := findKeySizeScores(a)
	scoredKeySizes := make([]scoredKeySize, len(keySizes))
	for i, ks := range keySizes {
		scoredKeySizes[i] = scoredKeySize{KeySize: ks, Score: scores[i]}
	}
	sort.Slice(scoredKeySizes, func(i, j int) bool {
		return scoredKeySizes[i].Score < scoredKeySizes[j].Score
	})
	finalKeySizes := make([]int, numBest)
	for i := 0; i < numBest; i++ {
		finalKeySizes[i] = scoredKeySizes[i].KeySize
	}
	return finalKeySizes
}

func findKeySizeScores(a []byte) ([]float64, []int) {
	scores := []float64{}
	keySizes := []int{}
	for keySize := 2; keySize < 40 && 2*keySize <= len(a); keySize++ {
		numBlocks := int(len(a)/keySize) - 1
		scoreTotal := 0
		for i := 0; i < numBlocks; i++ {
			score, err := hamming(a[i*keySize:(i+1)*keySize], a[(i+1)*keySize:(i+2)*keySize])
			if err != nil {
				panic(err)
			}
			scoreTotal += score

		}
		scores = append(scores, float64(scoreTotal)/float64(numBlocks)/float64(keySize))
		keySizes = append(keySizes, keySize)
	}
	return scores, keySizes
}

// B64ArrayToBytes Converts array of b64 encoded strings to array bytes
func B64ArrayToBytes(rawArray []string) []byte {
	decoded, err := b64.StdEncoding.DecodeString(strings.Join(rawArray, ""))
	if err != nil {
		panic(err)
	}
	return decoded
}

func loadChallenge6() []byte {
	return B64ArrayToBytes(FileToArray("/Users/dboren/dev/go/src/github.com/davidboren/cryptopals/data/set1/challenge_6.txt"))
}

// GetRepeatingXorCandidates returns a list of likely repeating xor keys and their scores
func GetRepeatingXorCandidates(arr []byte, numKeySizes int) ([][]byte, []float64) {
	xorKeyList := [][]byte{}
	xorKeyScores := []float64{}
	for _, keySize := range findBestKeySizes(arr, numKeySizes) {
		xorKey := []byte{}
		for j := 0; j < keySize; j++ {
			arr2 := []byte{}
			k := j
			for k < len(arr) {
				arr2 = append(arr2, arr[k])
				k += keySize
			}
			char, _, _ := MostLikelyXorChar(arr2)
			xorKey = append(xorKey, char)
		}
		xorKeyList = append(xorKeyList, xorKey)
		xorKeyScores = append(xorKeyScores, GetFreqencyLikelihood(string(RepeatingKeyXor(arr, xorKey))))
	}
	return xorKeyList, xorKeyScores

}

// DecodeRepeatingXor decode Xor using best score
func DecodeRepeatingXor(arr []byte, numKeySizes int) ([]byte, []byte) {
	var bestXorKey []byte
	var maxScore float64
	xorKeyList, xorKeyScores := GetRepeatingXorCandidates(arr, numKeySizes)
	isFirst := true
	for i, score := range xorKeyScores {
		if isFirst || score > maxScore {
			maxScore = score
			bestXorKey = xorKeyList[i]
			isFirst = false
		}
	}
	return bestXorKey, RepeatingKeyXor(arr, bestXorKey)
}
