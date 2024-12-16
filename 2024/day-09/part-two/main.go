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

	blockId--
	occurrences := 0
	for i := len(newLine) - 1; i >= 0 && blockId > 0; i-- {
		if newLine[i] == int64(blockId) {
			occurrences++
		} else if newLine[i] == int64(blockId-1) {
			freeSpaces := 0
			for j := 0; j < len(newLine) && j < i; j++ {
				if newLine[j] != int64(-1) {
					freeSpaces = 0
					continue
				}
				freeSpaces++
				if freeSpaces == occurrences {
					for k := 0; k < freeSpaces; k++ {
						newLine[j-k] = int64(blockId)
						newLine[i+1+k] = int64(-1)
					}
					break
				}
			}
			occurrences = 1
			blockId--
		} else if occurrences == 0 {
			continue
		} else {
			freeSpaces := 0
			for j := 0; j < len(newLine) && j <= i; j++ {
				if newLine[j] != int64(-1) {
					freeSpaces = 0
					continue
				}
				freeSpaces++
				if freeSpaces == occurrences {
					for k := 0; k < freeSpaces; k++ {
						newLine[j-k] = int64(blockId)
						newLine[i+1+k] = int64(-1)
					}
					break
				}
			}
			occurrences = 0
			blockId--
		}
	}

	for i := 0; i < len(newLine); i++ {
		if newLine[i] == int64(-1) {
			continue
		}

		checksum += uint64(newLine[i]) * uint64(i)
	}

	print(checksum)
}
