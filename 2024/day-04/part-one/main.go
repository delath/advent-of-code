package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string

	xmasCount := 0

	// Parse the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		input = append(input, line)
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			// horizontal front
			if len(input[0]) > j+3 && input[i][j] == 'X' && input[i][j+1] == 'M' && input[i][j+2] == 'A' && input[i][j+3] == 'S' {
				xmasCount++
			}

			// horizontal back
			if 0 <= len(input[0])-4-j && input[len(input)-1-i][len(input[0])-1-j] == 'X' && input[len(input)-1-i][len(input[0])-2-j] == 'M' && input[len(input)-1-i][len(input[0])-3-j] == 'A' && input[len(input)-1-i][len(input[0])-4-j] == 'S' {
				xmasCount++
			}

			// vertical down
			if len(input[0]) > j+3 && input[j][i] == 'X' && input[j+1][i] == 'M' && input[j+2][i] == 'A' && input[j+3][i] == 'S' {
				xmasCount++
			}

			// vertical up
			if 0 <= len(input[0])-4-j && input[len(input[0])-1-j][len(input)-1-i] == 'X' && input[len(input[0])-2-j][len(input)-1-i] == 'M' && input[len(input[0])-3-j][len(input)-1-i] == 'A' && input[len(input[0])-4-j][len(input)-1-i] == 'S' {
				xmasCount++
			}

			// diagonals
			if len(input[0]) > i+3 && len(input[0]) > j+3 && input[i][j] == 'X' && input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S' {
				xmasCount++
			}
			if 0 <= len(input)-4-i && 0 <= len(input)-4-j && input[len(input)-1-j][len(input)-1-i] == 'X' && input[len(input)-2-j][len(input)-2-i] == 'M' && input[len(input)-3-j][len(input)-3-i] == 'A' && input[len(input)-4-j][len(input)-4-i] == 'S' {
				xmasCount++
			}
			if len(input[0]) > j+3 && 0 <= i-3 && input[j][i] == 'X' && input[j+1][i-1] == 'M' && input[j+2][i-2] == 'A' && input[j+3][i-3] == 'S' {
				xmasCount++
			}
			if len(input[0]) > len(input)+2-i && 0 <= len(input)-4-j && input[len(input)-1-j][len(input)-1-i] == 'X' && input[len(input)-2-j][len(input)-i] == 'M' && input[len(input)-3-j][len(input)+1-i] == 'A' && input[len(input)-4-j][len(input)+2-i] == 'S' {
				xmasCount++
			}
		}
	}

	print(xmasCount)
}
