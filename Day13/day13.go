package main

import (
	"aoc-2024/utils"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 13")
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	partPtr := flag.Int("part", 0, "Choose part to run")
	flag.Parse()

	if *partPtr == 1 {
		part1(lines)
	} else if *partPtr == 2 {
		part2(lines)
	} else {
		part1(lines)
		part2(lines)
	}

}

func part1(lines []string) {
	var sum int
	for i := 0; i < len(lines); i += 4 {
		Ax, Ay := extractVals(lines[i])
		Bx, By := extractVals(lines[i+1])
		Px, Py := extractVals(lines[i+2])

		var countA, countB int
		for j := 0; j <= 100; j++ {
			restX := Px - Ax*j
			restY := Py - Ay*j
			if restX%Bx == 0 && restY%By == 0 && restX/Bx == restY/By {
				countA = j
				countB = restX / Bx
				break
			}
		}
		sum += countA*3 + countB*1
	}
	fmt.Printf("PART1 The total cost is %d\n", sum)
}

func part2(lines []string) {
	// Use Cramer rule instead
	var sum int
	for i := 0; i < len(lines); i += 4 {
		Ax, Ay := extractVals(lines[i])
		Bx, By := extractVals(lines[i+1])
		Px, Py := extractVals(lines[i+2])

		Px += 10000000000000
		Py += 10000000000000

		var A, B int
		D := Ax*By - Ay*Bx
		Da := Px*By - Py*Bx
		Db := Ax*Py - Px*Ay
		if D != 0 {
			A = Da / D
			B = Db / D
			if A > 0 && B > 0 && Da == A*D && Db == B*D {
				sum += A*3 + B*1
			}
		}
	}
	fmt.Printf("PART2 The total cost is %d\n", sum)
}

func extractVals(line string) (int, int) {
	re := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(line, -1)

	if len(matches) == 2 {
		x, err := strconv.Atoi(matches[0])
		utils.Check(err)
		y, err := strconv.Atoi(matches[1])
		utils.Check(err)
		return x, y
	} else {
		return -1, -1
	}
}
