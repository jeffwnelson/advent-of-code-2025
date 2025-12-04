package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(filename string) ([][]rune, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid, nil
}

func showGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func getCell(grid [][]rune, row, col int) rune {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return '.'
	}
	return grid[row][col]
}
