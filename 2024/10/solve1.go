package main

import (
	"fmt"
	"github.com/tire-fire/adventofcode/2024/lib"
)

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	// Parse the topographic map
	grid := parseMap(lines)

	// Find all trailheads (height 0 positions)
	trailheads := findTrailheads(grid)

	// Calculate the total score of all trailheads
	totalScore := 0
	for _, trailhead := range trailheads {
		totalScore += calculateTrailheadScore(grid, trailhead)
	}

	fmt.Println(totalScore)
}

// parseMap parses the input into a 2D grid of integers
func parseMap(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			grid[i][j] = int(char - '0')
		}
	}
	return grid
}

// findTrailheads finds all positions with height 0
func findTrailheads(grid [][]int) [][2]int {
	trailheads := [][2]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				trailheads = append(trailheads, [2]int{i, j})
			}
		}
	}
	return trailheads
}

// calculateTrailheadScore calculates the score for a single trailhead
func calculateTrailheadScore(grid [][]int, trailhead [2]int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// BFS to find reachable 9s
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	queue := [][2]int{trailhead}
	visited[trailhead[0]][trailhead[1]] = true

	reachableNines := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		curRow, curCol := current[0], current[1]
		curHeight := grid[curRow][curCol]

		for _, dir := range directions {
			newRow, newCol := curRow+dir[0], curCol+dir[1]

			// Check bounds
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				continue
			}

			// Check if already visited or not a valid step
			if visited[newRow][newCol] || grid[newRow][newCol] != curHeight+1 {
				continue
			}

			// Mark as visited and enqueue
			visited[newRow][newCol] = true
			queue = append(queue, [2]int{newRow, newCol})

			// Check if it's a 9
			if grid[newRow][newCol] == 9 {
				reachableNines++
			}
		}
	}

	return reachableNines
}

