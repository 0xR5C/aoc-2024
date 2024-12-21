package main

import (
	"aoc-2024/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type trailhead struct {
	x, y   int
	score  map[point]bool
	rating int
}

func main() {
	fmt.Println("Day 10")
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	hikingMap := make([][]int, len(lines))
	for i, line := range lines {
		for _, val := range line {
			strChar := string(val)
			height, err := strconv.Atoi(strChar)
			utils.Check(err)
			hikingMap[i] = append(hikingMap[i], height)
		}
	}

	partPtr := flag.Int("part", 0, "Choose part to run")
	flag.Parse()

	if *partPtr == 1 {
		part1(hikingMap)
	} else if *partPtr == 2 {
		part2(hikingMap)
	} else {
		part1(hikingMap)
		part2(hikingMap)
	}
}

func part1(hiking [][]int) {
	trailList := make([]trailhead, 0)
	for i, line := range hiking {
		for j := range line {
			if hiking[i][j] == 0 {
				scoreMap := make(map[point]bool, 0)
				trailList = append(trailList, trailhead{j, i, scoreMap, 0})
				testTrail(hiking, 0, j-1, i, &trailList[len(trailList)-1])
				testTrail(hiking, 0, j, i-1, &trailList[len(trailList)-1])
				testTrail(hiking, 0, j+1, i, &trailList[len(trailList)-1])
				testTrail(hiking, 0, j, i+1, &trailList[len(trailList)-1])
			}
		}
	}
	var sum int
	for _, val := range trailList {
		sum += len(val.score)
	}
	fmt.Printf("PART1 The total score of the trailheads is %d\n", sum)
}

func testTrail(hiking [][]int, prevVal, nextX, nextY int, head *trailhead) bool {
	if nextY >= len(hiking) || nextY < 0 || nextX >= len(hiking[0]) || nextX < 0 {
		return false
	} else if hiking[nextY][nextX] != prevVal+1 {
		return false
	} else {
		tempPoint := point{nextX, nextY}
		if _, ok := head.score[tempPoint]; hiking[nextY][nextX] == 9 && !ok {
			head.score[tempPoint] = true
			return true
		} else if hiking[nextY][nextX] == 9 {
			return false
		}

		testTrail(hiking, hiking[nextY][nextX], nextX-1, nextY, head)
		testTrail(hiking, hiking[nextY][nextX], nextX, nextY-1, head)
		testTrail(hiking, hiking[nextY][nextX], nextX+1, nextY, head)
		testTrail(hiking, hiking[nextY][nextX], nextX, nextY+1, head)
		return false
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// 											PART 2						   						  	//
//////////////////////////////////////////////////////////////////////////////////////////////////////

func part2(hiking [][]int) {
	trailList := make([]trailhead, 0)
	for i, line := range hiking {
		for j := range line {
			if hiking[i][j] == 0 {
				scoreMap := make(map[point]bool, 0)
				trailList = append(trailList, trailhead{j, i, scoreMap, 0})
				testTrailPt2(hiking, 0, j-1, i, &trailList[len(trailList)-1])
				testTrailPt2(hiking, 0, j, i-1, &trailList[len(trailList)-1])
				testTrailPt2(hiking, 0, j+1, i, &trailList[len(trailList)-1])
				testTrailPt2(hiking, 0, j, i+1, &trailList[len(trailList)-1])
			}
		}
	}
	var sum int
	for _, val := range trailList {
		sum += val.rating
	}
	fmt.Printf("PART2 The total rating of the trailheads is %d\n", sum)
}

func testTrailPt2(hiking [][]int, prevVal, nextX, nextY int, head *trailhead) bool {
	if nextY >= len(hiking) || nextY < 0 || nextX >= len(hiking[0]) || nextX < 0 {
		return false
	} else if hiking[nextY][nextX] != prevVal+1 {
		return false
	} else {
		if hiking[nextY][nextX] == 9 {
			head.rating += 1
			return true
		}

		testTrailPt2(hiking, hiking[nextY][nextX], nextX-1, nextY, head)
		testTrailPt2(hiking, hiking[nextY][nextX], nextX, nextY-1, head)
		testTrailPt2(hiking, hiking[nextY][nextX], nextX+1, nextY, head)
		testTrailPt2(hiking, hiking[nextY][nextX], nextX, nextY+1, head)
		return false
	}
}
