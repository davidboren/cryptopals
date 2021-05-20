package cryptopals

import (
	"errors"
	"strconv"
)

type MersenneTwister struct {
	arr       []int32
	index     int
	seeded    bool
	n         int
	w         int32
	f         int32
	r         int32
	d         int32
	u         int32
	s         int32
	t         int32
	b         int32
	c         int32
	l         int32
	a         int32
	m         int32
	lowerMask int32
	upperMask int32
}

var globalTwister *MersenneTwister

func initDefaultTwister() {
	twister := getMT19937Twister()
	globalTwister = &twister
}

func getMT19937Twister() MersenneTwister {
	n := 624
	var w, m, r int32 = 32, 397, 31
	var u, s, t, l int32 = 11, 7, 15, 18
	a, _ := strconv.ParseInt("9908B0DF", 16, 64)
	d, _ := strconv.ParseInt("FFFFFFFF", 16, 64)
	b, _ := strconv.ParseInt("9D2C5680", 16, 64)
	c, _ := strconv.ParseInt("EFC60000", 16, 64)
	var lowerMask int32 = (1 << r) - 1 // That is, the binary number of r 1's
	var upperMask int32 = ((1 << w) - 1) & (^lowerMask)
	var f int32 = 1812433253
	return MersenneTwister{make([]int32, n), 0, false, n, w, f, r, int32(d), u, s, t, int32(b), int32(c), l, int32(a), m, lowerMask, upperMask}
}

func seedMT(seedValue int32) {
	if globalTwister == nil {
		initDefaultTwister()
	}
	globalTwister.seedMT(seedValue)
}

// Initialize the generator from a seed
func (mt *MersenneTwister) seedMT(seed int32) {
	mt.index = mt.n
	mt.arr[0] = seed
	for i := 1; i < mt.n; i++ { // loop over each element
		mt.arr[i] = ((1 << mt.w) - 1) & (mt.f*(mt.arr[i-1]^(mt.arr[i-1]>>(mt.w-2))) + int32(i))
	}
	mt.seeded = true
}

// Extract a tempered value based on MT[index]
// Generate the next n values from the series x_i
func (mt *MersenneTwister) twist() {
	for i := 0; i < mt.n; i++ {
		x := (mt.arr[i] & mt.upperMask) + (mt.arr[(i+1)%mt.n] & mt.lowerMask)
		xA := x >> 1
		if (x % 2) != 0 { // lowest bit of x is 1
			xA = xA % mt.a
		}
		mt.arr[i] = mt.arr[(int32(i)+mt.m)%int32(mt.n)] ^ xA
	}
	mt.index = 0
}

// calling twist() every n numbers
func (mt *MersenneTwister) RandomInt() int32 {
	if !mt.seeded {
		panic(errors.New("Generator was never seeded"))
	}
	if mt.index == mt.n {
		mt.twist()
	}

	y := mt.arr[mt.index]
	y = y ^ ((y >> mt.u) & mt.d)
	y = y ^ ((y << mt.s) & mt.b)
	y = y ^ ((y << mt.t) & mt.c)
	y = y ^ (y >> mt.l)

	mt.index++
	return ((1 << mt.w) - 1) & (y)
}

func RandomInt() int32 {
	return globalTwister.RandomInt()
}
