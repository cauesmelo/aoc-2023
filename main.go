package main

import (
	"flag"
	"fmt"

	"github.com/cauesmelo/aoc-2023/solutions"
)

const LAST_DAY = 4

func main() {
	day := flag.Int("d", 0, "Number of the day to run")
	flag.Parse()

	if *day < 1 || *day > LAST_DAY {
		*day = LAST_DAY
	}

	// TODO: simplify logic here

	switch *day {
	case 1:
		fmt.Println("== Day 1 ==")
		fmt.Printf("Part 1 -> %d \nPart 2- > %d\n", solutions.Day1_part1(), solutions.Day1_part2())

	case 2:
		fmt.Println("== Day 2 ==")
		fmt.Printf("Part 1 -> %d \nPart 2- > %d\n", solutions.Day2_part1(), solutions.Day2_part2())

	case 3:
		fmt.Println("== Day 3 ==")
		fmt.Printf("Part 1 -> %d \nPart 2- > %d\n", solutions.Day3_part1(), solutions.Day3_part2())

	case 4:
		fmt.Println("== Day 4 ==")
		fmt.Printf("Part 1 -> %d \nPart 2- > %d\n", solutions.Day4_part1(), solutions.Day4_part2())
	}
}
