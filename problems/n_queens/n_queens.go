package main

import (
	"math/rand"
	gmo "github.com/cthulhu666/gmo/engine"
	"fmt"
)

var rnd = rand.New(rand.NewSource(0))
var populationSize = 40

var checkboard_size = 8

func main() {
	problem := NQueensProblem{checkboard_size: checkboard_size, rnd: rnd}
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
	rnd              *rand.Rand
	checkboard_size  int
}

func (p NQueensProblem) Evaluate(solution gmo.Solution) gmo.Evaluation {
	f := float64(solution.(board).rowClashes() + solution.(board).diagonalClashes())
	return gmo.Evaluation{Objectives: []float64{f}}
}

func (p NQueensProblem) RandomSolution() gmo.Solution {
	var arr []int
	for i := 0; i < p.checkboard_size; i++ {
		arr = append(arr, p.rnd.Intn(checkboard_size))
	}
	return board{arr}
}

type board struct {
	columns []int
}

func (b board) getColumns() []int {
	return append([]int(nil), b.columns...)
}

func (b board) Id() string {
	return "" // TODO
}

func (b board) Checksum() string {
	return "" // TODO
}

func (b board) rowClashes() int {
	m := make(map[int]int)
	for _, row := range b.columns {
		m[row]++
	}
	clashes := 0
	for _, count := range m {
		if count > 1 {
			clashes += count - 1
		}
	}
	return clashes
}

func (b board) diagonalClashes() int {
	clashes := 0
	d1 := mapWithIndex(b.columns, func(a, i int) int { return a + i })
	d2 := mapWithIndex(b.columns, func(a, i int) int { return a - i })
	fmt.Println(d1, d2)
	m1 := count(d1)
	m2 := count(d2)
	fmt.Println(m1, m2)
	for _, count := range m1 {
		if count > 1 {
			clashes += count - 1
		}
	}
	for _, count := range m2 {
		if count > 1 {
			clashes += count - 1
		}
	}
	return clashes
}

func mapWithIndex(arr []int, f func(a, i int) int) []int {
	rs := make([]int, len(arr))
	for i, a := range arr {
		rs[i] = f(a, i)
	}
	return rs
}

func count(arr []int) map[int]int {
	m := make(map[int]int)
	for _, a := range arr {
		m[a]++
	}
	return m
}