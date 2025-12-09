package main

import (
	"advent-of-code-2025/utils"
	"fmt"
)

func main() {
	jBoxes = parseInput("inputs/day08/input.txt")
	maxBoxesInCircuit = 1000 // sample.txt needs 10, input.txt needs 1000

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var jBoxes []JunctionBox
var maxBoxesInCircuit int

// Part 1 Problem:
//  1. Need to come up with a way to determine two junction boxes that are the closest together from a list
//  2. Then we need to find the next closest pair
//     2a. In finding the next pair, should one of those junctions boxes be already claimed by another circuit, then we'll add to that circuit.
//     2b. There maybe situations where the next closest pair is already in another circle, then nothing needs to happen
//
// 3. Once we have our list of circuits, we need to calculate how many junction boxes are in each circuit, this is our "size".
// 4. Multiply the 3 largest circuits by size together and that is our answer
func part1() int {
	var circuits []Circuit
	var closestPairs []ClosestJunctionBoxes
	connectedBoxes := make(map[string]bool)

	// Find our closest two points in our list (stop after maxBoxesInCircuit connections)
	for {
		jBox1Index, jBox2Index, closestPair := closestPoints(jBoxes, connectedBoxes)
		if jBox1Index == -1 {
			break // No more pairs available
		}

		// Store this pair
		closestPairs = append(closestPairs, closestPair)

		// Get our {x,y,z} from the pair
		jBox1 := closestPair.jBox1
		jBox2 := closestPair.jBox2

		// Get position of our jBoxes...
		jBox1Position := findjBox(jBox1, circuits)
		jBox2Position := findjBox(jBox2, circuits)

		// Start testing
		if jBox1Position == -1 && jBox2Position == -1 { // Neither jBox is in a circuit, create new circuit
			circuits = append(circuits, Circuit{junctionBox: []JunctionBox{jBox1, jBox2}})

		} else if jBox1Position == -1 { // Only jBox1 is new, add to jBox2's circuit
			circuits[jBox2Position].junctionBox = append(circuits[jBox2Position].junctionBox, jBox1)

		} else if jBox2Position == -1 { // Only jBox2 is new, add to jBox1's circuit
			circuits[jBox1Position].junctionBox = append(circuits[jBox1Position].junctionBox, jBox2)

		} else if jBox1Position != jBox2Position { // Found Both in different circuits, merge them
			circuits[jBox1Position].junctionBox = append(circuits[jBox1Position].junctionBox, circuits[jBox2Position].junctionBox...)
			circuits = append(circuits[:jBox2Position], circuits[jBox2Position+1:]...)
		}
		// If jBox1 == jBox2, then they are already in the same circuit, so nothing to do

		// Mark this pair as used
		pairKey := fmt.Sprintf("%d,%d", jBox1Index, jBox2Index)
		connectedBoxes[pairKey] = true

		maxBoxesInCircuit--
		if maxBoxesInCircuit <= 0 { // Check how many jBoxes are left in our circuit
			//showClosestPairs(closestPairs)
			break // We hit our limit, leave for loop
		}
	}

	// Find unused jBoxes and add each as its own circuit
	for _, jBox := range jBoxes {
		// If jBox search comes back empty, create a new circuit
		if findjBox(jBox, circuits) == -1 {
			circuits = append(circuits, Circuit{junctionBox: []JunctionBox{jBox}})
		}
	}

	// Sort circuits by size (largest to smallest)
	sortCircuitsBySize(circuits)

	//showCircuits(circuits)

	// Calculate answer
	solution := calculateCircuitSolution(circuits)
	return solution
}

func part2() int {
	return 0
}
