package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readInputFile(filename string) []string {
	exe, _ := os.Executable()
	dir := filepath.Dir(exe)

	file, _ := os.Open(filepath.Join(dir, filename))
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func readLine(line, key string) []int {
	line = strings.Split(line, key)[1]
	numbers := strings.Split(line, " ")
	var nums []int
	for _, numStr := range numbers {
		numInt, err := strconv.Atoi(numStr)
		if err == nil {
			nums = append(nums, numInt)
		}
	}
	return nums
}

func time2distance(holdTime, totalTime int) int {
	return (totalTime - holdTime) * holdTime
}

func readLine2(line, key string) int {
	line = strings.Split(line, key)[1]
	numbers := strings.Split(line, " ")
	number := ""
	for _, numStr := range numbers {
		number += numStr
	}
	num, _ := strconv.Atoi(number)
	return num
}

func main() {
	lines := readInputFile("06.txt")

	// Part 1
	times := readLine(lines[0], "Time:")
	records := readLine(lines[1], "Distance:")
	var res []int
	for i, t := range times {
		res = append(res, 0)
		for j := 0; j < t; j++ {
			d := time2distance(j, t)
			if d > records[i] {
				res[i] += 1
			}
		}
	}
	f := 1
	for _, n := range res {
		f *= n
	}
	fmt.Println(f)

	// Part 2
	times2 := readLine2(lines[0], "Time:")
	records2 := readLine2(lines[1], "Distance:")
	res2 := 0
	for j := 0; j < times2; j++ {
		d := time2distance(j, times2)
		if d > records2 {
			res2 += 1
		}
	}
	fmt.Println(res2)
}
