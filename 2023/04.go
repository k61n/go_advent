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

func getNumbers(card string, section int) []int {
	cardSlice := strings.Split(card, ":")
	numberSections := strings.Split(cardSlice[1], "|")
	numbers := strings.Split(numberSections[section], " ")
	var res []int
	for _, numStr := range numbers {
		numInt, err := strconv.Atoi(numStr)
		if err == nil {
			res = append(res, numInt)
		}
	}
	return res
}

func getWinNumbers(card string) []int {
	return getNumbers(card, 0)
}

func getCardNumbers(card string) []int {
	return getNumbers(card, 1)
}

func getMatches(slice0, slice1 []int) int {
	res := 0
	for i := 0; i < len(slice0); i++ {
		for j := 0; j < len(slice1); j++ {
			if slice1[j] == slice0[i] {
				res += 1
			}
		}
	}
	return res
}

func calcInstances(cardNum int, cards []string) int {
	res := 0
	score := getMatches(getWinNumbers(cards[cardNum-1]), getCardNumbers(cards[cardNum-1]))
	res += score
	for i := 1; i <= score; i++ {
		res += calcInstances(cardNum+i, cards)
	}
	return res
}

func main() {
	exe, _ := os.Executable()
	dir := filepath.Dir(exe)
	input := "04.txt"

	file, _ := os.Open(filepath.Join(dir, input))
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Part 1
	sum := 0
	for _, line := range lines {
		score := getMatches(getWinNumbers(line), getCardNumbers(line))
		if score > 0 {
			sum += int(math.Pow(2, float64(score-1)))
		}
	}
	fmt.Println(sum)

	// Part 2
	sum = 0
	for i, _ := range lines {
		sum += calcInstances(i+1, lines)
	}
	sum += len(lines)
	fmt.Println(sum)
}
