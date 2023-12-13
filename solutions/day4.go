package solutions

import (
	"math"
	"strings"

	"github.com/cauesmelo/aoc-2023/util"
)

type card struct {
	position int
	matchs   int
	copies   int
}

func procLine(line string) int {
	lineSplit := strings.Split(line, ": ")
	numbers := strings.Split(lineSplit[1], " | ")

	winStr := numbers[0]
	haveStr := numbers[1]

	win := util.GetNumbers(winStr)
	have := util.GetNumbers(haveStr)

	match := make([]int, 0)

	for _, n := range win {
		for _, j := range have {
			if n == j {
				match = append(match, j)
			}
		}
	}

	if len(match) == 0 {
		return 0
	}

	return int(math.Pow(2, float64(len(match)-1)))
}

func (AOC) Day4_part1() int {
	lines := util.GetInput(4, false)

	sum := 0

	for _, line := range lines {
		sum = sum + procLine(line)
	}

	return sum
}

func procLinev2(line string, index int) card {
	lineSplit := strings.Split(line, ": ")
	numbers := strings.Split(lineSplit[1], " | ")
	lineNumber := index + 1

	winStr := numbers[0]
	haveStr := numbers[1]

	win := util.GetNumbers(winStr)
	have := util.GetNumbers(haveStr)

	match := make([]int, 0)

	for _, n := range win {
		for _, j := range have {
			if n == j {
				match = append(match, j)
			}
		}
	}

	return card{
		position: lineNumber,
		copies:   0,
		matchs:   len(match),
	}
}

func procCopy(cards []card, i int) []card {
	multiplier := 1 + cards[i].copies
	matchsCounter := cards[i].matchs
	offSet := 1

	for matchsCounter > 0 {
		cards[i+offSet].copies = cards[i+offSet].copies + multiplier

		matchsCounter--
		offSet++
	}

	return cards
}

func (AOC) Day4_part2() int {
	lines := util.GetInput(4, false)

	cards := make([]card, 0)

	for i, line := range lines {
		cards = append(cards, procLinev2(line, i))
	}

	for i := range cards {
		cards = procCopy(cards, i)
	}

	total := 0

	for _, card := range cards {
		total = total + card.copies + 1
	}

	return total
}
