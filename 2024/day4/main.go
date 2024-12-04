package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	c, _ := os.ReadFile("./input.txt")
	cc := strings.Split(strings.TrimSpace(string(c)), "\n")

	part1(cc)
	part2(cc)
}

func part1(cc []string) {
	total := 0
	for y := range cc {
		line := cc[y]
		for x := 0; x < len(line); x++ {
			char := line[x]
			if char == 'X' {
				matches := check(x, y, cc)
				total += matches
			}
		}
	}

	fmt.Println("Part 1: ", total)
}

var dirs = [][]int{
	{0, -1},  // N
	{1, -1},  // NE
	{1, 0},   // E
	{1, 1},   // SE
	{0, 1},   // S
	{-1, 1},  // SW
	{-1, 0},  //W
	{-1, -1}, //NW
}

func check(x, y int, lines []string) int {
	total := 0
	for _, v := range dirs {
		if checkDir(x, y, lines, v[0], v[1]) {
			total++
		}
	}

	return total
}

func checkDir(x, y int, lines []string, xInc, yInc int) bool {
	if xInc < 0 && x < 3 {
		return false
	}

	if xInc > 0 && x > len(lines[y])-4 {
		return false
	}

	if yInc < 0 && y < 3 {
		return false
	}

	if yInc > 0 && y > len(lines[y])-4 {
		return false
	}

	s := ""
	for i := 1; i < 4; i++ {
		s += string(lines[y+yInc*i][x+xInc*i])
	}

	// matched the X to get here so just check the rest
	rs := s == "MAS"
	return rs
}

var masDirs = [][]int{
	{1, -1},  // NE
	{1, 1},   // SE
	{-1, 1},  // SW
	{-1, -1}, // NW
}

func checkMas(x, y int, lines []string, xInc, yInc int) bool {
	xInc = xInc * -1
	yInc = yInc * -1

	if x < 0 || y < 0 || y >= len(lines) || x >= len(lines[y]) {
		return false
	}

	if xInc < 0 && x < 2 {
		return false
	}

	if xInc > 0 && x > len(lines[y])-3 {
		return false
	}

	if yInc < 0 && y < 2 {
		return false
	}

	if yInc > 0 && y > len(lines)-3 {
		return false
	}

	s := string(lines[y][x])
	for i := 1; i < 3; i++ {
		s += string(lines[y+yInc*i][x+xInc*i])
	}

	// matched the X to get here so just check the rest
	rs := s == "MAS"
	return rs
}

func part2(cc []string) {
	total := 0
	for y := range cc {
		line := cc[y]
		for x := 0; x < len(line); x++ {
			char := line[x]
			if char == 'A' {
				matches := 0
				for _, v := range masDirs {
					match := checkMas(x+v[0], y+v[1], cc, v[0], v[1])
					if match {
						matches++
					}
				}

				if matches == 2 {
					total++
				}
			}
		}
	}

	fmt.Println("Part 2: ", total)
}
