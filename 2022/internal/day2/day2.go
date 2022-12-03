package day2

import (
	"os"
	"strings"
)

var scores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
	"A": 1,
	"B": 2,
	"C": 3,
}

func isWin(playA, playB string) bool {
	// rock loses to paper
	if playA == "A" && playB == "Y" {
		return true
	}

	// rock beats scissors
	if playA == "A" && playB == "Z" {
		return false
	}

	// paper beats rock
	if playA == "B" && playB == "X" {
		return false
	}

	// paper loses to scissors
	if playA == "B" && playB == "Z" {
		return true
	}

	// scissors loses to rock
	if playA == "C" && playB == "X" {
		return true
	}

	// scissors beats paper
	if playA == "C" && playB == "B" {
		return false
	}

	// draw
	return false
}

func Part1() int {
	c, _ := os.ReadFile("./input/day2")
	cc := strings.Split(string(c), "\n")
	totalScore := 0

	for _, v := range cc {
		plays := strings.Split(v, " ")
		if len(plays) < 2 {
			continue
		}

		playScore := scores[plays[1]]

		if scores[plays[0]] == playScore {
			totalScore = totalScore + 3 // draw
			totalScore = totalScore + playScore
		} else if isWin(plays[0], plays[1]) {
			totalScore += 6 // win
			totalScore = totalScore + playScore
		} else {
			totalScore = totalScore + playScore
		}
	}

	return totalScore
}

func Part2() int {
	c, _ := os.ReadFile("./input/day2")
	cc := strings.Split(string(c), "\n")
	totalScore := 0

	res := map[string][]int{
		"A": []int{2, 3, 1},
		"B": []int{3, 1, 2},
		"C": []int{1, 2, 3},
	}

	for _, v := range cc {
		plays := strings.Split(v, " ")
		if len(plays) < 2 {
			continue
		}

		switch plays[1] {
		case "X": // lose
			totalScore += res[plays[0]][1]
		case "Y": // draw
			totalScore += 3
			totalScore += res[plays[0]][2]
		case "Z": // win
			totalScore += 6
			totalScore += res[plays[0]][0]
		}
	}

	return totalScore

}
