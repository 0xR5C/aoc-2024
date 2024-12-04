package main

import (
	"aoc-2024/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(("Day 4"))
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Fields(input)

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	sum := 0
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'X' {
				// check horizontal
				if j < len(lines[i])-3 && lines[i][j+1] == 'M' && lines[i][j+2] == 'A' && lines[i][j+3] == 'S' {
					sum++
				}

				// check vertical
				if i < len(lines)-3 && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
					sum++
				}

				// check diagonal left
				if i < len(lines)-3 && j >= 3 && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
					sum++
				}

				// check diagonal left
				if i < len(lines)-3 && j < len(lines[i])-3 && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
					sum++
				}

			} else if lines[i][j] == 'S' {
				// check horizontal backwards
				if j < len(lines[i])-3 && lines[i][j+1] == 'A' && lines[i][j+2] == 'M' && lines[i][j+3] == 'X' {
					sum++
				}
				// check vertical
				if i < len(lines)-3 && lines[i+1][j] == 'A' && lines[i+2][j] == 'M' && lines[i+3][j] == 'X' {
					sum++
				}

				// check diagonal left
				if i < len(lines)-3 && j >= 3 && lines[i+1][j-1] == 'A' && lines[i+2][j-2] == 'M' && lines[i+3][j-3] == 'X' {
					sum++
				}

				// check diagonal left
				if i < len(lines)-3 && j < len(lines[i])-3 && lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'M' && lines[i+3][j+3] == 'X' {
					sum++
				}
			}
		}
	}

	fmt.Printf("PART1 'X-MAS' appears %d times\n", sum)
}

func part2(lines []string) {
	sum := 0
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'A' && i > 0 && i < len(lines)-1 && j > 0 && j < len(lines[i])-1 {
				//M left, S right
				if lines[i-1][j-1] == 'M' && lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S' && lines[i+1][j+1] == 'S' {
					sum++
				}
				// M up, S down
				if lines[i-1][j-1] == 'M' && lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M' && lines[i+1][j+1] == 'S' {
					sum++
				}
				// M right, S left
				if lines[i-1][j-1] == 'S' && lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M' && lines[i+1][j+1] == 'M' {
					sum++
				}
				// M down, S up
				if lines[i-1][j-1] == 'S' && lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S' && lines[i+1][j+1] == 'M' {
					sum++
				}
			}
		}
	}
	fmt.Printf("PART2 'X-MAS' appears %d times\n", sum)
}
