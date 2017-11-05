package main

import (
	"encoding/csv"
	"fmt"
	"github.com/montanaflynn/stats"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"sort"
)

type Evaluation struct {
	Constraints []float64
	Objectives  []float64
}

type Evaluator interface {
	Eval(Chromosome) Evaluation
}

type Config struct {
	PopulationSize    int
	MaxEvaluations    int
	MutationThreshold int
	TournamentSize    int
}

type Callback func(*Algorithm)
func EmptyCallback(*Algorithm) {
	// do nothing
}

type Algorithm struct {
	rand       *rand.Rand
	evaluator  Evaluator
	Population []Solution
	evals      int
	afterIterationCallback Callback
	Config
}

func NewAlgorithm(rand *rand.Rand, evaluator Evaluator, population []Solution, cfg Config, afterIterationCallback Callback) Algorithm {
	return Algorithm{
		rand: rand,
		evaluator: evaluator,
		Population: population,
		afterIterationCallback: afterIterationCallback,
		Config: cfg,
	}
}

func (a *Algorithm) Run() {
	for a.evals < a.MaxEvaluations {
		a.Iterate()
		a.afterIterationCallback(a)
		log.Printf("Evals:  %d\n", a.evals)
	}
}

func (a *Algorithm) Iterate() {
	//fmt.Printf("| Best: %s, %v, %v\n", a.population[0].ch.String(), a.population[0].constraints, a.population[0].objectives)
	//fmt.Printf("| Worst: %s, %v, %v\n", a.population[29].ch.String(), a.population[29].constraints, a.population[29].objectives)
	populationSize := len(a.Population)
	var offspring []Solution
	for i := 0; len(offspring) < populationSize; i++ {
		parents := a.selectParents(a.Population)
		child := a.evolve(parents)
		offspring = append(offspring, child)
	}
	EvaluateAll(offspring, a.evaluator)
	a.evals += len(offspring)
	a.Population = append(a.Population, offspring...)
	//fmt.Printf("-- before truncate: %d unique chromosomes\n", countUnique(a.population))
	a.Population = truncate(a.Population, a.Config.PopulationSize)
	storeStats(a.Population)
	//fmt.Printf("-- after truncate: %d unique chromosomes\n", countUnique(a.population))
	//fmt.Printf("Best: %s, %v, %v\n", a.population[0].ch.String(), a.population[0].constraints, a.population[0].objectives)
	//fmt.Printf("Worst: %s, %v, %v\n", a.population[29].ch.String(), a.population[29].constraints, a.population[29].objectives)
}

func EvaluateAll(solutions []Solution, evaluator Evaluator) {
	for _, s := range solutions {
		e := evaluator.Eval(s.Chromosome)
		copy(s.Constraints, e.Constraints)
		copy(s.Objectives, e.Objectives)
	}
}

var meanFitness []float64
var maxFitness []float64
var history [][]Solution

func storeStats(solutions []Solution) {
	var f []float64
	var historyLine []Solution
	for _, s := range solutions {
		obj1 := s.Objectives[0]
		obj2 := s.Objectives[1]
		fitness := math.Sqrt(obj1*obj1 + obj2*obj2)
		f = append(f, fitness)
		historyLine = append(historyLine, s)
	}
	history = append(history, historyLine)
	median, _ := stats.Median(f)
	max, _ := stats.Max(f)
	meanFitness = append(meanFitness, median)
	maxFitness = append(maxFitness, max)
}

// TODO: drop whole bunch of different stats to a separate dir
func DropStats() {
	//csvOut, err := os.Create("objective_median.csv")
	csvOut, err := os.Create("objective_max.csv")
	if err != nil {
		log.Fatal("Unable to open output")
	}
	w := csv.NewWriter(csvOut)
	defer csvOut.Close()

	for i, _ := range maxFitness {
		rec := []string{fmt.Sprintf("%f", maxFitness[i])}
		if err = w.Write(rec); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
}

func DropHistory() {
	csvOut, err := os.Create("history.csv")
	if err != nil {
		log.Fatal("Unable to open output")
	}
	w := csv.NewWriter(csvOut)
	defer csvOut.Close()

	for i, _ := range history {
		rec := record(history[i])
		if err = w.Write(rec); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
}

func record(line []Solution) []string {
	var r []string
	for _, s := range line {
		obj1 := s.Objectives[0]
		obj2 := s.Objectives[1]
		fitness := math.Sqrt(obj1*obj1 + obj2*obj2)
		r = append(r, fmt.Sprintf("%f", fitness))
	}
	return r
}

// TODO: avoiding mating a solution with itself
func (a *Algorithm) selectParents(population []Solution) [2]Solution {
	return [...]Solution{
		a.selectOneWithIncreasingTournamentSize(population, compoundComparator),
		a.selectOneWithIncreasingTournamentSize(population, compoundComparator),
	}
}

var selfMatingCounter int

func (a *Algorithm) evolve(parents [2]Solution) Solution {
	if reflect.DeepEqual(parents[0].Chromosome, parents[1].Chromosome) {
		selfMatingCounter++
	}
	ch := crossover(a.rand, parents[0].Chromosome, parents[1].Chromosome)
	ch.mutate(a.rand, a.Config.MutationThreshold)
	return NewSolution(ch)
}

// TODO: can we do without mutating the population?
func truncate(population []Solution, size int) []Solution {
	sort.Sort(SolutionSorting(population))
	return population[0:size]
}

func countUnique(population []Solution) int {
	m := make(map[string]bool)
	for _, s := range population {
		m[s.Chromosome.String()] = true
	}
	return len(m)
}

func (a *Algorithm) Result() NondominatedPopulation {
	p := NewNondominatedPopulation(compoundComparator)
	p.AddAll(a.Population)
	return p
}
