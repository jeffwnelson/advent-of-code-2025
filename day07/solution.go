package main

import (
	"advent-of-code-2025/utils"
)

func main() {
	grid1 = parseInput("inputs/day07/input.txt")
	grid2 = copyGrid(grid1)

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var grid1 [][]rune
var grid2 [][]rune
var beams [][]int

func part1() int {
	splitTotal := 0
	for row := 0; row < len(grid1); row++ {
		for col := 0; col < len(grid1[row]); col++ {
			splitTotal += checkCell(grid1, row, col)
		}
	}
	return splitTotal
}

func part2() int {
	// Initialize beams grid
	beams = make([][]int, len(grid2))
	for i := range beams {
		beams[i] = make([]int, len(grid2[i]))
	}

	timelines := 1
	for row := 0; row < len(grid2); row++ {
		for col := 0; col < len(grid2[row]); col++ {
			timelines += checkCellWithBeams(grid2, row, col)
		}
	}
	return timelines
}

// 3278 is too low
