package main

import (
	"fmt"
	"github.com/tire-fire/adventofcode/2024/lib"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func main() {
	lines, err := ReadInput()
	if err != nil {
		fmt.Fatalf("Failed to read input: %v", err)
	}

	rows := len(lines)

	grid := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		grid[i] = []rune(lines[i])
	}

	guardR, guardC, guardD := findGuardStart(grid)
	grid[guardR][guardC] = '.'

	visitedCount := simulatePatrol(grid, guardR, guardC, guardD)

	fmt.Println(visitedCount)
}

func findGuardStart(grid [][]rune) (int, int, Direction) {
	rows := len(grid)
	cols := len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			switch grid[r][c] {
			case '^':
				return r, c, Up
			case 'v':
				return r, c, Down
			case '<':
				return r, c, Left
			case '>':
				return r, c, Right
			}
		}
	}
	panic("No guard start found")
}

func simulatePatrol(grid [][]rune, guardR, guardC int, guardD Direction) int {
	rows := len(grid)
	cols := len(grid[0])
	visited := make(map[[2]int]struct{})

	// visited starting position
	visited[[2]int{guardR, guardC}] = struct{}{}

	for {
		nr, nc, nd := findNext(grid, guardR, guardC, guardD, rows, cols)
		if nr == -1 {
			break
		}
		guardR, guardC, guardD = nr, nc, nd
		visited[[2]int{guardR, guardC}] = struct{}{}
	}

	return len(visited)
}

func findNext(grid [][]rune, r, c int, d Direction, rows, cols int) (int, int, Direction) {
	// Put it in a loop if there's multiple obstacles
	for {
		nr, nc := stepForward(r, c, d)
		if outOfBounds(nr, nc, rows, cols) {
			return -1, -1, d
		}

		if grid[nr][nc] != '#' {
			// move forward.
			return nr, nc, d
		}

		// blocked, turn right
		d = turnRight(d)
	}
}

func turnRight(d Direction) Direction {
	return (d + 1) % 4
}

func stepForward(r, c int, d Direction) (int, int) {
	switch d {
	case Up:
		return r - 1, c
	case Right:
		return r, c + 1
	case Down:
		return r + 1, c
	case Left:
		return r, c - 1
	}
	return r, c
}

func outOfBounds(r, c, rows, cols int) bool {
	return r < 0 || c < 0 || r >= rows || c >= cols
}
