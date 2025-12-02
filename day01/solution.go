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
		position = utils.Modulo(position+distance, 100)
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
		step := 1
		if distance < 0 {
			step = -1
		}
		for i := 0; i < utils.Absolute(distance); i++ {
			position = utils.Modulo(position+step, 100)
			if position == 0 {
				password++
			}
		}
	}
	return password
}
