package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"math"
)

// TODO use arrays instead of multiple test methods for same function
func TestDistance(t *testing.T) {
	a := coordinates{1.0, 1.0}
	b := coordinates{1.0, 1.0}
	assert.Equal(t, 0.0, distance(a, b))
}

func TestDistance2(t *testing.T) {
	a := coordinates{1.0, 1.0}
	b := coordinates{2.0, 1.0}
	assert.Equal(t, 1.0, distance(a, b))
}

func TestDistance3(t *testing.T) {
	a := coordinates{1.0, 1.0}
	b := coordinates{2.0, 2.0}
	assert.Equal(t, math.Sqrt(2.0), distance(a, b))
}

func TestRouteLen(t *testing.T) {
	a := coordinates{1.0, 1.0}
	b := coordinates{2.0, 1.0}
	c := coordinates{2.0, 2.0}
	route := []coordinates{a, b, c}
	assert.Equal(t, 2.0, routeLen(route))
}

func TestHasCycle(t *testing.T) {
	route := []coordinates{
		{1.0, 1.0},
		{2.0, 2.0},
		{1.0, 1.0},
	}
	assert.True(t, hasCycle(route))
}
