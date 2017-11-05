package main

import "math/rand"

func (c *Chromosome) mutate(r *rand.Rand, threshold int) []bool {
	counter := 0
	for i := 0; i < len(c.Bits); i++ {
		if r.Intn(100) < threshold {
			c.Bits[i] = !c.Bits[i]
			counter++
		}
	}
	//fmt.Printf("Mutation changed %d bits\n", counter)
	return c.Bits
}

func crossover(r *rand.Rand, p1, p2 Chromosome) Chromosome {
	point := r.Intn(len(p1.Bits))
	//fmt.Printf("Crossover at position: %d\n", point)
	var bits []bool
	bits = append(bits, p1.Bits[0:point]...)
	bits = append(bits, p2.Bits[point:]...)
	//old broken code that was mutating parents:
	//bits := append(p1.bits[0:point], p2.bits[point:]...)
	return Chromosome{bits}
}
