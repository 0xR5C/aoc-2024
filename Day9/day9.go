package main

import (
	"aoc-2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 9")
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Split(input, "\n")

	part1(lines[0])
}

func part1(disk string) {
	fmt.Println("PART1")
	block := createBlocks(disk)

	for i, j := 0, lastNum(block); i < j; i++ {
		if block[i] == "." {
			temp := block[i]
			block[i] = block[j]
			block[j] = temp
			j = lastNum(block)
		}
	}

	var sum int
	for i, val := range block {
		if val == "." {
			break
		}
		num, err := strconv.Atoi(val)
		utils.Check(err)
		sum += i * num
	}
	fmt.Println(sum)
}

func createBlocks(disk string) []string {
	var index int
	block := make([]string, 0)
	for i, val := range disk {
		if i%2 == 0 {
			tempStr := string(val)
			indexStr := strconv.Itoa(index)
			tempNum, err := strconv.Atoi(tempStr)
			utils.Check(err)
			for j := 0; j < tempNum; j++ {
				block = append(block, indexStr)
			}
			index++
		} else {
			tempStr := string(val)
			tempNum, err := strconv.Atoi(tempStr)
			utils.Check(err)
			for j := 0; j < tempNum; j++ {
				block = append(block, ".")
			}
		}
	}
	return block
}

func lastNum(arr []string) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != "." {
			return i
		}
	}
	return -1
}
