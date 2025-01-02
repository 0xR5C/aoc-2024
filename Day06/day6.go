package main

import (
	"aoc-2024/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(("Day 6"))
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	mapGuard := make([][]string, len(lines))
	var guardX, guardY int
	for i, val := range lines {
		for j, val := range val {
			if val == '^' || val == '>' || val == '<' || val == 'v' {
				guardY = i
				guardX = j
			}
			mapGuard[i] = append(mapGuard[i], string(val))
		}
	}

	partPtr := flag.Int("part", 0, "Choose part to run")
	flag.Parse()

	if *partPtr == 1 {
		part1(mapGuard, guardX, guardY)
	} else if *partPtr == 2 {
		part2(mapGuard, guardX, guardY)
	} else {
		part1(mapGuard, guardX, guardY)
		part2(mapGuard, guardX, guardY)
	}
}

func part1(mapGuard [][]string, guardX, guardY int) {
	var distinct int

	mapPath := guardRoute(mapGuard, guardX, guardY)

	for i := range mapPath {
		for j := range mapPath[i] {
			if mapPath[i][j] == "X" {
				distinct++
			}
		}
	}

	fmt.Printf("PART1 Distinct steps of the guard are %d\n", distinct)
}

func part2(mapGuard [][]string, guardX, guardY int) {

	mapPath := guardRoute(mapGuard, guardX, guardY)
	var sum int
	for i := range mapPath {
		for j := range mapPath[i] {
			if mapPath[i][j] == "X" {
				if ok := checkNewObstruction(mapGuard, guardX, guardY, j, i); ok {
					sum++
				}

			}
		}
	}

	fmt.Printf("PART2 The different positions for the obstruction are %d\n", sum)
}

func guardRoute(mapGuard [][]string, guardX, guardY int) [][]string {
	// Create copy
	mapPath := [][]string{}
	for i := range mapGuard {
		mapPath = append(mapPath, []string{})
		mapPath[i] = append(mapPath[i], mapGuard[i]...)
	}

	// Find the path
	for guardX >= 0 && guardY >= 0 && guardY < len(mapGuard) && guardX < len(mapPath[guardY]) {
		guard := mapPath[guardY][guardX]
		switch guard {
		case "^":
			mapPath[guardY][guardX] = "X"
			guardY--
			if guardX < 0 || guardY < 0 || guardY >= len(mapPath) || guardX >= len(mapPath[guardY]) {
				break
			}
			if mapPath[guardY][guardX] == "#" {
				guardY++
				mapPath[guardY][guardX] = ">"
			} else {
				mapPath[guardY][guardX] = "^"
			}
		case "v":
			mapPath[guardY][guardX] = "X"
			guardY++
			if guardX < 0 || guardY < 0 || guardY >= len(mapPath) || guardX >= len(mapPath[guardY]) {
				break
			}
			if mapPath[guardY][guardX] == "#" {
				guardY--
				mapPath[guardY][guardX] = "<"
			} else {
				mapPath[guardY][guardX] = "v"
			}
		case ">":
			mapPath[guardY][guardX] = "X"
			guardX++
			if guardX < 0 || guardY < 0 || guardY >= len(mapPath) || guardX >= len(mapPath[guardY]) {
				break
			}
			if mapPath[guardY][guardX] == "#" {
				guardX--
				mapPath[guardY][guardX] = "v"
			} else {
				mapPath[guardY][guardX] = ">"
			}
		case "<":
			mapPath[guardY][guardX] = "X"
			guardX--
			if guardX < 0 || guardY < 0 || guardY >= len(mapPath) || guardX >= len(mapPath[guardY]) {
				break
			}
			if mapPath[guardY][guardX] == "#" {
				guardX++
				mapPath[guardY][guardX] = "^"
			} else {
				mapPath[guardY][guardX] = "<"
			}
		}
	}
	return mapPath
}

func checkNewObstruction(mapGuard [][]string, guardX, guardY, obsX, obsY int) bool {

	if guardX == obsX && guardY == obsY {
		return false
	}

	// Create copy
	mapPath := [][]string{}
	for i := range mapGuard {
		mapPath = append(mapPath, []string{})
		mapPath[i] = append(mapPath[i], mapGuard[i]...)
	}

	// Set new Obstruction
	mapPath[obsY][obsX] = "Ο"

	var steps int

	for ; guardX >= 0 && guardY >= 0 && guardY < len(mapPath) && guardX < len(mapPath[guardY]); steps++ {
		// Loop detection
		if steps >= 2*(len(mapPath)-1)*(len(mapPath[guardY])-1) {
			return true
		}

		// Traverse path
		guardPos := mapPath[guardY][guardX]
		switch guardPos {
		case "^":
			guardY--
			if guardX < 0 || guardY < 0 || guardY >= len(mapPath) || guardX >= len(mapPath[guardY]) {
				break
			}
			if mapPath[guardY][guardX] == "#" || mapPath[guardY][guardX] == "Ο" {
				guardY++
				mapPath[guardY][guardX] = ">"
			} else {
				mapPath[guardY+1][guardX] = "|"
				mapPath[guardY][guardX] = "^"
			}
		case "v":
			guardY++
			if guardX < 0 || guardY < 0 || guardY >= len(mapPath) || guardX >= len(mapPath[guardY]) {
				break
			}
			if mapPath[guardY][guardX] == "#" || mapPath[guardY][guardX] == "Ο" {
				guardY--
				mapPath[guardY][guardX] = "<"
			} else {
				mapPath[guardY-1][guardX] = "|"
				mapPath[guardY][guardX] = "v"
			}
		case ">":
			guardX++
			if guardX < 0 || guardY < 0 || guardY >= len(mapPath) || guardX >= len(mapPath[guardY]) {
				break
			}
			if mapPath[guardY][guardX] == "#" || mapPath[guardY][guardX] == "Ο" {
				guardX--
				mapPath[guardY][guardX] = "v"
			} else {
				mapPath[guardY][guardX-1] = "-"
				mapPath[guardY][guardX] = ">"
			}
		case "<":
			guardX--
			if guardX < 0 || guardY < 0 || guardY >= len(mapPath) || guardX >= len(mapPath[guardY]) {
				break
			}
			if mapPath[guardY][guardX] == "#" || mapPath[guardY][guardX] == "Ο" {
				guardX++
				mapPath[guardY][guardX] = "^"
			} else {
				mapPath[guardY][guardX+1] = "-"
				mapPath[guardY][guardX] = "<"
			}
		}

	}

	return false
}
