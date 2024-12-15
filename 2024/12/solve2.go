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

	// globals
	grid = make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	rows, cols = len(grid), len(grid[0])
	visited = make(map[Point]bool)

	totalPrice := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			point := Point{r, c}
			if !visited[point] {
				area, sides := calculateAreaAndSides(point)
				totalPrice += area * sides
			}
		}
	}

	fmt.Println(totalPrice)
}

type Point struct {
	row, col int
}

var Directions = []Point{
	{-1, 0}, // UP
	{0, 1},  // RIGHT
	{1, 0},  // DOWN
	{0, -1}, // LEFT
}

var (
	grid      [][]rune
	rows, cols int
	visited   map[Point]bool
)

func (p Point) Add(other Point) Point {
	return Point{
		row: p.row + other.row,
		col: p.col + other.col,
	}
}

// inbounds checks if a point is within the grid boundaries
func inbounds(p Point) bool {
	return p.row >= 0 && p.row < rows && p.col >= 0 && p.col < cols
}

func calculateSides(p Point) int {
	sides := 0
	plant := grid[p.row][p.col]

	for i, direction := range Directions {
		// Check in all directions
		adjacent := p.Add(direction)
		// If it's at the edge
		if !inbounds(adjacent) || grid[adjacent.row][adjacent.col] != plant {
			// Check counter clockwise 
			prev := p.Add(Directions[(i+3)%4])
			// If this point isn't the plant and isn't inbounds, it's the beginning of and edge
			isBeginEdge := !inbounds(prev) || grid[prev.row][prev.col] != plant

			// Also handle the case where it's an inverted corner
			// B
			// BB
			corner := adjacent.Add(Directions[(i+3)%4])
			isConcaveBeginEdge := inbounds(corner) && grid[corner.row][corner.col] == plant

			if isBeginEdge || isConcaveBeginEdge {
				sides++
			}
		}
	}
	return sides
}

func calculateAreaAndSides(start Point) (int, int) {
	area := 1
	sides := calculateSides(start)
	plant := grid[start.row][start.col]
	queue := []Point{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, direction := range Directions {
			neighbor := current.Add(direction)
			if inbounds(neighbor) && !visited[neighbor] && grid[neighbor.row][neighbor.col] == plant {
				area++
				sides += calculateSides(neighbor)
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return area, sides
}
