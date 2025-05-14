package main

import (
	"container/heap"
	"fmt"
	"strings"
	"github.com/tire-fire/adventofcode/2024/lib"
)

type Point struct {
	x, y int
}

type State struct {
	position  Point
	direction int // 0: East, 1: South, 2: West, 3: North
	score     int
	priority  int // score + heuristic for priority queue sorting
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func isValidPosition(grid [][]rune, point Point) bool {
	return point.x >= 0 && point.x < len(grid[0]) && point.y >= 0 && point.y < len(grid) && grid[point.y][point.x] != '#'
}

func heuristic(a, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

	directions := []Point{
		{x: 1, y: 0},  // East
		{x: 0, y: 1},  // South
		{x: -1, y: 0}, // West
		{x: 0, y: -1}, // North
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{position: start, direction: 0, score: 0, priority: heuristic(start, end)})

	visited := make(map[string]bool)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(State)

		// If we've reached the end, return the score
		if current.position == end {
			fmt.Println(current.score)
			return
		}

		// Generate a unique key for the current state
		key := fmt.Sprintf("%d,%d,%d", current.position.x, current.position.y, current.direction)
		if visited[key] {
			continue
		}
		visited[key] = true

		// Try moving forward
		newPos := Point{
			x: current.position.x + directions[current.direction].x,
			y: current.position.y + directions[current.direction].y,
		}
		if isValidPosition(grid, newPos) {
			heap.Push(pq, State{
				position: newPos,
				direction: current.direction,
				score: current.score + 1,
				priority: current.score + 1 + heuristic(newPos, end),
			})
		}

		// Try rotating clockwise and counterclockwise
		for _, turn := range []int{-1, 1} {
			newDir := (current.direction + turn + 4) % 4
			heap.Push(pq, State{
				position: current.position,
				direction: newDir,
				score: current.score + 1000,
				priority: current.score + 1000 + heuristic(current.position, end),
			})
		}
	}

	fmt.Println("No path found")
}

