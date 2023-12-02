package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getGameID(game string) int {
	parts := strings.Split(game, "Game ")
	num, _ := strconv.Atoi(parts[1])
	return num
}

func extractCubes(subset string, color string) int {
	parts := strings.Split(subset, ", ")
	num := 0
	for _, part := range parts {
		if strings.Contains(part, color) {
			res := strings.Split(part, " "+color)
			num, _ = strconv.Atoi(res[0])
			break
		}
	}
	return num
}

func condition(game map[string]int, maxcubes map[string]int) bool {
	ans := true
	for color, num := range game {
		if num > maxcubes[color] {
			ans = false
			break
		}
	}
	return ans
}

func powerOfSet(game map[string]int) int {
	res := 1
	for _, num := range game {
		res *= num
	}
	return res
}

func main() {
	exe, _ := os.Executable()
	dir := filepath.Dir(exe)
	input := "02.txt"

	file, _ := os.Open(filepath.Join(dir, input))
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		//fmt.Println(lines)
	}

	gamelist := make(map[int]map[string]int)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		id := getGameID(parts[0])
		gamelist[id] = make(map[string]int)
		gamelist[id]["blue"] = 0
		gamelist[id]["green"] = 0
		gamelist[id]["red"] = 0
		subsets := strings.Split(parts[1], "; ")
		for _, subset := range subsets {
			blue := extractCubes(subset, "blue")
			if blue > gamelist[id]["blue"] {
				gamelist[id]["blue"] = blue
			}
			green := extractCubes(subset, "green")
			if green > gamelist[id]["green"] {
				gamelist[id]["green"] = green
			}
			red := extractCubes(subset, "red")
			if red > gamelist[id]["red"] {
				gamelist[id]["red"] = red
			}
		}
	}

	maxcubes := make(map[string]int)
	maxcubes["blue"] = 14
	maxcubes["green"] = 13
	maxcubes["red"] = 12
	sum := 0
	sumofpowers := 0
	for id, game := range gamelist {
		if condition(game, maxcubes) {
			sum += id
		}
		sumofpowers += powerOfSet(game)
	}
	fmt.Println(sum)
	fmt.Println(sumofpowers)
}
