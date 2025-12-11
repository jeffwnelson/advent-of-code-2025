package main

import (
	"advent-of-code-2025/utils"
)

func main() {
	redTiles = parseInput("inputs/day09/input.txt")

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
	// Calculate the perimeter and save it
	// 'R' = Red tile, 'G' = Green tile
	var perimeter []Tile
	for _, p := range redTiles {
		perimeter = append(perimeter, Tile{Position: p, Type: 'R'})
	}

	for i := 0; i < len(redTiles); i++ {
		p1 := redTiles[i]
		p2 := redTiles[(i+1)%len(redTiles)] // wraps around

		xStep := 1
		if p2.X < p1.X {
			xStep = -1
		}
		yStep := 1
		if p2.Y < p1.Y {
			yStep = -1
		}

		if p1.X != p2.X {
			for x := p1.X + xStep; x != p2.X; x += xStep {
				perimeter = append(perimeter, Tile{Position: Point{X: x, Y: p1.Y}, Type: 'G'})
			}
		}
		if p1.Y != p2.Y {
			for y := p1.Y + yStep; y != p2.Y; y += yStep {
				perimeter = append(perimeter, Tile{Position: Point{X: p2.X, Y: y}, Type: 'G'})
			}
		}
	}

	return 0
}
