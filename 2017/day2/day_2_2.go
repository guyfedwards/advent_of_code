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

func getResult(row string) int {
	var res int
	arr := strings.Split(row, "\t")
	ints := makeInts(arr)

	for i, n := range ints {
		rest := getRest(i, ints)
		for _, num := range rest {
			if n%num == 0 {
				res = n / num
				break
			}
		}
		if res != 0 {
			break
		}
	}

	return res
}

func getRest(i int, arr []int) []int {
	a := make([]int, len(arr))
	copy(a, arr)
	before := a[:i]
	after := a[i+1:]

	return append(before, after...)
}

func makeInts(arr []string) []int {
	result := make([]int, len(arr))

	for i, v := range arr {
		result[i], _ = strconv.Atoi(v)
	}

	return result
}

func mod(a, b int) int {
	return a % b
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
		totals []int
	)

	rows := strings.Split(string(dat), "\n")

	for _, row := range rows[:len(rows)-1] {
		nums := getResult(row)

		totals = append(totals, nums)
	}

	fmt.Println(sum(totals))
}
