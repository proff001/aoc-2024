package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bitwiseXOR(regA, regB int) {
	a := strconv.FormatInt(int64(regA), 2)
	b := strconv.FormatInt(int64(regB), 2)
	out := ""

}

func main() {
	file, _ := os.Open("src/day17/test.txt")
	defer file.Close()

	r := bufio.NewReader(file)
	resultA := 0
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

	for i, op := range operations {
		operand := 0

		if i != len(operations)-1 {
			operand = operations[i+1]
		}

		switch op {
		case 0:
			denominator := 2 ^ operand
			fmt.Printf("Op 1, denominator: %d, out: %v\n", denominator, registers[0]/denominator)
			registers[0] = registers[0] / denominator
		case 1:

		case 2:

		case 3:

		case 4:

		case 5:

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

	fmt.Printf("Result A: %d\n", resultA)
	fmt.Printf("Result B: %d\n", resultB)
}
