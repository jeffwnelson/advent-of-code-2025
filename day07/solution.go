package main

import (
	"advent-of-code-2025/utils"
)

func main() {
	grid = parseInput("inputs/day07/input.txt")

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var grid [][]rune
var beams [][]int // Creating a beam map to track how many beams are each location

func part1() int {
	initBeams(grid) // not used for our part1 solution
	splitTotal := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			splits, _ := checkCell(grid, row, col)
			splitTotal += splits
		}
	}
	return splitTotal
}

func part2() int {
	initBeams(grid)
	timelinesTotal := 1
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			_, timelines := checkCell(grid, row, col)
			timelinesTotal += timelines
		}
	}
	return timelinesTotal
}
