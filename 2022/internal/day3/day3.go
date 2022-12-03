package day3

import (
	"os"
	"strings"
)

func Part1() int {
	c, _ := os.ReadFile("./input/day3")
	cc := strings.Split(string(c), "\n")

	allLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters := strings.Split(allLetters, "")

	total := 0
	for _, v := range cc {
		one := v[:len(v)/2]
		two := v[len(v)/2:]

		inBoth := make(map[string]bool)

		for _, vv := range one {
			for _, vvv := range two {
				if vvv == vv {
					inBoth[string(vv)] = true
				}
			}
		}

		for k, _ := range inBoth {
			for i, v := range letters {
				if v == k {
					total += i + 1
				}
			}
		}

	}

	return total
}

func findIndex(array []string, letter string) int {
	for i, v := range array {
		if string(v) == letter {
			return i
		}
	}
	return -1
}

func unique(array string) []string {
	letters := make(map[string]bool)
	for _, v := range array {
		letters[string(v)] = true
	}

	ret := []string{}
	for k, _ := range letters {
		ret = append(ret, k)
	}
	return ret
}

func Part2() int {
	c, _ := os.ReadFile("./input/day3")
	cc := strings.Split(string(c), "\n")
	totalScore := 0
	allLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters := strings.Split(allLetters, "")

	groupCounter := 1
	groupLetters := make(map[string]int)

	for _, v := range cc {

		uniqV := unique(v)
		for _, vv := range uniqV {
			groupLetters[string(vv)] = groupLetters[string(vv)] + 1
		}

		if groupCounter%3 == 0 && groupCounter != 0 {
			for k, v := range groupLetters {
				if v == 3 {
					ind := findIndex(letters, k)
					if ind != -1 {
						totalScore = totalScore + ind + 1
					}
				}
			}
			groupLetters = make(map[string]int)
		}

		groupCounter++
	}

	return totalScore
}
