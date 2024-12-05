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

	sum := 0
	hasRuled := false
	rules := make(map[int][]int)

	// Parse the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			hasRuled = true
		} else if hasRuled {
			parts := strings.Split(line, ",")
			update := make([]int, len(parts))

			for i := 0; i < len(parts); i++ {
				number, _ := strconv.Atoi(parts[i])
				update[i] = number
			}

			isValid := true
			for i := 0; i < len(update); i++ {
				for j := 0; j < len(rules[update[i]]); j++ {
					if contains(update[:i], rules[update[i]][j]) {
						isValid = false
					}
				}
			}

			if isValid {
				sum += update[len(update)/2]
			}

		} else {
			parts := strings.Split(line, "|")

			firstNumber, _ := strconv.Atoi(parts[0])
			secondNumber, _ := strconv.Atoi(parts[1])

			rules[firstNumber] = append(rules[firstNumber], secondNumber)
		}
	}

	print(sum)
}

func contains(arr []int, num int) bool {
	for _, n := range arr {
		if n == num {
			return true
		}
	}
	return false
}
