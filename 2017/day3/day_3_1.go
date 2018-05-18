package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(path string) (res int, err error) {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return 0, err
	}

	str := strings.TrimSpace(string(data))
	res, err = strconv.Atoi(str)

	if err != nil {
		return 0, err
	}

	return res, nil
}

func getDistance(n int) (x, y int) {
	var (
		total  int
		levels = 2
		side   = 1
	)

	// 277678

	for ok := true; ok; ok = total < n {
		side = side + 2
		fmt.Println("SIDE", side)
		total = side*4 + 4
		levels++
		fmt.Println("I", total)
	}

	fmt.Println("Levels: ", levels/2)

	return x, y
}

func addLayer() {

}

func main() {

	data, err := readFile("./input.txt")
	if err != nil {
		panic(err)
	}

	// data := 49
	fmt.Println(getDistance(data))
}
