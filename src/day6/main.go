package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("src/day6/test.txt")
	defer file.Close()

	r := bufio.NewReader(file)
	resultA := 0
	resultB := 0

	parsedMap := make([][]byte, 0)
	position := []int{0, 0}
	visitedPositions := make([][]byte, 0)

	for {
		lineBytes, _, err := r.ReadLine()

		if err != nil {
			break
		}

		fmt.Printf("Line: %q\n", lineBytes)

		row := make([]byte, len(lineBytes))

		copy(lineBytes, row)

		parsedMap = append(parsedMap, row)
	}

	for _, line := range parsedMap {
		fmt.Printf("Map: %q\n", line)
	}

	fmt.Printf("Result A: %d\n", resultA)
	fmt.Printf("Result B: %d\n", resultB)
}
