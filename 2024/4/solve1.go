package main

import (
	"fmt"
	"strings"
	"github.com/tire-fire/adventofcode/2024/lib"
)

var directions = [][2]int{
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 0},   // Down
	{-1, 0},  // Up
	{1, 1},   // Diagonal Down-Right
	{-1, -1}, // Diagonal Up-Left
	{1, -1},  // Diagonal Down-Left
	{-1, 1},  // Diagonal Up-Right
}

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

func findWordCount(grid [][]rune, word string) int {
	wordLen := len(word)
	count := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			for _, dir := range directions {
				if matchesWord(grid, word, row, col, dir, wordLen) {
					count++
				}
			}
		}
	}
	return count
}

func matchesWord(grid [][]rune, word string, row, col int, dir [2]int, wordLen int) bool {
	for i := 0; i < wordLen; i++ {
		newRow := row + i*dir[0]
		newCol := col + i*dir[1]
		if !isValid(grid, newRow, newCol) || grid[newRow][newCol] != rune(word[i]) {
			return false
		}
	}
	return true
}

func main() {
	lines, err := ReadInput()
	if err != nil {
		fmt.Fatalf("Failed to read input: %v", err)
	}

	grid := parseGrid(lines)
	count := findWordCount(grid, "XMAS")
	fmt.Println(count)
}
