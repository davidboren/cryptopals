package cryptopals

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

func findKeySizeScores(fullArr [][]byte) ([]float64, []int) {
	scores := []float64{}
	keySizes := []int{}
	maxKeySize := 0
	for _, a := range fullArr {
		maxKeySize = MaxInt(len(a), maxKeySize)
	}
	for keySize := 2; keySize <= 32 && keySize < maxKeySize; keySize++ {
		scoreTotal := 0
		numComparisons := 0
		allBlocks := [][]byte{}
		for _, a := range fullArr {
			for blockStart, blockEnd := 0, keySize; blockStart < len(a); blockStart, blockEnd = blockStart+keySize, blockEnd+keySize {
				if blockEnd >= len(a) {
					break
				}
				allBlocks = append(allBlocks, a[blockStart:blockEnd])
			}
		}
		for i := 0; i < len(allBlocks)-1; i++ {
			for j := i + 1; j < len(allBlocks); j++ {
				score, err := hamming(allBlocks[i], allBlocks[j])
				if err != nil {
					panic(err)
				}
				scoreTotal += score / keySize
				numComparisons++

			}
		}
		scores = append(scores, float64(scoreTotal)/float64(numComparisons))
		keySizes = append(keySizes, keySize)
	}
	return scores, keySizes
}

func findBestKeySizes(a [][]byte, numBest int) []int {
	scores, keySizes := findKeySizeScores(a)
	if numBest > len(keySizes) {
		numBest = len(keySizes)
	}
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
func GetRepeatingXorCandidates(fullArr [][]byte, numKeySizes int) ([][]byte, []float64) {
	xorKeyList := [][]byte{}
	xorKeyScores := []float64{}
	// if numKeySizes > len(fullArr[0]) {
	// 	numKeySizes = len(fullArr[0])
	// }
	maxLen := 0
	for _, arr := range fullArr {
		if len(arr) > maxLen {
			maxLen = len(arr)
		}
	}
	for _, keySize := range findBestKeySizes(fullArr, numKeySizes) {
		xorKey := []byte{}
		for j := 0; j < keySize; j++ {
			arr2 := []byte{}
			for k := j; k < maxLen; k += keySize {
				for _, arr := range fullArr {
					if k < len(arr) {
						arr2 = append(arr2, arr[k])
					}
				}
			}
			char, _, _ := MostLikelyXorChar(arr2)
			xorKey = append(xorKey, char)
		}
		xorKeyList = append(xorKeyList, xorKey)
		xordArr := []byte{}
		for _, arr := range fullArr {
			xordArr = append(xordArr, RepeatingKeyXor(arr, xorKey)...)
		}
		xorKeyScores = append(xorKeyScores, GetFreqencyLikelihood(string(xordArr)))
	}
	return xorKeyList, xorKeyScores
}

// BreakRepeatingXor decodes Xor using best score from a list of bytelists
func BreakRepeatingXorArrays(arr [][]byte, numKeySizes int) ([]byte, [][]byte) {
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
	xordResult := [][]byte{}
	for _, subArr := range arr {
		xordResult = append(xordResult, RepeatingKeyXor(subArr, bestXorKey))
	}
	return bestXorKey, xordResult
}

// BreakRepeatingXor decodes Xor using best score from a single bytelist
func BreakRepeatingXor(arr []byte, numKeySizes int) ([]byte, []byte) {
	bestXorKey, xoredResult := BreakRepeatingXorArrays([][]byte{arr}, numKeySizes)
	return bestXorKey, xoredResult[0]
}
