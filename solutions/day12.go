package solutions

import (
	"math"
	"regexp"
	"strings"

	"github.com/cauesmelo/aoc-2023/util"
)

func getDmgCombinations(cfg string) []string {
	re := regexp.MustCompile(`\?`)
	slots := re.FindAllStringIndex(cfg, -1)
	totalCombinations := int(math.Pow(2, float64(len(slots))))

	combinations := make([]string, totalCombinations)

	for cIdx := range combinations {
		currStr := cfg

		for sIdx := range slots {
			strIdx := slots[sIdx][0]

			factor := int(math.Pow(2, float64(sIdx)))
			flipStr := ((cIdx / factor) % 2) == 0

			if flipStr {
				currStr = util.ReplaceAtIndex(currStr, '.', strIdx)
			} else {
				currStr = util.ReplaceAtIndex(currStr, '#', strIdx)
			}
		}

		combinations[cIdx] = currStr
	}

	return combinations
}

func isValidCfg(cfg string, ns []int) bool {
	re := regexp.MustCompile(`#+`)
	matches := re.FindAllString(cfg, -1)

	if len(matches) != len(ns) {
		return false
	}

	for i, expectedLen := range ns {
		if len(matches[i]) != expectedLen {
			return false
		}
	}

	return true
}

func getDmgPossibilities(line string) int {
	lineSplt := strings.Split(line, " ")
	cfg := lineSplt[0]
	ns := util.GetNumbers(lineSplt[1])

	combinations := getDmgCombinations(cfg)

	result := 0

	for _, c := range combinations {
		if isValidCfg(c, ns) {
			result++
		}
	}

	return result
}

func (AOC) Day12_part1() int {
	lines := util.GetInput(12, false)

	sum := 0
	for _, line := range lines {
		sum += getDmgPossibilities(line)
	}

	return sum
}

func (AOC) Day12_part2() int {
	return 0
}
