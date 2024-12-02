package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day2")
	dat, err := os.ReadFile("input.txt")
	check(err)
	text := string(dat[:])
	reports := strings.Split(text, "\n")

	partPtr := flag.Int("part", 0, "Choose part to run")
	flag.Parse()

	if *partPtr == 1 {
		part1(reports)
	} else if *partPtr == 2 {
		part2(reports)
	} else {
		part1(reports)
		part2(reports)
	}
}

// PART 1
func part1(arr []string) {
	var safeReports int
	for _, rep := range arr {
		levels := strings.Fields(rep)
		numLevels := arrayAtoi(levels)
		if len(numLevels) == 0 {
			continue
		}
		ret, _ := checkLevel(numLevels)
		if ret == true {
			safeReports += 1
		}
	}
	fmt.Printf("PART1 Safe reports are %d\n", safeReports)
}

// PART 2
func part2(arr []string) {
	var safeReports int
	for _, rep := range arr {
		levels := strings.Fields(rep)
		numLevels := arrayAtoi(levels)
		if len(numLevels) == 0 {
			continue
		}
		ret, i := checkLevel(numLevels)
		if ret == true {
			safeReports += 1
		} else {
			for j := i; j >= 0; j-- {
				singleBad := removeElement(numLevels, j)
				ret, _ := checkLevel(singleBad)
				if ret == true {
					safeReports += 1
					break
				}
			}
		}
	}
	fmt.Printf("PART2 Safe reports are %d\n", safeReports)
}

// Utility Functions
// TODO - Create utils package

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func arrayAtoi(arr []string) []int {
	numArr := make([]int, len(arr))
	for i := range arr {
		val, ok := strconv.Atoi(arr[i])
		check(ok)
		numArr[i] = val
	}
	return numArr
}

func increasing(a, b int) bool {
	return a < b
}

func decreasing(a, b int) bool {
	return a > b
}

func checkDiff(a, b int) bool {
	if a > b {
		return a-b >= 1 && a-b <= 3
	} else if a < b {
		return b-a >= 1 && b-a <= 3
	} else {
		return false
	}
}

// Return true or false and index of problem if false (if true return -1)
func checkLevel(arr []int) (bool, int) {
	var compareFunc func(int, int) bool
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			continue
		} else if i == 1 {
			if arr[0] > arr[1] && checkDiff(arr[0], arr[1]) {
				compareFunc = decreasing
			} else if arr[0] < arr[1] && checkDiff(arr[0], arr[1]) {
				compareFunc = increasing
			} else {
				return false, 1
			}
		} else {
			if compareFunc(arr[i-1], arr[i]) == false || !checkDiff(arr[i-1], arr[i]) {
				return false, i
			}
		}
	}
	return true, -1
}

func removeElement(arr []int, index int) []int {
	newArr := make([]int, len(arr)-1)
	originalIndex := 0
	for i := range arr {
		if i != index {
			newArr[originalIndex] = arr[i]
			originalIndex += 1
		}
	}
	return newArr
}
