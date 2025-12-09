package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(filename string) []JunctionBox {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var jBoxes []JunctionBox
	for _, line := range lines {
		parts := strings.Split(line, ",")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		c, _ := strconv.Atoi(parts[2])
		jBoxes = append(jBoxes, JunctionBox{X: a, Y: b, Z: c})
	}

	return jBoxes
}

// EuclideanDistance calculates the shortest distance between 2 points in 3D space
// https://en.wikipedia.org/wiki/Euclidean_distance
// https://www.freecodecamp.org/news/euclidean-algorithm-in-golang/
func EuclideanDistance(p1, p2 JunctionBox) float64 {
	da := p2.X - p1.X
	db := p2.Y - p1.Y
	dc := p2.Z - p1.Z
	return math.Sqrt(float64(da*da + db*db + dc*dc))
}

func closestPoints(junctionBoxes []JunctionBox, usedPairs map[string]bool) (int, int, ClosestJunctionBoxes) {
	p1Index := -1
	p2Index := -1
	shortest := math.MaxFloat64

	for i := 0; i < len(junctionBoxes)-1; i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			// Create a key for this pair (order doesn't matter, smaller index first)
			pairKey := fmt.Sprintf("%d,%d", i, j)
			if usedPairs[pairKey] {
				continue
			}

			distance := EuclideanDistance(junctionBoxes[i], junctionBoxes[j])
			if distance < shortest {
				shortest = distance
				p1Index = i
				p2Index = j
			}
		}
	}

	entry := ClosestJunctionBoxes{}
	if p1Index != -1 {
		entry = ClosestJunctionBoxes{
			jBox1:    junctionBoxes[p1Index],
			jBox2:    junctionBoxes[p2Index],
			distance: shortest,
		}
	}
	return p1Index, p2Index, entry
}

// findjBox find needle in haystack, else return -1, return location of jBox in list of circuits
func findjBox(jBox JunctionBox, circuits []Circuit) int {
	for i, circuit := range circuits {
		for _, jb := range circuit.junctionBox {
			if jb == jBox {
				return i
			}
		}
	}
	return -1
}

// sortCircuitsBySize sorts circuits by number of points (largest first)
// https://stackoverflow.com/a/40932847
func sortCircuitsBySize(circuits []Circuit) {
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].junctionBox) > len(circuits[j].junctionBox)
	})
}

// calculateCircuitSolution returns the maths
func calculateCircuitSolution(circuits []Circuit) int {
	if len(circuits) < 3 {
		return 0
	}
	return len(circuits[0].junctionBox) * len(circuits[1].junctionBox) * len(circuits[2].junctionBox)
}

func showCircuits(circuits []Circuit) {
	for _, circuit := range circuits {
		fmt.Println(circuit)
	}
}

func showClosestPairs(boxes []ClosestJunctionBoxes) {
	for _, pair := range boxes {
		fmt.Println(pair)
	}
}

// multiplyXCoordinates returns the product of X coordinates from two junction boxes
func multiplyXCoordinates(jBox1, jBox2 JunctionBox) int {
	return jBox1.X * jBox2.X
}
