package cryptopals

// RandomIntGenerator requires a method for generating an int32
type RandomIntGenerator interface {
	genInt() int32
}

type MT19937Container struct {
	generator MersenneTwister
}

func (c *MT19937Container) genInt() int32 {
	return c.generator.RandomInt()
}
