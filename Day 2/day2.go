package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)





func main() {
	fmt.Println("day2")
	dat, err := os.ReadFile("input.txt")
	check(err)
	text := string(dat[:])
	reports := strings.Split(text, "\n")
	// TODO - Implement flags
	part1(reports)
	part2(reports)
}


// PART 1
func part1(arr []string) {
	var safeReports int
	for _, rep := range reports {
		levels := strings.Fields(rep)
		numLevels := arrayAtoi(levels)
		if len(numLevels) == 0 {
			continue
		}
		if checkLevel(numLevels) s= true {
			fmt.Println(numLevels)
			safeReports += 1
		}
	}
	fmt.Printf("PART1 Safe reports are %d\n", safeReports)
}



// PART 2
func part2(arr []string) {

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

func checkLevel(arr []int) bool {
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
				return false
			}
		} else {
			if compareFunc(arr[i-1], arr[i]) == false || !checkDiff(arr[i-1], arr[i]) {
				return false
			}
		}
	}
	return true
}