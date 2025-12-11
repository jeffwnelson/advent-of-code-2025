package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) []Point {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var points []Point
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{X: x, Y: y})
	}

	return points
}

func showAreas(results []AreaResult) {
	for _, r := range results {
		fmt.Printf("[%d,%d][%d,%d] = %d\n", r.P1.X, r.P1.Y, r.P2.X, r.P2.Y, r.Area)
	}
}

func showPerimeter(perimeter []Tile) {
	for _, tile := range perimeter {
		fmt.Printf("[%d,%d] %c\n", tile.Position.X, tile.Position.Y, tile.Type)
	}
}

func countTileTypes(perimeter []Tile) (red int, green int) {
	for _, tile := range perimeter {
		if tile.Type == 'R' {
			red++
		} else if tile.Type == 'G' {
			green++
		}
	}
	return red, green
}
