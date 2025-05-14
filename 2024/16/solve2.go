package main

import (
	"container/list"
	"fmt"
	"strings"
	"math"
	"github.com/tire-fire/adventofcode/2024/lib"
)

type Point struct {
	x, y int
}

// Use integers to represent directions
const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

var directions = []struct {
	delta    Point
	turnDirs [2]int
}{
	{Point{-1, 0}, [2]int{RIGHT, LEFT}}, // UP: move straight, can turn RIGHT or LEFT
	{Point{0, 1}, [2]int{DOWN, UP}},     // RIGHT
	{Point{1, 0}, [2]int{LEFT, RIGHT}},  // DOWN
	{Point{0, -1}, [2]int{UP, DOWN}},    // LEFT
}

type State struct {
	i, j      int            // Current position
	dir       int            // Direction: 0 (UP), 1 (RIGHT), 2 (DOWN), 3 (LEFT)
	score     int            // Accumulated score
	path      map[Point]bool // Set of visited points
}

func initializeValues(grid [][]rune) [][]map[int]struct {
	score int
	path  map[Point]bool
} {
	values := make([][]map[int]struct {
		score int
		path  map[Point]bool
	}, len(grid))

	for i := range grid {
		values[i] = make([]map[int]struct {
			score int
			path  map[Point]bool
		}, len(grid[i]))
		for j := range grid[i] {
			values[i][j] = make(map[int]struct {
				score int
				path  map[Point]bool
			})
			for d := 0; d < 4; d++ {
				values[i][j][d] = struct {
					score int
					path  map[Point]bool
				}{math.MaxInt, make(map[Point]bool)}
			}
		}
	}
	return values
}

func isValidPosition(grid [][]rune, point Point) bool {
	return point.x >= 0 && point.x < len(grid[0]) && point.y >= 0 && point.y < len(grid) && grid[point.y][point.x] != '#'
}

func copyPath(original map[Point]bool) map[Point]bool {
	newPath := make(map[Point]bool)
	for k := range original {
		newPath[k] = true
	}
	return newPath
}

func dfs(grid [][]rune, start, end Point) map[int]struct {
	score int
	path  map[Point]bool
} {
	values := initializeValues(grid)
	q := list.New()
	initialPath := map[Point]bool{start: true}
	q.PushBack(State{start.x, start.y, RIGHT, 0, initialPath})

	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		state := e.Value.(State)
	
i, j, d, score, path := state.i, state.j, state.dir, state.score, state.path

		// Bounds check and wall check
		if !isValidPosition(grid, Point{i, j}) {
			continue
		}

		// Update path and check current state
		path[Point{i, j}] = true
		current := values[i][j][d]
		if current.score < score {
			continue
		} else if current.score == score {
			for p := range path {
				current.path[p] = true
			}
		} else {
			newPath := copyPath(path)
			values[i][j][d] = struct {
				score int
				path  map[Point]bool
			}{score, newPath}
		}

		// Stop if end reached
		if i == end.x && j == end.y {
			continue
		}

		// Move straight
		move := directions[d]
		q.PushBack(State{i + move.delta.x, j + move.delta.y, d, score + 1, copyPath(path)})

		// Turn left and right
		for _, newDir := range move.turnDirs {
			q.PushBack(State{i, j, newDir, score + 1000, copyPath(path)})
		}
	}
	return values[end.x][end.y]
}

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}
	grid := make([][]rune, len(lines))
	var start, end Point

	for y, line := range lines {
		grid[y] = []rune(line)
		if strings.Contains(line, "S") {
			start = Point{x: strings.Index(line, "S"), y: y}
		}
		if strings.Contains(line, "E") {
			end = Point{x: strings.Index(line, "E"), y: y}
		}
	}
	score := dfs(grid, start, end)

	resp := 0
	minScore := math.MaxInt
	for _, v := range score {
		if v.score < minScore {
			resp = len(v.path)
			minScore = v.score
		}
	}

	fmt.Println("Tiles part of the best path:", resp)
}

