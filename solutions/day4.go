package solutions

import (
	"aoc23_Go/util"
	"strconv"
	"strings"
)

func processScratchCard(line string) ([]int, map[int]interface{}) {
	ourNums := make([]int, 0)
	winningNums := make(map[int]interface{})

	toProcess := strings.Split(line, ": ")[1]
	numbersStrings := strings.Split(toProcess, " | ")

	for _, num := range strings.Split(numbersStrings[1], " ") {
		if num == "" {
			continue
		}
		num_, _ := strconv.Atoi(num)
		ourNums = append(ourNums, num_)
	}

	for _, num := range strings.Split(numbersStrings[0], " ") {
		if num == "" {
			continue
		}
		num_, _ := strconv.Atoi(num)
		winningNums[num_] = new(interface{})
	}

	return ourNums, winningNums
}

func processAllCards(doCopies bool) int {
	reader, _ := util.GetInputFileReader(4, 1)
	defer reader.Close()

	score := 0
	copies := make([]int, 1)
	copies[0] = 1
	copyIdx := 0

	for reader.Next() {
		ourNums, winningNums := processScratchCard(reader.Line())
		cardScore := 0
		for _, num := range ourNums {
			_, isWinning := winningNums[num]
			if isWinning && !doCopies {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore *= 2
				}
			} else if isWinning {
				cardScore += 1
			}
		}
		if !doCopies {
			score += cardScore
		} else {
			for len(copies) < copyIdx+cardScore+1 {
				copies = append(copies, 1)
			}
			for j := 0; j < cardScore; j++ {
				copies[copyIdx+1+j] += copies[copyIdx]
			}
			copyIdx += 1
		}
	}

	if doCopies {
		return sumOfNums(copies)
	}
	return score
}

func Day4_1() int {
	return processAllCards(false)
}

func Day4_2() int {
	return processAllCards(true)
}
