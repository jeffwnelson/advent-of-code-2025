package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
)

func main() {
	var err error
	banks, err = parseInput("inputs/day03/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
		os.Exit(1)
	}

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var banks [][]int

func part1() int {
	totalOverJoltage := 0
	for _, bank := range banks {
		totalOverJoltage += largestJoltage(bank, 2)
	}
	return totalOverJoltage
}

func part2() int {
	totalOverJoltage := 0
	for _, bank := range banks {
		totalOverJoltage += largestJoltage(bank, 12)
	}
	return totalOverJoltage
}
