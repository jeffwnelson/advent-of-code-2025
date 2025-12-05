package main

import (
	"advent-of-code-2025/utils"
	"sort"
)

func main() {
	database, ingredients = parseInput("inputs/day05/input.txt")

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var database [][2]int
var ingredients []int

func part1() int {
	freshIngredients := 0
	for _, ingredient := range ingredients {
		for _, id := range database {
			if ingredient >= id[0] && ingredient <= id[1] {
				freshIngredients++
				break // We found a match, so stop checking the rest
			}
		}
	}

	return freshIngredients
}

func part2() int {
	// Sort by starting value
	// Found this, which saved me: https://stackoverflow.com/a/59897376
	sort.Slice(database, func(i, j int) bool {
		return database[i][0] < database[j][0]
	})

	for i := 0; i < len(database); {
		// If our next range starting value is <= current range ending value +1
		if database[i+1][0] <= database[i][1]+1 {
			// Make our current range's ending value the highest value between current range's starting and next range's ending value
			database[i][1] = max(database[i][1], database[i+1][1])
			// Replace the current range in the database by removing the old range entry and inserting the updated one, another save: https://stackoverflow.com/a/37335777
			database = append(database[:i+1], database[i+2:]...)
			// If not, continue to the next range
		} else {
			i++
		}
		// check if we have another range to process, if not end
		if i+1 >= len(database) {
			break
		}
	}

	totalIngredients := 0
	for _, ids := range database {
		totalIngredients += ids[1] - ids[0] + 1
	}

	return totalIngredients
}
