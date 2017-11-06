package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"math/rand"
	gmo "github.com/cthulhu666/gmo/engine"
)

func TestPMX(t *testing.T) {
	parents := []gmo.Solution{
		newRoute([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		newRoute([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}),
	}
	f, _ := pmx(rnd)
	children, _ := f(parents)
	assert.Len(t, children, 2)
}

func TestCombine2(t *testing.T) {
	a := newRoute([]int{1, 2, 3, 4, 5, 6, 7, 8})
	b := newRoute([]int{8, 7, 6, 5, 4, 3, 2, 1})

	arr := combine2(a, b, []int{2, 4})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, a.Points)
	assert.Equal(t, []int{8, 7, 6, 5, 4, 3, 2, 1}, b.Points)
	assert.Equal(t, []int{1, 2, 6, 5, 5, 6, 7, 8}, arr[0].Points)
	assert.Equal(t, []int{8, 7, 3, 4, 4, 3, 2, 1}, arr[1].Points)
}

func TestRandomLoci(t *testing.T) {
	rnd = rand.New(rand.NewSource(0))

	for i := 0; i < 100; i++ {
		loci := randomLoci(rnd, 2, 20)
		assert.Len(t, loci, 2)
		assert.True(t, loci[0] < loci[1], "! %d < %d", loci[0], loci[1])
	}
}

func TestRepair(t *testing.T) {
	a := newRoute([]int{1, 2, 6, 5, 5, 6, 7, 8, 8})
	b := newRoute([]int{9, 7, 3, 4, 4, 3, 2, 1, 9})

	arr := repair(a, b)
	assert.Equal(t, []int{1, 2, 6, 5, 3, 4, 7, 8, 9}, arr[0].(route).Points)
	assert.Equal(t, []int{9, 7, 3, 4, 6, 5, 2, 1, 8}, arr[1].(route).Points)
}

func TestFindDups(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 1, 4}
	b := []int{5, 6, 7, 2, 5, 6, 7}
	assert.Equal(t, map[int]pair{1: {0, 5}, 3: {2, 3}, 4: {4, 6}}, findDups(a))
	assert.Equal(t, map[int]pair{5: {0, 4}, 6: {1, 5}, 7: {2, 6}}, findDups(b))
}

func TestFindExchangePoints(t *testing.T) {
	a := map[int]pair{1: {0, 5}, 3: {2, 3}, 4: {4, 6}}
	b := map[int]pair{5: {0, 4}, 6: {1, 5}, 7: {2, 6}}
	assert.Equal(t, []pair{{5, 4}, {3, 5}, {6, 6}}, findExchangePoints(a, b))
}
