package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getRow(s string, lower, upper int) (row int) {
	letter := s[0:1]
	half := lower + ((upper - lower) / 2)

	if len(s) == 1 {
		if letter == "F" {
			return lower + 1
		}
		if letter == "B" {
			return upper + 1
		}
	}

	if letter == "F" {
		return getRow(s[1:], lower, half)
	}

	return getRow(s[1:], half, upper)
}

func binarySearch(s string, lower, upper int, lowerChar, upperChar string) (row int) {
	letter := s[0:1]
	half := lower + ((upper - lower) / 2)

	if len(s) == 1 {
		if letter == lowerChar {
			return lower
		}
		if letter == upperChar {
			return upper
		}
	}

	if letter == lowerChar {
		return binarySearch(s[1:], lower, half, lowerChar, upperChar)
	}

	return binarySearch(s[1:], half+1, upper, lowerChar, upperChar)
}

func getId(p string) int {
	row := binarySearch(p[0:7], 0, 127, "F", "B")

	column := binarySearch(p[7:], 0, 7, "L", "R")
	id := row*8 + column

	return id

}

func partOne(passes []string) (result int) {
	for _, p := range passes {
		if p == "" {
			continue
		}
		id := getId(p)

		if id > result {
			result = id
		}
	}

	return result
}

func partTwo(passes []string) (result int) {
	var ids []int
	for _, p := range passes {
		if p == "" {
			continue
		}

		ids = append(ids, getId(p))
	}

	sort.Ints(ids)

	for i := range ids {
		if i+1 >= len(ids) {
			break
		}
		if ids[i+1]-ids[i] != 1 {
			result = ids[i] + 1
			break
		}
	}
	return result
}

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	check(err)
	input := strings.Split(string(data), "\n")

	fmt.Printf("Result1: %v\n", partOne(input))
	fmt.Printf("Result2: %v\n", partTwo(input))
}
