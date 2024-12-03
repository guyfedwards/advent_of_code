package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	c, _ := os.ReadFile("./input.txt")
	cc := strings.Split(strings.TrimSpace(string(c)), "\n")

	part1(cc)
	part2(cc)
}

var mulrgx = regexp.MustCompile(`mul\(\d+,\d+\)`)
var mulcondrgx = regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
var digitRgx = regexp.MustCompile(`\d+`)

func part1(cc []string) {
	total := 0
	for _, v := range cc {
		matches := mulrgx.FindAllString(v, -1)
		for _, match := range matches {
			ds := digitRgx.FindAllString(match, -1)
			a, err := strconv.Atoi(ds[0])
			if err != nil {
				log.Fatal(err)
			}
			b, err := strconv.Atoi(ds[1])
			if err != nil {
				log.Fatal(err)
			}

			total += a * b
		}
	}

	fmt.Println("Part 1: ", total)
}

func part2(cc []string) {
	total := 0
	// join into one big string to match across \n
	joined := strings.Join(cc, "")
	matches := mulcondrgx.FindAllString(joined, -1)
	enabled := true
	for _, match := range matches {
		if match == "don't()" {
			enabled = false
			continue
		}

		if match == "do()" {
			enabled = true
			continue
		}

		if enabled {
			ds := digitRgx.FindAllString(match, -1)
			a, err := strconv.Atoi(ds[0])
			if err != nil {
				log.Fatal(err)
			}

			b, err := strconv.Atoi(ds[1])
			if err != nil {
				log.Fatal(err)
			}

			total += a * b
		}
	}

	fmt.Println("Part 2: ", total)
}
