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

func parseSequence(line string) []int {
	var res []int
	for _, l := range strings.Split(line, " ") {
		num, _ := strconv.Atoi(l)
		res = append(res, num)
	}
	return res
}

func reduceSequence(arr []int) []int {
	var res []int
	for i, n := range arr {
		if i > 0 {
			res = append(res, n-arr[i-1])
		}
	}
	return res
}

func allZero(slice []int) bool {
	res := true
	for _, n := range slice {
		if n != 0 {
			res = false
		}
	}
	return res
}

func sumSlice(slice []int) int {
	sum := 0
	for _, n := range slice {
		sum += n
	}
	return sum
}

func extrapolateNegative(slice []int) int {
	val := slice[len(slice)-1]
	for i := len(slice) - 2; i >= 0; i-- {
		val = slice[i] - val
	}
	return val
}

func main() {
	lines := readInputFile("09.txt")

	// Part 1&2
	sum1 := 0
	sum2 := 0
	for _, line := range lines {
		seq := parseSequence(line)
		var arr1 []int
		var arr2 []int
		for !allZero(seq) {
			arr1 = append(arr1, seq[len(seq)-1])
			arr2 = append(arr2, seq[0])
			seq = reduceSequence(seq)
		}
		sum1 += sumSlice(arr1)
		sum2 += extrapolateNegative(arr2)
	}
	fmt.Println(sum1, sum2)
}
