package main

import (
	"os"
	"log"
	"bufio"
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"github.com/google/uuid"
	"fmt"
	"crypto/md5"
	"path/filepath"
)

type Map struct {
	Distances [][]int
	Size int
}

type route struct {
	Points []int
	Length int
	id string
	checksum string
}

// perhaps this should be a default getter
func (r route) getPoints() []int {
	return append([]int(nil), r.Points...)
}

func newRoute(points []int) route {
	text := fmt.Sprintf("%x", points)
	checksum := md5.Sum([]byte(text))
	id := uuid.New()
	return route{points, len(points), id.String(), fmt.Sprintf("%x", checksum)}
}

func (r route) get(index int) int {
	if index < r.Length {
		return r.Points[index]
	} else if index == r.Length {
		return r.Points[0]
	} else {
		return -1
	}
}

func (r route) Id() string {
	return r.id
}

func (r route) Checksum() string {
	return r.checksum
}

func (r route) debugChecksum() string {
	text := fmt.Sprintf("%x", r.Points)
	checksum := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", checksum)
}

func newMap(distances [][]int) Map {
	return Map{Distances: distances, Size: len(distances)}
}

func (m Map) distance(route route) int {
	dist := 0
	for i := range route.Points {
		from, to := route.get(i), route.get(i+1)
		dist += m.Distances[from][to]
	}
	return dist
}

func loadData() Map {
	absPath, _ := filepath.Abs("p01_d.csv")
	csvFile, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var cities [][]int
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		panicOnError(error)
		cities = append(cities, parseLine(line))
	}
	return newMap(cities)
}

func parseLine(line []string) []int {
	var n []int
	for _, s := range line {
		i, err := strconv.Atoi(strings.TrimSpace(s))
		panicOnError(err)
		n = append(n, i)
	}
	return n
}
