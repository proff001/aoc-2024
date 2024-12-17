package main

import (
	"fmt"
	"io"
	"os"
)

func getSizeOfBlock(blocks []int, index int, rightToLeft bool) int {
	size := 0
	fileId := blocks[index]

	for {
		if (!rightToLeft && index+size > len(blocks)-1) || (rightToLeft && index-size < 0) {
			break
		}

		if !rightToLeft && fileId == blocks[index+size] {
			size++
		} else if rightToLeft && fileId == blocks[index-size] {
			size++
		} else {
			break
		}
	}

	return size
}

func main() {
	file, _ := os.Open("src/day9/input.txt")
	defer file.Close()

	lineBytes, _ := io.ReadAll(file)

	var resultA int64 = 0
	var resultB int64 = 0

	blocksA := make([]int, 0)
	blocksB := make([]int, 0)
	nextId := 0

	for i, c := range string(lineBytes) {
		num := int(c - '0')

		if i%2 == 0 {
			for i := 0; i < num; i++ {
				blocksA = append(blocksA, nextId)
				blocksB = append(blocksB, nextId)
			}

			if num != 0 {
				nextId++
			}
		} else {
			for i := 0; i < num; i++ {
				blocksA = append(blocksA, -1)
				blocksB = append(blocksB, -1)
			}
		}
	}

	for i := 0; i < len(blocksA)-1; i++ {
		shouldBreak := false

		if blocksA[i] == -1 {
			for j := len(blocksA) - 1; j > 0; j-- {
				if j == i {
					shouldBreak = true
					break
				}

				if blocksA[j] == -1 {
					continue
				}

				blocksA[i] = blocksA[j]
				blocksA[j] = -1

				break
			}
		}

		if shouldBreak {
			break
		}
	}

	for i := len(blocksB) - 1; i > 0; i-- {
		shouldBreak := false

		if blocksB[i] != -1 {
			sizeBlockA := getSizeOfBlock(blocksB, i, true)

			for j := 0; j < i-sizeBlockA; j++ {
				if j == i {
					shouldBreak = true
					break
				}

				if blocksB[j] != -1 {
					continue
				}

				sizeBlockB := getSizeOfBlock(blocksB, j, false)

				if sizeBlockA > sizeBlockB {
					j += sizeBlockB - 1
					continue
				}

				for k := 0; k < sizeBlockA; k++ {
					temp := blocksB[i-k]
					blocksB[i-k] = blocksB[j+k]
					blocksB[j+k] = temp
				}

				break
			}

			i -= sizeBlockA - 1
		}

		if shouldBreak {
			break
		}
	}

	for i, b := range blocksA {
		if b == -1 {
			continue
		}

		resultA += int64(i) * int64(b)
	}

	for i, b := range blocksB {
		if b == -1 {
			continue
		}

		resultB += int64(i) * int64(b)
	}

	fmt.Printf("Result A: %d\n", resultA)
	fmt.Printf("Result B: %d\n", resultB)
}
