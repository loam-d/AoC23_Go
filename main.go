package main

import (
	"aoc23_Go/solutions"
	"fmt"
)

func main() {
	fmt.Println("Advent of code 2023 in aoc23_Go\n")
	fmt.Println("Input [Day] to run or hit [Enter] to run most recent day:")
	var day string
	_, err := fmt.Scanln(&day)
	if err != nil {
		day = ""
	}
	switch day {
	case "1":
		fmt.Printf("day1:\n\t%d \n\t%d\n", solutions.Day1_1(), solutions.Day1_2())
	case "2":
	default:
		fmt.Printf("day2:\n\t%d \n\t%d\n", solutions.Day2_1(), solutions.Day2_2())
	}
}
