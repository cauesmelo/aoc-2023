package solutions

import (
	"github.com/cauesmelo/aoc-2023/util"
)

type dir struct {
	x int
	y int
}

func shootBeam(runes [][]rune, og pos, d dir) []pos {
	newPos := pos{og.x + d.x, og.y + d.y}
	res := []pos{}

	if newPos.x < 0 || newPos.y < 0 || newPos.x == len(runes[0]) || newPos.y == len(runes) {
		return res
	}

	if runes[newPos.y][newPos.x] == 'x' {
		return res
	}

	if runes[newPos.y][newPos.x] == 'I' && d.y != 0 {
		return res
	}

	if runes[newPos.y][newPos.x] == '_' && d.x != 0 {
		return res
	}

	res = append(res, newPos)

	if runes[newPos.y][newPos.x] == '-' {
		if d.x != 0 {
			res = append(res, shootBeam(runes, newPos, d)...)
		} else {
			res = append(res, shootBeam(runes, newPos, dir{1, 0})...)
			res = append(res, shootBeam(runes, newPos, dir{-1, 0})...)
		}
	}

	if runes[newPos.y][newPos.x] == '|' {
		if d.y != 0 {
			res = append(res, shootBeam(runes, newPos, d)...)
		} else {
			res = append(res, shootBeam(runes, newPos, dir{0, 1})...)
			res = append(res, shootBeam(runes, newPos, dir{0, -1})...)
		}
	}

	if runes[newPos.y][newPos.x] == '/' {
		if d.x == 1 {
			res = append(res, shootBeam(runes, newPos, dir{0, -1})...)
		}

		if d.x == -1 {
			res = append(res, shootBeam(runes, newPos, dir{0, 1})...)
		}

		if d.y == 1 {
			res = append(res, shootBeam(runes, newPos, dir{-1, 0})...)
		}

		if d.y == -1 {
			res = append(res, shootBeam(runes, newPos, dir{1, 0})...)
		}
	}

	if runes[newPos.y][newPos.x] == '\\' {
		if d.x == 1 {
			res = append(res, shootBeam(runes, newPos, dir{0, 1})...)
		}

		if d.x == -1 {
			res = append(res, shootBeam(runes, newPos, dir{0, -1})...)
		}

		if d.y == 1 {
			res = append(res, shootBeam(runes, newPos, dir{1, 0})...)
		}

		if d.y == -1 {
			res = append(res, shootBeam(runes, newPos, dir{-1, 0})...)
		}
	}

	if runes[newPos.y][newPos.x] == 'I' || runes[newPos.y][newPos.x] == '_' {
		runes[newPos.y][newPos.x] = 'x'
		res = append(res, shootBeam(runes, newPos, d)...)
	}

	if runes[newPos.y][newPos.x] == '.' {
		if d.x != 0 {
			runes[newPos.y][newPos.x] = '_'
		}

		if d.y != 0 {
			runes[newPos.y][newPos.x] = 'I'
		}

		res = append(res, shootBeam(runes, newPos, d)...)
	}

	return res
}

func calcEnergized(lines []string) int {
	runes := make([][]rune, len(lines))

	for i, line := range lines {
		runes[i] = []rune(line)
	}

	res := shootBeam(runes, pos{-1, 0}, dir{1, 0})

	m := map[pos]bool{}

	for _, p := range res {
		m[p] = true
	}

	return len(m)
}

func (AOC) Day16_part1() int {
	lines := util.GetInput(16, false)

	return calcEnergized(lines)
}

func (AOC) Day16_part2() int {
	// lines := util.GetInput(16, true)

	return 0
}
