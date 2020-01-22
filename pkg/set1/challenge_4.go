package set1

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func FileToArray(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := make([]string, 0)
	for scanner.Scan() {
		arr = append(arr, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return arr
}

// MostLikelyXorLine gets the line with the best xor score/character
func MostLikelyXorLine() (int, byte, float64, string) {
	var maxScore float64
	var bestLineNumber int
	var bestLineOutput string
	var bestLineChar byte
	isFirst := true
	for i, line := range loadChallenge4() {
		byteChar, score, xordLine := MostLikelyXorChar(HexToBytes(line))
		if isFirst || score > maxScore {
			maxScore = score
			bestLineNumber = i
			bestLineOutput = string(xordLine)
			bestLineChar = byteChar
			isFirst = false
		}
	}
	return bestLineNumber, bestLineChar, maxScore, bestLineOutput
}
func loadChallenge4() []string {
	return FileToArray("/Users/dboren/dev/go/src/github.com/davidboren/cryptopals/data/set1/challenge_4.txt")
}
