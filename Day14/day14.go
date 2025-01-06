package main

import (
	"aoc-2024/utils"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type robot struct {
	x, y, vx, vy int
}

const width = 101
const height = 103
const middleX = width / 2
const middleY = height / 2

func main() {
	fmt.Println("Day14")
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
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
	var robotList []robot
	for _, val := range lines {
		robotList = append(robotList, extractVals(val))
	}

	count, _ := calcPosition(robotList, 100)

	fmt.Printf("PART1 The safety factor after 100 seconds is %d\n", count)
}

func part2(lines []string) {
	var robotList []robot
	for _, val := range lines {
		robotList = append(robotList, extractVals(val))
	}

	var count, sec int
	overlap := true

	for i := 1; overlap; i++ {
		count, overlap = calcPosition(robotList, 1)
		if !overlap {
			sec = i
		}
	}

	fmt.Printf("PART2 The safety factor after %d seconds is %d\n", sec, count)
}

func extractVals(line string) robot {
	var temp robot
	re := regexp.MustCompile(`-?\d+`)

	matches := re.FindAllString(line, -1)

	if len(matches) == 4 {
		x, err := strconv.Atoi(matches[0])
		utils.Check(err)
		temp.x = x
		y, err := strconv.Atoi(matches[1])
		utils.Check(err)
		temp.y = y
		vx, err := strconv.Atoi(matches[2])
		utils.Check(err)
		temp.vx = vx
		vy, err := strconv.Atoi(matches[3])
		utils.Check(err)
		temp.vy = vy
		return temp
	}
	return temp
}

func calcPosition(robotList []robot, pos int) (int, bool) {
	overlap := false
	for j := range robotList {
		resX := (robotList[j].x + robotList[j].vx*pos) % width
		resY := (robotList[j].y + robotList[j].vy*pos) % height
		if resX < 0 {
			resX = width + resX
		}
		if resY < 0 {
			resY = height + resY
		}
		robotList[j].x = resX
		robotList[j].y = resY
	}

	tiles := make([][]int, 103)
	for i := range tiles {
		tiles[i] = make([]int, 101)
	}

	for _, val := range robotList {
		tiles[val.y][val.x]++
		if tiles[val.y][val.x] > 1 {
			overlap = true
		}
	}

	if !overlap {
		fmt.Println("Merry Christmas!")
		for i := range tiles {
			for j := range tiles[i] {
				if tiles[i][j] == 0 {
					fmt.Print(".")
				} else {
					fmt.Print(tiles[i][j])
				}
			}
			fmt.Println()
		}
	}

	count := 1
	count *= calcQuadrant(tiles, 0, middleX, 0, middleY)
	count *= calcQuadrant(tiles, 0, middleX, middleY+1, height)
	count *= calcQuadrant(tiles, middleX+1, width, 0, middleY)
	count *= calcQuadrant(tiles, middleX+1, width, middleY+1, height)

	return count, overlap
}

func calcQuadrant(tiles [][]int, x0, x1, y0, y1 int) int {
	var sum int
	for i := y0; i < y1; i++ {
		for j := x0; j < x1; j++ {
			sum += tiles[i][j]
		}
	}
	return sum
}
