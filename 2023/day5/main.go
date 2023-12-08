package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type all struct {
	seedToSoil   [][]int
	soilToFert   [][]int
	fertToWater  [][]int
	waterToLight [][]int
	lightToTemp  [][]int
	tempToHum    [][]int
	humToLoc     [][]int
}

func main() {
	c, _ := os.ReadFile("./input.txt")
	cc := strings.Split(strings.TrimSpace(string(c)), "\n\n")

	// part1(cc)
	part2(cc)
}

func part1(cc []string) {
	seeds := parseSeeds(cc[0])
	// [][]int [dest, source, range]
	seedToSoil := parseRanges(cc[1])
	soilToFert := parseRanges(cc[2])
	fertToWater := parseRanges(cc[3])
	waterToLight := parseRanges(cc[4])
	lightToTemp := parseRanges(cc[5])
	tempToHum := parseRanges(cc[6])
	humToLoc := parseRanges(cc[7])
	a := all{
		seedToSoil:   seedToSoil,
		soilToFert:   soilToFert,
		fertToWater:  fertToWater,
		waterToLight: waterToLight,
		lightToTemp:  lightToTemp,
		tempToHum:    tempToHum,
		humToLoc:     humToLoc,
	}

	var lowest int
	for _, seed := range seeds {
		loc := a.path(seed)

		if lowest == 0 || loc < lowest {
			lowest = loc
		}
	}

	fmt.Println(lowest)
}

func part2(cc []string) {
	seeds := parseSeeds(cc[0])
	// [][]int [dest, source, range]
	seedToSoil := parseRanges(cc[1])
	soilToFert := parseRanges(cc[2])
	fertToWater := parseRanges(cc[3])
	waterToLight := parseRanges(cc[4])
	lightToTemp := parseRanges(cc[5])
	tempToHum := parseRanges(cc[6])
	humToLoc := parseRanges(cc[7])
	a := all{
		seedToSoil:   seedToSoil,
		soilToFert:   soilToFert,
		fertToWater:  fertToWater,
		waterToLight: waterToLight,
		lightToTemp:  lightToTemp,
		tempToHum:    tempToHum,
		humToLoc:     humToLoc,
	}

	var lowest int
	fmt.Println(seeds)
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := start + seeds[i+1]
		ch := make(chan int, end-start)

		for ii := start; ii < end; ii++ {
			go func() {
				loc := a.path(ii)

				ch <- loc
			}()

			select {
			case loc := <-ch:
				if lowest == 0 || loc < lowest {
					fmt.Println("lowest")
					lowest = loc
				}
			}
		}
	}

	fmt.Println(lowest)
	fmt.Println()
}

// return location
func (a all) path(seed int) int {
	soil := convert(seed, a.seedToSoil)
	fert := convert(soil, a.soilToFert)
	water := convert(fert, a.fertToWater)
	light := convert(water, a.waterToLight)
	temp := convert(light, a.lightToTemp)
	hum := convert(temp, a.tempToHum)
	loc := convert(hum, a.humToLoc)

	return loc
}

func convert(in int, cs [][]int) int {
	res := in
	for _, v := range cs {
		dest := v[0]
		source := v[1]
		rang := v[2]

		if in >= source && in < source+rang {
			res = dest + (in - source)
			break
		}
	}

	return res
}

var digitRe = regexp.MustCompile(`\d+`)

func parseSeeds(in string) []int {
	var ret []int
	parts := digitRe.FindAllString(in, -1)
	for _, p := range parts {
		i, _ := strconv.Atoi(p)
		ret = append(ret, i)
	}

	return ret
}

func parseRanges(in string) [][]int {
	lines := strings.Split(in, "\n")
	var res = make([][]int, len(lines)-1)

	for i, line := range lines[1:] {
		parts := digitRe.FindAllString(line, -1)
		for _, p := range parts {
			num, _ := strconv.Atoi(p)
			res[i] = append(res[i], num)
		}
	}

	return res
}
