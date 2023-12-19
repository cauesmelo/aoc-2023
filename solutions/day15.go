package solutions

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/cauesmelo/aoc-2023/util"
)

type valOrder struct {
	val   int
	order int
}

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

func getHASHMAP(line string) int {
	cmds := strings.Split(line, ",")
	boxes := make([]map[string]valOrder, 256)
	order := 0

	re1 := regexp.MustCompile(`([a-z]+)`)
	re2 := regexp.MustCompile(`(=|-)`)
	re3 := regexp.MustCompile(`(\d+)`)

	for i := range boxes {
		boxes[i] = make(map[string]valOrder)
	}

	for _, cmd := range cmds {
		order++

		match1 := re1.FindString(cmd)
		match2 := re2.FindString(cmd)
		match3 := re3.FindString(cmd)

		label := match1
		posY := getHASH(label)
		action := rune(match2[0])

		if action == '-' {
			delete(boxes[posY], label)
			continue
		}

		if action == '=' {
			n, err := strconv.Atoi(match3)
			util.Check(err)

			_, exists := boxes[posY][label]
			if exists {
				boxes[posY][label] = valOrder{n, boxes[posY][label].order}
			} else {
				boxes[posY][label] = valOrder{n, order}
			}
		}
	}

	total := 0
	x := 0
	for y, box := range boxes {
		values := make([]valOrder, 0)

		for _, val := range box {
			values = append(values, val)
		}

		sort.Slice(values, func(i, j int) bool {
			return values[i].order < values[j].order
		})

		x = 0
		for _, val := range values {
			total += (y + 1) * (x + 1) * val.val
			x++
		}

	}

	return total
}

func (AOC) Day15_part1() int {
	lines := util.GetInput(15, false)
	l := lines[0]

	return getHASH(l)
}

func (AOC) Day15_part2() int {
	lines := util.GetInput(15, false)
	l := lines[0]

	return getHASHMAP(l)
}
