package solutions

import (
	"github.com/cauesmelo/aoc-2023/util"
)

type mCoord struct {
	xa int
	xb int
	ya int
	yb int
}

type cord struct {
	a int
	b int
}

func groupPatterns(lines []string) [][]string {
	patterns := make([][]string, 0)

	idx := 0
	for _, line := range lines {
		if line == "" {
			idx++
		} else {
			if len(patterns)-1 < idx {
				patterns = append(patterns, []string{})
			}

			patterns[idx] = append(patterns[idx], line)
		}
	}

	return patterns
}

func compareRows(pattern []string, ya int, yb int) bool {
	width := len(pattern[0])

	for x := 0; x < width; x++ {
		a := pattern[ya][x]
		b := pattern[yb][x]

		if a != b {
			return false
		}
	}

	return true
}

func scanPatternVertical(pattern []string, ignore *mCoord) *cord {
	height := len(pattern)

	for yStart := 0; yStart < height; yStart++ {
		ya := yStart
		yb := yStart + 1

		isMirror := true
		scans := 0

		for {
			if yb == height || !isMirror || ya < 0 {
				break
			}

			isMirror = compareRows(pattern, ya, yb)

			ya--
			yb++
			scans++
		}

		if scans == 0 {
			return nil
		}

		if isMirror {
			if ignore != nil && ignore.ya == yStart+1 && ignore.yb == yStart+2 {
				continue
			} else {
				return &cord{yStart + 1, yStart + 2}
			}
		}
	}

	return nil
}

func compareColumns(pattern []string, xa int, xb int) bool {
	height := len(pattern)

	for y := 0; y < height; y++ {
		a := pattern[y][xa]
		b := pattern[y][xb]

		if a != b {
			return false
		}
	}

	return true
}

func scanPatternHorizontal(pattern []string, ignore *mCoord) *cord {
	width := len(pattern[0])

	for xStart := 0; xStart < width; xStart++ {
		xa := xStart
		xb := xStart + 1

		isMirror := true
		scans := 0

		for {
			if xb == width || !isMirror || xa < 0 {
				break
			}

			isMirror = compareColumns(pattern, xa, xb)

			xa--
			xb++
			scans++
		}

		if scans == 0 {
			return nil
		}

		if isMirror {
			if ignore != nil && ignore.xa == xStart+1 && ignore.xb == xStart+2 {
				continue
			} else {
				return &cord{xStart + 1, xStart + 2}
			}
		}
	}

	return nil
}

func scanPattern(pattern []string, ignore *mCoord) (int, *mCoord) {
	cordH := scanPatternHorizontal(pattern, ignore)

	if cordH != nil {
		return cordH.a, &mCoord{cordH.a, cordH.b, 0, 0}
	}

	cordV := scanPatternVertical(pattern, ignore)
	if cordV != nil {
		return (100 * cordV.a), &mCoord{0, 0, cordV.a, cordV.b}
	}

	return 0, nil
}

func removeSmudge(pattern []string) int {
	_, coordToIgnore := scanPattern(pattern, nil)

	for y := 0; y < len(pattern); y++ {
		for x := 0; x < len(pattern[0]); x++ {
			if pattern[y][x] == '#' {
				pattern[y] = pattern[y][:x] + "." + pattern[y][x+1:]
				v, _ := scanPattern(pattern, coordToIgnore)

				if v != 0 {
					return v
				} else {
					pattern[y] = pattern[y][:x] + "#" + pattern[y][x+1:]
				}
			}
		}
	}

	return 0
}

func (AOC) Day13_part1() int {
	lines := util.GetInput(13, true)

	patterns := groupPatterns(lines)
	total := 0

	for _, pattern := range patterns {
		v, _ := scanPattern(pattern, nil)
		total += v
	}

	return total
}

func (AOC) Day13_part2() int {
	lines := util.GetInput(13, false)

	patterns := groupPatterns(lines)

	total := 0

	for _, pattern := range patterns {
		total += removeSmudge(pattern)
	}

	return total
}
