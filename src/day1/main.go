package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("src/day1/input.txt")
	defer file.Close();

	r := bufio.NewReader(file);

	left := make([]int, 0);
	right := make([]int, 0);

	for {
		line, _, err := r.ReadLine();

		if (err != nil) {
			break;
		}
		entries := strings.Split(string(line), "   ");
		val1, _ :=  strconv.Atoi(entries[0]);
		val2, _ :=  strconv.Atoi(entries[1]);

		left = append(left, val1);
		right = append(right, val2);
	}

	slices.Sort(left);
	slices.Sort(right);

	resultA := 0;
	resultB := 0;

	for i := 0; i < len(left); i++ {
		resultA += int(math.Abs(float64(left[i] - right[i])));

		count := 0;
		for j := 0; j < len(right); j++ {
			if (right[j] == left[i]) {
				count++;
			}
		}

		resultB += (left[i] * count)
	}

	fmt.Printf("Result A: %d\n", resultA);
	fmt.Printf("Result B: %d\n", resultB);
}