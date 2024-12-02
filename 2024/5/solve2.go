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
				fmt.Fatalf("Invalid rule format: %s", line)
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


func processIncorrectUpdates(rules map[int][]int, updates [][]int) int {
	total := 0

	for _, update := range updates {
		if !isValidUpdate(rules, update) {
			// Fix the order using topological sort
			fixedUpdate := topologicalSort(rules, update)
			// Calculate the middle page number
			middleIndex := len(fixedUpdate) / 2
			total += fixedUpdate[middleIndex]
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

func topologicalSort(rules map[int][]int, update []int) []int {
	// Create a graph from the rules, restricted to the pages in the update
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	// Initialize in-degree for all pages in the update
	for _, page := range update {
		inDegree[page] = 0
	}

	// Build graph and in-degree based on the update
	for x, dependencies := range rules {
		for _, y := range dependencies {
			if contains(update, x) && contains(update, y) {
				graph[x] = append(graph[x], y)
				inDegree[y]++
			}
		}
	}

	var sorted []int
	var queue []int

	// Start with pages with in-degree 0
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		fmt.Println("queue:", len(queue))
		// Pop a page from the queue
		current := queue[0]
		queue = queue[1:]

		// Add it to the sorted order
		sorted = append(sorted, current)

		// Decrease in-degree for its neighbors
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	lines, err := ReadInput()
	if err != nil {
		fmt.Fatalf("Failed to read input: %v", err)
	}

	rules, updates := parseInput(lines)

	result := processIncorrectUpdates(rules, updates)
	fmt.Println(result)
}

