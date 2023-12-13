package main

import (
	"flag"
	"fmt"

	"github.com/cauesmelo/aoc-2023/solutions"
)

const LAST_DAY = 9

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

	case 5:
		fmt.Println("== Day 5 ==")
		fmt.Printf("Part 1 -> %d \nPart 2- > %d\n", solutions.Day5_part1(), solutions.Day5_part2())

	case 6:
		fmt.Println("== Day 6 ==")
		fmt.Printf("Part 1 -> %d \nPart 2- > %d\n", solutions.Day6_part1(), solutions.Day6_part2())

	case 7:
		fmt.Println("== Day 7 ==")
		fmt.Printf("Part 1 -> %d \nPart 2- > %d\n", solutions.Day7_part1(), solutions.Day7_part2())

	case 8:
		fmt.Println("== Day 8 ==")
		fmt.Printf("Part 1 -> %d \nPart 2- > %d\n", solutions.Day8_part1(), solutions.Day8_part2())

	case 9:
		fmt.Println("== Day 9 ==")
		fmt.Printf("Part 1 -> %d \nPart 2- > %d\n", solutions.Day9_part1(), solutions.Day9_part2())
	}
}
