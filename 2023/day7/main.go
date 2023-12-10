package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards  []string
	counts map[string]int
	bid    int
	weight string
}

var cardWeights = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	// "J": 10, // part one
	"J": 0, // part two
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var handWeights = map[string]int{
	"five":    7,
	"four":    6,
	"full":    5,
	"three":   4,
	"twopair": 3,
	"onepair": 2,
	"high":    1,
}

func main() {
	c, _ := os.ReadFile("./input.txt")
	cc := strings.Split(strings.TrimSpace(string(c)), "\n")

	// part1(cc)
	part2(cc)
}

func part1(cc []string) {
	var hands []Hand
	for _, v := range cc {
		h := Hand{
			counts: make(map[string]int),
		}

		parts := strings.Split(v, " ")
		bid, _ := strconv.Atoi(parts[1])
		h.bid = bid

		cards := strings.Split(parts[0], "")
		for _, v := range cards {
			h.cards = append(h.cards, v)
			h.counts[v]++
		}

		hands = append(hands, h)
	}

	sort.SliceStable(hands, func(i, j int) bool { return !isStronger(hands[i], hands[j]) })

	result := 0
	for i, v := range hands {
		rank := i + 1
		fmt.Println(v.cards, v.bid, v.weight)
		result += rank * v.bid
	}

	fmt.Println(result)
}

func (h *Hand) getHandWeightJokers() int {
	var weight []string
	js := h.counts["J"]
	jss := js

	for k, c := range h.counts {
		// add the jokers as wildcards
		var v int
		v = c + js
		if k == "J" {
			v = v - 1
		}

		fmt.Printf("VVV: v: %d c: %d\n", v, c)

		if v == 5 {
			weight = append(weight, "five")
			fmt.Println(h.cards, h.bid, weight)
			break
		}

		if v == 4 {
			weight = append(weight, "four")
			fmt.Println(h.cards, h.bid, weight)
			continue
		}

		if contains(weight, "onepair") && v == 3 && c-js > 2 {
			fmt.Println(h.cards, h.bid, weight)

			weight = append(weight, "full")
			continue
		}

		if contains(weight, "three") && v >= 2 && c-js > 0 {
			fmt.Println(h.cards, h.bid, weight)

			weight = append(weight, "full")
			continue
		}

		if v == 3 {
			weight = append(weight, "three")
			fmt.Println(h.cards, h.bid, weight)
			jss = jss - c
			continue
		}

		if contains(weight, "onepair") && v == 2 {
			weight = append(weight, "twopair")
			fmt.Println(h.cards, h.bid, weight)
			continue
		}

		if v == 2 {
			weight = append(weight, "onepair")
			fmt.Println(h.cards, h.bid, weight)
			jss = jss - c
			continue
		}
	}

	highest := 0
	fmt.Println("FINALWEIGHT: ", weight)

	for _, v := range weight {
		w := handWeights[v]
		if w > highest {
			highest = w
		}
	}

	return highest
}
func contains(arr []string, v string) bool {
	for _, vv := range arr {
		if vv == v {
			return true
		}
	}
	return false
}

func (h *Hand) getHandWeight() int {
	weight := "high"
	for _, v := range h.counts {
		if v == 5 {
			weight = "five"
			fmt.Println(h.cards, h.bid, weight)
			break
		}

		if v == 4 {
			weight = "four"
			fmt.Println(h.cards, h.bid, weight)
			break
		}

		if (weight == "onepair" && v == 3) || (weight == "three" && v == 2) {
			weight = "full"
			fmt.Println(h.cards, h.bid, weight)
			break
		}

		if v == 3 {
			weight = "three"
			fmt.Println(h.cards, h.bid, weight)
		}

		if weight == "onepair" && v == 2 {
			weight = "twopair"
			fmt.Println(h.cards, h.bid, weight)
			break
		}

		if v == 2 {
			weight = "onepair"
			fmt.Println(h.cards, h.bid, weight)
		}
	}

	h.weight = weight
	return handWeights[weight]
}

func isStronger(a, b Hand) bool {
	aHandWeight := a.getHandWeight()
	bHandWeight := b.getHandWeight()

	if aHandWeight == bHandWeight {
		return isStrongerCards(a.cards, b.cards)
	}

	if aHandWeight > bHandWeight {
		return true
	}

	return false
}

func isStrongerJokers(a, b Hand) bool {
	aHandWeight := a.getHandWeightJokers()
	bHandWeight := b.getHandWeightJokers()

	if aHandWeight == bHandWeight {
		fmt.Println("COMPART")
		fmt.Println(a.cards, b.cards, aHandWeight, bHandWeight, isStrongerCards(a.cards, b.cards))
		fmt.Println("ENDCOMPART")
		return isStrongerCards(a.cards, b.cards)
	}

	if aHandWeight > bHandWeight {
		return true
	}

	return false
}

func isStrongerCards(a, b []string) bool {
	for i := range a {
		if cardWeights[a[i]] == cardWeights[b[i]] {
			continue
		}

		return cardWeights[a[i]] > cardWeights[b[i]]
	}

	return false
}

func part2(cc []string) {
	var hands []Hand
	for _, v := range cc {
		h := Hand{
			counts: make(map[string]int),
		}

		parts := strings.Split(v, " ")
		bid, _ := strconv.Atoi(parts[1])
		h.bid = bid

		cards := strings.Split(parts[0], "")
		for _, v := range cards {
			h.cards = append(h.cards, v)
			h.counts[v]++
		}

		hands = append(hands, h)
	}

	sort.SliceStable(hands, func(i, j int) bool { return !isStrongerJokers(hands[i], hands[j]) })

	result := 0
	for i, v := range hands {
		rank := i + 1
		fmt.Println(v.cards, v.bid, v.weight)
		result += rank * v.bid
	}

	fmt.Println(result)
}
