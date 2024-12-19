package main

import (
	"aoc-2024/utils"
	"flag"
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

	partPtr := flag.Int("part", 0, "Choose part to run")
	flag.Parse()

	if *partPtr == 1 {
		part1(lines[0])
	} else if *partPtr == 2 {
		part2(lines[0])
	} else {
		part1(lines[0])
		part2(lines[0])
	}
}

func part1(disk string) {
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
	fmt.Printf("PART1 The filesystem checksum is %d\n", sum)
}

func part2(disk string) {
	fmt.Println("PART2")
	block := createBlocks(disk)
	var consecutiveFile int
	var consecutiveFileVal string
	for i := lastNum(block); i >= 0; i-- {
		if block[i] == consecutiveFileVal {
			consecutiveFile++
		} else if block[i] != consecutiveFileVal {
			if consecutiveFileVal != "" {
				//
				index := findSpace(block, consecutiveFile, i+1)
				if index != -1 {
					for j, k := 0, i+1; j < consecutiveFile; j++ {
						temp := block[j+index]
						block[j+index] = block[k]
						block[k] = temp
						k++
					}
				}

			}
			if block[i] != "." {
				consecutiveFile = 1
				consecutiveFileVal = block[i]
			} else {
				consecutiveFile = 0
				consecutiveFileVal = ""
			}
		}

	}
	var sum int
	for i, val := range block {
		if val != "." {
			num, err := strconv.Atoi(val)
			utils.Check(err)
			sum += i * num
		}
	}
	fmt.Printf("PART2 The fileystem checksum is %d\n", sum)

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

func findSpace(arr []string, size, index int) int {
	var consecutiveSpace, iSpace int
	for i, val := range arr {
		if i > index {
			return -1
		}
		if val == "." && consecutiveSpace == 0 {
			consecutiveSpace++
			iSpace = i
		} else if val == "." {
			consecutiveSpace++
		} else {
			if consecutiveSpace >= size {
				return iSpace
			}
			consecutiveSpace = 0
		}
	}
	return -1
}
