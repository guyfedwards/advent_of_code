#!/usr/bin/env bash

set -e

d="day$1"

mkdir "$d"
touch "$d/main.go"
touch "$d/input.txt"
echo "package main

import (
  \"os\"
  \"strings\"
)

func main() {
  c, _ := os.ReadFile(\"./input.txt\")
  cc := strings.Split(string(c), \"\n\")

  part1(cc)
  part2(cc)
}

func part1(cc []string) {
  fmt.Println()
}

func part2(cc []string) {
  fmt.Println()
}
" > "$d/main.go"

cd "$d" || exit
