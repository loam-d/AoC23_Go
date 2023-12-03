package solutions

import (
	"aoc23_Go/util"
	"strconv"
	"unicode"
)

type EngineSchema struct {
	number_idxs map[util.Coordinate]int
	numbers     []int
	symbols     map[util.Coordinate]interface{}
	gears       map[util.Coordinate]interface{}
}

func adjCoords(coord util.Coordinate) []util.Coordinate {
	ret := make([]util.Coordinate, 0)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			ret = append(ret, util.Coordinate{coord.X + i, coord.Y + j})
		}
	}

	return ret
}

func findAdjNumsForCoord(coord util.Coordinate, engine *EngineSchema) []int {
	nums := make([]int, 0)
	poisoned_idxs := make(map[int]interface{})

	adj_nums := make(map[int]interface{})
	for _, adj_coord := range adjCoords(coord) {
		idx, ok := engine.number_idxs[adj_coord]
		if ok {
			adj_nums[idx] = new(interface{})
		}
	}
	for idx, _ := range adj_nums {
		_, ok := poisoned_idxs[idx]
		if !ok {
			nums = append(nums, engine.numbers[idx])
			poisoned_idxs[idx] = new(interface{})
		}
	}

	return nums
}

func findAdjNums(engine *EngineSchema) []int {
	nums := make([]int, 0)
	poisoned_idxs := make(map[int]interface{})
	for coord, _ := range engine.symbols {
		adj_nums := make(map[int]interface{})
		for _, adj_coord := range adjCoords(coord) {
			idx, ok := engine.number_idxs[adj_coord]
			if ok {
				adj_nums[idx] = new(interface{})
			}
		}
		for idx, _ := range adj_nums {
			_, ok := poisoned_idxs[idx]
			if !ok {
				nums = append(nums, engine.numbers[idx])
				poisoned_idxs[idx] = new(interface{})
			}
		}
	}

	return nums
}

func sumOfNums(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func findValidGearRatios(engine *EngineSchema) []int {
	ratios := make([]int, 0)
	for gear, _ := range engine.gears {
		nums := findAdjNumsForCoord(gear, engine)
		if len(nums) == 2 {
			ratios = append(ratios, nums[0]*nums[1])
		}
	}
	return ratios
}

func processEngineSchematic() *EngineSchema {
	reader, _ := util.GetInputFileReader(3, 1)
	defer reader.Close()

	engine := new(EngineSchema)
	engine.number_idxs = make(map[util.Coordinate]int)
	engine.numbers = make([]int, 0)
	engine.symbols = make(map[util.Coordinate]interface{})
	engine.gears = make(map[util.Coordinate]interface{})
	curr_y := 0
	for reader.Next() {
		line := reader.Line()
		num := ""
		num_coords := make([]util.Coordinate, 0)
		for curr_x, c := range line {
			if unicode.IsDigit(c) {
				num = num + string(c)
				num_coords = append(num_coords, util.Coordinate{curr_x, curr_y})
			} else {
				if c != '.' {
					engine.symbols[util.Coordinate{curr_x, curr_y}] = new(interface{})
				}
				if c == '*' {
					engine.gears[util.Coordinate{curr_x, curr_y}] = new(interface{})
				}
				if len(num) > 0 {
					true_num, _ := strconv.Atoi(num)
					engine.numbers = append(engine.numbers, true_num)
					for _, coord := range num_coords {
						engine.number_idxs[coord] = len(engine.numbers) - 1
					}
					num = ""
					num_coords = make([]util.Coordinate, 0)
				}
			}
		}
		if len(num) > 0 {
			true_num, _ := strconv.Atoi(num)
			engine.numbers = append(engine.numbers, true_num)
			for _, coord := range num_coords {
				engine.number_idxs[coord] = len(engine.numbers) - 1
			}
			num = ""
			num_coords = make([]util.Coordinate, 0)
		}
		curr_y += 1
	}

	return engine
}

func Day3_1() int {
	engine := processEngineSchematic()

	return sumOfNums(findAdjNums(engine))
}

func Day3_2() int {
	engine := processEngineSchematic()

	return sumOfNums(findValidGearRatios(engine))
}
