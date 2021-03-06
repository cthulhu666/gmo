package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"math/rand"
)

func TestTruncate(t *testing.T) {
	algorithm := Algorithm{truncationComparator: ObjectiveComparator}
	population := []Tuple{
		{nil, &Evaluation{[]float64{10.0}, []float64{0}}},
		{nil, &Evaluation{[]float64{5.0}, []float64{0}}},
		{nil, &Evaluation{[]float64{15.0}, []float64{0}}},
		{nil, &Evaluation{[]float64{7.0}, []float64{0}}},
		{nil, &Evaluation{[]float64{20.0}, []float64{0}}},
	}
	truncated := algorithm.truncate(population, 2)
	assert.Equal(t, 2, len(truncated))
	assert.Equal(t, []float64{5.0}, truncated[0].Objectives)
	assert.Equal(t, []float64{7.0}, truncated[1].Objectives)
}

func TestEvolve(t *testing.T) {
	rnd := rand.New(rand.NewSource(0))
	problem := testProblem{rnd}
	selection := TournamentSelection{Size: 2, Rnd: rnd, Comparator: ObjectiveComparator}

	operators := []Operator{}
	a := New(problem, selection, operators, Configuration{})

	parents := []Solution{
		problem.RandomSolution(),
		problem.RandomSolution(),
	}

	children := a.evolve(parents)
	assert.Len(t, children, 2)
}

type testProblem struct {
	rnd *rand.Rand
}

func (p testProblem) Evaluate(solution Solution) Evaluation {
	return Evaluation{}
}

func (p testProblem) RandomSolution() Solution {
	return testSolution{}
}

type testSolution struct {

}

func (b testSolution) Id() string {
	return ""
}

func (b testSolution) Checksum() string {
	return ""
}