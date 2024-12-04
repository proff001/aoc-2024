package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("src/day4/input.txt");
	defer file.Close();

	r := bufio.NewReader(file);
	resultA := 0;
	resultB := 0;

	lines := make([][]byte, 0);

	for {
		rLine, _, err := r.ReadLine();

		if (err != nil) {
			break;
		}

		line := make([]byte, len(rLine));

		copy(line, rLine);	
		
		lines = append(lines, line);
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i];

		for j := 0; j < len(line); j++ {
			if (line[j] == 'X') {
				if (j > 2) {
					if (line[j - 1] == 'M' && line[j - 2] == 'A' && line[j - 3] == 'S') {
						resultA++;
					}
				}

				if (j < len(line) - 3) {
					if (line[j + 1] == 'M' && line[j + 2] == 'A' && line[j + 3] == 'S') {
						resultA++;
					}
				}

				if (i > 2) {
					if (lines[i - 1][j] == 'M' && lines[i - 2][j] == 'A' && lines[i - 3][j] == 'S') {
						resultA++;
					}
				}

				if (i < len(lines) - 3) {
					if (lines[i + 1][j] == 'M' && lines[i + 2][j] == 'A' && lines[i + 3][j] == 'S') {
						resultA++;
					}
				}

				if (j > 2 && i > 2) {
					if (lines[i - 1][j - 1] == 'M' && lines[i - 2][j - 2] == 'A' && lines[i - 3][j - 3] == 'S') {
						resultA++;
					}
				}

				if (j < len(line) - 3 && i > 2) {
					if (lines[i - 1][j + 1] == 'M' && lines[i - 2][j + 2] == 'A' && lines[i - 3][j + 3] == 'S') {
						resultA++;
					}
				}

				if (j > 2 && i < len(lines) - 3) {
					if (lines[i + 1][j - 1] == 'M' && lines[i + 2][j - 2] == 'A' && lines[i + 3][j - 3] == 'S') {
						resultA++;
					}
				}

				if (j < len(line) - 3 && i < len(lines) - 3) {
					if (lines[i + 1][j + 1] == 'M' && lines[i + 2][j + 2] == 'A' && lines[i + 3][j + 3] == 'S') {	
						resultA++;
					}
				}
			}

			if (line[j] == 'A' && i > 0 && i < len(lines) - 1 && j > 0 && j < len(line) - 1) {
				topR, topL, botR, botL := lines[i - 1][j - 1], lines[i - 1][j + 1], lines[i + 1][j - 1], lines[i + 1][j + 1];

				if ((topR == 'M' && botL == 'S') || (topR == 'S' && botL == 'M')) {
					if ((topL == 'M' && botR == 'S') || (topL == 'S' && botR == 'M')) {
						resultB++;
					}
				}
			}
		}
	}

	fmt.Printf("Result A: %d\n", resultA);
	fmt.Printf("Result B: %d\n", resultB);
}