package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/tire-fire/adventofcode/2024/lib"
)

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return false // A single level cannot be increasing or decreasing
	}

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		// Check if the difference is out of the acceptable range
		if diff < -3 || (diff > -1 && diff < 1) || diff > 3 {
			return false
		}

		// Update the trend flags
		if diff < 0 {
			isIncreasing = false
		}
		if diff > 0 {
			isDecreasing = false
		}
	}

	// Safe if the levels are all increasing or all decreasing
	return isIncreasing || isDecreasing
}

func isSafeWithDampener(levels []int) bool {
	// If the report is already safe, return true
	if isSafe(levels) {
		return true
	}

	// Try removing each level to see if it becomes safe
	for i := 0; i < len(levels); i++ {
		// Create a new slice without the current level
		modified := make([]int, 0, len(levels)-1)
		modified = append(modified, levels[:i]...)
		modified = append(modified, levels[i+1:]...)
		if isSafe(modified) {
			return true
		}
	}

	return false
}

func countSafeReports(reports [][]int) int {
	safeCount := 0

	for _, levels := range reports {
		if isSafeWithDampener(levels) {
			safeCount++
		}
	}

	return safeCount
}

func parseReports(lines []string) ([][]int, error) {
	var reports [][]int

	for _, line := range lines {
		nums := strings.Fields(line)

		var levels []int
		for _, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			levels = append(levels, num)
		}
		reports = append(reports, levels)
	}

	return reports, nil
}

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	reports, err := parseReports(lines)
	if err != nil {
		panic("Failed to parse reports")
	}

	safeCount := countSafeReports(reports)
	fmt.Println(safeCount)
}
