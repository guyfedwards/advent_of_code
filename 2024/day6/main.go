package main

import (
	"fmt"
	"strings"
)

func main() {
	// c, _ := os.ReadFile("./input.txt\")
	// cc := strings.Split(strings.TrimSpace(string(c)), "\n")

	cc := strings.Split(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`, "\n")

	part1(cc)
	part2(cc)
}

type walker struct {
	start   []int
	dir     string
	cur     []int
	m       [][]string
	visited int
}

func (w *walker) load(in []string) {
	foundStart := false
	w.dir = "up"
	for i, v := range in {
		parts := strings.Split(v, "")
		w.m[i] = parts

		if !foundStart {
			for ii, part := range parts {
				if part == "^" {
					w.start = []int{ii, i}
					w.cur = w.start
					foundStart = true
				}
			}
		}
	}
}

var dirs = map[string][]int{
	"up":    []int{0, -1},
	"right": []int{1, 0},
	"down":  []int{0, 1},
	"left":  []int{-1, 0},
}

func (w *walker) rotate90() {
	switch w.dir {
	case "up":
		w.dir = "right"
	case "right":
		w.dir = "down"
	case "down":
		w.dir = "left"
	case "left":
		w.dir = "up"
	}
}

func (w *walker) walk() int {
	steps := 0

	done := false
	// loc := w.start
	for !done {
		fmt.Println("Step: ", steps, w.cur, w.dir)
		xy := dirs[w.dir]
		newX := w.cur[0] + xy[0]
		newY := w.cur[1] + xy[1]
		if newY < 0 || newY > len(w.m) || newX < 0 || newX > len(w.m[newY]) {
			fmt.Println("Done")
			done = true
			break
		}

		inFront := w.m[newY][newX]

		switch inFront {
		case "^":
			fallthrough
		case ".":
			w.cur = []int{newX, newY}
			steps++
		case "#":
			fmt.Println("Rotate")
			w.rotate90()
		}
	}
	// cur := w.start
	// for w.step(cur) {
	// 	steps+jkjk+
	// 	cur := step()
	// }

	return steps
}

// func (w *walker) step() bool {

// }

func part1(cc []string) {
	w := &walker{
		m: make([][]string, len(cc)),
	}
	w.load(cc)
	res := w.walk()
	fmt.Println("Part 1: ", res)
}

func part2(cc []string) {
	fmt.Println()
}
