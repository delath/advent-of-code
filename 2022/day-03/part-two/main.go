package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // until end of input, you can force it with ctrl+Z or ctrl+D.
		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()
		scanner.Scan()
		third := scanner.Text()

		found := false

		for i:=0;i<len([]rune(first)) && !found;i++ { // for every rune present in the first rucksack

			for j:=0;j<len([]rune(second)) && !found;j++ { // for every rune present in the second rucksack

				if first[i] == second[j]{ // both char are the same

					for k:=0;k<len([]rune(third)) && !found;k++{ // for every rune present in the third rucksack

						if first[i] == third[k] { // even the third char is the same as the other two

							sum += computePriority(first[i])

							found = true

						}

					}

				}

			}

		}

	}

	fmt.Printf("Sum of priorities is %d.\n", sum)
}

func computePriority(char byte) int {
	if char >= 97 && char <=122 {
		return int(char) - 96
	}else if char >= 65 && char <=90{
		return int(char) - 38
	}
	return -1
}