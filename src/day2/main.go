package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getAbsDiff(val1 int, val2 int) int {
	return int(math.Abs(float64(val1 - val2)))
}

func checkValues(levels []int) bool {
	ascending := (levels[0] < levels[len(levels)-1])

	for i := 0; i < len(levels)-1; i++ {
		val1, val2 := levels[i], levels[i+1]
		diff := getAbsDiff(val1, val2)

		if diff == 0 || diff > 3 {
			return false
		}

		if ascending && val1 > val2 {
			return false
		}

		if !ascending && val1 < val2 {
			return false
		}

	}

	return true
}

func main() {
	file, _ := os.Open("src/day2/input.txt")
	defer file.Close()

	r := bufio.NewReader(file)
	resultA := 0
	resultB := 0

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		strLevels := strings.Split(string(line), " ")
		levels := make([]int, len(strLevels))

		for i, s := range strLevels {
			levels[i], _ = strconv.Atoi(s)
		}

		if checkValues(levels) {
			resultA++
			resultB++
		} else {
			for i := 0; i < len(levels); i++ {
				dampenedLevels := make([]int, 0, len(levels)-1)
				dampenedLevels = append(dampenedLevels, levels[:i]...)
				dampenedLevels = append(dampenedLevels, levels[i+1:]...)

				if checkValues(dampenedLevels) {
					resultB++
					break
				}
			}
		}
	}

	fmt.Printf("Result A: %d\n", resultA)
	fmt.Printf("Result B: %d\n", resultB)
}
