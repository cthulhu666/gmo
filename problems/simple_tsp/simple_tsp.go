package main

import (
	"fmt"
	"log"
	"math/rand"
	gmo "github.com/cthulhu666/gmo/engine"
)

var rnd = rand.New(rand.NewSource(0))
var populationSize = 40

func main() {
	m := loadData()
	//fmt.Println(m)

	best := newRoute([]int{0, 12, 1, 14, 8, 4, 6, 2, 11, 13, 9, 7, 5, 3, 10})
	fmt.Printf("Best solution: %d\n", m.distance(best))

	problem := TravellingSalesmanProblem{rnd, m}
	selection := gmo.TournamentSelection{Size: 2, Rnd: rnd, Comparator: gmo.ObjectiveComparator}

	pmx, err := pmx(rnd); panicOnError(err)
	mut, err := mutation(1, rnd); panicOnError(err)

	operators := []gmo.Operator{pmx, mut}
	algorithm := gmo.New(problem, selection, operators)

	var population []gmo.Solution
	for i := 0; i < populationSize; i++ {
		population = append(population, problem.RandomSolution())
	}

	algorithm.Run(population)
}

func callback(population []gmo.Tuple) {
	//for i, s := range population {
	//	fmt.Println(i, s)
	//	realChecksum := s.Solution.(route).debugChecksum()
	//	if s.Solution.Checksum() != realChecksum {
	//		log.Panic("Whoaaa")
	//	}
	//}
}

func panicOnError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// TODO: make operetaors generic (act on Solution, not on route)

func mutation(threshold int, rnd *rand.Rand) (gmo.Operator, error) {
	if threshold < 0 || threshold > 100 {
		return nil, fmt.Errorf("Illegal threshold value")
	}
	f := func(solutions []gmo.Solution) ([]gmo.Solution, error) {
		var children []gmo.Solution
		for _, s := range solutions {
			if rnd.Intn(100) < threshold {
				//fmt.Println("*** mutation ***")
				children = append(children, swap(s.(route), rnd))
			} else {
				children = append(children, s)
			}
		}
		return children, nil
	}
	return f, nil
}

func swap(r route, rnd *rand.Rand) route {
	a, b := rnd.Intn(r.Length), rnd.Intn(r.Length)
	points := r.getPoints()
	points[a], points[b] = points[b], points[a]
	return newRoute(points)
}

func crossover(rnd *rand.Rand) (gmo.Operator, error) {
	f := func(solutions []gmo.Solution) ([]gmo.Solution, error) {
		locus := rnd.Intn(solutions[0].(route).Length)
		var children []gmo.Solution
		for _, c := range combine(solutions[0].(route), solutions[1].(route), locus) {
			children = append(children, c)
		}
		return children, nil
	}
	return f, nil
}

func combine(a, b route, locus int) []route {
	c, d := make([]int, a.Length), make([]int, a.Length)
	copy(c[:], a.Points[0:locus])
	copy(c[locus:], b.Points[locus:])
	copy(d[:], b.Points[0:locus])
	copy(d[locus:], a.Points[locus:])
	return []route{
		newRoute(c),
		newRoute(d),
	}
}
