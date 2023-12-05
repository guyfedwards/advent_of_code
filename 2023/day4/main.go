package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	c, _ := os.ReadFile("./input.txt")
	cc := strings.Split(strings.TrimSpace(string(c)), "\n")

	part1(cc)
	part2(cc)
}

var lineRe = regexp.MustCompile(`Card\s+(\d+):\s+(\d.*)\|\s+(\d.*)$`)
var digitRe = regexp.MustCompile(`\d+`)

func part1(cc []string) {
	total := 0
	for _, line := range cc {
		matches := lineRe.FindAllStringSubmatch(line, -1)
		winningNums := digitRe.FindAllString(matches[0][2], -1)
		gotNums := digitRe.FindAllString(matches[0][3], -1)

		num := numContains(winningNums, gotNums)
		total += double(num)
	}

	fmt.Println(total)
}

func double(max int) int {
	if max == 0 {
		return 0
	}

	total := 1
	for i := 1; i < max; i++ {
		total = total * 2
	}

	return total
}

func numContains(arr1, arr2 []string) int {
	var total int

	for _, v := range arr1 {
		if contains(arr2, v) {
			total++
		}
	}

	return total
}

func contains(arr []string, v string) bool {
	for _, s := range arr {
		if s == v {
			return true
		}
	}

	return false
}

func part2(cc []string) {
	total := 0
	cardMultiples := make(map[int]int)

	for _, line := range cc {
		matches := lineRe.FindAllStringSubmatch(line, -1)
		id, _ := strconv.Atoi(matches[0][1])
		winningNums := digitRe.FindAllString(matches[0][2], -1)
		gotNums := digitRe.FindAllString(matches[0][3], -1)

		num := numContains(winningNums, gotNums)

		copies := cardMultiples[id] + 1
		total += copies
		for i := 0; i < copies; i++ {
			for ii := 0; ii <= num; ii++ {
				cardMultiples[id+ii]++
			}
		}
	}

	fmt.Println(total)
}
