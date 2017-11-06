package main

import (
	"math/rand"
	"sort"
	gmo "github.com/cthulhu666/gmo/engine"
)

// works only for len(solutions) == 2
func pmx(rnd *rand.Rand) (gmo.Operator, error) {
	f := func(solutions []gmo.Solution) ([]gmo.Solution, error) {
		var children []gmo.Solution
		for _, c := range combine2(solutions[0].(route), solutions[1].(route), randomLoci(rnd, 2, solutions[0].(route).Length)) {
			children = append(children, c)
		}
		repaired := repair(children[0], children[1])
		return repaired, nil
	}
	return f, nil
}

func repair(a, b gmo.Solution) []gmo.Solution {
	c, d := a.(route).getPoints(), b.(route).getPoints()
	for _, p := range findExchangePoints(findDups(c), findDups(d)) {
		c[p.fst], d[p.snd] = d[p.snd], c[p.fst]
	}
	return []gmo.Solution{newRoute(c), newRoute(d)}
}

type pair struct{ fst, snd int }

func findDups(a []int) map[int]pair {
	tmp := make(map[int][]int)
	for i, e := range a {
		tmp[e] = append(tmp[e], i)
	}
	m := make(map[int]pair)
	for k, v := range tmp {
		if len(v) == 2 {
			m[k] = pair{v[0], v[1]}
		}
	}
	return m
}

func findExchangePoints(a, b map[int]pair) []pair {
	if len(a) != len(b) {
		panic("Both maps should have equal len")
	}
	xa := mapValues(a, func(m pair) int { return m.snd })
	xb := mapValues(b, func(m pair) int { return m.snd })
	return zip(xa, xb)
}

func mapValues(m map[int]pair, f func(m pair) int) []int {
	var a []int

	keys := make([]int, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		a = append(a, f(m[k]))
	}

	return a
}

func zip(a, b []int) []pair {
	var pairs []pair
	for i := range a {
		pairs = append(pairs, pair{a[i], b[i]})
	}
	return pairs
}

func randomLoci(rnd *rand.Rand, count int, len int) []int {
	// TODO: it can be slow for large `len`
	loci := rnd.Perm(len)[:count]
	sort.Ints(loci)
	return loci
}

func combine2(a, b route, locus []int) []route {
	c, d := make([]int, a.Length), make([]int, a.Length)

	i := 0
	// TODO: ensure we don't change `locus` if cap(locus) > len(locus)
	for n, loc := range append(locus, a.Length) {
		if n%2 == 0 {
			copy(c[i:loc], a.Points[i:loc])
			copy(d[i:loc], b.Points[i:loc])
		} else {
			copy(c[i:loc], b.Points[i:loc])
			copy(d[i:loc], a.Points[i:loc])
		}
		i = loc
	}

	return []route{
		newRoute(c),
		newRoute(d),
	}
}
