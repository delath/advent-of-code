package main

import (
	"bufio"
	"os"
)

type State struct {
	X, Y      int
	Direction byte
}

// This was a mess but it was fast to write (❁´◡`❁)
func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var floor [][]byte
	var xGuard int
	var yGuard int
	var direction byte
	validPositions := 0

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
	
	for i:=0; i<len(floor); i++ {
		for j:=0; j<len(floor[i]); j++ {
			if floor[i][j] != '.' {
				continue
			}
			
			// Create a deep copy of floor
			copiedFloor := make([][]byte, len(floor))
			for i := range floor {
				copiedFloor[i] = make([]byte, len(floor[i]))
				copy(copiedFloor[i], floor[i])
			}

			copiedXGuard := xGuard
			copiedYGuard := yGuard

			copiedDirection := direction

			copiedFloor[i][j] = '#'

			visited := make(map[State]bool)

			for {
				currentState := State{X: copiedXGuard, Y: copiedYGuard, Direction: copiedDirection}
				visited[currentState] = true
				if copiedDirection == '^' {
					if copiedXGuard-1 < 0 {
						break
					}
					if copiedFloor[copiedXGuard-1][copiedYGuard]=='#' {
						copiedFloor[copiedXGuard][copiedYGuard] = '+'
						copiedDirection='>'
					} else {
						if copiedFloor[copiedXGuard][copiedYGuard] == '-' || copiedFloor[copiedXGuard][copiedYGuard] == '+' {
							copiedFloor[copiedXGuard][copiedYGuard] = '+'
						} else {
							copiedFloor[copiedXGuard][copiedYGuard] = '|'
						}
						copiedXGuard--
					}
				} else if copiedDirection == '>' {
					if copiedYGuard+1 >= len(copiedFloor[copiedXGuard]) {
						break
					}
					if copiedFloor[copiedXGuard][copiedYGuard+1]=='#' {
						copiedFloor[copiedXGuard][copiedYGuard] = '+'
						copiedDirection='v'
					} else {
						if copiedFloor[copiedXGuard][copiedYGuard] == '|' || copiedFloor[copiedXGuard][copiedYGuard] == '+' {
							copiedFloor[copiedXGuard][copiedYGuard] = '+'
						} else {
							copiedFloor[copiedXGuard][copiedYGuard] = '-'
						}
						copiedYGuard++
					}
				} else if copiedDirection == 'v' {
					if copiedXGuard+1 >= len(copiedFloor) {
						break
					}
					if copiedFloor[copiedXGuard+1][copiedYGuard]=='#' {
						copiedFloor[copiedXGuard][copiedYGuard] = '+'
						copiedDirection='<'
					} else {
						if copiedFloor[copiedXGuard][copiedYGuard] == '-' || copiedFloor[copiedXGuard][copiedYGuard] == '+' {
							copiedFloor[copiedXGuard][copiedYGuard] = '+'
						} else {
							copiedFloor[copiedXGuard][copiedYGuard] = '|'
						}
						copiedXGuard++
					}
				} else if copiedDirection == '<' {
					if copiedYGuard-1 < 0 {
						break
					}
					if copiedFloor[copiedXGuard][copiedYGuard-1]=='#' {
						copiedFloor[copiedXGuard][copiedYGuard] = '+'
						copiedDirection='^'
					} else {
						if copiedFloor[copiedXGuard][copiedYGuard] == '|' || copiedFloor[copiedXGuard][copiedYGuard] == '+' {
							copiedFloor[copiedXGuard][copiedYGuard] = '+'
						} else {
							copiedFloor[copiedXGuard][copiedYGuard] = '-'
						}
						copiedYGuard--
					}
				}
				nextState := State{X: copiedXGuard, Y: copiedYGuard, Direction: copiedDirection}
				if visited[nextState] {	
					validPositions++
					break
				}
			}
		}

	}

	print(validPositions)
}
