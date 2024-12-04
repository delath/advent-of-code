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

		isSafe := checkConditions(fields)

		if isSafe {
			safeReportsCount++
		} else {
			for i := 0; i < len(fields); i++ {
				fieldsWithoutIndex := removeAtIndex(fields, i)
				isSafe = checkConditions(fieldsWithoutIndex)
				if isSafe {
					safeReportsCount++
					break
				}

			}
		}
	}
	print(safeReportsCount)
}

func checkConditions(fields []string) bool {
	isDecreasing := true
	isIncreasing := true
	isValid := true

	for i := 1; i < len(fields); i++ {
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

		if !isValid || (!isIncreasing && !isDecreasing) {
			return false
		}
	}

	return true
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func removeAtIndex(fields []string, index int) []string {
	if index < 0 || index >= len(fields) {
		return fields
	}
	newFields := make([]string, len(fields)-1)
	copy(newFields, fields[:index])
	copy(newFields[index:], fields[index+1:])
	return newFields
}
