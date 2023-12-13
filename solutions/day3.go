package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"

	"github.com/cauesmelo/aoc-2023/util"
)

type number struct {
	line       int
	value      int
	start      int
	end        int
	partNumber bool
}

func scanNumbers(line string, lineNumber int) []number {
	re := regexp.MustCompile(`\d+`)

	f := re.FindAllStringIndex(line, -1)

	numbers := make([]number, 0)

	for _, match := range f {
		valueStr := line[match[0]:match[1]]
		val, err := strconv.Atoi(valueStr)
		util.Check(err)

		n := number{
			line:       lineNumber,
			value:      val,
			start:      match[0],
			end:        match[1] - 1,
			partNumber: false,
		}

		numbers = append(numbers, n)
	}

	return numbers
}

func scanSymbols(line string) []int {
	pos := make([]int, 0)

	for p, char := range line {
		if !unicode.IsDigit(char) && char != '.' {
			pos = append(pos, p)
		}
	}

	return pos
}

func scanLines(linesStr []string) ([][]number, [][]int) {
	numbers := make([][]number, len(linesStr))
	symbols := make([][]int, len(linesStr))

	for i, line := range linesStr {
		numbers[i] = scanNumbers(line, i)
		symbols[i] = scanSymbols(line)
	}

	return numbers, symbols
}

func findAdj(numbers []number, symbolPos []int) []int {
	adj := make([]int, 0)

	fmt.Println("Numbers: ", numbers)
	fmt.Println("Symbols: ", symbolPos)

	for _, symbol := range symbolPos {
		for _, n := range numbers {
			if symbol >= n.start && symbol <= n.end {
				adj = append(adj, n.value)
			}
		}
	}

	fmt.Println("Adj: ", adj)

	return adj
}

func matchCurrLine(numbers []number, symbols []int) []int {
	values := make([]int, 0)

	for _, symbolPos := range symbols {
		for i, num := range numbers {
			if num.end == symbolPos-1 || num.start == symbolPos+1 {
				numbers[i].partNumber = true
				values = append(values, num.value)
			}
		}
	}

	return values
}

func matchAdjLine(numbers []number, symbols []int) []int {
	values := make([]int, 0)

	for _, symbolPos := range symbols {
		for i, num := range numbers {
			if symbolPos >= num.start-1 && symbolPos <= num.end+1 {
				numbers[i].partNumber = true
				values = append(values, num.value)
			}
		}
	}

	return values
}

func sumPartNumbers(numbers [][]number) int {
	sum := 0

	for _, line := range numbers {
		for _, n := range line {
			if n.partNumber {
				sum = sum + n.value
			}
		}
	}

	return sum
}

func (AOC) Day3_part1() int {
	linesStr := util.GetInput(3, false)

	numbers, symbols := scanLines(linesStr)

	for i := range linesStr {
		matchCurrLine(numbers[i], symbols[i])

		if i != 0 {
			matchAdjLine(numbers[i-1], symbols[i])
		}

		if i != len(linesStr)-1 {
			matchAdjLine(numbers[i+1], symbols[i])
		}
	}

	return sumPartNumbers(numbers)
}

func filterPossibleGears(linesStr []string, symbols [][]int) [][]int {
	gears := make([][]int, len(linesStr))

	for i := range symbols {
		for _, j := range symbols[i] {

			if linesStr[i][j] == '*' {
				gears[i] = append(gears[i], j)
			}
		}
	}

	return gears
}

func calcGears(prev []number, curr []number, next []number, gears []int) int {
	sum := 0

	for _, gear := range gears {
		fmt.Println("Cur gear pos: ", gear)
		parts := matchCurrLine(curr, []int{gear})
		parts = append(parts, matchAdjLine(prev, []int{gear})...)
		parts = append(parts, matchAdjLine(next, []int{gear})...)

		if len(parts) != 2 {
			continue
		}

		sum = sum + parts[0]*parts[1]
	}

	return sum
}

func (AOC) Day3_part2() int {
	linesStr := util.GetInput(3, false)

	numbers, symbols := scanLines(linesStr)
	gears := filterPossibleGears(linesStr, symbols)

	total := 0

	for i := range gears {
		var prev []number
		curr := numbers[i]
		var next []number

		if i != 0 {
			prev = numbers[i-1]
		}

		if i != len(linesStr)-1 {
			next = numbers[i+1]
		}

		total = total + calcGears(prev, curr, next, gears[i])
	}

	return total
}
