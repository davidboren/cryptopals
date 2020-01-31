package cryptopals

import (
	"fmt"
	"reflect"
	"sort"
)

// Slice is used for sorting arrays of arbitrary types and retrieving the sorted indices
type Slice struct {
	sort.Interface
	idx    []int
	values []int
}

// Swap allows the sort
func (s Slice) Swap(i, j int) {
	s.Interface.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

// NewSlice creates a new slice from a sort interface
func NewSlice(n sort.Interface) *Slice {
	s := &Slice{Interface: n, idx: make([]int, n.Len())}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}

// NewIntSlice creates a new Int slice
func NewIntSlice(n ...int) *Slice {
	return NewSlice(sort.IntSlice(n))
}

// NewFloat64Slice creates a new Float64 slice
func NewFloat64Slice(n ...float64) *Slice { return NewSlice(sort.Float64Slice(n)) }

// NewStringSlice creates a new String slice
func NewStringSlice(n ...string) *Slice { return NewSlice(sort.StringSlice(n)) }

// ReverseSlice reverses a slice
func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func hexArrayToBytes(sarr []string) [][]byte {
	arr := make([][]byte, len(sarr))
	for i, s := range sarr {
		arr[i] = HexToBytes(s)
	}
	return arr
}

func blocksEqual(b1 []byte, b2 []byte) bool {
	if len(b1) != len(b2) {
		panic(fmt.Errorf("Blocks 1 and 2 have lengths '%v' and '%v', respectively, which are not equal", len(b1), len(b2)))
	}
	allEqual := true
	for k, b := range b1 {
		if b != b2[k] {
			allEqual = false
			break
		}

	}
	return allEqual
}

func getDupCount(b []byte, blockSize int) int {
	count := 0
	totBlocks := int(len(b) / blockSize)
	for i := 0; i < totBlocks-1; i++ {
		block := b[i*blockSize : (i+1)*blockSize]
		for j := i + 1; j < totBlocks; j++ {
			compBlock := b[j*blockSize : (j+1)*blockSize]
			if blocksEqual(block, compBlock) {
				count++
			}

		}
	}
	return count
}

func countDupBlocks(arr [][]byte, blockSize int) []int {
	dupCounts := make([]int, len(arr))
	for l, b := range arr {
		if len(b)%blockSize != 0 {
			panic(fmt.Errorf("Line number '%v' does has numBytes '%v', which is not a multiple of blockSize '%v'", l, len(b), blockSize))

		}
		dupCounts[l] = getDupCount(b, blockSize)
	}
	return dupCounts

}

func getHighestDupIndices(arr [][]byte, blockSize int, numIndices int) ([]int, []int) {
	counts := countDupBlocks(arr, blockSize)
	s := NewIntSlice(counts...)
	sort.Sort(s)
	highest := s.idx[len(s.idx)-numIndices:]
	finalCounts := counts[len(counts)-numIndices:]
	ReverseSlice(highest)
	ReverseSlice(finalCounts)
	return highest, finalCounts

}

func loadChallenge8() [][]byte {
	return hexArrayToBytes(FileToArray("/Users/dboren/dev/go/src/github.com/davidboren/cryptopals/data/set1/challenge_8.txt"))
}
