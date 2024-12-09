package main

import (
	"fmt"
	"github.com/tire-fire/adventofcode/2024/lib"
)

type Segment struct {
	ID     int // -1 for free space
	Length int
}

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


func parseDiskMap(diskMap string) []Segment {
	segments := []Segment{}
	isFile := true
	fileID := 0

	for i := 0; i < len(diskMap); i++ {
		length := int(diskMap[i] - '0')
		if isFile {
			segments = append(segments, Segment{ID: fileID, Length: length})
			fileID++
		} else {
			segments = append(segments, Segment{ID: -1, Length: length})
		}
		isFile = !isFile
	}

	return segments
}


func compactDiskMap(segments []Segment) []Segment {
	for i := len(segments) - 1; i >= 0; i-- {
		segment := segments[i]
		if segment.ID == -1 {
			continue // Skip free space
		}

		// Find the leftmost span of free space
		for j := 0; j < i; j++ {
			if segments[j].ID == -1 && segments[j].Length >= segment.Length {
				// Move the file
				segments[j].Length -= segment.Length
				if segments[j].Length == 0 {
					segments = append(segments[:j], segments[j+1:]...) // Remove empty segment
					i-- // this bug caught me for so long
				}
				segments[i].ID = -1
				segments = append(segments[:j], append([]Segment{segment}, segments[j:]...)...)
				break
			}
		}
	}
	return segments
}

func calculateChecksum(segments []Segment) int64 {
	checksum := int64(0)
	position := 0
	freeSpace := false

	for _, segment := range segments {
		if freeSpace && segment.ID == -1 {
			position += segment.Length
			continue
		} else {
			freeSpace = false
		}

		if segment.ID != -1 {
			for i := 0; i < segment.Length; i++ {
				checksum += int64(position * segment.ID)
				position++
			}
		} else {
			position += segment.Length
			freeSpace = true
		}
	}

	return checksum
}

