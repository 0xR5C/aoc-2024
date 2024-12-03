package main

import (
	"aoc-2024/utils"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Day3")
	dat, err := os.ReadFile("input.txt")
	text := string(dat[:])
	utils.Check(err)
	fmt.Println(text)

	reMul := regexp.MustCompile(`mul\([0-9][0-9]{0,2},[0-9][0-9]{0,2}\)`)
	mul := reMul.FindAllString(text, -1)

	var sum int
	reNum := regexp.MustCompile(`[0-9][0-9]{0,2}`)
	for _, val := range mul {
		nums := utils.ArrayAtoi(reNum.FindAllString(val, -1))
		sum += nums[0] * nums[1]
	}
	fmt.Println(sum)
}
