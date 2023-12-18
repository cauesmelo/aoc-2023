package solutions

import (
	"math"
	"regexp"
	"strings"

	"github.com/cauesmelo/aoc-2023/util"
)

type cacheKey struct {
	cfg     string
	numHash [32]byte
}

func getDmgCombinations(cfg string) []string {
	re := regexp.MustCompile(`\?`)
	slots := re.FindAllStringIndex(cfg, -1)
	totalCombinations := uint64(math.Pow(2, float64(len(slots))))

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
	lines := util.GetInput(12, true)

	sum := 0
	for _, line := range lines {
		sum += getDmgPossibilities(line)
	}

	return sum
}

// Based on https://www.youtube.com/watch?v=g3Ms5e7Jdqo
func getDmgRec(cfg string, nums []int, cache map[cacheKey]int) int {
	if cfg == "" {
		if len(nums) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(nums) == 0 {
		if strings.Contains(cfg, "#") {
			return 0
		} else {
			return 1
		}
	}

	key := cacheKey{cfg, util.GetHashForSliceInt(nums)}

	v, ok := cache[key]
	if ok {
		return v
	}

	res := 0

	if strings.Contains(".?", string(cfg[0])) {
		res += getDmgRec(cfg[1:], nums, cache)
	}

	if strings.Contains("#?", string(cfg[0])) {
		if nums[0] <= len(cfg) && !strings.Contains(cfg[:nums[0]], ".") && (nums[0] == len(cfg) || cfg[nums[0]] != '#') {
			newIdx := nums[0] + 1

			if newIdx < len(cfg) {
				res += getDmgRec(cfg[newIdx:], nums[1:], cache)
			} else {
				res += getDmgRec("", nums[1:], cache)
			}
		}
	}

	cache[key] = res
	return res
}

func (AOC) Day12_part2() int {
	lines := util.GetInput(12, false)

	sum := 0

	cache := make(map[cacheKey]int)

	for _, line := range lines {
		lineSplt := strings.Split(line, " ")
		cfg := strings.Repeat(lineSplt[0]+"?", 5)
		cfg = strings.TrimSuffix(cfg, "?")
		nsOg := util.GetNumbers(lineSplt[1])
		ns := make([]int, 0)
		ns = append(ns, nsOg...)
		ns = append(ns, nsOg...)
		ns = append(ns, nsOg...)
		ns = append(ns, nsOg...)
		ns = append(ns, nsOg...)

		sum += getDmgRec(cfg, ns, cache)
	}

	return sum
}
