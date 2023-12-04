package solutions

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/cauesmelo/aoc-2023/util"
)

func getNumbers(line string) []int {
	re := regexp.MustCompile(`\d+`)

	f := re.FindAllStringIndex(line, -1)

	numbers := make([]int, 0)

	for _, match := range f {
		valueStr := line[match[0]:match[1]]
		val, err := strconv.Atoi(valueStr)
		util.Check(err)

		numbers = append(numbers, val)
	}

	return numbers
}

func procLine(line string) int {
	lineSplit := strings.Split(line, ": ")
	numbers := strings.Split(lineSplit[1], " | ")

	winStr := numbers[0]
	haveStr := numbers[1]

	win := getNumbers(winStr)
	have := getNumbers(haveStr)

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

func Day4_part1() int {
	lines := util.GetInput(4, false)

	sum := 0

	for _, line := range lines {
		sum = sum + procLine(line)
	}

	return sum
}

func Day4_part2() int {
	_ = util.GetInput(4, false)

	return 0
}
