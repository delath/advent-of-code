package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sum := 0
	var contents string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // until end of input, you can force it with ctrl+Z or ctrl+D.
		contents = scanner.Text()
		
		if !(strings.Compare(contents,"\n") == -1) { // new line

			firstHalf := contents[0:len([]rune(contents))/2]
			secondHalf := contents[len([]rune(contents))/2:len([]rune(contents))]

			found := false

			for i:=0;i<len([]rune(firstHalf)) && !found;i++ { // for every rune present in the first half

				for j:=0;j<len([]rune(secondHalf)) && !found;j++ { // for every rune present in the second half

					if firstHalf[i] == secondHalf[j]{ // when the recurring rune is found

						sum += computePriority(firstHalf[i])

						found = true
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