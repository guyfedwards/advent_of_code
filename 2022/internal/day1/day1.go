package day1

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1() int {
	c, _ := os.ReadFile("./input/day1")
	cc := strings.Split(string(c), "\n")

	maxTotal := 0
	runningTotal := 0

	for _, v := range cc {
		val, _ := strconv.Atoi(v)

		runningTotal = val + runningTotal

		if v == "" {
			if runningTotal >= maxTotal {
				maxTotal = runningTotal
			}
			runningTotal = 0
		}
	}

	return maxTotal
}

func Part2() int {
	c, _ := os.ReadFile("./input/day1")
	cc := strings.Split(string(c), "\n")

	maxTotal1 := 0
	maxTotal2 := 0
	maxTotal3 := 0
	allTotals := []int{}
	runningTotal := 0

	for _, v := range cc {
		val, _ := strconv.Atoi(v)

		runningTotal = val + runningTotal

		if v == "" {
			if runningTotal >= maxTotal1 {
				maxTotal1 = runningTotal
			} else if runningTotal >= maxTotal2 {
				maxTotal2 = runningTotal
			} else if runningTotal >= maxTotal3 {
				maxTotal3 = runningTotal
			}

			allTotals = append(allTotals, runningTotal)

			runningTotal = 0
		}
	}

	sort.Ints(allTotals)
	l := len(allTotals)

	return allTotals[l-1] + allTotals[l-3] + allTotals[l-2]
}
