package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	c, _ := os.ReadFile("./input.txt")
	cc := strings.Split(strings.TrimSpace(string(c)), "\n")

	part1(cc)
	part2(cc)
}

func getLineInts(line string) (int, int) {
	parts := strings.Split(line, "   ")
	left, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		panic(err)
	}

	return left, right
}

func part1(cc []string) {
	sorted := make([][]int, 2)
	for _, line := range cc {
		left, right := getLineInts(line)

		// TODO: insert efficiently instead of sorting post
		sorted[0] = append(sorted[0], left)
		sorted[1] = append(sorted[1], right)
	}

	sort.Ints(sorted[0])
	sort.Ints(sorted[1])

	totalDist := 0

	for i := range sorted[0] {
		totalDist += abs(sorted[0][i] - sorted[1][i])
	}

	fmt.Println(totalDist)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func part2(cc []string) {
	left := []int{}
	right := make(map[int]int)

	for _, line := range cc {
		l, r := getLineInts(line)
		left = append(left, l)
		right[r] += 1
	}

	total := 0
	for _, v := range left {
		m := right[v]
		total += m * v
	}

	fmt.Println(total)
}
