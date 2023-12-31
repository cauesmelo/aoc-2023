package solutions

import (
	"sort"
	"strconv"
	"strings"

	"github.com/cauesmelo/aoc-2023/util"
)

type handType uint8

const (
	HighCard handType = iota + 1
	OnePair
	TwoPair
	ThreeOf
	FullHouse
	FourOf
	FiveOf
)

var cardValue = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

var cardValue2 = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type bet struct {
	cards    string
	bet      int
	handType handType
	rank     int
}

func parseBets(lines []string) []bet {
	bets := make([]bet, 0)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		cards := parts[0]
		b, err := strconv.Atoi(parts[1])
		util.Check(err)

		newBet := bet{
			cards:    cards,
			bet:      b,
			handType: 0,
			rank:     0,
		}

		bets = append(bets, newBet)
	}

	return bets
}

func getHandPairs(bets []bet, joker bool) []bet {
	type pos struct {
		char rune
		occ  int
	}

	for idx := range bets {
		arr := [5]pos{}

		for _, let := range bets[idx].cards {

			for i := range arr {
				if arr[i].occ == 0 {
					arr[i].char = let
					arr[i].occ = 1
					break
				}

				if arr[i].char == let {
					arr[i].occ++
					break
				}

			}
		}

		sort.Slice(arr[:], func(i, j int) bool {
			return arr[i].occ > arr[j].occ
		})

		if joker {
			jokers := 0

			for x := range arr {
				if arr[x].char == 'J' {
					jokers = arr[x].occ
					arr[x].occ = 0
				}
			}

			if arr[0].char != 'J' {
				arr[0].occ = arr[0].occ + jokers
			} else {
				arr[1].occ = arr[1].occ + jokers
			}

			sort.Slice(arr[:], func(i, j int) bool {
				return arr[i].occ > arr[j].occ
			})
		}

		if arr[4].occ == 1 {
			bets[idx].handType = HighCard
			continue
		}

		if arr[3].occ == 1 {
			bets[idx].handType = OnePair
			continue
		}

		if arr[0].occ == 3 {
			if arr[1].occ == 2 {
				bets[idx].handType = FullHouse
				continue
			}

			bets[idx].handType = ThreeOf
			continue
		}

		if arr[0].occ == 2 && arr[1].occ == 2 {
			bets[idx].handType = TwoPair
			continue
		}

		if arr[0].occ == 4 {
			bets[idx].handType = FourOf
			continue
		}

		bets[idx].handType = FiveOf
	}

	return bets
}

func compareCards(a string, b string, v2 bool) bool {
	for i := range a {
		if a[i] == b[i] {
			continue
		}

		if v2 {
			return cardValue2[rune(a[i])] > cardValue2[rune(b[i])]
		}

		return cardValue[rune(a[i])] > cardValue[rune(b[i])]

	}

	return false
}

func getRanks(bets []bet, v2 bool) []bet {
	sort.Slice(bets, func(i, j int) bool {
		if bets[i].handType != bets[j].handType {
			return bets[i].handType > bets[j].handType
		}

		return compareCards(bets[i].cards, bets[j].cards, v2)
	})

	rank := len(bets)

	for i := range bets {
		bets[i].rank = rank
		rank--
	}

	return bets
}

func getTotal(bets []bet) int {
	sum := 0

	for _, bet := range bets {
		sum = sum + bet.rank*bet.bet
	}

	return sum
}

func (AOC) Day7_part1() int {
	lines := util.GetInput(7, false)

	bets := parseBets(lines)
	bets = getHandPairs(bets, false)
	bets = getRanks(bets, false)

	return getTotal(bets)
}

func (AOC) Day7_part2() int {
	lines := util.GetInput(7, false)

	bets := parseBets(lines)
	bets = getHandPairs(bets, true)
	bets = getRanks(bets, true)

	return getTotal(bets)
}
