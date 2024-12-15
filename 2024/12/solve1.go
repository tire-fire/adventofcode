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

	totalPrice := calculateTotalPrice()
	fmt.Println(totalPrice)
}

type Point struct {
	row, col int
}

var (
	grid    [][]rune
	rows, cols int
	visited map[Point]bool
)

var Directions = []Point{
	{-1, 0}, // UP
	{0, 1},  // RIGHT
	{1, 0},  // DOWN
	{0, -1}, // LEFT
}

func (p Point) Add(other Point) Point {
	return Point{
		row: p.row + other.row,
		col: p.col + other.col,
	}
}

func inbounds(p Point) bool {
	return p.row >= 0 && p.row < rows && p.col >= 0 && p.col < cols
}

func calculateAreaAndPerimeter(start Point, plant rune) (int, int) {
	stack := []Point{start}
	visited[start] = true
	area := 0
	perimeter := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		area++

		for _, direction := range Directions {
			neighbor := current.Add(direction)
			if !inbounds(neighbor) || grid[neighbor.row][neighbor.col] != plant {
				perimeter++
			} else if !visited[neighbor] {
				visited[neighbor] = true
				stack = append(stack, neighbor)
			}
		}
	}

	return area, perimeter
}

func calculateTotalPrice() int {
	totalPrice := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			point := Point{r, c}
			if !visited[point] {
				plant := grid[r][c]
				area, perimeter := calculateAreaAndPerimeter(point, plant)
				totalPrice += area * perimeter
			}
		}
	}

	return totalPrice
}

