package main

import (
	"fmt"
	"strconv"
)

// it is currently needed that a tilde (~) is passed at the end of the input
// it needs to pass the input two times too!
// this will surely be fixed in the future
func main() {
	var temp byte
	endReached:=false
	for i:=0; i<14; i++ { // to ignore first two rows
		fmt.Scanf("%c", &temp)
	}

	trueTotal := 0

	parentDir:=resolveDir(&trueTotal, &endReached)
	if parentDir <= 100000 {
		trueTotal+=parentDir
	}
	dirToBeDeletedAtLeast := 30000000 - (70000000 - parentDir)
	trueTotal=0
	endReached=false
	sizeToSetFree:=700000000
	resolveDirTwo(&trueTotal, &endReached, dirToBeDeletedAtLeast, &sizeToSetFree)

	fmt.Printf("Size that needs to be set free is %d.\n", sizeToSetFree)
}

func resolveDir(trueTotal *int, endReached *bool) int{
	var input string
	var temp byte
	totalSize := 0
	for {
		fmt.Scanf("%s", &input)

		if input=="dir" {
			readRow()
		}else if input=="$"{
			for i:=0; i<2; i++ {
				fmt.Scanf("%c", &temp)
			}
			if temp==100 { // if cd
				fmt.Scanf("%c", &temp)
				fmt.Scanf("%c", &temp)
				if temp==46{//if cd ..
					readRow()
					if totalSize <= 100000 {
						*trueTotal+=totalSize 
					}
					return totalSize
	
				}else{//if cd <something>
					readRow()
					totalSize+=resolveDir(trueTotal, endReached)	
				}
			}else if temp==115{//ls
				fmt.Scanf("%c", &temp)
				fmt.Scanf("%c", &temp)
			}
		}else{
			size, _ := strconv.Atoi(input)
			totalSize+=size
			if !readRow() {
				if totalSize <= 100000 {
					*trueTotal+=totalSize 
				}
				*endReached = true
				return totalSize
			}
		}
		if *endReached {
			if totalSize <= 100000 {
				*trueTotal+=totalSize 
			}
			return totalSize
		}
	}
}

func resolveDirTwo(trueTotal *int, endReached *bool, dirToBeDeletedAtLeast int, sizeToSetFree *int) int{
	var input string
	var temp byte
	totalSize := 0
	for {
		fmt.Scanf("%s", &input)

		if input=="dir" {
			readRow()
		}else if input=="$"{
			for i:=0; i<2; i++ {
				fmt.Scanf("%c", &temp)
			}
			if temp==100 { // if cd
				fmt.Scanf("%c", &temp)
				fmt.Scanf("%c", &temp)
				if temp==46{//if cd ..
					readRow()
					if totalSize <= 100000 {
						*trueTotal+=totalSize 
					}
					if (totalSize >= dirToBeDeletedAtLeast) && (*sizeToSetFree > totalSize){
						*sizeToSetFree = totalSize
					}
					return totalSize
	
				}else{//if cd <something>
					readRow()
					totalSize+=resolveDirTwo(trueTotal, endReached, dirToBeDeletedAtLeast, sizeToSetFree)	
				}
			}else if temp==115{//ls
				fmt.Scanf("%c", &temp)
				fmt.Scanf("%c", &temp)
			}
		}else{
			size, _ := strconv.Atoi(input)
			totalSize+=size
			if !readRow() {
				if totalSize <= 100000 {
					*trueTotal+=totalSize 
				}
				*endReached = true
				if (totalSize >= dirToBeDeletedAtLeast) && (*sizeToSetFree > totalSize){
					*sizeToSetFree = totalSize
				}
				return totalSize
			}
		}
		if *endReached {
			if totalSize <= 100000 {
				*trueTotal+=totalSize 
			}
			if (totalSize >= dirToBeDeletedAtLeast) && (*sizeToSetFree > totalSize){
				*sizeToSetFree = totalSize
			}
			return totalSize
		}
	}
}

func readRow() bool{ // returns true if new line reached, false if end of input reached
	var input byte
	for {
		fmt.Scanf("%c", &input)
		if input == 13 {
			fmt.Scanf("%c", &input)
			return true
		}else if input == 126{
			fmt.Scanf("%c", &input)
			return false
		}
	}
}
