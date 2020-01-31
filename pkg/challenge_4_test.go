package cryptopals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMostLikelyLine(t *testing.T) {
	bestLineNumber, bestLineChar, maxScore, bestLineOutput := MostLikelyXorLine()
	assert.Equal(t, "35", BytesToHex([]byte{bestLineChar}))
	assert.Equal(t, 170, bestLineNumber)
	// assert.InDelta(t, -108.99890802338807, maxScore, 0.000000001)
	assert.InDelta(t, -0.5213991281208517, maxScore, 0.000000001)
	assert.Equal(t, "Now that the party is jumping\n", bestLineOutput)
}
