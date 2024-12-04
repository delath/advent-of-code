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

	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[0])-1; j++ {
			if input[i][j] == 'A' {
				masCount := 0
				if input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S' {
					masCount++
				}
				if input[i+1][j+1] == 'M' && input[i-1][j-1] == 'S' {
					masCount++
				}
				if input[i+1][j-1] == 'M' && input[i-1][j+1] == 'S' {
					masCount++
				}
				if input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S' {
					masCount++
				}
				if masCount == 2 {
					xmasCount++
				}
			}
		}
	}

	print(xmasCount)
}
