package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func init() {
}

func TestChromosome_Randomize(t *testing.T) {
	r := rand.New(rand.NewSource(99))
	var ch Chromosome
	assert.Equal(t, []bool{true, true, false, false, false}, ch.Randomize(r, 5))
}

func TestMutate(t *testing.T) {
	r := rand.New(rand.NewSource(99))
	ch := Chromosome{[]bool{true, true, false, false, false}}
	assert.Equal(t, []bool{false, true, false, false, false}, ch.mutate(r, 30))
}

func TestCrossover(t *testing.T) {
	r := rand.New(rand.NewSource(99))
	p1 := Chromosome{[]bool{true, true, true, true, true}}
	p2 := Chromosome{[]bool{false, false, false, false, false}}
	assert.Equal(t, []bool{true, true, false, false, false}, crossover(r, p1, p2).Bits)
}

func TestChromosome_String(t *testing.T) {
	ch := Chromosome{[]bool{true, false}}
	assert.Equal(t, "10", ch.String())
}

func TestAlgorithm_selectParents(t *testing.T) {
	alg := Algorithm{rand: rand.New(rand.NewSource(123299))}
	population := []Solution{
		Solution{Constraints: []float64{5}, Objectives: []float64{1, 1}},
		Solution{Constraints: []float64{6}, Objectives: []float64{6, 6}},
		Solution{Constraints: []float64{5}, Objectives: []float64{2, 2}},
		Solution{Constraints: []float64{0}, Objectives: []float64{4, 4}},
		Solution{Constraints: []float64{0}, Objectives: []float64{1, 1}},
		Solution{Constraints: []float64{0}, Objectives: []float64{5, 5}},
		Solution{Constraints: []float64{1}, Objectives: []float64{5, 5}},
		Solution{Constraints: []float64{0}, Objectives: []float64{6, 6}},
		Solution{Constraints: []float64{0}, Objectives: []float64{0, 0}},
	}
	parents := alg.selectParents(population)

	t.Log(parents[0])
	t.Log(parents[1])
}

func TestAlgorithm_truncate(t *testing.T) {
	population := []Solution{
		Solution{Constraints: []float64{5}, Objectives: []float64{1, 1}},
		Solution{Constraints: []float64{6}, Objectives: []float64{6, 6}},
		Solution{Constraints: []float64{5}, Objectives: []float64{2, 2}},
		Solution{Constraints: []float64{0}, Objectives: []float64{4, 4}},
		Solution{Constraints: []float64{0}, Objectives: []float64{1, 1}},
		Solution{Constraints: []float64{0}, Objectives: []float64{5, 5}},
		Solution{Constraints: []float64{1}, Objectives: []float64{5, 5}},
		Solution{Constraints: []float64{0}, Objectives: []float64{6, 6}},
		Solution{Constraints: []float64{0}, Objectives: []float64{0, 0}},
	}
	survivors := truncate(population, 2)

	assert.Equal(t, 2, len(survivors))
	assert.EqualValues(t, 0, survivors[0].Constraints[0])
	assert.EqualValues(t, 0, survivors[1].Constraints[0])
	assert.EqualValues(t, 6, survivors[0].Objectives[0])
	assert.EqualValues(t, 5, survivors[1].Objectives[0])
}

type testEvaluator struct{}

func (testEvaluator) Eval(ch Chromosome) Evaluation {
	return Evaluation{[]float64{5}, []float64{15}}
}

func TestEvaluateAll(t *testing.T) {
	n := NewSolution(Chromosome{})
	population := []Solution{
		n,
	}

	EvaluateAll(population, testEvaluator{})

	assert.EqualValues(t, 5, population[0].Constraints[0])
}

func TestCountUnique(t *testing.T) {
	pop := []Solution{
		NewSolution(Chromosome{[]bool{false, false, false}}),
		NewSolution(Chromosome{[]bool{false, false, true}}),
		NewSolution(Chromosome{[]bool{false, false, false}}),
		NewSolution(Chromosome{[]bool{false, false, false}}),
		NewSolution(Chromosome{[]bool{false, false, true}}),
	}
	assert.Equal(t, 2, countUnique(pop))
}

//func TestAlgorithm_Iterate(t *testing.T) {
//	r := rand.New(rand.NewSource(666))
//	ship := itemStore.Find("Rifter")
//	decoder := NewDecoder(ship)
//	var pop []Solution
//	for i := 0; i < 10; i++ {
//		var ch Chromosome
//		ch.Randomize(r, decoder.chromosomeSize)
//		n := NewSolution(ch)
//		pop = append(pop, n)
//	}
//	alg := Algorithm{rand: r, decoder: decoder, Population: pop, evals: 0}
//	alg.evaluateAll(pop, Evaluate)
//	for i, s := range alg.Population {
//		t.Logf("%d: %s\n", i, s.String())
//	}
//	alg.Iterate()
//	t.Log("--------------------------")
//	for i, s := range alg.Population {
//		t.Logf("%d: %s\n", i, s.String())
//	}
//	alg.Iterate()
//	t.Log("--------------------------")
//	for i, s := range alg.Population {
//		t.Logf("%d: %s\n", i, s.String())
//	}
//	alg.Iterate()
//	t.Log("--------------------------")
//	for i, s := range alg.Population {
//		t.Logf("%d: %s\n", i, s.String())
//	}
//	ch := alg.Population[0].Chromosome
//	f := decoder.Decode(ch)
//	t.Log(f.ReadShipAttribute(ShieldCapacity))
//	t.Log(f.Validate())
//}

func TestSelectAndEvolve(t *testing.T) {
	r := rand.New(rand.NewSource(667))
	pop := []Solution{
		NewSolution(Chromosome{[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false}}),
		NewSolution(Chromosome{[]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true}}),
	}
	a := Algorithm{rand: r, Population: pop, evals: 0}
	for i, s := range a.Population {
		t.Logf("population %d: %s\n", i, s.String())
	}
	parents := a.selectParents(a.Population)
	for i, s := range parents {
		t.Logf("parents %d: %s\n", i, s.String())
	}
	child := a.evolve(parents)
	t.Logf("child: %s", child.String())
	for i, s := range a.Population {
		t.Logf("population %d: %s\n", i, s.String())
	}
}
