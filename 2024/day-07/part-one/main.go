package main

import (
	"bufio"
	"fmt"
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

	calibrationSum := 0

	// Parse the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")

		target, _ := strconv.Atoi(parts[0])
		numberStrings := re.FindAllString(parts[1], -1)

		var numbers []int
		for _, numStr := range numberStrings {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}

		if evaluate(numbers, target, 0, numbers[0]) {
			calibrationSum += target
		}
	}

	fmt.Println(calibrationSum)
}

func evaluate(numbers []int, target, index, current int) bool {
	if index == len(numbers)-1 {
		return current == target
	}

	// Addition
	if evaluate(numbers, target, index+1, current+numbers[index+1]) {
		return true
	}

	// Multiplication
	if evaluate(numbers, target, index+1, current*numbers[index+1]) {
		return true
	}

	return false
}
