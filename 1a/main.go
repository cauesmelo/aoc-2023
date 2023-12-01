package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")

	total := 0

	for _, line := range lines {
		total = total + calcLine(line)
	}
	fmt.Println(total)
}

func isNumber(r rune) bool {
	if r > 48 && r < 58 {
		return true
	}

	return false
}

func calcLine(line string) int {
	first := -1
	last := -1

	for _, char := range line {
		if isNumber(char) {
			n, err := strconv.Atoi(string(char))
			check(err)

			last = n

			if first == -1 {
				first = n
			}
		}
	}

	if last == -1 {
		last = first
	}

	str := fmt.Sprintf("%d%d", first, last)

	finalN, err := strconv.Atoi(str)
	check(err)

	return finalN
}
