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
	file, _ := os.Open("src/day2/input.txt")
	defer file.Close();

	r := bufio.NewReader(file);
	safe := 0;

	for {
		line, _, err := r.ReadLine();

		if (err != nil) {
			break;
		}

		levels := strings.Split(string(line), " ");
		
		directionSet := false;
		ascending := false;

		for i := 0; i < len(levels) - 1; i++ {
			lvl1, _ := strconv.Atoi(levels[i]);
			lvl2, _ := strconv.Atoi(levels[i + 1]);
			diff := int(math.Abs(float64(lvl1 - lvl2)));
			
			if (!directionSet) {
				if (lvl1 > lvl2) {
					ascending = false;
				} else {
					ascending = true;
				}
				
				directionSet = true;
			}

			if (diff > 0 && diff < 4) {
				if ((!ascending && lvl1 > lvl2) || ascending && lvl1 < lvl2) {
					if (i == len(levels) - 2) {
						safe++;
					}
				} else {
					break;
				}
			} else {
				break;
			}
		}
	}

	fmt.Printf("Result A: %d\n", safe);
	// fmt.Printf("Result B: %d\n", resultB);
}