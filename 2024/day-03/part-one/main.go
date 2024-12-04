package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)

	sum := 0

	// Parse the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			firstNumber, _ := strconv.Atoi(match[1])
			secondNumber, _ := strconv.Atoi(match[2])
			sum += firstNumber * secondNumber
		}
	}
	fmt.Println(sum)
}
