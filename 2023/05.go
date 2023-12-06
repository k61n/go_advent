package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseSeeds(lines []string) []int {
	subs := strings.Split(lines[0], "seeds: ")
	var res []int
	for _, numStr := range strings.Split(subs[1], " ") {
		numInt, err := strconv.Atoi(numStr)
		if err == nil {
			res = append(res, numInt)
		}
	}
	return res
}

func parseArray(lines []string) [][]int {
	var res [][]int
	for i, line := range lines {
		res = append(res, []int{})
		for _, numStr := range strings.Split(line, " ") {
			numInt, _ := strconv.Atoi(numStr)
			res[i] = append(res[i], numInt)
		}
	}
	return res
}

func parseMaps(lines []string) [][][]int {
	var res [][][]int
	var rows []int
	for j, line := range lines {
		if strings.Contains(line, "map:") {
			rows = append(rows, j)
		}
	}
	for i := 0; i < len(rows); i++ {
		if i != len(rows)-1 {
			res = append(res, parseArray(lines[rows[i]+1:rows[i+1]-1]))
		} else {
			res = append(res, parseArray(lines[rows[i]+1:]))
		}
	}
	return res
}

func seedToTarget(seed int, arr [][]int) int {
	for _, subarr := range arr {
		if subarr[1] <= seed && seed < subarr[1]+subarr[2] {
			seed = seed - subarr[1] + subarr[0]
			break
		}
	}
	return seed
}

func processSeeds(seeds []int, arr [][]int) []int {
	for i, _ := range seeds {
		seeds[i] = seedToTarget(seeds[i], arr)
	}
	return seeds
}

func findMin(arr []int) int {
	_min := math.MaxInt
	for _, num := range arr {
		if num < _min {
			_min = num
		}
	}
	return _min
}

func main() {
	exe, _ := os.Executable()
	dir := filepath.Dir(exe)
	input := "05.txt"

	file, _ := os.Open(filepath.Join(dir, input))
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Part 1
	seeds := parseSeeds(lines)
	maps := parseMaps(lines)
	for _, arr := range maps {
		seeds = processSeeds(seeds, arr)
	}
	fmt.Println(findMin(seeds))

	// Part 2
	seeds = parseSeeds(lines)
	var newseeds []int
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			newseeds = append(newseeds, j)
		}
	}
	for _, arr := range maps {
		newseeds = processSeeds(newseeds, arr)
	}
	fmt.Println(findMin(newseeds))
}
