package main

import (
	"fmt"
	"github.com/tire-fire/adventofcode/2024/lib"
	"strconv"
	"strings"
)

// Coordinate represents an (x, y) position on the grid
type Coordinate struct {
	x, y int
}

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	gridSize := 70
	corrupted := make(map[Coordinate]bool)

	// Parse input
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		corrupted[Coordinate{x, y}] = true
	}

	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if corrupted[Coordinate{x, y}] {
				fmt.Printf(string('#'))
			} else {
				fmt.Printf((string('.'))
			}
		}
	}


	// BFS to find the shortest path from (0,0) to (6,6)
	start := Coordinate{0, 0}
	goal := Coordinate{6, 6}

	directions := []Coordinate{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	type QueueNode struct {
		coord Coordinate
		steps int
	}

	visited := make(map[Coordinate]bool)
	queue := []QueueNode{{start, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// If we've reached the goal, return the steps
		if current.coord == goal {
			fmt.Println(current.steps)
			return
		}

		// Skip if already visited
		if visited[current.coord] {
			continue
		}
		visited[current.coord] = true

		// Explore neighbors
		for _, dir := range directions {
			next := Coordinate{
				current.coord.x + dir.x,
				current.coord.y + dir.y,
			}

			// Check if within bounds, not corrupted, and not visited
			if next.x >= 0 && next.x < gridSize && next.y >= 0 && next.y < gridSize &&
				!corrupted[next] && !visited[next] {
				queue = append(queue, QueueNode{next, current.steps + 1})
			}
		}
	}

	// If we exit the loop, there's no path
	fmt.Println("No path found")
}

