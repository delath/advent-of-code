package main

import (
	"bufio"
	"os"
)

type Coordinate struct {
	X, Y int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var newLine []uint64
	var checksum uint64
	var line string

	for scanner.Scan() {
		line = scanner.Text()
	}

	blockId := 0
	for i := 0; i < len(line); i = i + 2 {
		if i < len(line) {
			blockSize := int(line[i] - '0')
			for j := 0; j < blockSize; j++ {
				newLine = append(newLine, uint64(blockId))
			}
		}

		if i+1 < len(line) {
			freeSpace := int(line[i+1] - '0')
			for j := 0; j < freeSpace; j++ {
				newLine = append(newLine, uint64('.'))

			}
		}

		blockId++
	}

	isParsed := false
	for i := 0; i < len(newLine) && !isParsed; i++ {
		if newLine[i] != uint64('.') {
			continue
		}
		for j := len(newLine) - 1; j >= 0 && !isParsed; j-- {
			if i >= j {
				isParsed = true
			}
			if newLine[j] == uint64('.') || i >= j {
				continue
			}
			newLine[i] = newLine[j]
			newLine[j] = uint64('.')
			break
		}
	}

	for i := 0; i < len(newLine); i++ {
		if newLine[i] == uint64('.') {
			break
		}

		checksum += (newLine[i] * uint64(i))
	}

	print(checksum)
}