package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type Block struct {
	fileId   int
	position int
	size     int
}

func createBlock(id, pos int, size int) *Block {
	block := Block{fileId: id, position: pos, size: size}

	return &block
}

func printBlocks(blocks []*Block) {
	str := ""

	for _, b := range blocks {
		if b.fileId == -1 {
			str += "."
		} else {
			str += strconv.Itoa(b.fileId)
		}
	}

	fmt.Print(str, "\n")
}

func main() {
	file, _ := os.Open("src/day9/test.txt")
	defer file.Close()

	lineBytes, _ := io.ReadAll(file)

	var resultA int64 = 0
	resultB := 0

	blocks := make([]int, 0)
	blocks2 := make([]*Block, 0)
	nextId := 0

	for i, c := range string(lineBytes) {
		num := int(c - '0')

		if i%2 == 0 {
			for i := 0; i < num; i++ {
				blocks = append(blocks, nextId)
				blocks2 = append(blocks2, createBlock(nextId, len(blocks2), num))
			}

			if num != 0 {
				nextId++
			}
		} else {
			for i := 0; i < num; i++ {
				blocks = append(blocks, -1)
				blocks2 = append(blocks2, createBlock(-1, len(blocks2), num))
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

	printBlocks(blocks2)

	for i := 0; i < len(blocks2)-1; i++ {
		shouldBreak := false

		if blocks2[i].fileId == -1 {
			for j := len(blocks2) - 1; j > 0; j-- {
				if j == i {
					shouldBreak = true
					break
				}

				if blocks2[j].fileId == -1 || blocks2[i].size < blocks2[j].size {
					continue
				}

				fmt.Printf("A: %v, B: %v\n", blocks2[i], blocks2[j])

				amount := blocks2[j].size
				for k := 0; k < amount; k++ {
					temp := blocks2[i+k].position
					blocks2[i+k].position = blocks2[j-k].position
					blocks2[j-k].position = temp
				}

				sort.Slice(blocks2, func(i2, j2 int) bool {
					return blocks2[i2].position < blocks2[j2].position
				})

				break
			}
		}

		// if block.fileId == -1 {
		// 	for i := len(blocks) - 1; i > 0; i-- {
		// 		if i == bi {
		// 			shouldBreak = true
		// 			break
		// 		}

		// 		block2 := blocks[i]

		// 		if block2.fileId == -1 {
		// 			continue
		// 		}

		// 		newPos := block.position
		// 		block.position = block2.position
		// 		block2.position = newPos

		// 		sort.Slice(blocks, func(i, j int) bool {
		// 			return blocks[i].position < blocks[j].position
		// 		})

		// break
		// }
		// }

		if shouldBreak {
			break
		}
	}

	printBlocks(blocks2)

	for i, b := range blocks {
		if b == -1 {
			continue
		}

		resultA += int64(i) * int64(b)
	}

	fmt.Printf("Result A: %d\n", resultA)
	fmt.Printf("Result B: %d\n", resultB)
}
