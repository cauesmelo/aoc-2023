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

		if moves[mvIdx] == 'R' && m[curr].left != "" {
			curr = m[curr].right
		}

		mvIdx++
		count++
	}

	return count
}

func Day8_part1() int {
	lines := util.GetInput(8, false)

	tree := parseTree(lines[2:])

	moves := moveUntilZZZ(lines[0], tree)

	return moves
}

func Day8_part2() int {
	_ = util.GetInput(8, false)

	return 0
}
