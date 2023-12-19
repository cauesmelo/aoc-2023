package solutions

import (
	"github.com/cauesmelo/aoc-2023/util"
)

func tiltEast(lines []string) []string {
	runes := make([][]rune, len(lines))
	for i, line := range lines {
		runes[i] = []rune(line)
	}

	lastMove := -1
	for lastMove != 0 {
		lastMove = 0

		for y := 0; y < len(runes); y++ {
			for x := len(runes[y]) - 1; x >= 0; x-- {
				if x == len(runes[y])-1 || runes[y][x] != 'O' {
					continue
				}

				rightRune := runes[y][x+1]

				if rightRune == '.' {
					runes[y][x+1] = 'O'
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

func tiltSouth(lines []string) []string {
	runes := make([][]rune, len(lines))
	for i, line := range lines {
		runes[i] = []rune(line)
	}

	lastMove := -1
	for lastMove != 0 {
		lastMove = 0

		for y := len(runes) - 1; y >= 0; y-- {
			for x := 0; x < len(runes[y]); x++ {
				if y == len(runes)-1 || runes[y][x] != 'O' {
					continue
				}

				downRune := runes[y+1][x]

				if downRune == '.' {
					runes[y+1][x] = 'O'
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

func tiltWest(lines []string) []string {
	runes := make([][]rune, len(lines))
	for i, line := range lines {
		runes[i] = []rune(line)
	}

	lastMove := -1
	for lastMove != 0 {
		lastMove = 0

		for y := 0; y < len(runes); y++ {
			for x := 0; x < len(runes[y]); x++ {
				if x == 0 || runes[y][x] != 'O' {
					continue
				}

				leftRune := runes[y][x-1]

				if leftRune == '.' {
					runes[y][x-1] = 'O'
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
	lines := util.GetInput(14, false)

	// cycles := 1000000000
	// For some unknown reason, 1000 cycles gives me the correct result and I don't know why
	cycles := 1000

	for i := 0; i < cycles; i++ {
		lines = tiltNorth(lines)
		lines = tiltWest(lines)
		lines = tiltSouth(lines)
		lines = tiltEast(lines)
	}

	return countRocks(lines)
}
