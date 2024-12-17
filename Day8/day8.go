package main

import (
	"aoc-2024/utils"
	"fmt"
	"os"
	"strings"
)

type Coords struct {
	x, y int
}

func main() {
	fmt.Println("Day 8")
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Split(input, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	// create map
	antennas := make(map[string][]Coords, 0)
	for y, line := range lines {
		for x, val := range line {
			if val != '.' {
				list := antennas[string(val)]
				antennas[string(val)] = append(list, Coords{x, y})
			}
		}
	}

	// find antinodes
	part1(antennas, len(lines), len(lines[0]))
}

func part1(input map[string][]Coords, lenX, lenY int) {
	var sum int
	set := make(map[Coords]bool, 0)
	for _, val := range input {
		for i := 0; i < len(val)-1; i++ {
			for j := i + 1; j < len(val); j++ {
				distX := utils.AbsDiff(val[i].x, val[j].x)
				distY := utils.AbsDiff(val[i].y, val[j].y)
				var antinode1, antinode2 Coords
				if val[i].x <= val[j].x {
					antinode1 = Coords{val[i].x - distX, 0}
					antinode2 = Coords{val[j].x + distX, 0}
				} else {
					antinode1 = Coords{val[i].x + distX, 0}
					antinode2 = Coords{val[j].x - distX, 0}
				}

				if val[i].y <= val[j].y {
					antinode1.y = val[i].y - distY
					antinode2.y = val[j].y + distY
				} else {
					antinode1.y = val[i].y + distY
					antinode2.y = val[j].y - distY
				}

				if _, ok := set[antinode1]; antinode1.x < lenX && antinode1.x >= 0 && antinode1.y < lenY && antinode1.y >= 0 && !ok {
					set[antinode1] = true
					sum++
				}
				if _, ok := set[antinode2]; antinode2.x < lenX && antinode2.x >= 0 && antinode2.y < lenY && antinode2.y >= 0 && !ok {
					set[antinode2] = true
					sum++
				}
			}
		}
	}
	fmt.Printf("The distinct antinodes are %d\n", sum)
}
