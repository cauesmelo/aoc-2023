package solutions

import (
	"github.com/cauesmelo/aoc-2023/util"
)

type race struct {
	time int
	dist int
}

func getRaces(lines []string) []race {
	times := getNumbers(lines[0])
	dists := getNumbers(lines[1])

	races := make([]race, 0)

	for i := range times {
		races = append(races, race{
			time: times[i],
			dist: dists[i],
		})
	}

	return races
}

func getWaysToBeat(race race) int {
	holdFor := 0

	ways := 0

	for holdFor <= race.time {
		timeLeft := race.time - holdFor
		res := timeLeft * holdFor

		if res > race.dist {
			ways++
		}

		holdFor++
	}

	return ways
}

func Day6_part1() int {
	lines := util.GetInput(6, false)

	races := getRaces(lines)

	res := 0

	for _, race := range races {
		ways := getWaysToBeat(race)

		if res == 0 {
			res = ways
			continue
		}

		res = res * ways
	}

	return res
}

func Day6_part2() int {
	// _ := util.GetInput(6, true)

	return 0
}
