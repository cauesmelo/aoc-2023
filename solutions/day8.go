package solutions

import (
	"regexp"

	"github.com/cauesmelo/aoc-2023/util"
)

type node struct {
	left  string
	right string
}

func parseTree(lines []string) map[string]node {
	re := regexp.MustCompile("[A-Za-z]+")

	treeMap := make(map[string]node)

	for _, line := range lines {
		m := re.FindAllString(line, -1)

		nodeRoot := m[0]
		nodeLeft := m[1]
		nodeRight := m[2]

		treeMap[nodeRoot] = node{
			left:  nodeLeft,
			right: nodeRight,
		}
	}

	return treeMap
}

func moveUntilZZZ(moves string, m map[string]node) int {
	count := 0

	curr := "AAA"
	mvIdx := 0

	for curr != "ZZZ" {
		if mvIdx >= len(moves) {
			mvIdx = 0
		}

		if moves[mvIdx] == 'L' && m[curr].left != "" {
			curr = m[curr].left
		}

		if moves[mvIdx] == 'R' && m[curr].right != "" {
			curr = m[curr].right
		}

		mvIdx++
		count++
	}

	return count
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i], integers[i+1:])
	}

	return result
}

func moveUntilZZZSim(moves string, m map[string]node) int {
	count := 0

	ghosts := make([]string, 0)
	steps := make([]int, 0)

	for k := range m {
		if k[2] == 'A' {
			ghosts = append(ghosts, k)
		}
	}

	mvIdx := 0

	for {
		for i, g := range ghosts {
			if g[2] == 'Z' {
				ghosts = append(ghosts[:i], ghosts[i+1:]...)
				steps = append(steps, count)
			}
		}

		if len(ghosts) == 0 {
			break
		}

		if mvIdx >= len(moves) {
			mvIdx = 0
		}

		for i, curr := range ghosts {
			if moves[mvIdx] == 'L' && m[curr].left != "" {
				ghosts[i] = m[curr].left
			}

			if moves[mvIdx] == 'R' && m[curr].right != "" {
				ghosts[i] = m[curr].right
			}
		}

		mvIdx++
		count++
	}

	return LCM(steps[0], steps[1], steps[2:])
}

func (AOC) Day8_part1() int {
	lines := util.GetInput(8, false)
	tree := parseTree(lines[2:])
	moves := moveUntilZZZ(lines[0], tree)

	return moves
}

func (AOC) Day8_part2() int {
	lines := util.GetInput(8, false)

	tree := parseTree(lines[2:])
	moves := moveUntilZZZSim(lines[0], tree)

	return moves
}
