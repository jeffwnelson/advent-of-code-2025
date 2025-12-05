package main

import (
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) ([][2]int, []int) {
	data, _ := os.ReadFile(filename)

	sections := strings.Split(string(data), "\n\n")

	// Ingredient IDs
	var database [][2]int
	for _, line := range strings.Split(strings.TrimSpace(sections[0]), "\n") {
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			database = append(database, [2]int{start, end})
		}
	}

	// Available Ingredient IDs
	var ingredients []int
	if len(sections) > 1 {
		for _, line := range strings.Split(strings.TrimSpace(sections[1]), "\n") {
			id, _ := strconv.Atoi(line)
			ingredients = append(ingredients, id)
		}
	}

	return database, ingredients
}
