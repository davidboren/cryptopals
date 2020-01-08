package set1

import (
	"math"
	"strings"
)

var minFrequency = 0.0001
var englishOccurrences = map[string]float64{
	"E": 12.0,
	"T": 9.10,
	"A": 8.12,
	"O": 7.68,
	"I": 7.31,
	"N": 6.95,
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

// GetFreqencyLikelihood gets the loglikelihood of a sequence of english letters
func GetFreqencyLikelihood(s string) float64 {
	if len(englishFrequencies) == 0 {
		setLikelihoods()
	}
	s = strings.ToUpper(s)
	countMap := make(map[string]int)
	for k := range englishFrequencies {
		if _, ok := countMap[k]; !ok {
			countMap[k] = strings.Count(s, k)
		}

	}
	logLikelihood := 0.0
	for k, count := range countMap {
		freq, ok := englishFrequencies[k]
		if !ok {
			freq = minFrequency
		}
		logLikelihood = logLikelihood + float64(count)*math.Log(freq)

	}
	return logLikelihood
}

// MostLikelyXorChar returns the character most likely used to xor a given string
// func MostLikelyXorChar(s string) (string, float64, string) {
// 	maxChar := ""
// 	maxLikelihood := 0.0
// 	maxXord := ""
// 	chars := make([]string, 0)
// 	for i := 0; i <= 255; i++ {
// 		chars = append(chars, string([]byte{byte(i)}))
// 	}
// 	for _, char := range chars {
// 		xord := XorSingleChar(s, char)
// 		ll := GetFreqencyLikelihood(xord)
// 		if ll > maxLikelihood {
// 			maxLikelihood = ll
// 			maxChar = char
// 			maxXord = xord
// 		}
// 	}
// 	return maxChar, maxLikelihood, maxXord
// }
