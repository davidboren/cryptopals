package set1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChallenge8(t *testing.T) {
	fullArray := loadChallenge8()
	indices, dupCounts := getHighestDupIndices(fullArray, 16, len(fullArray))
	assert.Equal(t, 132, indices[0])
	assert.Equal(t, 15, dupCounts[0])
	assert.Equal(t, 9, dupCounts[1])
}
