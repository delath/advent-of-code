package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var player rune
	var opponent rune
	var endChecker rune
	totalScore := 0

	reader := bufio.NewReader(os.Stdin)

	for endChecker!=13 { // until interrupt key

		opponent,_,_ = reader.ReadRune()
		reader.ReadRune()
		player,_,_ = reader.ReadRune()
		reader.ReadRune()
		endChecker,_,_ = reader.ReadRune()

		diff := conv(player) - conv(opponent)

		if diff == 0 { // draw

			totalScore += 3 + conv(player)
			
		}else if diff == 1 || diff == -2 { // won

			totalScore += 6 + conv(player)

		}else if diff == -1 || diff == 2 { // lost

			totalScore += conv(player)

		}
	}

	fmt.Printf("Total score: %d\n", totalScore)
}

func conv(char rune) int {
	switch char {

		//Rock
		case 'A':
			return 1
		case 'X':
			return 1

		//Paper
		case 'B':
			return 2
		case 'Y':
			return 2

		//Scissors
		case 'C':
			return 3
		case 'Z':
			return 3

	}
	return -1
}