package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse the file line by line
	for scanner.Scan() {
		line := scanner.Text()

	}

}
