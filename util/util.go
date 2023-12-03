package util

import (
	"fmt"
	"os"
	"strings"
)

func GetInput(day int, test bool) []string {
	partN := 2
	if test {
		partN = 1
	}

	fileName := fmt.Sprintf("d%d_%d.txt", day, partN)

	dat, err := os.ReadFile("./input/" + fileName)
	Check(err)

	lines := strings.Split(string(dat), "\n")

	return lines
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
