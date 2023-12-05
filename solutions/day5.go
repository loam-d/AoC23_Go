package solutions

import (
	"aoc23_Go/util"
	"strconv"
	"strings"
	"sync"
)

type SeedMap struct {
	seeds     []int
	locations [][]rangeFunc
}

func seedToLoc(seed int, locations [][]rangeFunc) int {
	currSrc := seed
	for _, locationMap := range locations {
		currSrc = destFromRangeFuncs(currSrc, locationMap)
	}
	return currSrc
}

func getLowestLocation(map_ SeedMap, sickoMode bool) int {
	var minloc int

	if sickoMode {
		wg := sync.WaitGroup{}
		minChan := make(chan int)
		maxParallel := 10
		semaphore := make(chan struct{}, maxParallel)
		for i := 1; i < len(map_.seeds); i += 2 {
			wg.Add(1)
			go func(i int) {
				semaphore <- struct{}{}
				println("Start Go routine")
				var threadMin int
				for seed := 0; seed < map_.seeds[i]; seed++ {
					currSrc := seedToLoc(seed+map_.seeds[i-1], map_.locations)
					if seed == 0 {
						threadMin = currSrc
					} else {
						threadMin = min(minloc, currSrc)
					}
				}
				println(threadMin)
				minChan <- threadMin
				wg.Done()
				<-semaphore
			}(i)
		}
		wg.Wait()
		minloc = <-minChan
		for len(minChan) > 0 {
			minloc = min(minloc, <-minChan)
		}
	} else {
		for i, seed := range map_.seeds {
			loc := seedToLoc(seed, map_.locations)
			if i == 0 {
				minloc = loc
			} else {
				minloc = min(minloc, loc)
			}
		}
	}

	return minloc
}

type rangeFunc func(int) (bool, int, int)

func destFromRangeFuncs(src int, funcs []rangeFunc) int {
	for _, fn := range funcs {
		inrng, srcStart, destStart := fn(src)
		if inrng {
			return destStart + (src - srcStart)
		}
	}

	// This should be unreachable but I don;t want to set up errors
	return -1
}

func processMap() SeedMap {
	reader, _ := util.GetInputFileReader(5, 1)
	defer reader.Close()

	seeds := make([]int, 0)
	locations := make([][]rangeFunc, 0)
	locationsIdx := -1

	for reader.Next() {
		line := reader.Line()
		if strings.Contains(line, "seeds") {
			seedstr := strings.Split(line, " ")[1:]
			for _, seed := range seedstr {
				seednum, _ := strconv.Atoi(seed)
				seeds = append(seeds, seednum)
			}
		} else if strings.Contains(line, ":") {
			locations = append(locations, make([]rangeFunc, 0))
			locationsIdx += 1
		} else if len(line) > 0 {
			locationstr := strings.Split(line, " ")
			locationnums := make([]int, 0)
			for _, num := range locationstr {
				numm, _ := strconv.Atoi(num)
				locationnums = append(locationnums, numm)
			}
			dest := locationnums[0]
			src := locationnums[1]
			rng := locationnums[2]

			locations[locationsIdx] = append(locations[locationsIdx], func(srcd int) (bool, int, int) {
				return srcd >= src && srcd < src+rng, src, dest
			})
		} else if locationsIdx > -1 {
			// add all the unmapped ones, using an identity function
			locations[locationsIdx] = append(locations[locationsIdx], func(srcd int) (bool, int, int) {
				return true, srcd, srcd
			})
		}
	}
	return SeedMap{seeds: seeds, locations: locations}
}

func Day5_1() int {
	map_ := processMap()
	return getLowestLocation(map_, false)
}

func Day5_2() int {
	map_ := processMap()
	return getLowestLocation(map_, true)
}
