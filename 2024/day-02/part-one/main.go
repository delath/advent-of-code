package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeReportsCount := 0

	// Parse the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		isDecreasing := true
		isIncreasing := true
		isValid := true

		for i := 1; i < len(fields) && isValid && (isIncreasing || isDecreasing); i++ {
			currentNumber, _ := strconv.Atoi(fields[i])
			previousNumber, _ := strconv.Atoi(fields[i-1])

			if currentNumber > previousNumber {
				isDecreasing = false
			}
			if currentNumber < previousNumber {
				isIncreasing = false
			}

			difference := absDiff(currentNumber, previousNumber)
			if difference < 1 || difference > 3 {
				isValid = false
			}
		}

		if isValid && (isIncreasing || isDecreasing) {
			safeReportsCount++
		}
	}

	print(safeReportsCount)
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
