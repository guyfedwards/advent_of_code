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
	half := len(arg) / 2

	var nums []int

	for i := 0; i < len(arg); i++ {
		var y int

		x, _ := strconv.Atoi(string(arg[i]))

		if i < half {
			y, _ = strconv.Atoi(string(arg[i+half]))
		} else {
			z := i + half - len(arg)
			y, _ = strconv.Atoi(string(arg[z]))
		}

		if x == y {
			nums = append(nums, x)
		}
	}

	fmt.Println(sum(nums))
}
