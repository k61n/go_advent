package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
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

func parseLine(line string) (string, int) {
	hand := strings.Split(line, " ")
	num, _ := strconv.Atoi(hand[1])
	return hand[0], num
}

func cardToNum(card string, withJ bool) int {
	cards := make(map[string]int)
	cards["T"] = 10
	if withJ {
		cards["J"] = 1
	} else {
		cards["J"] = 11
	}
	cards["Q"] = 12
	cards["K"] = 13
	cards["A"] = 14
	num, err := strconv.Atoi(card)
	if err == nil {
		return num
	} else {
		return cards[card]
	}
}

func mapHand(hand string, withJ bool) map[string]int {
	handMap := make(map[string]int)
	for _, l := range hand {
		handMap[string(l)] += 1
	}
	if withJ {
		jVal := handMap["J"]
		delete(handMap, "J")
		maxVal := 0
		maxKey := ""
		for key, value := range handMap {
			if value >= maxVal {
				maxVal = value
				maxKey = key
			}
		}
		handMap[maxKey] += jVal
	}
	return handMap
}

func getCounts(hand map[string]int) []int {
	var counts []int
	for _, value := range hand {
		counts = append(counts, value)
	}
	return counts
}

func rankHand(hand string, withJ bool) int {
	handMap := mapHand(hand, withJ)
	counts := getCounts(handMap)
	switch slices.Max(counts) {
	case 5:
		return 6
	case 4:
		return 5
	case 3:
		if slices.Contains(counts, 2) {
			return 4
		} else {
			return 3
		}
	case 2:
		n := 0
		for _, count := range counts {
			if count == 2 {
				n += 1
			}
		}
		return n
	}
	return 0
}

func compare(smaller, bigger string, withJ bool) bool {
	rank0 := rankHand(smaller, withJ)
	rank1 := rankHand(bigger, withJ)
	if rank0 == rank1 {
		for i := 0; i < len(smaller); i++ {
			if cardToNum(string(smaller[i]), withJ) != cardToNum(string(bigger[i]), withJ) {
				return cardToNum(string(smaller[i]), withJ) < cardToNum(string(bigger[i]), withJ)
			}
		}
	}
	return rank0 < rank1
}

func sortStr(arr []string, low, high int, withJ bool) {
	if low < high {
		pivotIndex := split(arr, low, high, withJ)
		sortStr(arr, low, pivotIndex-1, withJ)
		sortStr(arr, pivotIndex+1, high, withJ)
	}
}

func split(arr []string, low, high int, withJ bool) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if compare(arr[j], pivot, withJ) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	lines := readInputFile("07.txt")

	var hands []string
	var bids []int
	handToBid := make(map[string]int)
	for _, line := range lines {
		hand, bid := parseLine(line)
		hands = append(hands, hand)
		bids = append(bids, bid)
		handToBid[hand] = bid
	}

	// Part 1
	sortStr(hands, 0, len(hands)-1, false)
	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * handToBid[hand]
	}
	fmt.Println(sum)

	// Part 2
	sortStr(hands, 0, len(hands)-1, true)
	sum = 0
	for i, hand := range hands {
		sum += (i + 1) * handToBid[hand]
	}
	fmt.Println(sum)
}
