package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	gmo "github.com/cthulhu666/gmo/engine"
)

func TestCrossover(t *testing.T) {
	parents := []gmo.Solution{
		newRoute([]int{1, 2, 3, 4}),
		newRoute([]int{5, 6, 7, 8}),
	}
	f, _ := crossover(rnd)
	children, _ := f(parents)
	assert.Len(t, children, 2)
	assert.Equal(t, []int{1, 2, 3, 4}, parents[0].(route).Points)
	assert.Equal(t, []int{5, 6, 7, 8}, parents[1].(route).Points)
}

func TestCombine(t *testing.T) {
	a := newRoute([]int{1, 2, 3, 4})
	b := newRoute([]int{5, 6, 7, 8})

	arr := combine(a, b, 2)
	assert.Equal(t, []int{1, 2, 3, 4}, a.Points)
	assert.Equal(t, []int{5, 6, 7, 8}, b.Points)
	assert.Equal(t, []int{1, 2, 7, 8}, arr[0].Points)
	assert.Equal(t, []int{5, 6, 3, 4}, arr[1].Points)
}
