package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	max := 0
	sum := 0
	var temp string
	var convertedTemp int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // until end of input, you can force it with ctrl+Z or ctrl+D.
		temp = scanner.Text()
		
		if strings.Compare(temp,"\n") == -1 { // end of elf detected

			if max < sum {
				max = sum
			}
			sum = 0

		}else{ // still the same elf

			convertedTemp, _ = strconv.Atoi(temp)
			sum += convertedTemp
			temp = scanner.Text()

		}

	}

	fmt.Printf("Elf is carrying %d calories", max)

}