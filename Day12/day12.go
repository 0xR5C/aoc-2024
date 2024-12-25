package main

import (
	"aoc-2024/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type region struct {
	area, perimeter, edges int
	val                    string
	plots                  map[point]bool
}

func main() {
	fmt.Println("Day 12")
	dat, err := os.ReadFile("input.txt")
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Fields(input)
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	// Create farm from input
	farm := make([][]string, len(lines))
	for i := range lines {
		farm[i] = make([]string, len(lines[i]))
		for j, val := range lines[i] {
			farm[i][j] = string(val)
		}
	}

	partPtr := flag.Int("part", 0, "Choose part to run")
	flag.Parse()

	if *partPtr == 1 {
		part1(farm)
	} else if *partPtr == 2 {
		part2(farm)
	} else {
		part1(farm)
		part2(farm)
	}
}

func part1(farm [][]string) {
	visited := make([][]bool, len(farm))
	for i := range farm {
		visited[i] = make([]bool, len(farm[i]))
	}

	arr := make([]region, 0)

	for i := range farm {
		for j := range farm[i] {
			if visited[i][j] {
				continue
			}
			reg := region{area: 1, val: farm[i][j], plots: make(map[point]bool)}
			reg.plots[point{j, i}] = true
			reg = bfs(farm, visited, j, i, reg)
			arr = append(arr, reg)
		}
	}

	var sum int
	for i := range arr {
		sum += arr[i].area * arr[i].perimeter
	}
	fmt.Printf("PART1 The total price is %d\n", sum)
}

func part2(farm [][]string) {
	visited := make([][]bool, len(farm))
	for i := range farm {
		visited[i] = make([]bool, len(farm[i]))
	}

	arr := make([]region, 0)

	for i := range farm {
		for j := range farm[i] {
			if visited[i][j] {
				continue
			}
			reg := region{area: 1, val: farm[i][j], plots: make(map[point]bool)}
			reg.plots[point{j, i}] = true
			reg = bfs(farm, visited, j, i, reg)
			arr = append(arr, reg)
		}
	}

	var sum int
	for i := range arr {
		sum += arr[i].area * arr[i].edges
	}
	fmt.Printf("PART2 The total price is %d\n", sum)
}

func bfs(farm [][]string, visited [][]bool, x, y int, reg region) region {
	var perimeter int
	visited[y][x] = true

	// In order to find corners for edges, check which of these directions are part of an edge (have a perimeter)
	var ym1bool, yp1bool, xm1bool, xp1bool bool

	xm1 := point{x - 1, y}
	if x-1 >= 0 && !visited[xm1.y][xm1.x] && farm[y][x] == farm[xm1.y][xm1.x] {
		reg.area += 1
		reg.plots[xm1] = true
		reg = bfs(farm, visited, xm1.x, xm1.y, reg)
	} else if x-1 < 0 || farm[y][x] != farm[xm1.y][xm1.x] {
		perimeter++
		xm1bool = true
	}
	xp1 := point{x + 1, y}
	if x+1 < len(farm[0]) && !visited[xp1.y][xp1.x] && farm[y][x] == farm[xp1.y][xp1.x] {
		reg.area += 1
		reg.plots[xp1] = true
		reg = bfs(farm, visited, xp1.x, xp1.y, reg)
	} else if x+1 >= len(farm[0]) || farm[y][x] != farm[xp1.y][xp1.x] {
		perimeter++
		xp1bool = true
	}
	ym1 := point{x, y - 1}
	if y-1 >= 0 && !visited[ym1.y][ym1.x] && farm[y][x] == farm[ym1.y][ym1.x] {
		reg.area += 1
		reg.plots[ym1] = true
		reg = bfs(farm, visited, ym1.x, ym1.y, reg)
	} else if y-1 < 0 || farm[y][x] != farm[ym1.y][ym1.x] {
		perimeter++
		ym1bool = true
	}
	yp1 := point{x, y + 1}
	if y+1 < len(farm) && !visited[yp1.y][yp1.x] && farm[y][x] == farm[yp1.y][yp1.x] {
		reg.area += 1
		reg.plots[yp1] = true
		reg = bfs(farm, visited, yp1.x, yp1.y, reg)
	} else if y+1 >= len(farm) || farm[y][x] != farm[yp1.y][yp1.x] {
		perimeter++
		yp1bool = true
	}
	reg.perimeter += perimeter

	// Find corners, first part is exterior edges(ez case), second part is interior edges
	if (ym1bool && xm1bool) || (y-1 >= 0 && x-1 >= 0 && !ym1bool && !xm1bool && farm[y][x] != farm[y-1][x-1]) {
		reg.edges++
	}
	if (xm1bool && yp1bool) || (y+1 < len(farm) && x-1 >= 0 && !xm1bool && !yp1bool && farm[y][x] != farm[y+1][x-1]) {

		reg.edges++
	}
	if (yp1bool && xp1bool) || (y+1 < len(farm) && x-1 < len(farm[y]) && !yp1bool && !xp1bool && farm[y][x] != farm[y+1][x+1]) {

		reg.edges++
	}
	if (xp1bool && ym1bool) || (y-1 >= 0 && x+1 < len(farm) && !xp1bool && !ym1bool && farm[y][x] != farm[y-1][x+1]) {
		reg.edges++
	}

	return reg
}
