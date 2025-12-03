package main

import (
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) ([][]int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	banks := make([][]int, len(lines))

	for i, line := range lines {
		bank := make([]int, len(line))
		for j, ch := range line {
			bank[j] = int(ch - '0')
		}
		banks[i] = bank
	}

	return banks, nil
}

func largestJoltage(bank []int, count int) int {
	positions := make([]int, count) // Create a new slide of size "count"
	startingIndex := 0              // Start at 0

	for i := 0; i < count; i++ {
		positionsNeeded := count - i - 1            // How many positions are left to fill.
		lastPosition := len(bank) - positionsNeeded // Last position is size of our bank
		largestIndex := startingIndex               // Largest index is our first position in the array
		largestValue := bank[startingIndex]         // Largest value is our first number in the array

		for j := startingIndex; j < lastPosition; j++ {
			if bank[j] > largestValue { // If value at bank is larger than our current known largest...
				largestValue = bank[j] //   Save this value as our new largestValue
				largestIndex = j       //   Save index as our new largestIndex
			}
		}
		positions[i] = largestIndex      // Save the largestIndex to our i position
		startingIndex = largestIndex + 1 // Update our startingIndex to +1 from our largestIndex
	}

	// We should now have an array of positions, "count" in size, sorted from largest and decreasing

	text := ""
	for _, position := range positions { // For each position in our array...
		text += strconv.Itoa(bank[position]) // Get the value from our bank in that position and concat to our text string
	}
	value, _ := strconv.Atoi(text) // Convert our string to an int
	return value
}
