package main

import (
	"fmt"
	"strings"
	"github.com/tire-fire/adventofcode/2024/lib"
)

func parseGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(strings.TrimSpace(line))
	}
	return grid
}

func isValid(grid [][]rune, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func findWordCount(grid [][]rune) int {
	count := 0
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			fmt.Printf("Checking center at (%d, %d)", row, col)
			if matchesWord(grid, row, col) {
				count++
			}
		}
	}
	return count
}

func matchesWord(grid [][]rune, row, col int) bool {
	var topLeft, topRight, center, bottomLeft, bottomRight rune

	if isValid(grid, row-1, col-1) {
		topLeft = grid[row-1][col-1]
	}
	if isValid(grid, row-1, col+1) {
		topRight = grid[row-1][col+1]
	}
	if isValid(grid, row, col) {
		center = grid[row][col]
	}
	if isValid(grid, row+1, col-1) {
		bottomLeft = grid[row+1][col-1]
	}
	if isValid(grid, row+1, col+1) {
		bottomRight = grid[row+1][col+1]
	}

	fmt.Printf("Top-left: %c, Top-right: %c, Center: %c, Bottom-left: %c, Bottom-right: %c",
		topLeft, topRight, center, bottomLeft, bottomRight)

	// Center must be 'A'
	if center != 'A' {
		fmt.Printf("Center (%d, %d) is not 'A', skipping.", row, col)
		return false
	}

	// Check top and bottom diagonals
	if (topLeft == 'M' || topLeft == 'S') &&
		(topRight == 'M' || topRight == 'S') &&
		(bottomLeft == 'M' || bottomLeft == 'S') &&
		(bottomRight == 'M' || bottomRight == 'S') &&
		(topLeft != bottomRight) && // Opposite diagonals
		(topRight != bottomLeft) {
		fmt.Printf("Matched X-MAS pattern at center (%d, %d)", row, col)
		return true
	}

	return false
}

func main() {
	lines, err := ReadInput()
	if err != nil {
		fmt.Fatalf("Failed to read input: %v", err)
	}

	grid := parseGrid(lines)
	count := findWordCount(grid)
	fmt.Println(count)
}

