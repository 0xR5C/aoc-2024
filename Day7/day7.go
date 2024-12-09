package main

import (
	"aoc-2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	test  int
	count int
	nums  []int
}

func main() {
	fmt.Println("Day 7")
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	// Create input array
	equations := make([]Equation, len(lines))
	for i := range lines {
		split := strings.Split(lines[i], ":")
		tempTest, err := strconv.Atoi(split[0])
		utils.Check(err)
		tempNums := strings.Fields(split[1])
		equations[i] = Equation{tempTest, 0, utils.ArrayAtoi(tempNums)}
	}
	part1(equations)

}

func part1(input []Equation) {
	var sum int
	for i := range input {
		checkOperator(&input[i], input[i].nums[0], 1)
		if input[i].count != 0 {
			sum += input[i].test
		}
	}
	fmt.Println(sum)
}

func checkOperator(eq *Equation, sum, i int) bool {
	if len(eq.nums)-1 == i {
		// Final element of Operation
		tempPlus := sum + eq.nums[i]
		tempMul := sum * eq.nums[i]
		if tempPlus == eq.test {
			eq.count++
		}

		if tempMul == eq.test {
			eq.count++
		}

		if tempMul == eq.test || tempPlus == eq.test {
			return true
		} else {
			return false
		}

	} else {
		tempPlus := sum + eq.nums[i]
		tempMul := sum * eq.nums[i]
		okPlus := checkOperator(eq, tempPlus, i+1)
		okMul := checkOperator(eq, tempMul, i+1)
		if (tempPlus <= eq.test && okPlus) || (tempMul <= eq.test && okMul) {
			return true
		}

		return false

	}

}
