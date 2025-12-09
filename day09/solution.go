package main

import (
	"advent-of-code-2025/utils"
	"fmt"
)

func main() {
	redTiles = parseInput("inputs/day09/sample.txt")

	utils.Run("Part 1", part1)
	utils.Run("Part 2", part2)
}

var redTiles []Point

func part1() int {
	//showTiles(redTiles)
	largestArea := 0
	for i := 0; i < len(redTiles); i++ {
		for j := 0; j < len(redTiles); j++ {
			p1 := redTiles[i]
			p2 := redTiles[j]
			width := utils.Absolute(p2.X-p1.X) + 1
			height := utils.Absolute(p2.Y-p1.Y) + 1
			area := width * height
			//fmt.Printf("[%d,%d] - [%d,%d] = %d", p1.X, p1.Y, p2.X, p2.Y, area)
			if area > largestArea {
				largestArea = area
				//fmt.Printf(" <-- Largest\n")
			}
		}
	}
	return largestArea
}

func part2() int {
	var greenTiles []Point

	for i := 0; i < len(redTiles)-1; i++ {
		p1 := redTiles[i]
		p2 := redTiles[i+1]

		// Determine direction for X
		xStep := 1
		if p2.X < p1.X {
			xStep = -1
		}

		// Determine direction for Y
		yStep := 1
		if p2.Y < p1.Y {
			yStep = -1
		}

		// Add all points between p1 and p2 (along both axes)
		if p1.X != p2.X {
			for x := p1.X + xStep; x != p2.X; x += xStep {
				greenTiles = append(greenTiles, Point{X: x, Y: p1.Y})
			}
		}
		if p1.Y != p2.Y {
			for y := p1.Y + yStep; y != p2.Y; y += yStep {
				greenTiles = append(greenTiles, Point{X: p2.X, Y: y})
			}
		}
	}

	fmt.Println(greenTiles)
	return len(greenTiles)
}
