package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	c, _ := os.ReadFile("./input.txt")
	cc := strings.Split(strings.TrimSpace(string(c)), "\n\n")

	rulesRaw := strings.Split(cc[0], "\n")
	updatesRaw := strings.Split(cc[1], "\n")

	rules := make([][]int, len(rulesRaw))
	for i, v := range rulesRaw {
		parts := strings.Split(v, "|")
		a, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		rules[i] = append(rules[i], a)

		b, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		rules[i] = append(rules[i], b)
	}

	updates := make([][]int, len(updatesRaw))
	for i, v := range updatesRaw {
		parts := strings.Split(v, ",")
		for _, part := range parts {
			p, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal(err)
			}
			updates[i] = append(updates[i], p)
		}
	}

	parts(rules, updates)
}

type rentry struct {
	Before []int
	After  []int
}

func parts(rules, updates [][]int) {
	// build a map of rules
	rulesD := make(map[int]rentry)
	for _, rule := range rules {
		// first rule
		val, ok := rulesD[rule[0]]
		if !ok {
			rulesD[rule[0]] = rentry{}
		}
		val.Before = append(val.Before, rule[1])
		rulesD[rule[0]] = val

		// second rule
		val, ok = rulesD[rule[1]]
		if !ok {
			rulesD[rule[1]] = rentry{}
		}
		val.After = append(val.After, rule[0])
		rulesD[rule[1]] = val
	}

	result := 0
	result2 := 0
	for _, update := range updates {
		legit := true
		for i, num := range update {
			actualBefore := update[:i+1]
			actualAfter := update[i:]
			if anyOverlap(actualBefore, rulesD[num].Before) ||
				anyOverlap(actualAfter, rulesD[num].After) {
				legit = false
				break
			}
		}

		if legit {
			middle := getMiddle(update)
			result += middle
		}

		if !legit {
			u := sort(update, rulesD)
			m := getMiddle(u)
			result2 += m
		}
	}

	fmt.Println("Part 1: ", result)
	fmt.Println("Part 2: ", result2)
}

func sort(u []int, rulesD map[int]rentry) []int {
	num := 0
	sorted := u
	for i := 0; i < len(u); i++ {
		if i == len(u)-1 {
			break
		}

		a := u[i]
		b := u[i+1]
		if IndexOf(rulesD[a].After, b) != -1 {
			num++
			sorted = sorted[:i]
			sorted = append(sorted, b, a)
			sorted = append(sorted, sorted[i+2:]...)
		} else {
			sorted = sorted[:i]
			sorted = append(sorted, a, b)
			sorted = append(sorted, sorted[i+2:]...)
		}
	}

	if num != 0 {
		sorted = sort(sorted, rulesD)
	}

	return sorted
}

func getMiddle(in []int) int {
	if len(in) == 0 {
		return -1
	}

	return in[len(in)/2]
}

func anyOverlap(a, b []int) bool {
	res := false
outer:
	for _, v := range a {
		for _, vv := range b {
			if v == vv {
				res = true
				break outer
			}
		}
	}
	return res
}

func IndexOf(arr []int, v int) int {
	for i, vv := range arr {
		if v == vv {
			return i
		}
	}

	return -1
}
