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

	fmt.Println(lines)
}
