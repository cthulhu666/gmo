package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistance(t *testing.T) {
	p := problem(t)
	route := newRoute([]int{0, 1, 2})
	assert.Equal(t, 166, p.m.distance(route))
}

func TestValidateSolution(t *testing.T) {
	p := problem(t)
	route := p.RandomSolution().(route)
	assert.True(t, p.validateSolution(route))
}

func TestValidateSolutionHasNoSideEffects(t *testing.T) {
	p := problem(t)
	route := p.RandomSolution().(route)
	points := append([]int(nil), route.Points...)
	p.validateSolution(route)
	assert.Equal(t, points, route.Points)
}

func TestChecksum(t *testing.T) {
	r := newRoute([]int{11, 5, 2, 4, 14, 8, 12, 0, 13, 3, 6, 9, 10, 7, 1})
	assert.Equal(t, "da32d691be350bba555c10e24200c3a0", r.checksum)
}

func problem(t *testing.T) TravellingSalesmanProblem {
	t.Helper()
	m := loadData()
	return TravellingSalesmanProblem{rnd, m}
}
