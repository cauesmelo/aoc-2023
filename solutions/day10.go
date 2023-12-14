package solutions

import (
	"fmt"
	"slices"

	"github.com/cauesmelo/aoc-2023/util"
)

type Direction uint8

const (
	None Direction = iota
	Up
	Right
	Down
	Left
)

type pos struct {
	x int
	y int
}

func getWays(curr pos, lines []string) []Direction {
	mv := make([]Direction, 0)

	switch lines[curr.y][curr.x] {
	case '|':
		mv = append(mv, Up)
		mv = append(mv, Down)
	case '-':
		mv = append(mv, Left)
		mv = append(mv, Right)
	case 'L':
		mv = append(mv, Up)
		mv = append(mv, Right)
	case 'J':
		mv = append(mv, Up)
		mv = append(mv, Left)
	case '7':
		mv = append(mv, Down)
		mv = append(mv, Left)
	case 'F':
		mv = append(mv, Right)
		mv = append(mv, Down)
	case 'S':
		mv = append(mv, Up)
		mv = append(mv, Left)
		mv = append(mv, Right)
		mv = append(mv, Down)
	}

	return mv
}

func getDir(a, b pos) Direction {
	diff := pos{
		x: a.x - b.x,
		y: a.y - b.y,
	}

	if diff.x > 0 {
		return Right
	}

	if diff.x < 0 {
		return Left
	}

	if diff.y > 0 {
		return Down
	}

	if diff.y < 0 {
		return Up
	}

	return None
}

func getStart(lines []string) pos {
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				return pos{x, y}
			}
		}
	}

	panic("No start found.")
}

func getNextFrom(curr pos, prev pos, lines []string) pos {
	max := len(lines) - 1
	from := getDir(prev, curr)
	waysCurr := getWays(curr, lines)

	for _, w := range waysCurr {
		if w == from {
			continue
		}

		mv := None
		future := pos{
			x: curr.x,
			y: curr.y,
		}

		if w == Up && curr.y > 0 {
			future.y -= 1
			mv = Down
		}

		if w == Right && curr.x < max {
			future.x += 1
			mv = Left
		}

		if w == Down && curr.y < max {
			future.y += 1
			mv = Up
		}

		if w == Left && curr.x > 0 {
			future.x -= 1
			mv = Right
		}

		waysFuture := getWays(future, lines)

		if slices.Contains(waysFuture, mv) {
			return future
		}
	}

	panic(fmt.Sprintf("No way found on x:%d y:%d\n", curr.x, curr.y))
}

func getFarthests(s pos, lines []string) int {
	prev := s
	curr := s
	steps := 0

	done := false
	for !done {
		n := getNextFrom(curr, prev, lines)
		prev = curr
		curr = n
		steps++

		if lines[curr.y][curr.x] == 'S' {
			done = true
		}
	}

	return steps
}

func getPath(s pos, lines []string) []pos {
	prev := s
	curr := s
	path := make([]pos, 0)

	done := false
	for !done {
		n := getNextFrom(curr, prev, lines)
		prev = curr
		curr = n
		path = append(path, n)

		if lines[curr.y][curr.x] == 'S' {
			done = true
		}
	}

	return path
}

func isInside(line string, x int, y int, m map[pos]pos) bool {
	cross := 0
	corner := '0'

	for i := x + 1; i < len(line); i++ {
		_, isPath := m[pos{i, y}]

		if !isPath {
			continue
		}

		c := line[i]
		if c == '|' {
			cross++
			continue
		}

		if c != '-' && c != 'S' {
			if corner != '0' {
				if corner == 'L' && c == '7' {
					cross++
				} else if corner == 'F' && c == 'J' {
					cross++
				}

				corner = '0'
			} else {
				corner = rune(c)
			}
		}
	}

	return cross%2 == 1
}

func countInner(lines []string, path []pos) int {
	count := 0
	m := make(map[pos]pos)
	for _, p := range path {
		m[pos{x: p.x, y: p.y}] = p
	}

	for y, line := range lines {
		for x := range line {
			_, isPath := m[pos{x, y}]

			if isPath {
				continue
			}

			if isInside(line, x, y, m) {
				count++
			}
		}
	}

	return count
}

func (AOC) Day10_part1() int {
	lines := util.GetInput(10, false)
	s := getStart(lines)

	return getFarthests(s, lines) / 2
}

func (AOC) Day10_part2() int {
	lines := util.GetInput(10, false)
	s := getStart(lines)
	path := getPath(s, lines)

	return countInner(lines, path)
}
