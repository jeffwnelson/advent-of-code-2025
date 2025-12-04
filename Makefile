## When nothing is passed, default to help
.DEFAULT_GOAL := help

## Force to run even if conditions are true
FORCE:

## Show available commands
help:
	@echo "Usage: make <command>"
	@echo ""
	@echo "Commands:"
	@echo "  day<XX>      Run solution for day XX (e.g., make day01)"
	@echo "  new<XX>      Create new day XX with template (e.g., make new02)"
	@echo "  delete<XX>   Delete day XX and its inputs (e.g., make delete02)"
	@echo "  help         Show this help message"

## Run a day: make day01
day%: FORCE
	@go run ./day$*

## Create a new day: make new01
new%: FORCE
	@mkdir -p day$* inputs/day$*
	@touch inputs/day$*/input.txt inputs/day$*/sample.txt
	@if [ ! -f day$*/solution.go ]; then \
		echo 'package main\n\nimport "advent-of-code-2025/utils"\n\nfunc main() {\n\t// TODO: parse input from ../inputs/day$*/input.txt\n\n\tutils.Run("Part 1", part1)\n\tutils.Run("Part 2", part2)\n}\n\nfunc part1() int {\n\treturn 0\n}\n\nfunc part2() int {\n\treturn 0\n}' > day$*/solution.go; \
	fi
	@echo "Created day$*/ and inputs/day$*/"

## Delete a day: make delete01
delete%: FORCE
	@read -p "Delete day$* and inputs/day$*? [y/N] " confirm && [ "$$confirm" = "y" ] && rm -rf day$* inputs/day$* && echo "Deleted day$*/ and inputs/day$*/" || echo "Cancelled"
