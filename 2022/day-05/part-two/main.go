package main

import (
	"fmt"
)

func main() {
	area := make([]byte, 0)
	var charHolder byte
	rowLength := 0
	columnLength := 0
	lengthFound := false
	crates := make([]byte, 0)
	var from int
	var to int
	var numToBeMoved int
	var temp byte
	pile := make([]byte, 0)
	pileSize := 0

	for {
		fmt.Scanf("%c", &charHolder) // first char of new line

		if charHolder == 13 || charHolder == 10 { // if end of construction area
			fmt.Scanf("%c", &charHolder)
			break
		}else{
			columnLength++
		}

		area = append(area, charHolder)

		for charHolder != 13 { // until \r\n
			fmt.Scanf("%c", &charHolder)
			area = append(area, charHolder)

			if !lengthFound {
				rowLength++
				if charHolder == 13 {

					lengthFound = true
				}
			}
		}
		fmt.Scanf("%c", &charHolder) // remove \n
	}

	rowLength++
	columnLength--
	for j:=columnLength-1; j>=0; j-- {
		for i:=1+(j*rowLength); i<rowLength*(1+j); i=i+4 {
			crates = append(crates, area[i])
		}	
	}

	rowLength = rowLength / 4 // rowLength of the new array

	for {
		fmt.Scanf("move %d from %d to %d", &numToBeMoved, &from, &to)

		for k:=0; k<numToBeMoved; k++ { // for X times, where X is the number got from input and put in numToBeMoved

			for i:=columnLength-1; i>=0; i-- { // until i find the highest item of the column

				if crates[(from-1)+(i*rowLength)] != 32 { // item found

					if pileSize < (k+1) {
						pile = append(pile, crates[(from-1)+(i*rowLength)])
						pileSize++
					}else{
						pile[k] = crates[(from-1)+(i*rowLength)]
					}
					
					crates[(from-1)+(i*rowLength)] = 32
					break
				}

			}
	
		}

		for k:=0; k<numToBeMoved; k++ { // for X times, where X is the number got from input and put in numToBeMoved

			for i:=columnLength-1; i>=0; i-- { // until i find the highest item of the column
	
	
				if crates[(to-1)+(i*rowLength)] != 32 || i==0 { // item found or ground floor reached
	
					if (i+1) == columnLength {
						for j:=0; j<rowLength; j++ {
							crates = append(crates, 32)
						}
						columnLength++
					}
					crates[(to-1)+((i+1)*rowLength)] = pile[numToBeMoved-1-k]
					break
					
				}
			}
	
		}

		fmt.Scanf("%c", &temp)
		if temp == 13 {
			break
		}
	}

	fmt.Printf("The top crates are ")
	for i:=0; i<rowLength; i++ {
		for j:=columnLength-1; j>=0; j-- { // until i find the highest item of the column
			if crates[(i)+(j*rowLength)] != 32 { // item found
				fmt.Printf("%c", crates[(i)+(j*rowLength)])
				break
			}
		}
	}
	fmt.Printf("\n")
}
