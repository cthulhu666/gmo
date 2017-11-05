package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstraintComparator(t *testing.T) {
	a := Solution{Constraints: []float64{5.0}}
	b := Solution{Constraints: []float64{15.0}}
	assert.Equal(t, -1, constraintComparator(a, b))
}

func TestObjectiveComparator(t *testing.T) {
	a := Solution{Objectives: []float64{5.0}}
	b := Solution{Objectives: []float64{15.0}}
	assert.Equal(t, 1, objectiveComparator(a, b))
}

func TestParetoObjectiveComparator(t *testing.T) {
	a := Solution{Objectives: []float64{5.0, 3.0}}
	b := Solution{Objectives: []float64{15.0, 6.0}}
	assert.Equal(t, 1, paretoObjectiveComparator(a, b))

	a = Solution{Objectives: []float64{5.0, 6.0}}
	b = Solution{Objectives: []float64{15.0, 3.0}}
	assert.Equal(t, 0, paretoObjectiveComparator(a, b))
}
