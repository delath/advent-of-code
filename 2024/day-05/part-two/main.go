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
	var incorrectUpdates [][]int

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
					isPresent, _ := contains(update[:i], rules[update[i]][j])
					if isPresent {
						isValid = false
					}
				}
			}

			if !isValid {
				incorrectUpdates = append(incorrectUpdates, update)
			}

		} else {
			parts := strings.Split(line, "|")

			firstNumber, _ := strconv.Atoi(parts[0])
			secondNumber, _ := strconv.Atoi(parts[1])

			rules[firstNumber] = append(rules[firstNumber], secondNumber)
		}

	}

	for i := 0; i < len(incorrectUpdates); i++ {
		isFixed := false
		fixedUpdate := incorrectUpdates[i]
		for !isFixed {
			fixedUpdate, isFixed = fix(rules, fixedUpdate)
		}
		sum += fixedUpdate[len(fixedUpdate)/2]
	}

	print(sum)
}

func contains(arr []int, num int) (bool, int) {
	for i, n := range arr {
		if n == num {
			return true, i
		}
	}
	return false, -1
}

func fix(rules map[int][]int, update []int) ([]int, bool) {
	isValid := true
	for i := 0; i < len(update); i++ {
		for j := 0; j < len(rules[update[i]]); j++ {
			isPresent, index := contains(update[:i], rules[update[i]][j])
			if isPresent {
				temp := update[i]
				update[i] = update[index]
				update[index] = temp
				isValid = false
				break
			}
		}
		if !isValid {
			break
		}
	}
	if !isValid {
		return update, false
	} else {
		return update, true
	}
}
