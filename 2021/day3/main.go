package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func e(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func walk(a []string, s string, ind int) (filtered []string) {
	for _, v := range a {
		if v[ind] != s[ind] {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// type Counts [5]map[string]
func walkb(lines []string, s string, charInd int) []string {
	counts := [5]map[string]int{{}, {}, {}, {}, {}}
	mostCommonBit := s[charInd]
	var result []string
	var common []string
	var commonInv []string

	fmt.Printf("len lines %v", len(lines))
	for _, line := range lines {
		ss := strings.Split(line, "")

		for index, v := range ss {
			if v == "0" {
				counts[index]["0"]++
			}

			if v == "1" {
				counts[index]["1"]++
			}
		}
	}

	for _, v := range counts {
		if v["0"] == v["1"] {
			common = append(common, "1")
		} else if v["0"] > v["1"] {
			common = append(common, "0")
			commonInv = append(commonInv, "1")
		} else {
			common = append(common, "1")
			commonInv = append(commonInv, "0")
		}
	}

	// part 1
	commonS := strings.Join(common, "")
	// commonSs := strings.Join(commonInv, "")
	fmt.Printf("COMMMON: %s\n", commonS)

	for _, line := range lines {
		if line[charInd] == mostCommonBit {
			result = append(result, line)
		}
	}

	if len(result) == 1 {
		fmt.Printf("return %v\n", result)
		return result
	}

	return walkb(result, s, charInd+1)
}

func main() {
	file, err := os.Open("./input.txt")
	e(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// counts := [12]map[string]int{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}}
	counts := [5]map[string]int{{}, {}, {}, {}, {}}
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		v := scanner.Text()
		ss := strings.Split(v, "")

		for index, v := range ss {
			if v == "0" {
				counts[index]["0"]++
			}

			if v == "1" {
				counts[index]["1"]++
			}
		}
	}

	var common []string
	var commonInv []string

	for _, v := range counts {
		if v["0"] == v["1"] {
			common = append(common, "1")
		} else if v["0"] > v["1"] {
			common = append(common, "0")
			commonInv = append(commonInv, "1")
		} else {
			common = append(common, "1")
			commonInv = append(commonInv, "0")
		}
	}

	// part 1
	s := strings.Join(common, "")
	ss := strings.Join(commonInv, "")
	fmt.Printf("COMMMON: %s\n", s)
	bin, _ := strconv.ParseInt(s, 2, 64)
	epsilon := bin ^ 0b111111111111

	// part2 redo
	oxy, _ := strconv.ParseInt(walkb(lines, s, 0)[0], 2, 64)
	co2, _ := strconv.ParseInt(walkb(lines, ss, 0)[0], 2, 64)

	fmt.Printf("Result oxy: %v\n", oxy)
	fmt.Printf("Result co2: %v\n", co2)
	fmt.Printf("Result1: %d\n", bin*epsilon)
	fmt.Printf("Result part2: %v\n", oxy*co2)
}
