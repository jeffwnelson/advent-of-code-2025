package main

import (
	"os"
	"strings"
)

func parseInput(filename string) [][]rune {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	return grid
}

func copyGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for i, row := range grid {
		newGrid[i] = make([]rune, len(row))
		copy(newGrid[i], row)
	}
	return newGrid
}

func getCell(grid [][]rune, row, col int) rune {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
		return '.'
	}
	return grid[row][col]
}

func tachyonCheck(grid [][]rune, row, col int) bool {
	tachyon := false
	if getCell(grid, row, col) == '^' {
		tachyon = true
	}
	return tachyon
}

// splitCount = number of splits (for part 1)
// timelines = beam count at each split (for part 2)
func checkCell(grid [][]rune, row, col int) (int, int) {
	splitCount := 0
	timelines := 0
	beamCount := getBeam(row-1, col)

	if getCell(grid, row-1, col) == 'S' {
		beamCount = 1
	}

	if beamCount > 0 {
		if tachyonCheck(grid, row, col) {
			addBeam(row, col-1, beamCount)
			addBeam(row, col+1, beamCount)
			splitCount = 1
			timelines = beamCount
		} else {
			addBeam(row, col, beamCount)
		}
	}
	return splitCount, timelines
}

// Create the same size map as grid to store beam values
func initBeams(grid [][]rune) {
	beams = make([][]int, len(grid))
	for i := range beams {
		beams[i] = make([]int, len(grid[i]))
	}
}

func getBeam(row, col int) int {
	// Check if beam could be out of bounds
	if row < 0 || row >= len(beams) || col < 0 || col >= len(beams[row]) {
		return 0
	}
	return beams[row][col]
}

func addBeam(row, col, count int) {
	beams[row][col] += count
}
