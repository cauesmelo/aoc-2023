package solutions

import (
	"slices"

	"github.com/cauesmelo/aoc-2023/util"
)

func getGalaxies(lines []string) []pos {
	galaxies := make([]pos, 0)

	for y, line := range lines {
		for x, ch := range line {
			if ch == '#' {
				galaxies = append(galaxies, pos{x, y})
			}
		}
	}

	return galaxies
}

func getExpandedSpace(lines []string) ([]int, []int) {
	rows := make([]int, 0)
	cols := make([]int, 0)

	for y, line := range lines {
		for x, ch := range line {
			if ch != '.' {
				break
			}

			if x == len(line)-1 {
				rows = append(rows, y)
			}
		}
	}

	width := len(lines[0])

	for x := 0; x < width; x++ {
		for y := 0; y < len(lines); y++ {
			if lines[y][x] != '.' {
				break
			}

			if y == len(lines)-1 {
				cols = append(cols, x)
			}
		}
	}

	return rows, cols
}

func getWidthBetween(g_a, g_b int, cols []int) int {
	a := g_a
	b := g_b

	if g_b < g_a {
		a = g_b
		b = g_a
	}

	w := 0

	for i := a; i < b; i++ {
		if slices.Contains(cols, i) {
			w++
		}

		w++
	}

	return w
}

func getHeightBetween(g_a, g_b int, rows []int) int {
	a := g_a
	b := g_b

	if g_b < g_a {
		a = g_b
		b = g_a
	}

	h := 0

	for i := a; i < b; i++ {
		if slices.Contains(rows, i) {
			h++
		}

		h++
	}

	return h
}

func getSumSPGalaxies(galaxies []pos, rows []int, cols []int) int {
	sum := 0

	for gan, g_a := range galaxies {
		gs := galaxies[gan+1:]

		for _, g_b := range gs {
			w := getWidthBetween(g_a.x, g_b.x, cols)
			h := getHeightBetween(g_a.y, g_b.y, rows)

			sum += w + h
		}
	}

	return sum
}

func (AOC) Day11_part1() int {
	lines := util.GetInput(11, false)

	galaxies := getGalaxies(lines)

	eRows, eCols := getExpandedSpace(lines)

	return getSumSPGalaxies(galaxies, eRows, eCols)
}

func (AOC) Day11_part2() int {
	return 0
}
