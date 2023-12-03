package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	ID    int
	round [3]round
}

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

func validateRounds(rounds []round) bool {
	for _, round := range rounds {
		if round.red > 12 {
			return false
		}

		if round.green > 13 {
			return false
		}

		if round.blue > 14 {
			return false
		}
	}

	return true
}

func isValid(line string) (int, bool) {
	lineParts := strings.Split(line, ":")
	gameIdentifier := lineParts[0]
	gameSetsStr := lineParts[1]

	ID := getNumber(gameIdentifier)

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

	return ID, validateRounds(rounds)
}

func main() {
	dat, err := os.ReadFile("./input2.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")

	total := 0

	for _, line := range lines {
		ID, valid := isValid(line)

		if valid {
			total = total + ID
		}
	}

	fmt.Println(total)
}
