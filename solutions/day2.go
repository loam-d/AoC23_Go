package solutions

import (
	"aoc23_Go/util"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	id     int
	colors []map[string]int
}

func isGamePossible(prospect map[string]int, test map[string]int) bool {
	for color, num_cubes := range prospect {
		if num_cubes > test[color] {
			return false
		}
	}
	return true
}

func calcPower(colors []map[string]int) int {
	maxColor := make(map[string]int)
	for _, bag := range colors {
		for color, cnt := range bag {
			maxCnt, ok := maxColor[color]
			if !ok || ok && maxCnt < cnt {
				maxColor[color] = cnt
			}
		}
	}

	pow := 1
	for _, cnt := range maxColor {
		pow *= cnt
	}
	return pow
}

func readGameLine(line string) Game {
	res1 := strings.Split(line, ": ")
	gameId, _ := strconv.Atoi(strings.Split(res1[0], " ")[1])
	colors := make([]map[string]int, 0)
	for _, play := range strings.Split(res1[1], "; ") {
		cubes := make(map[string]int)
		for _, cube := range strings.Split(play, ", ") {
			res2 := strings.Split(cube, " ")
			quant, _ := strconv.Atoi(res2[0])
			color := res2[1]
			cubes[color] = quant
		}
		colors = append(colors, cubes)
	}
	return Game{id: gameId, colors: colors}
}

func processGameLines(getPower bool) int {
	reader, err := util.GetInputFileReader(2, 1)
	if err != nil {
		fmt.Print("PANIK day 2 1")
		return -1
	}
	defer reader.Close()

	sum := 0
	bag := make(map[string]int)

	bag["red"] = 12
	bag["green"] = 13
	bag["blue"] = 14
ALL_BAGS:
	for reader.Next() {
		game := readGameLine(reader.Line())
		if !getPower {
			for _, color := range game.colors {
				if !isGamePossible(color, bag) {
					continue ALL_BAGS
				}
			}
			sum += game.id
		} else {
			sum += calcPower(game.colors)
		}
	}

	return sum
}

func Day2_1() int {
	return processGameLines(false)
}

func Day2_2() int {
	return processGameLines(true)
}
