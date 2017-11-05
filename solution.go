package main

import "fmt"

type Solution struct {
	Chromosome  Chromosome
	Constraints []float64
	Objectives  []float64
}

func NewSolution(ch Chromosome) Solution {
	return Solution{
		Chromosome:  ch,
		Constraints: []float64{0.0},
		Objectives:  []float64{0.0, 0.0},
	}
}

func (s *Solution) String() string {
	return fmt.Sprintf("%s %v %v", s.Chromosome.String(), s.Constraints, s.Objectives)
}

// sorting

type SolutionSorting []Solution

func (s SolutionSorting) Len() int {
	return len(s)
}
func (s SolutionSorting) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SolutionSorting) Less(i, j int) bool {
	return compoundComparator(s[i], s[j]) < 0
}
