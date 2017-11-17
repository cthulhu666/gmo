package main

import (
	"math/rand"
	gmo "github.com/cthulhu666/gmo/engine"
	"fmt"
	"crypto/md5"
	"github.com/google/uuid"
	"log"
)

var rnd = rand.New(rand.NewSource(0))
var populationSize = 500

var checkboard_size = 8

func main() {
	problem := NQueensProblem{checkboard_size: checkboard_size, rnd: rnd}
	selection := gmo.TournamentSelection{Size: 4, Rnd: rnd, Comparator: gmo.ObjectiveComparator}

	crossover, err := crossover(rnd); panicOnError(err)
	mut, err := mutation(1, rnd); panicOnError(err)

	operators := []gmo.Operator{crossover, mut}
	algorithm := gmo.New(problem, selection, operators)

	population := initialPopulation(problem)

	algorithm.Run(population)
}

// TODO: move to gmo.Algorithm?
func initialPopulation(p gmo.Problem) []gmo.Solution {
	var population []gmo.Solution
	for i := 0; i < populationSize; i++ {
		population = append(population, p.RandomSolution())
	}
	return population
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
	return newBoard(arr)
}

type board struct {
	columns []int
	id string
	checksum string
}

func newBoard(columns []int) board {
	text := fmt.Sprintf("%x", columns)
	checksum := md5.Sum([]byte(text))
	id := uuid.New()
	return board{columns, id.String(), fmt.Sprintf("%x", checksum)}
}

func (b board) Columns() []int {
	return append([]int(nil), b.columns...)
}

func (b board) Id() string {
	return b.id
}

func (b board) Checksum() string {
	return b.checksum
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
	return diagonalClashes(b, 1) + diagonalClashes(b, -1)
}

// dir can be either 1 or -1
func diagonalClashes(b board, dir int) int {
	clashes := 0
	d := mapWithIndex(b.columns, func(a, i int) int { return a + (i * dir) })
	for _, count := range count(d) {
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


func panicOnError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func crossover(rnd *rand.Rand) (gmo.Operator, error) {
	f := func(solutions []gmo.Solution) ([]gmo.Solution, error) {
		locus := rnd.Intn(checkboard_size)
		var children []gmo.Solution
		for _, c := range combine(solutions[0].(board), solutions[1].(board), locus) {
			children = append(children, c)
		}
		return children, nil
	}
	return f, nil
}

func combine(a, b board, locus int) []board {
	c, d := make([]int, checkboard_size), make([]int, checkboard_size)
	// TODO: avoid using `columns` directly
	copy(c[:], a.columns[0:locus])
	copy(c[locus:], b.columns[locus:])
	copy(d[:], b.columns[0:locus])
	copy(d[locus:], a.columns[locus:])
	return []board{
		newBoard(c),
		newBoard(d),
	}
}

func mutation(threshold int, rnd *rand.Rand) (gmo.Operator, error) {
	if threshold < 0 || threshold > 100 {
		return nil, fmt.Errorf("Illegal threshold value")
	}
	f := func(solutions []gmo.Solution) ([]gmo.Solution, error) {
		var children []gmo.Solution
		for _, s := range solutions {
			if rnd.Intn(100) < threshold {
				children = append(children, swap(s.(board), rnd))
			} else {
				children = append(children, s)
			}
		}
		return children, nil
	}
	return f, nil
}

func swap(r board, rnd *rand.Rand) board {
	a, b := rnd.Intn(checkboard_size), rnd.Intn(checkboard_size)
	points := r.Columns()
	points[a], points[b] = points[b], points[a]
	return newBoard(points)
}
