package main

import (
	"fmt"
)

func main() {
	hasDuplicate := false
	counter := 14
	var lastFour [14]byte
	for i:=0; i<len(lastFour); i++{
		fmt.Scanf("%c", &lastFour[i])
	}

	for {
		for i:=0; i<len(lastFour) && !hasDuplicate; i++ {
			for j:=(i+1); j<len(lastFour) && !hasDuplicate; j++ {
				if lastFour[i] == lastFour[j] {
					hasDuplicate = true
				}
			}
		}
		
		if !hasDuplicate {
			fmt.Printf("first start-of-packet marker is detected at %d.\n", counter)
			break
		}
		
		for i:=1; i<len(lastFour); i++ {
			lastFour[i-1]=lastFour[i]
		}

		fmt.Scanf("%c", &lastFour[len(lastFour)-1])
		hasDuplicate = false
		counter++
	}
}