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

func checkSlope(all []string, h, v int) (result int) {
	var horizontal int

	for i := 0; i < len(all); i += v {
		if i+v >= len(all) {
			break
		}
		row := strings.Split(all[i+v], "")
		if len(row) == 0 {
			break
		}

		horizontal = horizontal + h
		position := row[horizontal%len(row)]

		if position == "#" {
			result++
		}
	}
	return result
}

func partOne(rows []string) (result int) {
	return checkSlope(rows, 3, 1)
}

func partTwo(rows []string) (result int) {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	var trees []int

	for _, slope := range slopes {
		trees = append(trees, checkSlope(rows, slope[0], slope[1]))
	}

	result = trees[0]
	for i := 1; i < len(trees); i++ {
		result = result * trees[i]
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
