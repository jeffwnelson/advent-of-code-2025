package main

import (
	"bufio"
	"os"
	"strconv"
)

func parseInput(filename string) ([]Instruction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var instructions []Instruction
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		rotation := string(line[0])
		travel, _ := strconv.Atoi(line[1:])

		instructions = append(instructions, Instruction{
			Rotation: rotation,
			Travel:   travel,
		})
	}

	return instructions, nil
}

func countZerosCrossed(oldPosition, newPosition int) int {
	count := 0
	if oldPosition < newPosition { // Moving right
		for pos := oldPosition + 1; pos <= newPosition; pos++ {
			if pos%100 == 0 {
				count++
			}
		}
	} else { // Moving left
		for pos := oldPosition - 1; pos >= newPosition; pos-- {
			if ((pos%100)+100)%100 == 0 {
				count++
			}
		}
	}
	return count
}
