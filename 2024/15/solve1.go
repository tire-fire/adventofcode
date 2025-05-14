package main

import (
    "fmt"
    "github.com/tire-fire/adventofcode/2024/lib"
    "strings"
)

type coord struct {
    r, c int
}

func main() {
    lines, err := lib.ReadInput()
    if err != nil {
        panic("Failed to read input")
    }

    // Separate warehouse map from moves
    // The puzzle states that the map and the moves will be provided.
    // Typically, you'd parse until you hit an empty line or something similar.
    // For the sake of this demonstration, let's assume:
    // - The warehouse map lines come first.
    // - Then one or more lines of moves come after.
    //
    // Adjust parsing logic depending on your actual input format.

    var warehouse []string
    var movesLines []string

    // Find the boundary between the warehouse and moves
    // Assume warehouse lines start and end with '#' rows, and after that all are moves.
    inMap := true
    for _, line := range lines {
        if inMap {
            // If we hit a line that doesn't look like warehouse, we switch
            // However, the puzzle states moves are after the warehouse.
            // Let's assume that once we hit a line with no '#' or something similar we switch.
            // Better yet, we know the warehouse is enclosed by walls (#),
            // so after the top and bottom rows (which are all '#'), the moves start.
            if len(line) > 0 && strings.Contains(line, "#") {
                warehouse = append(warehouse, line)
            } else {
                // Switch to moves
                inMap = false
                if line != "" {
                    movesLines = append(movesLines, line)
                }
            }
        } else {
            if line != "" {
                movesLines = append(movesLines, line)
            }
        }
    }

    // Join all moves into a single string (ignore newlines)
    movesStr := strings.Join(movesLines, "")
    movesStr = strings.ReplaceAll(movesStr, "\n", "")

    // Convert warehouse to a more mutable structure
    // We'll use a slice of slices of runes.
    grid := make([][]rune, len(warehouse))
    for i := range warehouse {
        grid[i] = []rune(warehouse[i])
    }

    // Find the robot's initial position
    var robot coord
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[r]); c++ {
            if grid[r][c] == '@' {
                robot = coord{r, c}
                grid[r][c] = '.' // remove the '@' and consider it empty space
                break
            }
        }
    }

    // Directions map
    dirMap := map[rune]coord{
        '^': {-1, 0},
        'v': {1, 0},
        '<': {0, -1},
        '>': {0, 1},
    }

    // Function to check if a cell is within the warehouse bounds
    inBounds := func(r, c int) bool {
        return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
    }

    // Process each move
    for _, move := range movesStr {
        d := dirMap[move]
        nextR := robot.r + d.r
        nextC := robot.c + d.c

        if !inBounds(nextR, nextC) || grid[nextR][nextC] == '#' {
            // Can't move into a wall or out of bounds - do nothing
            continue
        }

        if grid[nextR][nextC] == 'O' {
            // Need to try to push boxes
            // Find the chain of boxes in this direction
            boxPositions := []coord{}
            rr, cc := nextR, nextC
            for inBounds(rr, cc) && grid[rr][cc] == 'O' {
                boxPositions = append(boxPositions, coord{rr, cc})
                rr += d.r
                cc += d.c
            }
            // Now rr, cc is either out of bounds or not 'O'
            if !inBounds(rr, cc) || grid[rr][cc] == '#' {
                // Can't push because next space after boxes is wall/out of bounds
                continue
            }
            if grid[rr][cc] == 'O' {
                // More boxes that continue and can't push them all if blocked at the end
                continue
            }
            // If empty '.', we can push the entire chain of boxes forward
            // Move them from last to first
            grid[rr][cc] = 'O'
            for i := len(boxPositions) - 1; i > 0; i-- {
                prev := boxPositions[i-1]
                grid[boxPositions[i].r][boxPositions[i].c] = 'O'
                _ = prev // not needed explicitly
            }
            // The first box position moves into what was empty
            grid[boxPositions[0].r][boxPositions[0].c] = '.'
            // Now move the robot
            grid[robot.r][robot.c] = '.'
            robot.r, robot.c = nextR, nextC
        } else {
            // Just move the robot if empty
            grid[robot.r][robot.c] = '.'
            robot.r, robot.c = nextR, nextC
        }
    }

    // After all moves, sum up GPS coordinates of all boxes
    // GPS coordinate = 100*r + c
    sum := 0
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[r]); c++ {
            if grid[r][c] == 'O' {
                sum += 100*r + c
            }
        }
    }

    fmt.Println(sum)
}

