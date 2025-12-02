package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) ([]ProductIDRange, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	line := strings.TrimSpace(string(content))
	parts := strings.Split(line, ",")

	var productIDRanges []ProductIDRange
	for _, part := range parts {
		ids := strings.Split(part, "-")
		firstID, _ := strconv.Atoi(ids[0])
		lastID, _ := strconv.Atoi(ids[1])
		productIDRanges = append(productIDRanges, ProductIDRange{
			FirstID: firstID,
			LastID:  lastID,
		})
	}

	return productIDRanges, nil
}

func splitEven(n int) (string, string) {
	s := fmt.Sprintf("%d", n)
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func isDigitLengthEven(n int) bool {
	s := fmt.Sprintf("%d", n)
	return len(s)%2 == 0
}

func isRepeatedPatternInt(id int) bool {
	value := fmt.Sprintf("%d", id) // Convert int to string
	size := len(value)             // Get length of our string

	for i := 1; i <= size/2; i++ {
		if size%i != 0 {
			continue
		}

		pattern := value[:i]
		if strings.Repeat(pattern, size/i) != value {
			continue
		}

		return true
	}

	return false
}
