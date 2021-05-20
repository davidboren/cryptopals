package cryptopals

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChallenge22(t *testing.T) {
	var seedSearchDays int32 = 5
	rand1 := getMT19937Twister()
	rand2 := getMT19937Twister()
	randomSeed := int32(time.Now().Unix())
	rand1.seedMT(randomSeed)

	firstInt := rand1.RandomInt()
	var detectedSeed int32 = 0

	startSeed := int32(time.Now().Unix()) - seedSearchDays*24*3600
	j := 0
	for i := startSeed; i < int32(time.Now().Unix()); i++ {
		rand2.seedMT(int32(i))
		newFirstInt := rand2.RandomInt()
		if firstInt == newFirstInt {
			detectedSeed = int32(i)

		}
		j++
	}
	assert.Equal(t, detectedSeed, randomSeed)

}
