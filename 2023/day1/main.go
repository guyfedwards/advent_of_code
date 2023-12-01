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

func part1(cc []string) {
	total := 0
	re := regexp.MustCompile(`\d`)
	for _, v := range cc {
		if v == "" {
			continue
		}

		s := re.FindAllString(v, -1)
		f := s[0]
		l := s[len(s)-1]

		n, _ := strconv.Atoi(f + l)

		total += n
	}
	fmt.Println(total)
}

func part2(cc []string) {
	nums := map[string]int{
		"one":   1,
		"1":     1,
		"two":   2,
		"2":     2,
		"three": 3,
		"3":     3,
		"four":  4,
		"4":     4,
		"five":  5,
		"5":     5,
		"six":   6,
		"6":     6,
		"seven": 7,
		"7":     7,
		"eight": 8,
		"8":     8,
		"nine":  9,
		"9":     9,
	}

	total := 0
	re := regexp.MustCompile(`(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|(\d)`)
	for _, v := range cc {
		if v == "" {
			continue
		}

		var ns []string
		for i := range v {
			parts := re.FindAllString(v[i:], -1)
			ns = append(ns, parts...)
		}

		f := nums[ns[0]]
		l := nums[ns[len(ns)-1]]

		n, err := strconv.Atoi(fmt.Sprintf("%d%d", f, l))
		if err != nil {
			fmt.Println("fuck: %w", err)
			return
		}

		total += n
	}
	fmt.Println(total)
}
