package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("src/day4/test.txt");
	defer file.Close();

	r := bufio.NewReader(file);
	resultA := 0;
	resultB := 0;

	lines := make([][]byte, 0);

	for {
		line, _, err := r.ReadLine();

		if (err != nil) {
			break;
		}

		lines = append(lines, line);
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i];

		for j := 0; j < len(line); j++ {
			fmt.Printf("Line: %q", string(lines[i]));
			if (line[j] == 'X') {
				fmt.Printf("Found X: %d On line: %d\n", j, i)
				fmt.Printf("Char: %q, Is X: %t\n", line[j], line[j] == 'X')
				
				if (j > 3) {
					fmt.Printf("Found XMAS Backwards: %d\n", j)
					if (line[j - 1] == 'M' && line[j - 2] == 'A' && line[j - 3] == 'S') {
						resultA++;
					}
				}
				
				if (j < len(line) - 3) {
					fmt.Printf("Found XMAS Forwards: %d\n", j)
					if (line[j + 1] == 'M' && line[j + 2] == 'A' && line[j + 3] == 'S') {
						resultA++;
					}
				}
				
				if (i > 3) {
					fmt.Printf("Found XMAS Upwards: %d\n", j)
					if (lines[i - 1][j] == 'M' && lines[i - 2][j] == 'A' && lines[i - 3][j] == 'S') {
						resultA++;
					}
				}
				
				if (i < len(line) - 3) {
					fmt.Printf("Found XMAS Downwards: %d\n", j)
					if (lines[i + 1][j] == 'M' && lines[i + 2][j] == 'A' && lines[i - 3][j] == 'S') {
						resultA++;
					}
				}
			}
		}
	}

	fmt.Printf("Result A: %d\n", resultA);
	fmt.Printf("Result B: %d\n", resultB);
}