package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var jumpMap map[int]int

var program []int

var potentialSolutions []int

// This is a complete mess and it was complex af due to bit shifting and stuff, so i'm not gonna refactor it
func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var registerA, registerB, registerC int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Register A:") {
			registerA = 0
		} else if strings.HasPrefix(line, "Register B:") {
			registerB = -1
		} else if strings.HasPrefix(line, "Register C:") {
			registerC = -1
		} else if strings.HasPrefix(line, "Program:") {
			programStr := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), ",")
			for _, numStr := range programStr {
				num, _ := strconv.Atoi(numStr)
				program = append(program, num)
			}
		}
	}

	jumpMap = precomputeJumps(program)

	findRegisterA(registerA, registerB, registerC, len(program)-2, len(program)-1)

	for _, solution := range potentialSolutions {
		var output []int
		registerA = solution
		registerB = 0
		registerC = 0

		for i := 0; i < len(program); i = i + 2 {
			opcode := program[i]
			operand := program[i+1]

			if opcode == 0 {
				comboValue := getComboValue(operand, registerA, registerB, registerC)

				result := registerA >> comboValue

				registerA = result
			} else if opcode == 1 {
				registerB = registerB ^ operand
			} else if opcode == 2 {
				comboValue := getComboValue(operand, registerA, registerB, registerC)

				registerB = comboValue % 8
			} else if opcode == 3 {
				if registerA != 0 {
					i = operand - 2
					continue
				}
			} else if opcode == 4 {
				registerB = registerB ^ registerC
			} else if opcode == 5 {
				comboValue := getComboValue(operand, registerA, registerB, registerC)

				output = append(output, comboValue%8)
			} else if opcode == 6 {
				comboValue := getComboValue(operand, registerA, registerB, registerC)

				result := registerA / (1 << comboValue)

				registerB = result
			} else if opcode == 7 {
				comboValue := getComboValue(operand, registerA, registerB, registerC)

				result := registerA / (1 << comboValue)

				registerC = result
			}
		}

		if arraysEqual(output, program) {
			print(solution)
			print("\n")
		}
	}
}

func arraysEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func findRegisterA(registerA, registerB, registerC int, index int, programIndex int) {
	for i := index; i >= 0; i -= 2 {
		opcode := program[i]
		operand := program[i+1]

		if opcode == 0 {
			if operand >= 4 { // this never happens :D
				panic("This fucks up registerA, how do i revert?")
			}
			comboValue := getComboValue(operand, registerA, registerB, registerC)

			if registerA == 0 {
				for j := 1; j < (1 << comboValue); j++ {
					findRegisterA(j, registerB, registerC, i-2, programIndex)
					jumpTarget, exists := jumpMap[i]
					if exists {
						findRegisterA(j, registerB, registerC, jumpTarget, programIndex)
					}
				}

				return
			} else {
				for j := 0; j < (1 << comboValue); j++ {
					findRegisterA((registerA<<3)|j, registerB, registerC, i-2, programIndex)
					jumpTarget, exists := jumpMap[i]
					if exists {
						findRegisterA((registerA<<3)|j, registerB, registerC, jumpTarget, programIndex)
					}
				}

				return
			}
		} else if opcode == 1 {
			// ignore cos i dun care bro
		} else if opcode == 2 {
			if operand == 4 {
				comboValue := getComboValue(operand, registerA, registerB, registerC)

				isValidBranch := parseProgram(registerA, comboValue%8, registerC, i+2, programIndex)
				if !isValidBranch {
					return
				}
			}
		} else if opcode == 3 {
			// This has already been inverted :D
		} else if opcode == 4 {
			// ignore cos i dun care bro
		} else if opcode == 5 {
			if operand != 5 && operand != 6 {
				comboValue := getComboValue(operand, registerA, registerB, registerC)

				if programIndex < 0 || program[programIndex] != comboValue%8 {
					return
				}
				//print(comboValue % 8)
				//print(" ")
			} else {
				//print("X ")
				//print(" ")
			}
			programIndex--
		} else if opcode == 6 {
			// NOT CALLED
		} else if opcode == 7 {
			// ignore cos i dun care bro
		}

		jumpTarget, exists := jumpMap[i]
		if exists && registerA != 0 && programIndex > -1 {
			findRegisterA(registerA, registerB, registerC, jumpTarget, programIndex)
		}

	}

	if (index >= 0 || programIndex == -1) && registerA != 0 {
		potentialSolutions = append(potentialSolutions, registerA)
	}
}

func parseProgram(registerA, registerB, registerC int, index int, programIndex int) bool {
	for i := index; i < len(program); i = i + 2 {
		opcode := program[i]
		operand := program[i+1]

		if opcode == 0 {
			comboValue := getComboValue(operand, registerA, registerB, registerC)

			result := registerA >> comboValue

			registerA = result
		} else if opcode == 1 {
			registerB = registerB ^ operand
		} else if opcode == 2 {
			comboValue := getComboValue(operand, registerA, registerB, registerC)

			registerB = comboValue % 8
		} else if opcode == 3 {
			if registerA != 0 {
				i = operand - 2
				continue
			}
		} else if opcode == 4 {
			registerB = registerB ^ registerC
		} else if opcode == 5 {
			programIndex++
			expectedResult := program[programIndex]

			comboValue := getComboValue(operand, registerA, registerB, registerC)

			if comboValue%8 != expectedResult {
				return false
			}
		} else if opcode == 6 {
			comboValue := getComboValue(operand, registerA, registerB, registerC)

			result := registerA / (1 << comboValue)

			registerB = result
		} else if opcode == 7 {
			comboValue := getComboValue(operand, registerA, registerB, registerC)

			result := registerA / (1 << comboValue)

			registerC = result
		}
	}
	return true
}

func getComboValue(operand, registerA, registerB, registerC int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registerA
	case 5:
		return registerB
	case 6:
		return registerC
	default:
		panic("Invalid operand")
	}
}

func precomputeJumps(program []int) map[int]int {
	jumpMap := make(map[int]int)

	for i := 0; i < len(program); i += 2 {
		opcode := program[i]
		if opcode == 3 && i != len(program) {
			jumpTarget := program[i+1]
			jumpMap[jumpTarget] = i
		}
	}

	return jumpMap
}
