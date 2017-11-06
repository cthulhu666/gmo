package engine

import (
	"sort"
	"fmt"
	"log"
)

type Algorithm struct {
	Problem
	Selection
	Operators []Operator
	generation int
	callback Callback
	//best Tuple
	truncationComparator Comparator
}

func New(problem Problem, selection Selection, operators []Operator) Algorithm {
	return Algorithm{Problem: problem, Selection: selection, Operators: operators, truncationComparator: ObjectiveComparator}
}

type Callback = func(population []Tuple)

type Selection interface {
	selectOne(population []Tuple) Tuple
}

type Comparator = func(*Evaluation, *Evaluation) int

type Problem interface {
	Evaluate(solution Solution) Evaluation
	RandomSolution() Solution
}

type Operator = func ([]Solution) ([]Solution, error)

type Evaluation struct {
	Objectives 	[]float64
	Constraints []float64
}

type Solution interface {
	Id() string
	Checksum() string
}

type Tuple struct{Solution; *Evaluation}

func (t Tuple) String() string {
	return fmt.Sprintf("%v, %.0f", t.Solution, t.Objectives)
}

func (a Algorithm) iterate(population []Tuple) []Tuple {
	evaluatedSolutions := a.evaluateAll(population)
	populationSize := len(population)

	var offspring []Tuple
	//selectionRegistrar := make(map[string]int)
	solutionRegistrar := make(map[string]bool)

	for i := 0; len(offspring) < populationSize; i++ {
		parents := a.selectParents(evaluatedSolutions)
		//for _, p := range parents {
		//	selectionRegistrar[p.Id()]++
		//}
		children := a.evolve(parents)
		for _, c := range children {
			if !solutionRegistrar[c.Checksum()] {
				offspring = append(offspring, Tuple{c, nil})
				solutionRegistrar[c.Checksum()] = true
			}
		}
	}
	evaluatedOffspring := a.evaluateAll(offspring)
	evaluatedSolutions = append(evaluatedSolutions, evaluatedOffspring...)

	//a.callback(evaluatedSolutions)

	return a.truncate(evaluatedSolutions, populationSize)
}

func (a Algorithm) selectParents(population []Tuple) []Solution {
	parents := [2]Solution{
		a.Selection.selectOne(population).Solution,
		a.Selection.selectOne(population).Solution,
	}
	return parents[:]
}

func (a Algorithm) evolve(parents []Solution) []Solution {
	tmp := parents
	for _, f := range a.Operators {
		children, err := f(tmp)
		panicOnError(err)
		tmp = children
	}
	return tmp
}

func (a Algorithm) evaluateAll(population []Tuple) []Tuple {
	var tuples []Tuple
	for _, s := range population {
		if s.Evaluation != nil {
			tuples = append(tuples, s)
		} else {
			eval := a.Problem.Evaluate(s.Solution)
			tuple := Tuple{s.Solution, &eval}
			tuples = append(tuples, tuple)
		}
	}
	return tuples
}

func (a Algorithm) terminated() bool {
	// a.best.Objectives[0] <= 291 ||
	return a.generation >= 50
}

func (a Algorithm) Run(initialPopulation []Solution) {
	var population []Tuple
	for _, s := range initialPopulation {
		population = append(population, Tuple{s, nil})
	}
	for {
		if a.terminated() {
			break
		}
		fmt.Printf("Starting generation %d\n", a.generation)
		population = a.iterate(population) // FIXME
		best := population[0]
		fmt.Printf("Best solution: %s\n", best)
		a.generation++
	}
}

func (a Algorithm) truncate(population []Tuple, size int) []Tuple {
	// TODO: use sort.Slice ?
	sort.Sort(SolutionSorting{population, a.truncationComparator})
	return population[0:size]
}

type SolutionSorting struct {
	data []Tuple
	comparator Comparator
}
func (s SolutionSorting) Len() int {
	return len(s.data)
}
func (s SolutionSorting) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}
func (s SolutionSorting) Less(i, j int) bool {
	return s.comparator(s.data[i].Evaluation, s.data[j].Evaluation) < 0
}


func panicOnError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
