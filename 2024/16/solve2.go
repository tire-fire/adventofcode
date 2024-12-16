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

type Path struct {
	currentNode Point
	direction   int
	score       int
	priority    int
	path        []Point
}

type PriorityQueue []Path

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Path))
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

func aStar(grid [][]rune, start, end Point, directions []Point, minScore int) map[Point]bool {
	visited := make(map[Point]int)
	tileMarks := make(map[Point]bool)
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Path{
		currentNode: start,
		direction:   0,
		score:       0,
		priority:    heuristic(start, end),
		path:        []Point{start},
	})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Path)

		// If we've reached the end and the score matches minScore, mark the path
		if current.currentNode == end {
			if current.score == minScore {
				for _, point := range current.path {
					tileMarks[point] = true
				}
				fmt.Printf("Found path with score: %d\n", current.score)
			}
			continue
		}

		// Skip nodes if a better score has already been recorded
		if bestScore, found := visited[current.currentNode]; found && current.score > bestScore {
			continue
		}
		// Update the visited map to include paths with scores equal to the best score
		visited[current.currentNode] = current.score

		// Explore neighbors
		for newDir, dir := range directions {
			newNode := Point{x: current.currentNode.x + dir.x, y: current.currentNode.y + dir.y}
			if !isValidPosition(grid, newNode) {
				continue
			}

			// Calculate the new score, adding 1000 if there's a turn
			turnCost := 0
			if newDir != current.direction {
				turnCost = 1000
			}
			newScore := current.score + 1 + turnCost

			// Only push paths that do not exceed minScore
			if newScore <= minScore {
				heap.Push(pq, Path{
					currentNode: newNode,
					direction:   newDir,
					score:       newScore,
					priority:    newScore + heuristic(newNode, end),
					path:        append([]Point{}, append(current.path, newNode)...),
				})
			}
		}
	}

	return tileMarks
}

func main() {
	lines, err := lib.ReadLines("exampleinput")
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

	// First, determine the minimum score using A*
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Path{
		currentNode: start,
		direction:   0,
		score:       0,
		priority:    heuristic(start, end),
	})

	minScore := -1
	visited := make(map[Point]int)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Path)

		if current.currentNode == end {
			if minScore == -1 || current.score < minScore {
				minScore = current.score
			}
			continue
		}

		if bestScore, found := visited[current.currentNode]; found && current.score >= bestScore {
			continue
		}
		visited[current.currentNode] = current.score

		for newDir, dir := range directions {
			newNode := Point{x: current.currentNode.x + dir.x, y: current.currentNode.y + dir.y}
			if !isValidPosition(grid, newNode) {
				continue
			}

			turnCost := 0
			if newDir != current.direction {
				turnCost = 1000
			}
			newScore := current.score + 1 + turnCost

			heap.Push(pq, Path{
				currentNode: newNode,
				direction:   newDir,
				score:       newScore,
				priority:    newScore + heuristic(newNode, end),
			})
		}
	}

	// Now, find all paths using A* with the determined minScore
	tileMarks := aStar(grid, start, end, directions, minScore)

	// Count the tiles that are part of any best path
	total := 0
	for y := range grid {
		for x := range grid[y] {
			if tileMarks[Point{x: x, y: y}] {
				total++
				grid[y][x] = 'O'
			}
		}
	}

	fmt.Println("Total tiles part of any best path:", total)

	// Optionally, print the grid with marked paths
	for _, line := range grid {
		fmt.Println(string(line))
	}
}

