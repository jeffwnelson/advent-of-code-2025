package main

import (
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	Operation rune
	Start     int
	End       int
}

func parseInput(filename string) ([][]rune, []Problem) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")

	var grid [][]rune
	for _, line := range lines[:len(lines)-1] {
		grid = append(grid, []rune(line))
	}

	operators := []rune(lines[len(lines)-1])

	// Find operator positions and create problems
	var operatorPositions []int
	for i, ch := range operators {
		if ch == '+' || ch == '*' {
			operatorPositions = append(operatorPositions, i)
		}
	}

	var problems []Problem
	for i, opPos := range operatorPositions {
		start := opPos
		var end int
		if i+1 < len(operatorPositions) {
			end = operatorPositions[i+1] - 1 // -1 for separator
		} else {
			end = len(grid[0])
		}
		problems = append(problems, Problem{
			Operation: operators[opPos],
			Start:     start,
			End:       end,
		})
	}

	return grid, problems
}

// getRowNumbers extracts one number per row from a column range (start to end)
func getRowNumbers(grid [][]rune, start, end int) []int {
	var numbers []int // accumulates the number from each row

	for row := 0; row < len(grid); row++ { // iterate through each row
		e := end                // local copy of end for bounds adjustment
		if e > len(grid[row]) { // bounds check: clamp end to row length
			e = len(grid[row])
		}
		if start >= len(grid[row]) { // skip if start is beyond this row
			continue
		}

		segment := string(grid[row][start:e]) // extract substring for this problem's columns
		trimmed := strings.TrimSpace(segment) // remove leading/trailing spaces

		if trimmed != "" { // if there's a number in this row
			num, _ := strconv.Atoi(trimmed) // parse the whole number
			numbers = append(numbers, num)  // add to our list
		}
	}

	return numbers
}

// getColNumbers extracts numbers by reading columns right-to-left within a range
func getColNumbers(grid [][]rune, start, end int) []int {
	var numbers []int // accumulates numbers from each column

	for col := end - 1; col >= start; col-- { // iterate columns right-to-left
		num := columnConcatenate(grid, col) // get number from this column's digits
		if num > 0 {                        // skip empty columns (no digits)
			numbers = append(numbers, num) // add to our list
		}
	}

	return numbers
}

// columnConcatenate reads digits vertically from a single column (top to bottom)
func columnConcatenate(grid [][]rune, col int) int {
	var digits string // accumulates digit characters as we read down the column

	for row := 0; row < len(grid); row++ { // iterate through each row (top to bottom)
		if col < len(grid[row]) { // bounds check: ensure column exists in this row
			char := grid[row][col] // get the character at this row/column position

			if char >= '0' && char <= '9' { // check if it's a digit (0-9)
				digits += string(char) // append digit to our string
			}
		}
	}

	if digits == "" { // no digits found in this column
		return 0
	}

	num, _ := strconv.Atoi(digits) // convert to integer
	return num
}

// maths applies the operator to our numbers
func maths(numbers []int, operation string) int {
	result := 0 // default for adding
	if operation == "*" {
		result = 1 // default for multiplying (since multiplying by 0 is just 0)
	}

	for _, num := range numbers { // iterate through all numbers
		if operation == "+" {
			result += num // add to running total
		} else if operation == "*" {
			result *= num // multiply to running total
		}
	}

	return result
}
