package solutions

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/cauesmelo/aoc-2023/util"
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

func isValid_part1(line string) (int, bool) {
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

func Day2_part1() int {
	lines := util.GetInput(2, false)

	total := 0

	for _, line := range lines {
		ID, valid := isValid_part1(line)

		if valid {
			total = total + ID
		}
	}

	return total
}

func isValid_part2(line string) int {
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

func Day2_part2() int {
	lines := util.GetInput(2, false)

	total := 0

	for _, line := range lines {
		sum := isValid_part2(line)

		total = total + sum
	}

	return total
}
