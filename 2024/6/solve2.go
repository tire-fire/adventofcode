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
	cols := len(lines[0])

	grid := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		grid[i] = []rune(lines[i])
	}

	guardR, guardC, guardD := findGuardStart(grid)
	grid[guardR][guardC] = '.'

	visited, visitedCount := simulatePatrol(grid, guardR, guardC, guardD)

	fmt.Println(visitedCount)

	loopCount := 0
	for pos, isCandidate := range visited {
		if !isCandidate {
			continue
		}
		r := pos / cols
		c := pos % cols
		if r == guardR && c == guardC {
			continue
		}
		if grid[r][c] == '.' {
			grid[r][c] = '#'
			if causesLoop(grid, guardR, guardC, guardD, rows, cols) {
				loopCount++
			}
			grid[r][c] = '.'
		}
	}

	fmt.Println(loopCount)
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

func simulatePatrol(grid [][]rune, guardR, guardC int, guardD Direction) ([]bool, int) {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([]bool, rows*cols)
	rounds := 0

	// visited starting position
	visited[guardR*cols+guardC] = true

	for {
		nr, nc, nd := findNext(grid, guardR, guardC, guardD, rows, cols)
		if nr == -1 {
			break
		}
		guardR, guardC, guardD = nr, nc, nd
		visited[guardR*cols+guardC] = true
		rounds++
	}
	return visited, rounds
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

func causesLoop(grid [][]rune, guardR, guardC int, guardD Direction, rows, cols int) bool {
	// store seen states in a boolean slice
	// stateIdx = ((r*cols)+c)*4 + int(d)
	maxStates := rows * cols * 4
	seen := make([]bool, maxStates)

	for {
		nr, nc, nd := findNext(grid, guardR, guardC, guardD, rows, cols)
		if nr == -1 {
			// guard leaves the area
			return false
		}
		guardR, guardC, guardD = nr, nc, nd

		stateIdx := ((guardR*cols)+guardC)*4 + int(guardD)
		if seen[stateIdx] {
			// loop detected
			return true
		}
		seen[stateIdx] = true
	}
}

