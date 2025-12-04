package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
)

func main() {
	var err error
	grid, err = parseInput("inputs/day04/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
		os.Exit(1)
	}

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var grid [][]rune

func part1() int {
	rolls := 0
	maxAdjacentRolls := 3
	for row, rowData := range grid {
		for col := range rowData {
			if grid[row][col] == '@' {
				adjacentRolls := 0

				// check top left
				if getCell(grid, row-1, col-1) == '@' {
					adjacentRolls++
				}

				// check top
				if getCell(grid, row-1, col) == '@' {
					adjacentRolls++
				}

				// check top right
				if getCell(grid, row-1, col+1) == '@' {
					adjacentRolls++
				}

				// check left
				if getCell(grid, row, col-1) == '@' {
					adjacentRolls++
				}

				// check right
				if getCell(grid, row, col+1) == '@' {
					adjacentRolls++
				}

				// check bottom left
				if getCell(grid, row+1, col-1) == '@' {
					adjacentRolls++
				}

				// check bottom
				if getCell(grid, row+1, col) == '@' {
					adjacentRolls++
				}

				// check bottom right
				if getCell(grid, row+1, col+1) == '@' {
					adjacentRolls++
				}

				if adjacentRolls <= maxAdjacentRolls {
					rolls += 1
				}
			}
		}
	}
	return rolls
}

func part2() int {
	totalRolls := 0
	maxAdjacentRolls := 3
	for {
		rolls := 0
		for row, rowData := range grid {
			for col := range rowData {
				if grid[row][col] == '@' {
					adjacentRolls := 0

					// check top left
					if getCell(grid, row-1, col-1) == '@' {
						adjacentRolls++
					}
					// check top
					if getCell(grid, row-1, col) == '@' {
						adjacentRolls++
					}

					// check top right
					if getCell(grid, row-1, col+1) == '@' {
						adjacentRolls++
					}

					// check left
					if getCell(grid, row, col-1) == '@' {
						adjacentRolls++
					}

					// check right
					if getCell(grid, row, col+1) == '@' {
						adjacentRolls++
					}

					// check bottom left
					if getCell(grid, row+1, col-1) == '@' {
						adjacentRolls++
					}

					// check bottom
					if getCell(grid, row+1, col) == '@' {
						adjacentRolls++
					}

					// check bottom right
					if getCell(grid, row+1, col+1) == '@' {
						adjacentRolls++
					}

					if adjacentRolls <= maxAdjacentRolls {
						grid[row][col] = '.'
						rolls++
					}
				}
			}
		}
		if rolls == 0 {
			break
		}
		totalRolls += rolls
	}
	return totalRolls
}
