package util

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetHashForSliceInt(list []int) [32]byte {
	var buffer bytes.Buffer

	for i := range list {
		buffer.WriteString(strconv.Itoa(list[i]))
		buffer.WriteString("0")
	}
	return sha256.Sum256(buffer.Bytes())
}

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func GetNumbers(line string) []int {
	re := regexp.MustCompile(`-?\d+`)

	f := re.FindAllStringIndex(line, -1)

	numbers := make([]int, 0)

	for _, match := range f {
		valueStr := line[match[0]:match[1]]
		val, err := strconv.Atoi(valueStr)
		Check(err)

		numbers = append(numbers, val)
	}

	return numbers
}

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
