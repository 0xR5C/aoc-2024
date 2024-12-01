package main

import (
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

	fmt.Printf("Similarity index is: %d\n", similarity)
}
