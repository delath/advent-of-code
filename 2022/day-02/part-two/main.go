package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var roundEnd rune
	var opponent rune
	var endChecker rune
	totalScore := 0

	reader := bufio.NewReader(os.Stdin)

	for endChecker!=13 { // until end of input, you can force it with ctrl+Z or ctrl+D at the end of your input.

		opponent,_,_ = reader.ReadRune()
		reader.ReadRune()
		roundEnd,_,_ = reader.ReadRune()
		reader.ReadRune()
		endChecker,_,_ = reader.ReadRune()

		player := getPlayer(conv(opponent))

		if conv(roundEnd) == 3 { // draw

			totalScore += conv(roundEnd) + conv(opponent)
			
		}else if conv(roundEnd) == 6 { // won

			totalScore += conv(roundEnd) + player

		}else if conv(roundEnd) == 0 { // lost

			totalScore += flip(player)

		}
	}

	fmt.Printf("Total score: %d\n", totalScore)
}

func conv(char rune) int {
	switch char {

		//Rock
		case 'A':
			return 1

		//Paper
		case 'B':
			return 2

		//Scissors
		case 'C':
			return 3


		//Lose
		case 'X':
			return 0

		//Draw
		case 'Y':
			return 3

		//Win
		case 'Z':
			return 6

	}
	return -1
}

func getPlayer(opponent int) int {
	if opponent == 3 {
		return 1
	}
	return opponent+1
}

func flip(player int) int {
	if player+1 == 4 {
		return 1
	}
	return player+1
}