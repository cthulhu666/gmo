package main

import (
	"sort"
	gmo "github.com/cthulhu666/gmo/engine"
	"math/rand"
)

type TravellingSalesmanProblem struct {
	rand *rand.Rand
	m Map
}

func (t TravellingSalesmanProblem) Evaluate(solution gmo.Solution) gmo.Evaluation {
	route := solution.(route)
	var constraint float64
	if t.validateSolution(route) {
		constraint = 0
	} else {
		constraint = 10
	}
	objective := float64(t.m.distance(route))
	return gmo.Evaluation{Constraints: []float64{constraint}, Objectives: []float64{objective}}
}

func (t TravellingSalesmanProblem) RandomSolution() gmo.Solution {
	list := t.rand.Perm(t.m.Size)
	return newRoute(list)
}

// checks if each city is visited exactly once
func (t TravellingSalesmanProblem) validateSolution(route route) bool {
	if t.m.Size != route.Length {
		return false
	}
	sortedPoints := sortInts(route.Points)
	if sortedPoints[0] == 0 && sortedPoints[route.Length-1] == route.Length-1 {
		return true
	}
	return false
}

func sortInts(arr []int) []int {
	copy := append([]int(nil), arr...)
	sort.Ints(copy)
	return copy
}
