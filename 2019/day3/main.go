package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type lineMap struct {
	v int
	h int
}

func mapLine(line string, l *lineMap) {
	for _, instruction := range strings.Split(line, ",") {
		dir := string(instruction[0])
		amount, err := strconv.Atoi(instruction[1:])
		if err != nil {
			log.Fatal(err)
		}

		if dir == "U" {
			fmt.Println("Up")
			l.v += amount
		}

		if dir == "R" {
			fmt.Println("Right")
			l.h += amount
		}

		if dir == "D" {
			fmt.Println("Down")
			l.v -= amount
		}

		if dir == "L" {
			fmt.Println("Left")
			l.h -= amount
		}
	}
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	line1 := lineMap{}
	line2 := lineMap{}

	mapLine(lines[0], &line1)
	mapLine(lines[1], &line2)
	fmt.Printf("%v", line1)
	fmt.Printf("%v", line2)
}
