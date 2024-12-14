package main

import (
	"bufio"
	"os"
)

type Coordinate struct {
	X, Y int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var city [][]rune
	signals := make(map[rune][]Coordinate)
	antinodes := make(map[Coordinate]bool)

	y := 0
	// Parse the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		runesLine := []rune(line)

		city = append(city, runesLine)

		for i := 0; i < len(runesLine); i++ {
			if runesLine[i] != '.' {
				signals[runesLine[i]] = append(signals[runesLine[i]], Coordinate{X: i, Y: y})
			}
		}
		y++
	}

	for _, coordinates := range signals {
		for i := 0; i < len(coordinates); i++ {
			for j := 0; j < len(coordinates); j++ {
				if i == j {
					continue
				}
				firstCoordinate := coordinates[i]
				secondCoordinate := coordinates[j]
				antinodes[firstCoordinate] = true
				antinodes[secondCoordinate] = true
				for {
					if 2*secondCoordinate.X-firstCoordinate.X < 0 || 2*secondCoordinate.X-firstCoordinate.X >= len(city[0]) || 2*secondCoordinate.Y-firstCoordinate.Y < 0 || 2*secondCoordinate.Y-firstCoordinate.Y >= len(city) {
						break
					}
					antinodes[Coordinate{X: 2*secondCoordinate.X - firstCoordinate.X, Y: 2*secondCoordinate.Y - firstCoordinate.Y}] = true

					temp := secondCoordinate
					secondCoordinate.X = 2*secondCoordinate.X - firstCoordinate.X
					secondCoordinate.Y = 2*secondCoordinate.Y - firstCoordinate.Y
					firstCoordinate = temp
				}
			}
		}
	}

	print(len(antinodes))
}
