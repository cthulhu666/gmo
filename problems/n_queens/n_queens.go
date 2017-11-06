package main

import (
	"math/rand"
	gmo "github.com/cthulhu666/gmo/engine"
)

var rnd = rand.New(rand.NewSource(0))
var populationSize = 40

var checkboard_size = 8
var n_queens = 8

func main() {
	problem := NQueensProblem{checkboard_size: checkboard_size, number_of_queens: n_queens}
	selection := gmo.TournamentSelection{Size: 2, Rnd: rnd, Comparator: gmo.ObjectiveComparator}

	//pmx, err := pmx(rnd); panicOnError(err)
	//mut, err := mutation(1, rnd); panicOnError(err)

	operators := []gmo.Operator{}
	algorithm := gmo.New(problem, selection, operators)

	population := initialPopulation()

	algorithm.Run(population)
}

func initialPopulation() []gmo.Solution {
	return []gmo.Solution{}
}

type NQueensProblem struct {
	checkboard_size int
	number_of_queens int
}

func (p NQueensProblem) Evaluate(solution Solution) Evaluation {
	// TODO
}

func (p NQueensProblem) RandomSolution() Solution {
	// TODO
}
