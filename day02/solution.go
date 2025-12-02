package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
)

func main() {
	var err error
	productIDRanges, err = parseInput("inputs/day02/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
		os.Exit(1)
	}

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var productIDRanges []ProductIDRange

func part1() int {
	totalInvalidIDs := 0
	for _, id := range productIDRanges {
		checkID := id.FirstID
		for {
			if checkID > id.LastID {
				break
			}

			if isDigitLengthEven(checkID) {
				a, b := splitEven(checkID)

				if a == b {
					totalInvalidIDs = totalInvalidIDs + checkID
				}
			}
			checkID++
		}
	}

	return totalInvalidIDs
}

func part2() int {
	totalInvalidIDs := 0
	for _, id := range productIDRanges {
		checkID := id.FirstID
		for {
			if checkID > id.LastID {
				break
			}

			if isRepeatedPatternInt(checkID) {
				totalInvalidIDs = totalInvalidIDs + checkID
			}
			checkID++
		}
	}

	return totalInvalidIDs
}
