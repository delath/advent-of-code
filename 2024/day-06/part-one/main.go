package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var floor [][]byte
	var xGuard int
	var yGuard int
	var direction byte

	// Parse the file line by line
	for scanner.Scan() {
		line := []byte(scanner.Text()) // Convert string to []byte

		floor = append(floor, line)

		for i:=0; i<len(line); i++ {
			if line[i]=='^' || line[i]=='>' || line[i]=='v' || line[i]=='<' {
				direction=line[i]
				xGuard=len(floor)-1
				yGuard=i
			}
		}
	}
	
	for {
		if direction == '^' {
			if xGuard-1 < 0 {
				floor[xGuard][yGuard] = 'X'
				break
			}
			if floor[xGuard-1][yGuard]=='#' {
				direction='>'
			} else {
				floor[xGuard][yGuard] = 'X'
				xGuard--
			}
		} else if direction == '>' {
			if yGuard+1 >= len(floor[xGuard]) {
				floor[xGuard][yGuard] = 'X'
				break
			}
			if floor[xGuard][yGuard+1]=='#' {
				direction='v'
			} else {
				floor[xGuard][yGuard] = 'X'
				yGuard++
			}
		} else if direction == 'v' {
			if xGuard+1 >= len(floor) {
				floor[xGuard][yGuard] = 'X'
				break
			}
			if floor[xGuard+1][yGuard]=='#' {
				direction='<'
			} else {
				floor[xGuard][yGuard] = 'X'
				xGuard++
			}
		} else if direction == '<' {
			if yGuard-1 < 0 {
				floor[xGuard][yGuard] = 'X'
				break
			}
			if floor[xGuard][yGuard-1]=='#' {
				direction='^'
			} else {
				floor[xGuard][yGuard] = 'X'
				yGuard--
			}
		}
	}

	XCount := 0

	for i:=0; i<len(floor); i++ {
		for j:=0; j<len(floor[i]); j++ {
			if floor[i][j]=='X' {
				XCount++
			}
		}
	}

	print(XCount)
}
