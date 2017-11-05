package main

// https://ecs.victoria.ac.nz/foswiki/pub/Main/TechnicalReportSeries/ECSTR09-10.pdf
func (a *Algorithm) selectOne(population []Solution, comparator func(Solution, Solution) int) Solution {
	winner := a.sample(population)
	for i := 0; i < a.Config.TournamentSize; i++ {
		candidate := a.sample(population)
		if comparator(winner, candidate) > 0 {
			winner = candidate
		}
	}
	return winner
}

func (a *Algorithm) selectOneWithIncreasingTournamentSize(population []Solution, comparator func(Solution, Solution) int) Solution {
	winner := a.sample(population)
	var tournamentSize int
	switch {
	case float64(a.evals)/float64(a.Config.MaxEvaluations) < 0.1:
		tournamentSize = 1
	case float64(a.evals)/float64(a.Config.MaxEvaluations) < 0.5:
		tournamentSize = 2
	default:
		tournamentSize = 3 // a.Config.TournamentSize
	}
	for i := 0; i < tournamentSize; i++ {
		candidate := a.sample(population)
		if comparator(winner, candidate) > 0 {
			winner = candidate
		}
	}
	return winner
}

func (a *Algorithm) sample(population []Solution) Solution {
	i := a.rand.Int31n(int32(len(population)))
	return population[i]
}
