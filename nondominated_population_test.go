package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNondominatedPopulationAdd(t *testing.T) {
	n := NewNondominatedPopulation(compoundComparator)
	s := Solution{Constraints: []float64{0.0}, Objectives: []float64{0.0, 0.0}}
	assert.True(t, n.Add(s))
	assert.Len(t, n.Population, 1)
	s2 := Solution{Constraints: []float64{1.0}, Objectives: []float64{0.0, 0.0}}
	assert.False(t, n.Add(s2))
	assert.Len(t, n.Population, 1)
	assert.Equal(t, s, n.Population[0])
	s3 := Solution{Constraints: []float64{0.0}, Objectives: []float64{1.0, 0.0}}
	assert.True(t, n.Add(s3))
	assert.Len(t, n.Population, 1)
	assert.Equal(t, s3, n.Population[0])
	//s4 := Solution{constraints: []float64{0.0}, objectives: []float64{0.0, 1.0}}
	//assert.True(t, n.Add(s4))
	//assert.Len(t, n.Population, 2)
	s5 := Solution{Constraints: []float64{0.0}, Objectives: []float64{0.0, 1.0}}
	assert.False(t, n.Add(s5))
}

func TestNondominatedPopulationAdd2(t *testing.T) {

}
