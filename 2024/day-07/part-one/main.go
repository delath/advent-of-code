package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`\d+`)

	// Parse the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")

		result, _ := strconv.Atoi(parts[0])
		numberStrings := re.FindAllString(parts[1], -1)

		var numbers []int
		for _, numStr := range numberStrings {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}

		operations := make([]rune, len(numbers)-1)
		for i := range operations {
			operations[i] = '+'
		}

		isValid := false

		for {
			temporaryResult := 0

			for i := 0; i < len(operations); i++ {
				if operations[i] == '+' {
					temporaryResult += numbers[i] + numbers[i+1]
				} else if operations[i] == '*' {
					temporaryResult += numbers[i] * numbers[i+1]
				}
			}

			if temporaryResult == result {
				isValid = true
				break
			}

			// add a single multiplier
			i := 0
			for {
				if operations[i+1] == '*' || i == len(operations)-1 {
					operations[i] = '*'
					break
				}
				i++
			}

		}
	}

}
