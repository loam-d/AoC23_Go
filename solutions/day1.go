package solutions

import (
	"aoc23_Go/util"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func replaceNums(check string) string {
	check = strings.Replace(check, "one", "o1e", -1)
	check = strings.Replace(check, "two", "t2o", -1)
	check = strings.Replace(check, "three", "t3e", -1)
	check = strings.Replace(check, "four", "f4r", -1)
	check = strings.Replace(check, "five", "f5e", -1)
	check = strings.Replace(check, "six", "s6x", -1)
	check = strings.Replace(check, "seven", "s7n", -1)
	check = strings.Replace(check, "eight", "e8t", -1)
	check = strings.Replace(check, "nine", "n9e", -1)
	return check
}

func getSumOfCalibration(check_number_strings bool) int {
	reader, err := util.GetInputFileReader(1, 1)
	if err != nil {
		fmt.Print("PANIK day 1 1")
		return -1
	}
	defer reader.Close()

	sum := 0
	for reader.Next() {
		line := reader.Line()
		first := 'a'
		var last rune
		if check_number_strings {
			line = replaceNums(line)
		}
		for _, c := range line {
			if unicode.IsDigit(c) {
				if first == 'a' {
					first = c
				}
				last = c
			}
		}
		calibnum, _ := strconv.Atoi(fmt.Sprintf("%s%s", string(first), string(last)))
		sum += calibnum
	}

	return sum
}

func Day1_1() int {
	return getSumOfCalibration(false)
}

func Day1_2() int {
	return getSumOfCalibration(true)
}
