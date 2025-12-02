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