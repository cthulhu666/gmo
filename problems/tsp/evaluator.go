package main

import (
	"math"
)

type Evaluator struct {
	Cities []City
}

//func (Evaluator) Eval(x interface{}) gmo.Evaluation {
//	sequence := x.([]int)
//
//	return gmo.Evaluation{}
//}

func distance(a, b coordinates) float64 {
	distance := 0.0
	for i := 0; i < 2; i++ {
		distance += (float64(a[i]) - float64(b[i]))*(float64(a[i])-float64(b[i]))
	}
	return math.Sqrt(distance)
}

func routeLen(route []coordinates) float64 {
	d := 0.0
	for i := 1; i < len(route); i++ {
		d += distance(route[i-1], route[i])
	}
	return d
}

func hasCycle(route []coordinates) bool {
	return route[0] == route[len(route)-1]
}
