package day4

import (
	"os"
	"strconv"
	"strings"
)

func Part1() (result int) {
	c, _ := os.ReadFile("./input/day4")
	cc := strings.Split(string(c), "\n")

	for _, v := range cc {
		// ["1-3", "2-3"]
		ranges := strings.Split(v, ",")
		var all []int

		for _, r := range ranges {
			// ["1", "3"]
			lr := strings.Split(r, "-")
			if len(lr) < 2 {
				break
			}

			// 1
			l, _ := strconv.Atoi(lr[0])
			// 3
			r, _ := strconv.Atoi(lr[1])

			// [1, 3 ...]
			all = append(all, l, r)
		}

		if len(all) > 0 && (all[0] >= all[2] && all[1] <= all[3] || all[2] >= all[0] && all[3] <= all[1]) {
			result++
		}

	}

	return result
}

func Part2() (result int) {
	c, _ := os.ReadFile("./input/day4")
	cc := strings.Split(string(c), "\n")

	for _, v := range cc {
		// ["1-3", "2-3"]
		ranges := strings.Split(v, ",")
		var all []int

		for _, r := range ranges {
			// ["1", "3"]
			lr := strings.Split(r, "-")
			if len(lr) < 2 {
				break
			}

			// 1
			l, _ := strconv.Atoi(lr[0])
			// 3
			r, _ := strconv.Atoi(lr[1])

			// [1, 3 ...]
			all = append(all, l, r)
		}

		if len(all) > 0 && ((all[0] >= all[2] && all[0] <= all[3]) || (all[1] >= all[2] && all[1] <= all[3]) || (all[2] >= all[0] && all[2] <= all[1]) || (all[3] >= all[0] && all[3] <= all[1])) {
			result++
		}

	}

	return result
}
