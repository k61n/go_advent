package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
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

func parsePath(lines []string) string {
	return lines[0]
}

func parseInstructions(lines []string) map[string][]string {
	res := make(map[string][]string)
	for _, line := range lines[2:] {
		sub := strings.Split(line, " = (")
		left := strings.Split(sub[1], ", ")[0]
		right := strings.Split(sub[1], ", ")[1][:3]
		res[sub[0]] = append(res[sub[0]], left, right)
	}
	return res
}

func findStartingNodes(lines []string) []string {
	var res []string
	for _, line := range lines[2:] {
		node := strings.Split(line, " = (")[0]
		if string(node[2]) == "A" {
			res = append(res, node)
		}
	}
	return res
}

func solveSingle(cur string, end string, lines []string) int {
	dir := make(map[string]int)
	dir["L"] = 0
	dir["R"] = 1
	instr := parseInstructions(lines)
	step := 0
	for true {
		for _, i := range parsePath(lines) {
			cur = instr[cur][dir[string(i)]]
			step += 1
		}
		if end == "ZZZ" {
			if cur == end {
				break
			}
		} else {
			if string(cur[2]) == "Z" {
				break
			}
		}
	}
	return step
}

func checkEndCondition(nodes []string) bool {
	res := true
	for _, node := range nodes {
		if string(node[2]) != "Z" {
			res = false
			break
		}
	}
	return res
}

func main() {
	lines := readInputFile("08.txt")

	// Part 1
	fmt.Println(solveSingle("AAA", "ZZZ", lines))

	// Part 2
	nodes := findStartingNodes(lines)
	var nums []int
	for _, node := range nodes {
		nums = append(nums, solveSingle(node, "**Z", lines))
	}
	var lcm big.Int
	lcm.SetInt64(int64(nums[0]))
	for i := 1; i < len(nums); i++ {
		num := big.NewInt(int64(nums[i]))
		gcd := new(big.Int).GCD(nil, nil, &lcm, num)
		lcm.Mul(&lcm, num)
		lcm.Div(&lcm, gcd)
	}
	fmt.Println(lcm.String())
}
