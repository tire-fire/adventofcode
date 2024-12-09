package main

import (
	"fmt"
	"github.com/tire-fire/adventofcode/2024/lib"
)

func main() {
	lines, err := lib.ReadInput()
	//lines, err := lib.ReadLines("exampleinput")
	if err != nil {
		panic("Failed to read input")
	}

	diskMap := lines[0]
	blocks := parseDiskMap(diskMap)
	compacted := compactDiskMap(blocks)
	checksum := calculateChecksum(compacted)
	fmt.Println(checksum)
}

func parseDiskMap(diskMap string) []rune {
	var blocks []rune
	isFile := true
	fileID := 0

	for i := 0; i < len(diskMap); i++ {
		length := int(diskMap[i] - '0')
		if isFile {
			// Add fileIDs to blocks
			for j := 0; j < length; j++ {
				blocks = append(blocks, rune(fileID) + '0')
			}
			fileID++
		} else {
			// Add free space blocks
			for j := 0; j < length; j++ {
				blocks = append(blocks, '.')
			}
		}
		isFile = !isFile
	}

	return blocks
}

func compactDiskMap(blocks []rune) []rune {
	for i := len(blocks) - 1; i >= 0; i-- {
		// Only move if it's a file block
		if blocks[i] != '.' {
			// Find the leftmost free space
			for j := 0; j < i; j++ {
				if blocks[j] == '.' {
					// Move the block
					blocks[j] = blocks[i]
					blocks[i] = '.'
					break
				}
			}
		}
	}
	return blocks
}

func calculateChecksum(blocks []rune) int {
	checksum := 0

	for i, block := range blocks {
		if block != '.' {
			fileID := int(block - '0')
			checksum += i * fileID
		}
	}

	return checksum
}
