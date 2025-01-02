package main

import (
	"aoc-2024/utils"
	"flag"
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

	partPtr := flag.Int("part", 0, "Choose part to run")
	flag.Parse()

	if *partPtr == 1 {
		part1(equations)
	} else if *partPtr == 2 {
		part2(equations)
	} else {
		part1(equations)
		part2(equations)
	}

}

func part1(input []Equation) {
	var sum int
	for i := range input {
		checkOperator(&input[i], input[i].nums[0], 1)
		if input[i].count != 0 {
			sum += input[i].test
		}
	}
	fmt.Printf("PART1 The total calibration is %d\n", sum)
}

func part2(input []Equation) {
	var sum int
	for i := range input {
		checkOperatorPart2(&input[i], input[i].nums[0], 1)
		if input[i].count != 0 {
			sum += input[i].test
		}
	}
	fmt.Printf("PART2 The total calibration with || is %d\n", sum)
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

func checkOperatorPart2(eq *Equation, sum, i int) bool {
	if len(eq.nums)-1 == i {
		// Final element of Operation
		tempPlus := sum + eq.nums[i]
		tempMul := sum * eq.nums[i]

		tempConc, err := strconv.Atoi(strconv.Itoa(sum) + strconv.Itoa(eq.nums[i]))
		utils.Check(err)

		if tempPlus == eq.test {
			eq.count++
		}

		if tempMul == eq.test {
			eq.count++
		}

		if tempConc == eq.test {
			eq.count++
		}

		if tempMul == eq.test || tempPlus == eq.test || tempConc == eq.test {
			return true
		} else {
			return false
		}

	} else {
		tempPlus := sum + eq.nums[i]
		tempMul := sum * eq.nums[i]

		tempConc, err := strconv.Atoi(strconv.Itoa(sum) + strconv.Itoa(eq.nums[i]))
		utils.Check(err)

		okPlus := checkOperatorPart2(eq, tempPlus, i+1)
		okMul := checkOperatorPart2(eq, tempMul, i+1)
		okConc := checkOperatorPart2(eq, tempConc, i+1)

		if (tempPlus <= eq.test && okPlus) || (tempMul <= eq.test && okMul) || (tempConc <= eq.test && okConc) {
			return true
		}

		return false

	}

}
