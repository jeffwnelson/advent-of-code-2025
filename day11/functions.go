package main

import (
	"os"
	"strings"
)

func parseInput(filename string) map[string][]string {
	file, _ := os.ReadFile(filename)

	devices := make(map[string][]string)
	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		deviceName := parts[0]
		outputs := strings.Fields(parts[1])
		devices[deviceName] = outputs
	}

	return devices
}

// memo caches path counts between node pairs to avoid redundant calculations
var cache = make(map[string]int)

func countPaths(start, end string, devices map[string][]string) int {
	// Check if we've already seen this path
	// Switch to Memoization / Caching to help:
	// https://fullstackdojo.medium.com/memoization-in-golang-a946acd10829
	key := start + ">" + end
	if cached, ok := cache[key]; ok {
		return cached
	}

	// Base case: reached the destination
	if start == end {
		return 1
	}

	// Recursively count paths through each neighbor
	count := 0
	for _, next := range devices[start] {
		count += countPaths(next, end, devices)
	}

	// Cache the result before returning
	cache[key] = count
	return count
}

func countPathsContains(start, end, contains1, contains2 string, devices map[string][]string) int {
	// Count paths that go through both contains1 and contains2 (in either order)
	// Path: start -> contains1 -> contains2 -> end
	dacPaths := countPaths(start, contains1, devices) *
		countPaths(contains1, contains2, devices) *
		countPaths(contains2, end, devices)

	// Path: start -> contains2 -> contains1 -> end
	fftPaths := countPaths(start, contains2, devices) *
		countPaths(contains2, contains1, devices) *
		countPaths(contains1, end, devices)

	return dacPaths + fftPaths
}
