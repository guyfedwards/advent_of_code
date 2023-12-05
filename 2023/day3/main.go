package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	c, _ := os.ReadFile("./input.txt")
	cc := strings.Split(strings.TrimSpace(string(c)), "\n")

	// part1(cc)
	part2(cc)
}

var digitRe = regexp.MustCompile(`\d+`)
var symRe = regexp.MustCompile(`[^\d.]`)
var astRe = regexp.MustCompile(`\*`)

func clampMin(a int) int {
	if a <= 0 {
		return 0
	}
	return a
}

func clampMax(a, max int) int {
	if a >= max {
		return max
	}
	return a
}

func part1(cc []string) {
	var result int
	for i := range cc {
		line := cc[i]
		numbers := digitRe.FindAllStringIndex(line, -1)
		for _, v := range numbers {
			// check current
			if symRe.MatchString(string(line[clampMin(v[0]-1)])) || symRe.MatchString(string(line[clampMax(v[1], len(line)-1)])) {
				slice := line[v[0]:v[1]]
				num, _ := strconv.Atoi(slice)
				result += num
			}

			// check above
			if i != 0 {
				preLine := cc[i-1]
				for ii := clampMin(v[0] - 1); ii < clampMax(v[1]+1, len(preLine)); ii++ {
					if symRe.MatchString(string(preLine[ii])) {
						slice := line[v[0]:v[1]]
						num, _ := strconv.Atoi(slice)
						result += num
						break
					}
				}
			}

			//check below
			if i != len(cc)-1 {
				postLine := cc[clampMax(i+1, len(cc)-1)]
				for ii := clampMin(v[0] - 1); ii < clampMax(v[1]+1, len(postLine)); ii++ {
					if symRe.MatchString(string(postLine[ii])) {
						slice := line[v[0]:v[1]]
						num, _ := strconv.Atoi(slice)
						result += num
						break
					}
				}
			}
		}
	}

	fmt.Println(result)
}

func part2(cc []string) {
	var result int
	var linesDs = make([][][]int, len(cc))

	for i := range cc {
		line := cc[i]
		numbers := digitRe.FindAllStringIndex(line, -1)
		linesDs[i] = numbers

		if i < 2 {
			continue
		}

		preLine := linesDs[i-2]
		midLine := linesDs[i-1]
		postLine := linesDs[i]

		asterisks := astRe.FindAllStringIndex(cc[i-1], -1)

		for _, v := range asterisks {
			ind := v[0]

			topLineNums := getLineNums(ind, preLine, cc[i-2])
			midLineNums := getLineNums(ind, midLine, cc[i-1])
			botLineNums := getLineNums(ind, postLine, cc[i])

			var nums []int
			tot := 0

			nums = append(nums, uniq(topLineNums)...)
			nums = append(nums, uniq(midLineNums)...)
			nums = append(nums, uniq(botLineNums)...)

			if len(nums) == 2 {
				tot += nums[0] * nums[1]
			}

			result += tot
		}
	}

	fmt.Println(result)
}

func uniq(arr []int) []int {
	var seen = make(map[int]bool)
	var ret []int
	for _, v := range arr {
		if seen[v] {
			continue
		}

		seen[v] = true

		ret = append(ret, v)
	}

	return ret
}

func getLineNums(i int, list [][]int, lineString string) []int {
	var nums []int
	for _, v := range list {
		for ii := i - 1; ii < i+2; ii++ {
			if ii >= v[0] && ii < v[1] {
				slice := lineString[v[0]:v[1]]
				num, _ := strconv.Atoi(slice)
				nums = append(nums, num)
			}
		}
	}

	return nums
}
