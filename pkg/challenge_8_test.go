package cryptopals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChallenge8(t *testing.T) {
	fullArray := loadChallenge8()
	indices, dupCounts := getHighestDupIndices(fullArray, 16, len(fullArray))
	assert.Equal(t, 132, indices[0])
	assert.Equal(t, 12, dupCounts[0])
	assert.Equal(t, 0, dupCounts[1])
}
