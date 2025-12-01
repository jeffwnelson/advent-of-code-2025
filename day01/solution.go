package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"os"
)

func main() {
	var err error
	instructions, err = parseInput("inputs/day01/input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
		os.Exit(1)
	}

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var instructions []Instruction

func part1() int {
	position := 50
	password := 0

	for _, instruction := range instructions {
		distance := instruction.Travel
		if instruction.Rotation == "L" {
			distance = -distance
		}
		position = ((position+distance)%100 + 100) % 100
		if position == 0 {
			password++
		}
	}
	return password
}

func part2() int {
	position := 50
	password := 0

	for _, instruction := range instructions {
		distance := instruction.Travel
		if instruction.Rotation == "L" {
			distance = -distance
		}
		newPosition := position + distance
		password += countZerosCrossed(position, newPosition)
		position = ((newPosition % 100) + 100) % 100
	}
	return password
}
