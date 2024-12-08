package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/tire-fire/adventofcode/2024/lib"
)

func parseInput(lines []string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	var updates [][]int
	isUpdateSection := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			isUpdateSection = true
			continue
		}

		if !isUpdateSection {
			// Parse rules
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				panic("Invalid rule format")
			}
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			rules[x] = append(rules[x], y)
		} else {
			// Parse updates
			parts := strings.Split(line, ",")
			var update []int
			for _, p := range parts {
				num, _ := strconv.Atoi(p)
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates
}

func processUpdates(rules map[int][]int, updates [][]int) int {
	total := 0

	for _, update := range updates {
		if isValidUpdate(rules, update) {
			// Calculate the middle page number
			middleIndex := len(update) / 2
			total += update[middleIndex]
		}
	}

	return total
}

func isValidUpdate(rules map[int][]int, update []int) bool {
	pagePositions := make(map[int]int)

	// Map each page to its position in the update
	for i, page := range update {
		pagePositions[page] = i
	}

	// Validate all rules applicable to the update
	for x, dependencies := range rules {
		for _, y := range dependencies {
			posX, existsX := pagePositions[x]
			posY, existsY := pagePositions[y]

			// Ignore rules where one or both pages are not in the update
			if !existsX || !existsY {
				continue
			}

			// Check if the rule is violated
			if posX >= posY {
				return false
			}
		}
	}

	return true
}

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	rules, updates := parseInput(lines)

	result := processUpdates(rules, updates)
	fmt.Println(result)
}
