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
	cc := strings.Split(string(c), "\n")

	part1(cc)
	part2(cc)
}

type Game struct {
	id    int
	blue  int
	red   int
	green int
	sets  []Game
}

type Set struct {
	blue  int
	red   int
	green int
}

var digitRe = regexp.MustCompile(`\d+`)
var digitLetterRe = regexp.MustCompile(`\d+\s\w`)

func part1(cc []string) {
	total := 0

	for _, line := range cc {
		if line == "" {
			continue
		}

		gamePossible := true
		id, _ := strconv.Atoi(digitRe.FindString(line))
		setsstr := strings.Split(line, ":")
		sets := strings.Split(setsstr[1], ";")

		for _, s := range sets {
			counts := make(map[string]int)
			dc := digitLetterRe.FindAllString(s, -1)

			for _, part := range dc {
				v := strings.Split(part, " ")
				d, _ := strconv.Atoi(v[0])
				c := v[1]
				counts[c] = d
			}

			if counts["r"] > 12 || counts["g"] > 13 || counts["b"] > 14 {
				gamePossible = false
			}
		}

		if gamePossible {
			total += id
		}
	}

	fmt.Println(total)
}

func part2(cc []string) {
	total := 0

	for _, line := range cc {
		if line == "" {
			continue
		}

		mins := make(map[string]int)
		// id, _ := strconv.Atoi(digitRe.FindString(line))
		setsstr := strings.Split(line, ":")
		sets := strings.Split(setsstr[1], ";")

		for _, s := range sets {
			counts := make(map[string]int)
			dc := digitLetterRe.FindAllString(s, -1)

			for _, part := range dc {
				v := strings.Split(part, " ")
				d, _ := strconv.Atoi(v[0])
				c := v[1]
				counts[c] = d
			}

			if mins["r"] == 0 || counts["r"] > mins["r"] {
				mins["r"] = counts["r"]
			}
			if mins["g"] == 0 || counts["g"] > mins["g"] {
				mins["g"] = counts["g"]
			}
			if mins["b"] == 0 || counts["b"] > mins["b"] {
				mins["b"] = counts["b"]
			}
		}

		total += mins["r"] * mins["g"] * mins["b"]
	}

	fmt.Println(total)
}
