package solutions

import (
	"math"
	"strings"

	"github.com/cauesmelo/aoc-2023/util"
)

type seq struct {
	start int
	end   int
	diff  int
}

type set = []seq

func scanSeeds(line string) []int {
	return getNumbers(line)
}

func scanSeq(line string) seq {
	numbers := getNumbers(line)

	source := numbers[1]
	dest := numbers[0]
	size := numbers[2]

	return seq{
		start: source,
		end:   source + size - 1,
		diff:  dest - source,
	}
}

func getLocation(seed int, sets []set) int {
	curr := seed

	for _, set := range sets {
		var seqSelected seq
		found := false

		for _, seq := range set {
			if curr >= seq.start && curr <= seq.end {
				seqSelected = seq
				found = true
			}
		}

		if found {
			curr = curr + seqSelected.diff
		}
	}

	return curr
}

func Day5_part1() int {
	solution := math.Inf(1)
	lines := util.GetInput(5, false)

	seeds := scanSeeds(lines[0])
	sets := make([]set, 0)

	for i, line := range lines {
		if i == 0 || len(line) == 0 {
			continue
		}

		if strings.Contains(line, "map:") {
			sets = append(sets, make([]seq, 0))
			continue
		}

		lastSetIdx := len(sets) - 1

		sets[lastSetIdx] = append(sets[lastSetIdx], scanSeq(line))
	}

	seedsResult := make([]int, 0)

	for _, seed := range seeds {
		seedsResult = append(seedsResult, getLocation(seed, sets))
	}

	for _, res := range seedsResult {
		if res < int(solution) {
			solution = float64(res)
		}
	}

	return int(solution)
}

func Day5_part2() int {
	// _ = util.GetInput(4, false)

	total := 0

	return total
}
