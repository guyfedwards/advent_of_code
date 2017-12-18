package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getLargestSmallest(row string) (int, int) {
	var (
		largest  int
		smallest int
	)

	arr := strings.Split(row, "\t")

	smallest, _ = strconv.Atoi(arr[0])

	for i := 0; i < len(arr); i++ {
		num, _ := strconv.Atoi(arr[i])

		if num > largest {
			largest = num
		}

		if num < smallest {
			smallest = num
		}
	}

	return largest, smallest
}

func getDiff(l, s int) int {
	return l - s
}

func sum(a []int) int {
	var total int

	for _, v := range a {
		total += v
	}

	return total
}

func main() {
	dat, err := ioutil.ReadFile("./input.txt")

	check(err)

	var (
		diffs []int
	)

	rows := strings.Split(string(dat), "\n")

	for _, row := range rows[:len(rows)-1] {
		l, s := getLargestSmallest(row)

		diffs = append(diffs, getDiff(l, s))
	}

	fmt.Println(sum(diffs))
}
