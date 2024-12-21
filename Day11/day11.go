package main

import (
	"aoc-2024/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Day 11")
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Fields(input)
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

func part1(line []string) {
	start := time.Now()
	mapLine := make(map[string]int)
	for _, val := range line {
		mapLine[val] += 1
	}

	for i := 0; i < 25; i++ {
		mapLine = blink(mapLine)
	}
	var sum int
	for _, val := range mapLine {
		sum += val
	}
	elapsed := time.Since(start)
	fmt.Printf("PART1 The number of stones after 25 blinks is %d\n", sum)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}

func part2(line []string) {
	start := time.Now()
	mapLine := make(map[string]int)
	for _, val := range line {
		mapLine[val] += 1
	}

	for i := 0; i < 75; i++ {
		mapLine = blink(mapLine)
	}
	var sum int
	for _, val := range mapLine {
		sum += val
	}
	elapsed := time.Since(start)
	fmt.Printf("PART2 The number of stones after 75 blinks is %d\n", sum)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}

func blink(line map[string]int) map[string]int {
	tempMap := make(map[string]int)
	for stone, count := range line {
		if count > 0 {
			if stone == "0" {
				tempMap["1"] += count
			} else if len(stone)%2 == 0 {
				var leftPart, rightPart string
				for j, run := range stone {
					if j < len(stone)/2 {
						leftPart = leftPart + string(run)
					} else {
						rightPart = rightPart + string(run)
					}
				}
				rightPart = strings.TrimLeft(rightPart, "0")
				if len(rightPart) == 0 {
					// if rightpart is empty, it's a zero
					rightPart += "0"
				}
				tempMap[rightPart] += count
				tempMap[leftPart] += count
			} else {
				num, err := strconv.Atoi(stone)
				utils.Check(err)
				num *= 2024
				tempMap[strconv.Itoa(num)] += count
			}
			line[stone] = 0
		}
	}
	return tempMap
}
