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
	secondRowMap := make(map[int]int)
	var num int

	// Parse the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		num, _ = strconv.Atoi(fields[0])
		firstRow = append(firstRow, num)

		num, _ = strconv.Atoi(fields[1])
		secondRowMap[num]++
	}

	sort.Ints(firstRow)

	sum := 0

	for i := 0; i < len(firstRow); i++ {
		sum += firstRow[i] * secondRowMap[firstRow[i]]
	}

	print(sum)
}
