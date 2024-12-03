package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	c, _ := os.ReadFile("./input.txt")
	cc := strings.Split(strings.TrimSpace(string(c)), "\n")

	part1(cc)
	part2(cc)
}

func part1(cc []string) {
	totalSafe := 0

	for _, report := range cc {
		levels := makeInts(strings.Split(report, " "))
		dir := ""
		safe := true

		for i, l := range levels {
			if i == len(levels)-1 {
				break
			}

			if outofbounds(levels[i+1], l) {
				safe = false
				break
			}

			if levels[i+1] < l {
				if dir == "" {
					dir = "desc"
				} else if dir == "asc" {
					safe = false
					break
				}
			}

			if levels[i+1] > l {
				if dir == "" {
					dir = "asc"
				} else if dir == "desc" {
					safe = false
					break
				}
			}
		}

		if safe {
			totalSafe++
		}
	}

	fmt.Println("Part 1:", totalSafe)
}

func part2(cc []string) {
	totalSafe := 0

	for _, report := range cc {
		levels := makeInts(strings.Split(report, " "))

		safe := true
		for i := range levels {
			safe = test(sansIdx(levels, i))
			if safe {
				break
			}
		}

		if safe {
			totalSafe++
		}
	}

	fmt.Println("Part 2: ", totalSafe)
}

func sansIdx(levels []int, idx int) []int {
	ret := []int{}
	ret = append(ret, levels[:idx]...)
	ret = append(ret, levels[idx+1:]...)
	return ret
}

func test(levels []int) bool {
	dir := ""
	for i := 0; i < len(levels); i++ {
		if i == len(levels)-1 {
			break
		}

		l := levels[i]
		if outofbounds(levels[i+1], l) {
			return false
		}

		if levels[i+1] < l {
			if dir == "" {
				dir = "desc"
			} else if dir == "asc" {
				return false
			}
		}

		if levels[i+1] > l {
			if dir == "" {
				dir = "asc"
			} else if dir == "desc" {
				return false
			}
		}
	}

	return true
}

func makeInts(in []string) (out []int) {
	for _, v := range in {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("makeints strconv")
		}
		out = append(out, i)
	}

	return out
}

// returns if absolute diff is greater than 1, less than 3
func outofbounds(a, b int) bool {
	diff := abs(a - b)
	if diff < 1 || diff > 3 {
		return true
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
