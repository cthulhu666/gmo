package main

func compoundComparator(a, b Solution) int {
	//fmt.Printf("a = %s\nb = %s\n", a.String(), b.String())
	c := constraintComparator(a, b)
	switch {
	case c < 0:
		return -1
	case c > 0:
		return +1
	}
	return paretoObjectiveComparator(a, b)
}

// less = better
func constraintComparator(a, b Solution) int {
	return sgn(a.Constraints[0] - b.Constraints[0])
}

// more = better
func objectiveComparator(a, b Solution) int {
	return sgn(b.Objectives[0] - a.Objectives[0])
}

//ParetoObjectiveComparator = proc do |s1, s2|
//	dominate1 = false
//	dominate2 = false
//
//	s1.number_of_objectives.times do |i|
//	if s1.objectives[i] < s2.objectives[i]
//		dominate1 = true
//		next 0 if dominate2
//		elsif s1.objectives[i] > s2.objectives[i]
//		dominate2 = true
//		next 0 if dominate1
//		end
//	end
//
//	next 0 if dominate1 == dominate2
//	next -1 if dominate1
//	next 1
//end
const NumberOfObjectives = 2

func paretoObjectiveComparator(a, b Solution) int {
	dominate1, dominate2 := false, false
	for i := 0; i < NumberOfObjectives; i++ {
		if a.Objectives[i] > b.Objectives[i] {
			dominate1 = true
			if dominate2 {
				return 0
			}
		} else if a.Objectives[i] < b.Objectives[i] {
			dominate2 = true
			if dominate1 {
				return 0
			}
		}
	}
	var retVal int
	switch {
	case dominate1 == dominate2:
		retVal = 0
	case dominate1:
		retVal = -1
	default:
		retVal = 1
	}
	return retVal
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
