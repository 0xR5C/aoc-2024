package main

import (
	"aoc-2024/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//TODO: Make code readable

func main() {
	fmt.Println(("Day 5"))
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Split(input, "\n")
	if len(lines[len(lines)-1]) == 0 {
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
	rules, manuals := makeRuleMap(lines)
	var sum int
	for _, val := range manuals {
		man := strings.Split(val, ",")
		manNum := utils.ArrayAtoi(man)

		flag := true
		for i, val := range manNum {
			if _, present := rules[val]; present {
				ruleOk, _ := checkRule(manNum[:], rules[val], i)
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

func part2(lines []string) {
	// Find incorrect rules - Duplicate code, try to fix this
	rules, manuals := makeRuleMap(lines)
	var incorrectManuals [][]int
	for _, val := range manuals {
		man := strings.Split(val, ",")
		manNum := utils.ArrayAtoi(man)

		for i, val := range manNum {
			if _, present := rules[val]; present {
				ruleOk, _ := checkRule(manNum[:], rules[val], i)
				if !ruleOk {
					incorrectManuals = append(incorrectManuals, manNum)
					break
				}
			}
		}
	}

	// Now reorder each manual
	var sum int
	for _, manual := range incorrectManuals {
		flag := true
		for i := 0; i < len(manual); i++ {
			num := manual[i]
			if _, present := rules[num]; present {
				ruleOk, wrongNumIndex := checkRule(manual, rules[num], i)
				if !ruleOk {
					flag = false
					temp := manual[i]
					manual[i] = manual[wrongNumIndex]
					manual[wrongNumIndex] = temp
					i--
				}

			}
		}
		if !flag {
			sum += manual[len(manual)/2]
		}
	}
	fmt.Printf("PART2 The incorrect updates reordered have a middle page number sum of %d\n", sum)
}

// General Functions

func makeRuleMap(lines []string) (map[int][]int, []string) {
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
	return rules, manuals
}

func checkRule(input, rule []int, index int) (bool, int) {
	inputSet := map[int]int{}
	for i, val := range input {
		inputSet[val] = i
	}

	for _, val := range rule {
		j, ok := inputSet[val]

		if ok && j > index {
			return false, j
		}
	}

	return true, 0
}
