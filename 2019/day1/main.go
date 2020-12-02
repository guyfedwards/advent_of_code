package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calcFuel(mass int) int {
	third := math.Floor(float64(mass) / 3)
	fuel := int(third - 2)

	if fuel > 0 {
		fuel += calcFuel(fuel)
	}

	if fuel <= 0 {
		return 0
	}

	return int(fuel)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		fuel := calcFuel(i)

		total += fuel
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d\n", total)
}
