// Some reading: https://www.quora.com/What-is-the-Pareto-dominance-concept-in-multi-objective-optimization

package main

import (
	"math"
	"reflect"
)

type Comparator func(a, b Solution) int

type NondominatedPopulation struct {
	comparator Comparator
	Population []Solution
}

func NewNondominatedPopulation(c Comparator) NondominatedPopulation {
	return NondominatedPopulation{comparator: c}
}

func (p *NondominatedPopulation) AddAll(solutions []Solution) {
	for _, s := range solutions {
		p.Add(s)
	}
}

func (p *NondominatedPopulation) Add(sol Solution) bool {
	//fmt.Println("-----")
	for i := range p.Population {
		s := p.Population[i]
		flag := p.comparator(sol, s)
		//fmt.Printf("%v %v %d\n", s, sol, flag)
		if flag < 0 {
			p.Population = remove(p.Population, i)
		} else if flag > 0 {
			return false
		} else if isDuplicate(s, sol) {
			return false
		}
	}
	p.Population = append(p.Population, sol)
	return true
}

func remove(s []Solution, i int) []Solution {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

const EPS = 1e-10

func isDuplicate(a, b Solution) bool {
	return reflect.DeepEqual(a.Chromosome.Bits, b.Chromosome.Bits) || distance(a, b) < EPS
}

func distance(a, b Solution) float64 {
	distance := 0.0
	for i := 0; i < NumberOfObjectives; i++ {
		distance += math.Pow(a.Objectives[i]-b.Objectives[i], 2.0)
	}
	return math.Sqrt(distance)
}
