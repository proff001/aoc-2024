package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("src/day3/input.txt");
	defer file.Close();

	r := bufio.NewReader(file);
	resultA := 0;
	resultB := 0;

	for {
		lineBytes, _, err := r.ReadLine();

		if (err != nil) {
			break;
		}		
	}

	fmt.Printf("Result A: %d\n", resultA);
	fmt.Printf("Result B: %d\n", resultB);
}