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
		key := parts[0]

		children := regexp.MustCompile(`(\d)\s([\w\s]+bag)`).FindAllStringSubmatch(parts[1], -1)
		// children := strings.Split(parts[1], ",")
		fmt.Println(key, "=====", children)
		for _, child := range children {
			i, err := strconv.Atoi(child[1])
			check(err)

			if rulesMap[key] == nil {
				rulesMap[key] = make(map[string]int)
			}

			rulesMap[key][child[2]] = i
			// for _, c := range child {
			// 	fmt.Printf("%+v\n",  c)
			// }
		}

		// fmt.Println(parts)
	}

	for k, v := range rulesMap {
		fmt.Println(k, v)
	}

	// key := s
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
