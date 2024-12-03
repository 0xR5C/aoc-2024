package main

import (
	"aoc-2024/utils"
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(2, 3), utils.Works())
}
