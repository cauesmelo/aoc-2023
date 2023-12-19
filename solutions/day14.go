package solutions

import "github.com/cauesmelo/aoc-2023/util"

func tiltNorth(lines []string) []string {
	runes := make([][]rune, len(lines))
	for i, line := range lines {
		runes[i] = []rune(line)
	}

	lastMove := -1
	for lastMove != 0 {
		lastMove = 0

		for y := 0; y < len(runes); y++ {
			for x := 0; x < len(runes[y]); x++ {
				if y == 0 || runes[y][x] != 'O' {
					continue
				}

				upRune := runes[y-1][x]

				if upRune == '.' {
					runes[y-1][x] = 'O'
					runes[y][x] = '.'
					lastMove += 1
				}
			}
		}
	}

	newLines := make([]string, len(runes))
	for i, line := range runes {
		newLines[i] = string(line)
	}

	return newLines
}

func countRocks(lines []string) int {
	sum := 0

	for y, line := range lines {
		for _, r := range line {
			if r == 'O' {
				tot := len(line)
				sum += tot - y
			}
		}
	}

	return sum
}

func (AOC) Day14_part1() int {
	lines := util.GetInput(14, false)
	res := tiltNorth(lines)

	return countRocks(res)
}

func (AOC) Day14_part2() int {
	// lines := util.GetInput(14, true)

	total := 0

	return total
}
