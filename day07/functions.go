package main

import (
	"fmt"
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

func showGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func saveGrid(grid [][]rune, filename string) {
	var lines []string
	for _, row := range grid {
		lines = append(lines, string(row))
	}
	os.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
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

func splitBeam(grid [][]rune, row, col int) {
	grid[row][col-1] = '|'
	grid[row][col+1] = '|'
}

// Check if the cell directly above us is an "S"
func checkCell(grid [][]rune, row, col int) int {
	splitCount := 0
	if getCell(grid, row-1, col) == 'S' {
		if tachyonCheck(grid, row, col) {
			splitBeam(grid, row, col)
			splitCount++
		} else {
			grid[row][col] = '|'
		}
	} else if getCell(grid, row-1, col) == '|' {
		if tachyonCheck(grid, row, col) {
			splitBeam(grid, row, col)
			splitCount++
		} else {
			grid[row][col] = '|'
		}
	}
	return splitCount
}

// checkCellWithBeams tracks beam counts for accurate timeline calculation
func checkCellWithBeams(grid [][]rune, row, col int) int {
	timelines := 0
	beamCount := getBeam(row-1, col)

	if getCell(grid, row-1, col) == 'S' {
		beamCount = 1
	}

	if beamCount > 0 {
		if tachyonCheck(grid, row, col) {
			// Each beam splits into 2, adding beamCount new timelines
			addBeam(row, col-1, beamCount)
			addBeam(row, col+1, beamCount)
			timelines = beamCount
		} else {
			// Beams continue straight down
			addBeam(row, col, beamCount)
		}
	}
	return timelines
}

func getBeam(row, col int) int {
	if row < 0 || row >= len(beams) || col < 0 || col >= len(beams[row]) {
		return 0
	}
	return beams[row][col]
}

func addBeam(row, col, count int) {
	if row < 0 || row >= len(beams) || col < 0 || col >= len(beams[row]) {
		return
	}
	beams[row][col] += count
}
