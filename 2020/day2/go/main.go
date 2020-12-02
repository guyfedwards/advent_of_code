package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseEntry(raw string) (min, max int, letter, password string) {
	var err error
	parts := regexp.MustCompile(`[-\s:]`).Split(raw, -1)

	min, err = strconv.Atoi(parts[0])
	check(err)
	max, err = strconv.Atoi(parts[1])
	check(err)
	letter = strings.Trim(parts[2], " ")
	password = strings.Trim(strings.Join(parts[3:], ""), " ")

	return min, max, letter, password
}

func newEntry(min, max int, letter, password string) *entry {
	return &entry{
		min:      min,
		max:      max,
		letter:   letter,
		password: password,
	}
}

type entry struct {
	min      int
	max      int
	letter   string
	password string
}

func getEntries(raw string) (result []*entry) {
	arr := strings.Split(raw, "\n")

	for _, r := range arr {
		if r == "" {
			continue
		}
		min, max, letter, password := parseEntry(r)

		result = append(result, newEntry(min, max, letter, password))
	}

	return result
}

func partOne(in []*entry) (result int) {
	for _, e := range in {
		c := strings.Count(e.password, e.letter)
		if c >= e.min && c <= e.max {
			result++
		}
	}

	return result
}

func partTwo(in []*entry) (result int) {
	for _, e := range in {
		firstChar := string(e.password[e.min-1])
		secondChar := string(e.password[e.max-1])

		if (firstChar == e.letter && secondChar != e.letter) || (firstChar != e.letter && secondChar == e.letter) {
			result++
		}
	}

	return result
}

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	check(err)
	input := getEntries(string(data))

	fmt.Printf("Result1: %v\n", partOne(input))
	fmt.Printf("Result2: %v\n", partTwo(input))
}
