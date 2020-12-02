package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func genArray(code []string) [][]int {
	var rows [][]int
	var currentRow = 0
	rows = append(rows, []int{})

	for i, item := range code {
		itemi, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}

		if i%4 == 0 && i != 0 {
			currentRow++
			rows = append(rows, []int{})
		}

		rows[currentRow] = append(rows[currentRow], itemi)
	}

	return rows
}

func genCodeAsInts(code []string) []int {
	var codeAsInts []int

	for _, s := range code {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		codeAsInts = append(codeAsInts, i)
	}
	return codeAsInts
}

func isAllowedOperator(i int) bool {
	if i == 1 || i == 2 || i == 99 {
		return true
	}

	return false
}

func run(content []byte, input1 int, input2 int) string {
	code := strings.Split(strings.TrimSpace(string(content)), ",")
	code[1] = strconv.Itoa(input1)
	code[2] = strconv.Itoa(input2)

	rows := genArray(code)
	codeAsInts := genCodeAsInts(code)

	// codeAsInts[1] = input1
	// codeAsInts[2] = input2

	for i, _ := range rows {
		row := rows[i]

		if !isAllowedOperator(row[0]) {
			continue
		}

		if row[0] == 99 {
			break
		}

		var result int

		if row[0] == 1 {
			result = codeAsInts[row[1]] + codeAsInts[row[2]]
		}

		if row[0] == 2 {
			result = codeAsInts[row[1]] * codeAsInts[row[2]]
		}

		code[row[3]] = strconv.Itoa(result)
		rows = genArray(code)
		codeAsInts = genCodeAsInts(code)
	}

	return code[0]
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var finished = false

	part1 := run(content, 12, 2)
	var inputs []int

	for i := 0; i < 100; i++ {
		if finished {
			break
		}

		for ii := 0; ii < 100; ii++ {
			result := run(content, i, ii)
			fmt.Printf("%v\n", result)

			if result == "19690720" {
				finished = true
				fmt.Printf("%v %d %d\n", result, i, ii)
				inputs = []int{i, ii}
				break
			}
		}
	}

	fmt.Printf("Result part 1: %v\n", part1)
	fmt.Printf("Inputs part 2: %v\n", inputs)
	fmt.Printf("Result part 2: %v\n", 100*inputs[0]+inputs[1])
}
