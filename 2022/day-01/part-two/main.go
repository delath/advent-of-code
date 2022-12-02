package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	var temp string
	var convertedTemp int

	var firstMax int
	var secondMax int
	thirdMax := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // until end of input. (you can force it with ctrl+Z or ctrl+D)
		temp = scanner.Text()
		
		if strings.Compare(temp,"\n") == -1 { // end of elf detected

			if firstMax < sum {
				thirdMax = secondMax
				secondMax = firstMax
				firstMax = sum
			} else if secondMax < sum {
				thirdMax = secondMax
				secondMax = sum
			} else if thirdMax < sum {
				thirdMax = sum
			}
			
			sum = 0

		}else{ // still the same elf

			convertedTemp, _ = strconv.Atoi(temp)
			sum += convertedTemp
			temp = scanner.Text()

		}
	}

	fmt.Printf("Those elves are carrying a total of %d calories", firstMax + secondMax + thirdMax)
}