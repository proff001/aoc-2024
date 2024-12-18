package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("src/day17/input.txt")
	defer file.Close()

	r := bufio.NewReader(file)
	resultA := ""
	resultB := 0

	registers := make([]int, 3)
	operations := make([]int, 0)

	for i := 0; i < 5; i++ {
		lineBytes, _, err := r.ReadLine()

		if err != nil {
			break
		}

		str := string(lineBytes)
		num := strings.Split(str, ": ")

		if i < 3 {
			num, _ := strconv.Atoi(num[1])
			registers[i] = num
		} else if i == 4 {
			nums := strings.Split(num[1], ",")

			for j := 0; j < len(nums); j++ {
				num, _ := strconv.Atoi(nums[j])
				operations = append(operations, num)
			}
		}
	}

	for ptr := 0; ptr < len(operations)-1; ptr += 2 {
		operation := operations[ptr]
		operand := operations[ptr+1]
		combo := operand

		switch operand {
		case 4:
			combo = registers[0]
		case 5:
			combo = registers[1]
		case 6:
			combo = registers[2]
		}

		switch operation {
		case 0:
			denominator := int(math.Pow(2, float64(operand)))
			fmt.Printf("Op 0, denominator: %d, out: %v\n", denominator, registers[0]/denominator)
			registers[0] = registers[0] / denominator
		case 1:
			fmt.Printf("Op 1, regb: %b, oprand: %b, out: %b\n", registers[1], operand, registers[1]^operand)
			registers[1] = registers[1] ^ operand
		case 2:
			fmt.Printf("Op 2, operand: %d, out: %d\n", operand, operand%8)
			registers[1] = operand % 8
		case 3:
			if registers[0] != 0 {
				ptr = operand - 2
			}
		case 4:
			fmt.Printf("Op 4, regb: %b, regc: %b, out: %b\n", registers[1], registers[2], registers[1]^registers[2])
			registers[1] = registers[1] ^ registers[2]
		case 5:
			resultA += strconv.Itoa(combo % 8)
		case 6:
			denominator := 2 ^ operand
			fmt.Printf("Op 6, denominator: %d, out: %v\n", denominator, registers[0]/denominator)
			registers[1] = registers[0] / denominator
		case 7:
			denominator := 2 ^ operand
			fmt.Printf("Op 7, denominator: %d, out: %v\n", denominator, registers[0]/denominator)
			registers[2] = registers[0] / denominator
		}
	}

	fmt.Printf("Registers: %v\nOperations: %v\n", registers, operations)

	fmt.Printf("Result A: %q\n", resultA)
	fmt.Printf("Result B: %d\n", resultB)
}
