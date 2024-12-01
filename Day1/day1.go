package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printArr(arr []string) {
	for _, val := range arr {
		fmt.Println(val)
	}
}

func distInt(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	text := string(dat[:])
	split := strings.Fields(text)
	list1 := make([]string, 0)
	list2 := make([]string, 0)
	for i, val := range split {
		if i%2 == 0 {
			list1 = append(list1, val)
		} else {
			list2 = append(list2, val)
		}
	}
	slices.Sort(list1)
	slices.Sort(list2)

	partPtr := flag.Int("part", 0, "Choose part to run")
	flag.Parse()

	if *partPtr == 1 {
		part1(list1, list2)
	} else if *partPtr == 2 {
		part2(list1, list2)
	} else {
		part1(list1, list2)
		part2(list1, list2)
	}

}

func part1(list1, list2 []string) {
	var dist int
	for i := 0; i < len(list1); i++ {
		num1, err := strconv.Atoi(list1[i])
		check(err)
		num2, err := strconv.Atoi(list2[i])
		check(err)
		dist += distInt(num1, num2)
	}
	fmt.Printf("PART1 Total distance is: %d\n", dist)
}

func part2(list1, list2 []string) {
	m := make(map[string]int)
	for i := 0; i < len(list1); i++ {
		m[list1[i]] = 0
	}
	var similarity int
	for _, val := range list2 {
		if _, ok := m[val]; ok {
			m[val] += 1
		}
	}

	for key := range m {
		num, err := strconv.Atoi(key)
		check(err)
		similarity += num * m[key]
	}

	fmt.Printf("PART2 Similarity score is: %d\n", similarity)
}
