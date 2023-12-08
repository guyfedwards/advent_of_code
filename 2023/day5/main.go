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
	r := parse(cc[0], cc[1])
	r2 := parse2(cc[0], cc[1])

	part1(r)
	part2(r2)
}

type Race struct {
	time   int
	record int
}

var digitRe = regexp.MustCompile(`\d+`)

func parse2(line1, line2 string) Race {
	times := digitRe.FindAllString(line1, -1)
	records := digitRe.FindAllString(line2, -1)

	t, _ := strconv.Atoi(strings.Join(times, ""))
	d, _ := strconv.Atoi(strings.Join(records, ""))

	return Race{time: t, record: d}
}

func parse(line1, line2 string) []Race {
	times := digitRe.FindAllString(line1, -1)
	records := digitRe.FindAllString(line2, -1)

	var res []Race
	for i := range times {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(records[i])

		res = append(res, Race{time: t, record: d})
	}

	return res
}

func part1(rs []Race) {
	res := 1
	for _, r := range rs {
		newRecord := 0
		for i := 0; i < r.time; i++ {
			speed := i
			dist := speed * (r.time - i)
			if dist > r.record {
				newRecord++
			}
		}
		res = res * newRecord
	}
	fmt.Println("RES: ", res)
}

func part2(r Race) {
	res := 1
	newRecord := 0
	for i := 0; i < r.time; i++ {
		speed := i
		dist := speed * (r.time - i)
		if dist > r.record {
			newRecord++
		}
	}
	res = res * newRecord
	fmt.Println("RES: ", res)
}
