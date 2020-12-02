package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getIntSlice(s []byte) (result []int) {
	for _, line := range strings.Split(string(s), "\n") {
		if line != "" {
			i, err := strconv.Atoi(line)
			check(err)

			result = append(result, i)
		}
	}

	return result
}

func includes(all []int, value int) bool {
	for _, a := range all {
		if a == value {
			return true
		}
	}

	return false
}

func partOne(all []int) (result int) {
	for _, n := range all {
		remainder := 2020 - n

		if includes(all, remainder) {
			result = n * remainder
			break
		}
	}

	return result
}

func partTwo(all []int) (result int) {
outer:
	for _, n := range all {
		for _, nn := range all {
			for _, nnn := range all {
				if n+nn+nnn == 2020 {
					result = n * nn * nnn
					break outer
				}
			}
		}
	}
	return result
}

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	check(err)
	input := getIntSlice(data)

	fmt.Printf("Result1: %v\n", partOne(input))
	fmt.Printf("Result2: %v\n", partTwo(input))
}
