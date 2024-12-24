package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var registerA, registerB, registerC int
	var program []int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Register A:") {
			registerA, _ = strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1]))
		} else if strings.HasPrefix(line, "Register B:") {
			registerB, _ = strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1]))
		} else if strings.HasPrefix(line, "Register C:") {
			registerC, _ = strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1]))
		} else if strings.HasPrefix(line, "Program:") {
			programStr := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), ",")
			for _, numStr := range programStr {
				num, _ := strconv.Atoi(numStr)
				program = append(program, num)
			}
		}
	}

	var output []int

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

	for i := 0; i < len(output); i++ {
		print(output[i])
		if i != len(output)-1 {
			print(",")
		}
	}
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
