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

func main() {
	file, err := os.Open("./input.txt")
	e(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	horizontal := 0
	vertical := 0
	aim := 0

	for scanner.Scan() {
		v := scanner.Text()
		ss := strings.Split(v, " ")
		cmd := ss[0]
		num := ss[1]

		amount, err := strconv.Atoi(num)
		e(err)

		switch cmd {
		case "forward":
			{
				horizontal += amount
				vertical += amount * aim
			}
		case "down":
			{
				aim += amount
			}
		case "up":
			{
				aim -= amount
			}
		}
	}

	result := horizontal * vertical

	fmt.Printf("Total: %d", result)
}
