package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(filename string) []Machine {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var machines []Machine
	for _, line := range lines {
		if line == "" {
			continue
		}

		var machine Machine

		// Parse lights: [...]
		lightsRe := regexp.MustCompile(`\[([^\]]*)\]`)
		lightsMatch := lightsRe.FindStringSubmatch(line)
		if len(lightsMatch) > 1 {
			machine.Lights = lightsMatch[1]
		}

		// Parse buttons: multiple (...)
		buttonsRe := regexp.MustCompile(`\(([^)]*)\)`)
		buttonsMatches := buttonsRe.FindAllStringSubmatch(line, -1)
		for _, match := range buttonsMatches {
			if len(match) > 1 {
				var nums []int
				parts := strings.Split(match[1], ",")
				for _, p := range parts {
					n, _ := strconv.Atoi(strings.TrimSpace(p))
					nums = append(nums, n)
				}
				machine.Buttons = append(machine.Buttons, nums)
			}
		}

		// Parse joltage: {...}
		joltageRe := regexp.MustCompile(`\{([^}]*)\}`)
		joltageMatch := joltageRe.FindStringSubmatch(line)
		if len(joltageMatch) > 1 {
			parts := strings.Split(joltageMatch[1], ",")
			for _, p := range parts {
				n, _ := strconv.Atoi(strings.TrimSpace(p))
				machine.Joltage = append(machine.Joltage, n)
			}
		}

		machines = append(machines, machine)
	}

	return machines
}
