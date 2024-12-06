package main

import (
	"aoc-2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//TODO: Make code readable

func main() {
	fmt.Println(("Day 4"))
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Split(input, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	part1(lines)
}

func part1(lines []string) {
	rules := make(map[int][]int)
	var manuals []string
	for i := range lines {
		if len(lines[i]) == 0 {
			manuals = lines[i+1:]
			break
		}
		numsStr := strings.Split(lines[i], "|")
		num1, ok := strconv.Atoi(numsStr[0])
		utils.Check(ok)
		num2, ok := strconv.Atoi(numsStr[1])
		utils.Check(ok)

		_, present := rules[num2]
		if !present {
			rules[num2] = make([]int, 0)
		}
		rules[num2] = append(rules[num2], num1)

	}
	var sum int
	for _, val := range manuals {
		man := strings.Split(val, ",")
		manNum := utils.ArrayAtoi(man)

		flag := true
		for i, val := range manNum {
			if _, present := rules[val]; present {
				ruleOk := checkRule(manNum[:], rules[val], i)
				if !ruleOk {
					flag = false
					break
				}
			}
		}
		if flag {
			sum += manNum[len(manNum)/2]
		}
	}

	fmt.Printf("PART1 The correct updates have a middle page number sum of %d\n", sum)
}

func checkRule(input, rule []int, index int) bool {
	inputSet := map[int]int{}
	for i, val := range input {
		inputSet[val] = i
	}

	for _, val := range rule {
		j, ok := inputSet[val]

		if ok && j > index {
			return false
		}
	}

	return true
}
