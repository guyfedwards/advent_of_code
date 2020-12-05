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

func validateByr(s string) bool {
	i, err := strconv.Atoi(s)
	if len(s) != 4 || err != nil {
		return false
	}

	if i >= 1920 && i <= 2002 {
		return true
	}

	return false
}

func validateIyr(s string) bool {
	i, err := strconv.Atoi(s)
	if len(s) != 4 || err != nil {
		return false
	}

	if i >= 2010 && i <= 2020 {
		return true
	}

	return false
}

func validateEyr(s string) bool {
	i, err := strconv.Atoi(s)
	if len(s) != 4 || err != nil {
		return false
	}

	if i >= 2020 && i <= 2030 {
		return true
	}

	return false
}

func validateHgt(s string) bool {
	prefix, err := strconv.Atoi(s[0 : len(s)-2])
	if err != nil {
		return false
	}

	suffix := s[len(s)-2:]
	if suffix == "cm" && prefix >= 150 && prefix <= 193 {
		return true
	}

	if suffix == "in" && prefix >= 59 && prefix <= 76 {
		return true
	}

	return false
}

func validateHcl(s string) bool {
	prefix := s[0:1]
	suffix := s[1:]
	isValidChars := regexp.MustCompile(`^[a-f0-9]+$`).MatchString(suffix)

	if prefix == "#" && len(suffix) == 6 && isValidChars {
		return true
	}

	return false
}

func validateEcl(s string) bool {
	enum := []string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}

	for _, e := range enum {
		if s == e {
			return true
		}
	}

	return false
}

func validatePid(s string) bool {
	_, err := strconv.Atoi(s)

	if err != nil || len(s) != 9 {
		return false
	}

	return true
}

func containsAllFields(s string) bool {
	fields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	result := true

	for _, f := range fields {
		if !strings.Contains(s, f) {
			result = false
			break
		}
	}

	return result
}

type validator = func(string) bool

func validatePassport(fields []string) bool {
	validators := map[string]validator{
		"byr": validateByr,
		"iyr": validateIyr,
		"eyr": validateEyr,
		"hgt": validateHgt,
		"hcl": validateHcl,
		"ecl": validateEcl,
		"pid": validatePid,
	}

	passMap := map[string]string{}
	for _, f := range fields {
		parts := strings.Split(f, ":")
		passMap[parts[0]] = parts[1]
	}

	for k, v := range passMap {
		if k == "cid" {
			continue
		}
		if !validators[k](v) {
			return false
		}
	}

	return true
}

func partOne(passports []string) (result int) {
	for _, p := range passports {
		if containsAllFields(p) {
			result++
		}
	}
	return result
}

func partTwo(passports []string) (result int) {
	for _, p := range passports {
		if containsAllFields(p) && validatePassport(strings.Fields(p)) {
			result++
		}
	}
	return result
}

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	check(err)
	input := strings.Split(string(data), "\n\n")

	fmt.Printf("Result1: %v\n", partOne(input))
	fmt.Printf("Result2: %v\n", partTwo(input))
}
