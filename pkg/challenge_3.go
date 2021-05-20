package cryptopals

import (
	"fmt"
	"math"
	"strings"
)

var minFrequency = 0.000000000000000001
var englishOccurrences = map[string]float64{
	"E": 12.0,
	"T": 9.10,
	"A": 8.12,
	"O": 7.68,
	"I": 7.31,
	"N": 6.95,
	" ": 6,
	"S": 6.28,
	"R": 6.02,
	"H": 5.92,
	"D": 4.32,
	"L": 3.98,
	"U": 2.88,
	"C": 2.71,
	"M": 2.61,
	"F": 2.30,
	"Y": 2.11,
	"W": 2.09,
	"G": 2.03,
	"P": 1.82,
	"B": 1.49,
	"V": 1.11,
	"K": 0.69,
	"X": 0.17,
	"Q": 0.11,
	"J": 0.10,
	"Z": 0.07,
	",": 0.25,
	"!": 0.05,
	".": 1.45,
	"'": 1.45,
	"/": 0.05,
	"?": 0.15,
}
var englishFrequencies = make(map[string]float64)

func setLikelihoods() {
	fullRate := 0.0
	for _, rate := range englishOccurrences {
		fullRate = fullRate + rate
	}
	for k, rate := range englishOccurrences {
		englishFrequencies[k] = rate / fullRate

	}
}

// GetCounts gets the counts of characters in a string
func GetCounts(s string) map[string]int {
	s = strings.ToUpper(s)
	countMap := make(map[string]int)
	for _, char := range s {
		if _, ok := countMap[string(char)]; !ok {
			countMap[string(char)] = strings.Count(s, string(char))
		}

	}
	return countMap
}

// GetFreqencyLikelihood gets the loglikelihood of a sequence of english letters
func GetFreqencyLikelihood(s string) float64 {
	if len(englishFrequencies) == 0 {
		setLikelihoods()
	}
	likelihood := 0.0
	countMap := GetCounts(s)
	for k, count := range countMap {
		freq, ok := englishFrequencies[k]
		if !ok {
			freq = minFrequency
		}
		// likelihood -= math.Sqrt(math.Pow((float64(count)/float64(len(s)) - freq), 2))
		likelihood += float64(count) * math.Log(freq)

	}
	return likelihood
}

// MostLikelyXorChar returns the character most likely used to xor a given string
func MostLikelyXorChar(b []byte) (byte, float64, []byte) {
	v1 := []byte(string("I"))
	v2 := []byte(string("C"))
	v3 := []byte(string("F"))
	b1 := v1[0] ^ b[0]
	b2 := v2[0] ^ b[1]
	b3 := v3[0] ^ b[2]
	c1 := b1 ^ b[0]
	c2 := b2 ^ b[1]
	c3 := b3 ^ b[2]
	secret := getSecret()
	fmt.Sprintf("b1: %v, b2: %v, b3: %v, Secret: %v", b1, b2, b3, secret)
	fmt.Sprintf("c1: %v, c2: %v, c3: %v, Secret: %v", c1, c2, c3, secret)

	var maxChar byte
	maxLikelihood := 0.0
	maxXord := []byte{}
	for i := 0; i < 255; i++ {
		if byte(i) == b3 {
			fmt.Sprintf("c1: %v, c2: %v, c3: %v, Secret: %v", c1, c2, c3, secret)

		}
		byteChar := byte(i)
		xord := XorBytes(b, []byte{byteChar})
		ll := GetFreqencyLikelihood(string(xord))
		if i == 0 || ll > maxLikelihood {
			maxLikelihood = ll
			maxChar = byteChar
			maxXord = xord
		}
	}
	return maxChar, maxLikelihood, maxXord
}
