package engine

import "math/rand"

type TournamentSelection struct {
	Rnd        *rand.Rand
	Size       int
	Comparator func(*Evaluation, *Evaluation) int
}

// https://ecs.victoria.ac.nz/foswiki/pub/Main/TechnicalReportSeries/ECSTR09-10.pdf
func (t TournamentSelection) selectOne(population []Tuple) Tuple {
	winner := t.sample(population)
	for i := 0; i < t.Size; i++ {
		candidate := t.sample(population)
		if t.Comparator(winner.Evaluation, candidate.Evaluation) > 0 {
			winner = candidate
		}
	}
	return winner
}

func (t TournamentSelection) sample(population []Tuple) Tuple {
	i := t.Rnd.Int31n(int32(len(population)))
	return population[i]
}

func CompoundComparator(a, b *Evaluation) int {
	c := ConstraintComparator(a, b)
	switch {
	case c < 0:
		return -1
	case c > 0:
		return +1
	}
	return ObjectiveComparator(a, b)
}

// less = better
func ConstraintComparator(a, b *Evaluation) int {
	return sgn(a.Constraints[0] - b.Constraints[0])
}

// less = better
func ObjectiveComparator(a, b *Evaluation) int {
	return sgn(a.Objectives[0] - b.Objectives[0])
}

func sgn(a float64) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}
