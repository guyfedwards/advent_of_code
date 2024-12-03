package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed cookie.txt
var cookie string

func main() {
	day := os.Args[1]
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%s/input", day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Cookie", strings.TrimSpace(cookie))
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	outPath := fmt.Sprintf("%s/day%s", wd, day)

	err = os.Mkdir(outPath, 0700)
	if err != nil {
		log.Fatal(err)
	}

	fd, err := os.OpenFile(fmt.Sprintf("%s/input.txt", outPath), os.O_CREATE|os.O_RDWR, 0700)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fd.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	fd.Close()

	fd, err = os.OpenFile(fmt.Sprintf("%s/main.go", outPath), os.O_CREATE|os.O_RDWR, 0700)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fd.WriteString(`
package main

import (
	"fmt"
  "strings"
)

func main() {
  // c, _ := os.ReadFile("./input.txt\")
  // cc := strings.Split(strings.TrimSpace(string(c)), "\n")

  cc := strings.Split(, "\n")

  part1(cc)
  part2(cc)
}

func part1(cc []string) {
  fmt.Println()
}

func part2(cc []string) {
  fmt.Println()
}
	`)

	if err != nil {
		log.Fatal(err)
	}
}
