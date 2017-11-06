package main

/* To make things more funny, we aim to travel all the US states visiting each state only once,
minimizing distance travelled and also Levenshtein distance between cities' names */

import (
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"strings"
	"strconv"
	"fmt"
	"math/rand"
	//gmo "github.com/cthulhu666/gmo/engine"
)

type City struct {
	Coordinates coordinates
	City        string
	County      string
	State       string
}

type coordinates = [2]float32

//func callback(algorithm *gmo.Algorithm) {
//	fmt.Printf("Finished iteration: %d", algorithm.MaxEvaluations)
//}

var rnd *rand.Rand

func main() {
	cities := loadCities()
	fmt.Printf("Loaded %d cities\n", len(cities))

	rnd = rand.New(rand.NewSource(0))
	//config := gmo.Config{PopulationSize: 10, MaxEvaluations: 100, MutationThreshold: 5, TournamentSize: 3}
	//evaluator := Evaluator{cities}
	//population := initialPopulation(config.PopulationSize)
	//
	//algorithm := gmo.NewAlgorithm(rand, evaluator, population, config, callback)
	//algorithm.Run()

}

//func initialPopulation(size int) []gmo.Solution {
//	var population []gmo.Solution
//	for i := 0; i < size; i++ {
//		n := randomSolution()
//		population = append(population, n)
//	}
//	return population
//}
//
//func randomSolution() gmo.Solution {
//	s := randomValidCitySequence()
//	gmo.NewSolution(s)
//}
//
//func (e Evaluator) randomValidCitySequence() []int {
//	routeLen := 48 // FIXME
//}

func loadCities() []City {
	csvFile, err := os.Open("problems/tsp/usa115475_cities.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '|'
	var cities []City
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		cities = append(cities, City{
			Coordinates: parseCoordinates(line[0]),
			City:        line[1],
			County:      line[2],
			State:       line[3],
		})
	}
	return cities
}

func parseCoordinates(s string) coordinates {
	split := strings.Split(strings.TrimSpace(s), "  ")
	if len(split) != 2 {
		log.Fatalf("Wrong coordinates format: %s", s)
	}
	var result [2]float32
	lat, _ := strconv.ParseFloat(strings.TrimSpace(split[0]), 64)
	long, _ := strconv.ParseFloat(strings.TrimSpace(split[1]), 64)
	result[0] = float32(lat)
	result[1] = float32(long)
	return result
}
