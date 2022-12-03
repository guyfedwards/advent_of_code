package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func e(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}

func main() {
	file, err := os.Open("./input.txt")
	e(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	var nums []int

	index := 0

	for scanner.Scan() {
		v := scanner.Text()
		i, err := strconv.Atoi(v)
		e(err)

		nums = append(nums, i)

		// return if first few numbers
		if index < 3 {
			index++
			continue
		}

		windowA := sum([]int{nums[index-1], nums[index-2], nums[index-3]})
		windowB := sum([]int{nums[index], nums[index-1], nums[index-2]})

		if windowA < windowB {
			total++
		}

		index++
	}

	fmt.Printf("Total: %d", total)
}
