package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Block struct {
	fileId   int
	position int
}

func createBlock(id, pos int) *Block {
	block := Block{fileId: id, position: pos}

	return &block
}

func getSizeOfBlock(blocks []*Block, index int, rightToLeft bool) int {
	size := 0
	fileId := blocks[index].fileId

	for {
		if (!rightToLeft && index+size > len(blocks)-1) || (rightToLeft && index-size < 0) {
			break
		}

		if !rightToLeft && fileId == blocks[index+size].fileId {
			size++
		} else if rightToLeft && fileId == blocks[index-size].fileId {
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

	blocks := make([]int, 0)
	blocks2 := make([]*Block, 0)
	nextId := 0

	for i, c := range string(lineBytes) {
		num := int(c - '0')

		if i%2 == 0 {
			for i := 0; i < num; i++ {
				blocks = append(blocks, nextId)
				blocks2 = append(blocks2, createBlock(nextId, len(blocks2)))
			}

			if num != 0 {
				nextId++
			}
		} else {
			for i := 0; i < num; i++ {
				blocks = append(blocks, -1)
				blocks2 = append(blocks2, createBlock(-1, len(blocks2)))
			}
		}
	}

	for i := 0; i < len(blocks)-1; i++ {
		shouldBreak := false

		if blocks[i] == -1 {
			for j := len(blocks) - 1; j > 0; j-- {
				if j == i {
					shouldBreak = true
					break
				}

				if blocks[j] == -1 {
					continue
				}

				blocks[i] = blocks[j]
				blocks[j] = -1

				break
			}
		}

		if shouldBreak {
			break
		}
	}

	for i := len(blocks2) - 1; i > 0; i-- {
		shouldBreak := false

		if blocks2[i].fileId != -1 {
			sizeBlockA := getSizeOfBlock(blocks2, i, true)

			for j := 0; j < i-sizeBlockA; j++ {
				if j == i {
					shouldBreak = true
					break
				}

				if blocks2[j].fileId != -1 {
					continue
				}

				sizeBlockB := getSizeOfBlock(blocks2, j, false)

				if sizeBlockA > sizeBlockB {
					j += sizeBlockB - 1
					continue
				}

				for k := 0; k < sizeBlockA; k++ {
					temp := blocks2[i-k].position
					blocks2[i-k].position = blocks2[j+k].position
					blocks2[j+k].position = temp
				}

				sort.Slice(blocks2, func(i2, j2 int) bool {
					return blocks2[i2].position < blocks2[j2].position
				})

				break
			}

			i -= sizeBlockA - 1
		}

		if shouldBreak {
			break
		}
	}

	for i, b := range blocks {
		if b == -1 {
			continue
		}

		resultA += int64(i) * int64(b)
	}

	for i, b := range blocks2 {
		if b.fileId == -1 {
			continue
		}

		resultB += int64(i) * int64(b.fileId)
	}

	fmt.Printf("Result A: %d\n", resultA)
	fmt.Printf("Result B: %d\n", resultB)
}
