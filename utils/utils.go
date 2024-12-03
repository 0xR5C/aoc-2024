package utils

import (
	"strconv"
)

func Works() int {
	return 0
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ArrayAtoi(arr []string) []int {
	numArr := make([]int, len(arr))
	for i := range arr {
		val, ok := strconv.Atoi(arr[i])
		Check(ok)
		numArr[i] = val
	}
	return numArr
}
