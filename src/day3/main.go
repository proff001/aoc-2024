package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("src/day3/input.txt")
	defer file.Close()

	regexNums := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	regexDos := regexp.MustCompile(`do\(\)`)
	regexDonts := regexp.MustCompile(`don\'t\(\)`)

	r := bufio.NewReader(file)
	resultA := 0
	resultB := 0

	do := true

	for {
		lineBytes, _, err := r.ReadLine()

		if err != nil {
			break
		}

		line := string(lineBytes)

		mulS := regexNums.FindAllStringSubmatchIndex(line, -1)
		doS := regexDos.FindAllStringSubmatchIndex(line, -1)
		dontS := regexDonts.FindAllStringSubmatchIndex(line, -1)

		instructions := make([][]int, 0)

		instructions = append(instructions, mulS...)
		instructions = append(instructions, doS...)
		instructions = append(instructions, dontS...)

		for i := 0; i < len(instructions)-1; i++ {
			for j := 0; j < len(instructions)-i-1; j++ {
				if instructions[j][0] > instructions[j+1][0] {
					instructions[j], instructions[j+1] = instructions[j+1], instructions[j]
				}
			}
		}

		for _, instruction := range instructions {
			str := line[instruction[0]:instruction[1]]

			if str == "do()" {
				do = true
			} else if str == "don't()" {
				do = false
			} else {
				numbers := regexNums.FindStringSubmatch(str)

				if len(numbers) != 3 {
					continue
				}

				num1, err1 := strconv.Atoi(numbers[1])
				num2, err2 := strconv.Atoi(numbers[2])

				if err1 != nil && err2 != nil {
					continue
				}

				resultA += num1 * num2

				if do {
					resultB += num1 * num2
				}
			}
		}
	}

	fmt.Printf("Result A: %d\n", resultA)
	fmt.Printf("Result B: %d\n", resultB)
}
