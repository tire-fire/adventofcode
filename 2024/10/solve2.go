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

	grid := parseMap(lines)

	// Find all trailheads (height 0 positions)
	trailheads := findTrailheads(grid)

	// Calculate the total rating of all trailheads
	totalRating := 0
	for _, trailhead := range trailheads {
		totalRating += calculateTrailheadRating(grid, trailhead)
	}

	fmt.Println("Total Rating:", totalRating)
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

// calculateTrailheadRating calculates the rating (number of distinct trails) for a trailhead
func calculateTrailheadRating(grid [][]int, trailhead [2]int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(row, col, currentHeight int) int
	dfs = func(row, col, currentHeight int) int {
		// If out of bounds, or not the next valid step, return 0
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] != currentHeight {
			return 0
		}

		// If we've reached height 9, this is a valid trail
		if currentHeight == 9 {
			return 1
		}

		// Temporarily mark this cell to prevent revisiting
		originalHeight := grid[row][col]
		grid[row][col] = -1

		// Explore all four directions
		trails := 0
		for _, dir := range directions {
			trails += dfs(row+dir[0], col+dir[1], currentHeight+1)
		}

		// Restore the original height for future paths
		grid[row][col] = originalHeight
		return trails
	}

	// Start DFS from the trailhead
	return dfs(trailhead[0], trailhead[1], 0)
}

