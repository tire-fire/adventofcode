package main

import (
    "fmt"
    "strings"

    "github.com/tire-fire/adventofcode/2024/lib"
)

func main() {
    lines, err := lib.ReadInput()
    if err != nil {
        panic("failed to read input")
    }

    input := strings.Join(lines, "\n")
    parts := strings.Split(input, "\n\n")
    if len(parts) < 2 {
        panic("input format incorrect: expected map and moves separated by blank line")
    }

    mapData := parts[0]
    movesData := parts[1]

    mapLines := strings.Split(mapData, "\n")

    walls := make(map[[2]int]bool)
    leftBoxes := make(map[[2]int]bool)
    rightBoxes := make(map[[2]int]bool)
    var robot [2]int

    // Parse the map
    for i, L := range mapLines {
        for j, c := range L {
            switch c {
            case '#':
                walls[[2]int{i, 2*j}] = true
                walls[[2]int{i, 2*j+1}] = true
            case 'O':
                leftBoxes[[2]int{i, 2*j}] = true
                rightBoxes[[2]int{i, 2*j+1}] = true
            case '@':
                robot = [2]int{i, 2*j}
            case '.':
                // do nothing
            default:
                // Unexpected char
            }
        }
    }

    // Split moves by line, then each line by char
    moveLines := strings.Split(movesData, "\n")

    inWalls := func(r, c int) bool {
        return walls[[2]int{r, c}]
    }
    inLeft := func(r, c int) bool {
        return leftBoxes[[2]int{r, c}]
    }
    inRight := func(r, c int) bool {
        return rightBoxes[[2]int{r, c}]
    }

    for _, moveLine := range moveLines {
        for _, move := range moveLine {
            if move == '^' {
                // Up
                chain := make(map[[2]int]bool)
                boxesToProcess := make(map[[2]int]bool)
                i, j := robot[0], robot[1]
                i -= 1
                if inWalls(i, j) {
                    // no move
                    continue
                } else if inLeft(i, j) {
                    boxesToProcess[[2]int{i, j}] = true
                    boxesToProcess[[2]int{i, j+1}] = true
                } else if inRight(i, j) {
                    boxesToProcess[[2]int{i, j}] = true
                    boxesToProcess[[2]int{i, j-1}] = true
                }

                wallFlag := false
                chain[[2]int{robot[0], robot[1]}] = true
                for len(boxesToProcess) > 0 {
                    var b [2]int
                    for k := range boxesToProcess {
                        b = k
                        break
                    }
                    delete(boxesToProcess, b)
                    chain[b] = true
                    bi, bj := b[0], b[1]
                    bi -= 1
                    if inWalls(bi, bj) {
                        wallFlag = true
                        continue
                    } else if inLeft(bi, bj) {
                        boxesToProcess[[2]int{bi, bj}] = true
                        boxesToProcess[[2]int{bi, bj+1}] = true
                    } else if inRight(bi, bj) {
                        boxesToProcess[[2]int{bi, bj}] = true
                        boxesToProcess[[2]int{bi, bj-1}] = true
                    }
                }
                if wallFlag {
                    continue
                }

                // Move chain up one space
                var oldLeftBoxes, newLeftBoxes, oldRightBoxes, newRightBoxes [] [2]int
                var newRobot [2]int
                for b := range chain {
                    if b == robot {
                        newRobot = [2]int{robot[0]-1, robot[1]}
                    } else if inLeft(b[0], b[1]) {
                        oldLeftBoxes = append(oldLeftBoxes, b)
                        newLeftBoxes = append(newLeftBoxes, [2]int{b[0]-1, b[1]})
                    } else if inRight(b[0], b[1]) {
                        oldRightBoxes = append(oldRightBoxes, b)
                        newRightBoxes = append(newRightBoxes, [2]int{b[0]-1, b[1]})
                    }
                }

                robot = newRobot
                // update leftBoxes and rightBoxes
                for _, b := range oldLeftBoxes {
                    delete(leftBoxes, b)
                }
                for _, b := range newLeftBoxes {
                    leftBoxes[b] = true
                }
                for _, b := range oldRightBoxes {
                    delete(rightBoxes, b)
                }
                for _, b := range newRightBoxes {
                    rightBoxes[b] = true
                }

            } else if move == '>' {
                // Right
                chain := [][2]int{robot}
                i, j := robot[0], robot[1]
                j += 1
                for inLeft(i, j) {
                    chain = append([][2]int{{i, j+1}, {i, j}}, chain...)
                    j += 2
                }

                if inWalls(i, j) {
                    continue
                }

                length := len(chain)
                // stepping by 2 means take pairs
                for x := 0; x < length-1; x += 2 {
                    rightBracket := chain[x]   // since they 'b' in python was the loop variable
                    leftBracket := [2]int{rightBracket[0], rightBracket[1]-1}

                    // Move box one to the right
                    // remove right_bracket from rightBoxes
                    delete(rightBoxes, rightBracket)
                    // add new right bracket
                    rightBoxes[[2]int{rightBracket[0], rightBracket[1]+1}] = true
                    // remove leftBracket from leftBoxes
                    delete(leftBoxes, leftBracket)
                    // add new left bracket
                    leftBoxes[[2]int{leftBracket[0], leftBracket[1]+1}] = true
                }

                robot = [2]int{robot[0], robot[1] + 1}

            } else if move == 'v' {
                // Down
                chain := make(map[[2]int]bool)
                boxesToProcess := make(map[[2]int]bool)
                i, j := robot[0], robot[1]
                i += 1
                if inWalls(i, j) {
                    continue
                } else if inLeft(i, j) {
                    boxesToProcess[[2]int{i, j}] = true
                    boxesToProcess[[2]int{i, j+1}] = true
                } else if inRight(i, j) {
                    boxesToProcess[[2]int{i, j}] = true
                    boxesToProcess[[2]int{i, j-1}] = true
                }

                wallFlag := false
                chain[[2]int{robot[0], robot[1]}] = true
                for len(boxesToProcess) > 0 {
                    var b [2]int
                    for k := range boxesToProcess {
                        b = k
                        break
                    }
                    delete(boxesToProcess, b)
                    chain[b] = true
                    bi, bj := b[0], b[1]
                    bi += 1
                    if inWalls(bi, bj) {
                        wallFlag = true
                        continue
                    } else if inLeft(bi, bj) {
                        boxesToProcess[[2]int{bi, bj}] = true
                        boxesToProcess[[2]int{bi, bj+1}] = true
                    } else if inRight(bi, bj) {
                        boxesToProcess[[2]int{bi, bj}] = true
                        boxesToProcess[[2]int{bi, bj-1}] = true
                    }
                }
                if wallFlag {
                    continue
                }

                // Move chain down one space
                var oldLeftBoxes, newLeftBoxes, oldRightBoxes, newRightBoxes [] [2]int
                var newRobot [2]int
                for b := range chain {
                    if b == robot {
                        newRobot = [2]int{robot[0] + 1, robot[1]}
                    } else if inLeft(b[0], b[1]) {
                        oldLeftBoxes = append(oldLeftBoxes, b)
                        newLeftBoxes = append(newLeftBoxes, [2]int{b[0] + 1, b[1]})
                    } else if inRight(b[0], b[1]) {
                        oldRightBoxes = append(oldRightBoxes, b)
                        newRightBoxes = append(newRightBoxes, [2]int{b[0] + 1, b[1]})
                    }
                }

                robot = newRobot
                for _, b := range oldLeftBoxes {
                    delete(leftBoxes, b)
                }
                for _, b := range newLeftBoxes {
                    leftBoxes[b] = true
                }
                for _, b := range oldRightBoxes {
                    delete(rightBoxes, b)
                }
                for _, b := range newRightBoxes {
                    rightBoxes[b] = true
                }

            } else {
                // '<' left
                chain := [][2]int{robot}
                i, j := robot[0], robot[1]
                j -= 1
                for inRight(i, j) {
                    chain = append([][2]int{{i, j-1}, {i, j}}, chain...)
                    j -= 2
                }

                if inWalls(i, j) {
                    continue
                }

                // Move chain left one space
                // similar logic to '>' move
                length := len(chain)
                // chain[:-1:2] every second element up to the second last
                for x := 0; x < length-1; x += 2 {
                    leftBracket := chain[x]
                    // The corresponding right bracket is leftBracket[1]+1
                    rightBracket := [2]int{leftBracket[0], leftBracket[1] + 1}

                    // Move box one to the left
                    delete(leftBoxes, leftBracket)
                    leftBoxes[[2]int{leftBracket[0], leftBracket[1] - 1}] = true
                    delete(rightBoxes, rightBracket)
                    rightBoxes[[2]int{rightBracket[0], rightBracket[1] - 1}] = true
                }

                robot = [2]int{robot[0], robot[1] - 1}
            }
        }
    }

    // Compute GPS sum
    total := 0
    for b := range leftBoxes {
        total += 100*b[0] + b[1]
    }
    fmt.Println(total)
}

