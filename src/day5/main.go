package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("src/day5/test.txt");
	defer file.Close();

	r := bufio.NewReader(file);
	resultA := 0;
	resultB := 0;

	parsingPages := false;
	rules := make([][]int, 0);
	pages := make([][]int ,0);

	for {
		lineBytes, _, err := r.ReadLine();

		if (err != nil) {
			break;
		}

		if (len(lineBytes) < 1) {
			parsingPages = true;
			continue;
		}

		if (!parsingPages) {
			vals := strings.Split(string(lineBytes), "|");
			num1, _ := strconv.Atoi(vals[0]);
			num2, _ := strconv.Atoi(vals[1]);
			rule := []int{num1, num2};

			rules = append(rules, rule);
		} else {
			tmpPages := make([]int ,0);
			nums := strings.Split(string(lineBytes), ",");

			for _, n := range nums {
				num, _ := strconv.Atoi(n);
				tmpPages = append(tmpPages, num);
			}

			pages = append(pages, tmpPages)
		}
	}

	sort.Slice(pages, func(i, j int) bool {
		 
	})
	
	fmt.Printf("Rules %v\n", rules);
	fmt.Printf("Pages %v\n", pages);

	fmt.Printf("Result A: %d\n", resultA);
	fmt.Printf("Result B: %d\n", resultB);
}