package main

import (
	"flag"
	"fmt"
	"reflect"

	"github.com/cauesmelo/aoc-2023/solutions"
)

const CURRENT_DAY = 13

func part(n int) int {
	aoc := reflect.ValueOf(solutions.AOC{})
	name := fmt.Sprintf("Day%d_part%d", CURRENT_DAY, n)
	m := aoc.MethodByName(name)

	return int(m.Call(nil)[0].Int())
}

func main() {
	day := flag.Int("d", 0, "Number of the day to run")
	flag.Parse()

	if *day < 1 || *day > CURRENT_DAY {
		*day = CURRENT_DAY
	}

	fmt.Printf("== Day %d ==\n", *day)
	fmt.Printf("Part 1 -> %d \nPart 2- > %d\n", part(1), part(2))
}
