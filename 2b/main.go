package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type round struct {
	red   int
	green int
	blue  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getNumber(s string) int {
	re := regexp.MustCompile(`\d+`)

	match := re.FindString(s)

	num, err := strconv.Atoi(match)
	check(err)

	return num
}

func isValid(line string) int {
	lineParts := strings.Split(line, ":")
	gameSetsStr := lineParts[1]

	gameSets := strings.Split(gameSetsStr, ";")

	rounds := make([]round, 0)

	for _, set := range gameSets {
		balls := strings.Split(set, ",")

		s := round{}

		for _, ball := range balls {
			n := getNumber(ball)

			if strings.Contains(ball, "red") {
				s.red = n
			} else if strings.Contains(ball, "green") {
				s.green = n
			} else if strings.Contains(ball, "blue") {
				s.blue = n
			}
		}

		rounds = append(rounds, s)
	}

	fewestN := round{}

	for _, round := range rounds {
		if fewestN.red < round.red {
			fewestN.red = round.red
		}

		if fewestN.green < round.green {
			fewestN.green = round.green
		}

		if fewestN.blue < round.blue {
			fewestN.blue = round.blue
		}
	}

	sum := fewestN.red * fewestN.green * fewestN.blue

	return sum
}

func main() {
	dat, err := os.ReadFile("./input2.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")

	total := 0

	for _, line := range lines {
		sum := isValid(line)

		total = total + sum
	}

	fmt.Println(total)
}
