package main

import "strings"
import "math/rand"

// TODO: use some more optimal representation of a bitset
type Chromosome struct {
	Bits []bool
}

func (c *Chromosome) String() string {
	var s []string
	for i := 0; i < len(c.Bits); i++ {
		if c.Bits[i] {
			s = append(s, "1")
		} else {
			s = append(s, "0")
		}
	}
	return strings.Join(s, "")
}

func (c *Chromosome) Randomize(r *rand.Rand, len int) []bool {
	c.Bits = make([]bool, len)
	for i := 0; i < len; i++ {
		if r.Intn(2) == 1 {
			c.Bits[i] = true
		}
	}
	return c.Bits
}
