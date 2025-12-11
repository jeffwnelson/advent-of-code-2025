package main

import "advent-of-code-2025/utils"

func main() {
	devices = parseInput("inputs/day11/input.txt")

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var devices map[string][]string

func part1() int {
	return countPaths("you", "out", devices)
}

func part2() int {
	return countPathsContains("svr", "out", "dac", "fft", devices)
}
