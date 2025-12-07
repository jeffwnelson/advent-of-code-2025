package main

import (
	"advent-of-code-2025/utils"
)

func main() {
	grid, problems = parseInput("inputs/day06/input.txt")

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var grid [][]rune
var problems []Problem

func part1() int {
	total := 0
	for _, problem := range problems {
		numbers := getRowNumbers(grid, problem.Start, problem.End)
		total += maths(numbers, string(problem.Operation))
	}
	return total
}

func part2() int {
	total := 0
	// Process right-to-left
	for i := len(problems) - 1; i >= 0; i-- {
		problem := problems[i]
		numbers := getColNumbers(grid, problem.Start, problem.End)
		total += maths(numbers, string(problem.Operation))
	}
	return total
}
