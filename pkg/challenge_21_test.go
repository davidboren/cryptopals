package cryptopals

import (
	"strconv"
	"testing"
)

func printBit(t *testing.T, i int) {
	t.Logf("\nVal: %016b | %v", i, i)
}

func toInt(s string) int {
	v, _ := strconv.ParseInt(s, 2, 64)
	return int(v)
}

func TestChallenge21(t *testing.T) {
	seedMT(0)
	// t.Logf("%v", globalTwister)
	intSet := make(map[int32]struct{})
	for i := 0; i < 100; i++ {
		rand := RandomInt()
		// t.Logf("%v", rand)
		if _, ok := intSet[rand]; ok {
			t.Fatalf("Found dup random number %v after %v tries", rand, i+1)
		}
		intSet[rand] = struct{}{}
	}
	// printBit(t, (1<<5)-1)
	// printBit(t, (2 << 5))
	// printBit(t, (8>>2)-1)
	// printBit(t, (1>>5)^((1<<5)-1))
	// printBit(t, ((1 << 4) - 1))
	// printBit(t, ((1<<4)-1)&toInt("111111010"))
	// printBit(t, toInt("111111010"))
	// printBit(t, 2&2)
	// t.Fail()

}
