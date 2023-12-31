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
		fmt.Printf("day2:\n\t%d \n\t%d\n", solutions.Day2_1(), solutions.Day2_2())
	case "3":
		fmt.Printf("day3:\n\t%d \n\t%d\n", solutions.Day3_1(), solutions.Day3_2())
	case "4":
		fmt.Printf("day4:\n\t%d \n\t%d\n", solutions.Day4_1(), solutions.Day4_2())
	case "5":
	default:
		fmt.Printf("day5:\n\t%d \n\t%d\n", solutions.Day5_1(), solutions.Day5_2())
	}
}
