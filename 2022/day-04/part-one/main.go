package main

import (
	"fmt"
)

func main() {
	var beginFirst int
	var endFirst int
	var beginSecond int
	var endSecond int
	var temp byte
	sum := 0

	for { // until end of input, you can force it with ctrl+Z or ctrl+D.
		fmt.Scanf("%d-%d,%d-%d", &beginFirst, &endFirst, &beginSecond, &endSecond)

		if (beginFirst <= beginSecond && endFirst >= endSecond) || (beginSecond <= beginFirst && endSecond >= endFirst) {
			sum++
		}

		fmt.Scanf("%c",&temp)
		if temp == 13 {
			break
		}
	}

	fmt.Printf("Number of assignment with the condition is %d.\n", sum)
}