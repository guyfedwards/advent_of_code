package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(arr []int) int {
	var total int

	for i := 0; i < len(arr); i++ {
		total = arr[i] + total
	}

	return total
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No arg passed")
		os.Exit(1)
	}
	arg := os.Args[1]

	var nums []int

	for i := 0; i < len(arg); i++ {
		var y int

		x, _ := strconv.Atoi(string(arg[i]))

		if i == len(arg)-1 {
			y, _ = strconv.Atoi(string(arg[0]))
		} else {
			y, _ = strconv.Atoi(string(arg[i+1]))
		}

		if x == y {
			nums = append(nums, x)
		}
	}

	fmt.Println(sum(nums))
}
