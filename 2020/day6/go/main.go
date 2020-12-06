package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func contains(arr []string, l string) bool {
	for _, ll := range arr {
		if ll == l {
			return true
		}
	}

	return false
}

func countGroup(group string) (result int) {
	seen := []string{}

	for _, l := range group {
		if !contains(seen, string(l)) {
			result++
			seen = append(seen, string(l))
		}
	}

	return result
}

func sum(arr []int) (result int) {
	for _, v := range arr {
		result += v
	}

	return result
}

func countAllGroup(group string) (result int) {
	answers := map[string]int{}
	numInGroup := len(strings.Split(strings.Trim(group, "\n"), "\n"))
	group = strings.Join(strings.Split(group, "\n"), "")

	for _, l := range group {
		answers[string(l)] = answers[string(l)] + 1
	}

	for _, v := range answers {
		if v == numInGroup {
			result++
		}
	}

	return result
}

func partOne(answers []string) int {
	groups := []int{}

	for _, p := range answers {
		groups = append(groups, countGroup(strings.Join(strings.Split(p, "\n"), "")))
	}

	return sum(groups)
}

func partTwo(answers []string) int {
	groups := []int{}

	for _, p := range answers {
		groups = append(groups, countAllGroup(p))
	}

	return sum(groups)
}

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	check(err)
	input := strings.Split(string(data), "\n\n")

	fmt.Printf("Result1: %v\n", partOne(input))
	fmt.Printf("Result2: %v\n", partTwo(input))
}
