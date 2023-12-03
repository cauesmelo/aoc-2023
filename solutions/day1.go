package solutions

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/cauesmelo/aoc-2023/util"
)

type lineData struct {
	Num int
	Pos int
}

func parseSpelled(slice string) int {
	sliceLen := len(slice)

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for pos, n := range numbers {
		numberLen := len(n)

		if sliceLen >= numberLen && slice[0:numberLen] == n {
			return pos + 1
		}
	}

	return 0
}

func parseLine(line string) []lineData {
	ld := make([]lineData, 0)

	for pos, r := range line {
		char := string(r)

		if isNumber(r) {
			num, err := strconv.Atoi(char)
			check(err)

			ld = append(ld, lineData{
				Num: num,
				Pos: pos,
			})
		} else {
			res := parseSpelled(line[pos:])

			if res > 0 {
				ld = append(ld, lineData{
					Num: res,
					Pos: pos,
				})
			}
		}
	}

	return ld
}

func isNumber(r rune) bool {
	if r > 48 && r < 58 {
		return true
	}

	return false
}

func calcLine_1(line string) int {
	first := -1
	last := -1

	for _, char := range line {
		if isNumber(char) {
			n, err := strconv.Atoi(string(char))
			util.Check(err)

			last = n

			if first == -1 {
				first = n
			}
		}
	}

	if last == -1 {
		last = first
	}

	str := fmt.Sprintf("%d%d", first, last)

	finalN, err := strconv.Atoi(str)
	util.Check(err)

	return finalN
}

func Day1_part1() int {
	lines := util.GetInput(1, false)

	total := 0

	for _, line := range lines {
		total = total + calcLine_1(line)
	}

	return total
}

func calcLine_2(line []lineData) int {
	first := -1
	last := -1

	first = line[0].Num
	last = line[len(line)-1].Num

	str := fmt.Sprintf("%d%d", first, last)

	finalN, err := strconv.Atoi(str)
	check(err)

	return finalN
}

func Day1_part2() int {

	lines := util.GetInput(1, false)

	parsedLines := make(chan []lineData, len(lines))
	var wg sync.WaitGroup

	for _, line := range lines {
		wg.Add(1)

		go func(l string) {
			defer wg.Done()

			parsedLines <- parseLine(l)
		}(line)
	}

	go func() {
		wg.Wait()
		close(parsedLines)
	}()

	total := 0

	for line := range parsedLines {
		total = total + calcLine_2(line)
	}

	return total
}
