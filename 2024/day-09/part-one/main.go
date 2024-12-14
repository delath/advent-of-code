package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var newLine []int64
	var checksum uint64
	var line string

	for scanner.Scan() {
		line = scanner.Text()
	}

	blockId := 0
	isBlockSize := true
	for i := 0; i < len(line); i++ {
		if isBlockSize {
			blockSize := int(line[i] - '0')
			for j := 0; j < blockSize; j++ {
				newLine = append(newLine, int64(blockId))
			}
			isBlockSize = false
			blockId++
		} else {
			freeSpace := int(line[i] - '0')
			for j := 0; j < freeSpace; j++ {
				newLine = append(newLine, int64(-1))

			}
			isBlockSize = true
		}
	}

	isParsed := false
	for i := 0; i < len(newLine) && !isParsed; i++ {
		if newLine[i] != int64(-1) {
			continue
		}
		for j := len(newLine) - 1; j >= 0 && !isParsed; j-- {
			if i >= j {
				isParsed = true
			}
			if newLine[j] == int64(-1) || i >= j {
				continue
			}
			newLine[i] = newLine[j]
			newLine[j] = int64(-1)
			break
		}
	}

	for i := 0; i < len(newLine); i++ {
		if newLine[i] == int64(-1) {
			break
		}

		checksum += uint64(newLine[i]) * uint64(i)
	}

	print(checksum)
}
