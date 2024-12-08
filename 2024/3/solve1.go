package main

import (
	"fmt"
	"regexp"
	"strconv"
	"github.com/tire-fire/adventofcode/2024/lib"
)

func parseAndCalculate(lines []string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	totalSum := 0

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			// Extract X and Y from the match and convert to integers
			x, err1 := strconv.Atoi(match[1])
			y, err2 := strconv.Atoi(match[2])

			if err1 == nil && err2 == nil {
				totalSum += x * y
			}
		}
	}

	return totalSum
}

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	result := parseAndCalculate(lines)
	fmt.Println(result)
}

