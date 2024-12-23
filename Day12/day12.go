package main

import (
	"aoc-2024/utils"
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type region struct {
	area, perimeter int
	val             string
	plots           map[point]bool
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

	farm := make([][]string, len(lines))
	for i := range lines {
		farm[i] = make([]string, len(lines[i]))
		for j, val := range lines[i] {
			farm[i][j] = string(val)
		}
	}
	part1(farm)
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
			reg := region{area: 1, val: farm[i][j]}
			reg = bfs(farm, visited, j, i, reg)
			arr = append(arr, reg)
		}
	}

	var sum int
	for i := range arr {
		sum += arr[i].area * arr[i].perimeter
	}
	fmt.Printf("The total price is %d\n", sum)
}

func bfs(farm [][]string, visited [][]bool, x, y int, reg region) region {
	var perimeter int
	visited[y][x] = true
	xm1 := point{x - 1, y}
	if x-1 >= 0 && !visited[xm1.y][xm1.x] && farm[y][x] == farm[xm1.y][xm1.x] {
		reg.area += 1
		reg = bfs(farm, visited, xm1.x, xm1.y, reg)
	} else if x-1 < 0 || farm[y][x] != farm[xm1.y][xm1.x] {
		perimeter++
	}
	xp1 := point{x + 1, y}
	if x+1 < len(farm[0]) && !visited[xp1.y][xp1.x] && farm[y][x] == farm[xp1.y][xp1.x] {
		reg.area += 1
		reg = bfs(farm, visited, xp1.x, xp1.y, reg)
	} else if x+1 >= len(farm[0]) || farm[y][x] != farm[xp1.y][xp1.x] {
		perimeter++
	}
	ym1 := point{x, y - 1}
	if y-1 >= 0 && !visited[ym1.y][ym1.x] && farm[y][x] == farm[ym1.y][ym1.x] {
		reg.area += 1
		reg = bfs(farm, visited, ym1.x, ym1.y, reg)
	} else if y-1 < 0 || farm[y][x] != farm[ym1.y][ym1.x] {
		perimeter++
	}
	yp1 := point{x, y + 1}
	if y+1 < len(farm) && !visited[yp1.y][yp1.x] && farm[y][x] == farm[yp1.y][yp1.x] {
		reg.area += 1
		reg = bfs(farm, visited, yp1.x, yp1.y, reg)
	} else if y+1 >= len(farm) || farm[y][x] != farm[yp1.y][yp1.x] {
		perimeter++
	}
	reg.perimeter += perimeter
	return reg
}
