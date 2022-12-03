package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partOne(rules []string) (result int) {
	rulesMap := map[string]map[string]int{}
	for _, rule := range rules {
		if rule == "" {
			continue
		}
		parts := strings.Split(rule, "contain")
		containingBag := parts[0]

		children := regexp.MustCompile(`(\d)\s([\w\s]+bag)`).FindAllStringSubmatch(parts[1], -1)
		fmt.Println(containingBag, "=====", children)

		for _, child := range children {
			i, err := strconv.Atoi(child[1])
			check(err)

			if rulesMap[containingBag] == nil {
				rulesMap[containingBag] = make(map[string]int)
			}

			rulesMap[containingBag][child[2]] = i
		}

		// fmt.Println(parts)
	}

	for k, v := range rulesMap {

		fmt.Println(k, v)
	}

	return result
}

func partTwo(rules []string) (result int) {
	return result
}

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	check(err)
	input := strings.Split(string(data), "\n")

	fmt.Printf("Result1: %v\n", partOne(input))
	fmt.Printf("Result2: %v\n", partTwo(input))
}
