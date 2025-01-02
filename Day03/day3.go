package main

import (
	"aoc-2024/utils"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Day3")
	dat, err := os.ReadFile("input.txt")
	text := string(dat[:])
	utils.Check(err)

	partPtr := flag.Int("part", 0, "Choose part to run")
	flag.Parse()

	if *partPtr == 1 {
		part1(text)
	} else if *partPtr == 2 {
		part2(text)
	} else {
		part1(text)
		part2(text)
	}

}

func part1(arr string) {
	reMul := regexp.MustCompile(`mul\([0-9][0-9]{0,2},[0-9][0-9]{0,2}\)`)
	mul := reMul.FindAllString(arr, -1)

	var sum int
	reNum := regexp.MustCompile(`[0-9][0-9]{0,2}`)
	for _, val := range mul {
		nums := utils.ArrayAtoi(reNum.FindAllString(val, -1))
		sum += nums[0] * nums[1]
	}
	fmt.Printf("PART1 The sum of the multiplications is %d\n", sum)
}

func part2(arr string) {
	reMul := regexp.MustCompile(`(mul\([0-9][0-9]{0,2},[0-9][0-9]{0,2}\))|(do(n't){0,1}\(\))`)
	mul := reMul.FindAllString(arr, -1)

	var sum int
	// Multiply the sum by 1 if do() is last, by 0 if don't()
	doVal := 1
	reNum := regexp.MustCompile(`[0-9][0-9]{0,2}`)
	for _, val := range mul {
		if val == "do()" {
			doVal = 1
		} else if val == "don't()" {
			doVal = 0
		} else {
			nums := utils.ArrayAtoi(reNum.FindAllString(val, -1))
			sum += doVal * nums[0] * nums[1]
		}
	}
	fmt.Printf("PART2 The sum of the multiplications is %d\n", sum)
}
