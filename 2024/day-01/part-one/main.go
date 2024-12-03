package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var firstRow []int
	var secondRow []int
	var num int

	// Parse the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		num, _ = strconv.Atoi(fields[0])
		firstRow = append(firstRow, num)

		num, _ = strconv.Atoi(fields[1])
		secondRow = append(secondRow, num)
	}

	sort.Ints(firstRow)
	sort.Ints(secondRow)

	sum := 0

	for i := 0; i < len(firstRow); i++ {
		sum += absDiff(firstRow[i], secondRow[i])
	}
	
	print(sum)
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
