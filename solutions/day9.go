package solutions

import (
	"github.com/cauesmelo/aoc-2023/util"
)

func getSequences(lines []string) [][]int {
	ns := make([][]int, 0)

	for _, line := range lines {
		ns = append(ns, getNumbers(line))
	}

	return ns
}

func isAllZeroes(seq []int) bool {
	for _, i := range seq {
		if i != 0 {
			return false
		}
	}

	return true
}

func extrapolate(seq []int) int {
	if isAllZeroes(seq) {
		return 0
	}

	prev := seq[0]
	var curr int

	diff := make([]int, 0)

	for i := 1; i < len(seq); i++ {
		curr = seq[i]

		diff = append(diff, curr-prev)
		prev = seq[i]
	}

	return curr + extrapolate(diff)
}

func Day9_part1() int {
	lines := util.GetInput(9, false)

	seqs := getSequences(lines)

	sum := 0
	for _, seq := range seqs {
		sum += extrapolate(seq)
	}

	return sum
}

func Day9_part2() int {
	_ = util.GetInput(9, true)

	return 0
}
