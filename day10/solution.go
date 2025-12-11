package main

import (
	"advent-of-code-2025/utils"
)

func main() {
	machines = parseInput("inputs/day10/input.txt")

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var machines []Machine
var solutions [][]Solution // solutions for each machine

func part1() int {
	solutions = make([][]Solution, len(machines))

	for i, m := range machines {
		numButtons := len(m.Buttons)
		pressed := make([]bool, numButtons)

		// Try all possible combinations of button presses
		for {
			// Start with all lights off
			lights := make([]rune, len(m.Lights))
			for light := range lights {
				lights[light] = '.'
			}

			// Store button presses
			var pressedButtons []int

			// Check each button to see if it's pressed in this combination
			for button := 0; button < numButtons; button++ {
				if pressed[button] {
					// button is pressed, toggle the light
					for _, pos := range m.Buttons[button] {
						if lights[pos] == '.' {
							lights[pos] = '#'
						} else {
							lights[pos] = '.'
						}
					}
					pressedButtons = append(pressedButtons, button)
				}
			}

			// Check if this button combo matches the target light pattern
			if string(lights) == m.Lights {
				solutions[i] = append(solutions[i], Solution{
					Buttons: pressedButtons,
					Count:   len(pressedButtons),
				})
			}

			// Increment like a binary counter
			carry := true
			for j := 0; carry && j < numButtons; j++ {
				if pressed[j] {
					pressed[j] = false
				} else {
					pressed[j] = true
					carry = false
				}
			}
			if carry {
				break // overflowed, done with all combinations
			}
		}

		//fmt.Printf("Machine %d: found %d solutions\n", i, len(solutions[i]))
		//for _, solution := range solutions[i] {
		//	fmt.Printf("  %d buttons: %v\n", solution.Count, solution.Buttons)
		//}
	}

	// Sum the fewest button presses for each machine
	total := 0
	for _, machineSolutions := range solutions {
		// Find the solution with minimum button presses
		minPresses := machineSolutions[0].Count
		for _, solution := range machineSolutions[1:] {
			if solution.Count < minPresses {
				minPresses = solution.Count
			}
		}
		total += minPresses
	}
	return total
}

func part2() int {
	return 0
}
