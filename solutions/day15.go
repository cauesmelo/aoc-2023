package solutions

import "github.com/cauesmelo/aoc-2023/util"

func getHASH(line string) int {

	total := 0
	curr := 0

	for _, ch := range line {
		if ch == ',' {
			total += curr
			curr = 0
			continue
		}

		curr += int(ch)

		curr *= 17
		curr %= 256
	}

	return total + curr
}

func (AOC) Day15_part1() int {
	lines := util.GetInput(15, false)
	l := lines[0]

	return getHASH(l)
}

func (AOC) Day15_part2() int {
	// lines := util.GetInput(14, false)

	return 0
}
