package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func notContains(search string, slice []string) bool {
	for _, l := range search {
		for _, str := range slice {
			if str != string(l) {
				return true
			}
		}
	}
	return false
}

func contains(search string, slice []string) bool {
	for _, l := range search {
		for _, str := range slice {
			if str == string(l) {
				return true
			}
		}
	}
	return false
}

func extractNumbers(line string) map[string][]int {
	numbers := make(map[string][]int)
	number := ""
	conversion := false
	for i := 0; i < len(line); i++ {
		_, err := strconv.Atoi(string(line[i]))
		if err == nil {
			number += string(line[i])
			conversion = true
		} else {
			if conversion {
				num, _ := strconv.Atoi(number)
				numbers[number] = append(numbers[number], num)
				numbers[number] = append(numbers[number], i-len(number))
				number = ""
				conversion = false
			}
		}
		if i == len(line)-1 && conversion {
			num, _ := strconv.Atoi(number)
			numbers[number] = append(numbers[number], num)
			numbers[number] = append(numbers[number], i-len(number)+1)
			number = ""
			conversion = false
		}
	}
	return numbers
}

func lengthOfMap(sample map[string][]int) int {
	if len(sample) == 1 {
		for _, value := range sample {
			return len(value) / 2
		}
	}
	return len(sample)
}

func surroundLine(lines []string, col0 int, col1 int, row int) string {
	res := "."
	if row > 0 {
		if col0 > 0 {
			res += string(lines[row-1][col0-1])
		}
		res += lines[row-1][col0 : col1+1]
		if col1 < len(lines[0])-1 {
			res += string(lines[row-1][col1+1])
		}
	}
	res += "."
	if len(lines)-1 > row {
		if col0 > 0 {
			res += string(lines[row+1][col0-1])
		}
		res += lines[row+1][col0 : col1+1]
		if col1 < len(lines[0])-1 {
			res += string(lines[row+1][col1+1])
		}
	}
	res += "."
	if col0 > 0 {
		res += string(lines[row][col0-1])
	}
	res += "."
	if col1 < len(lines[0])-1 {
		res += string(lines[row][col1+1])
	}
	res += "."
	return res
}

func main() {
	exe, _ := os.Executable()
	dir := filepath.Dir(exe)
	input := "03.txt"

	file, _ := os.Open(filepath.Join(dir, input))
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var excluded []string
	excluded = append(excluded, ".")
	var included []string
	included = append(included, "*")

	sum := 0
	for i, line := range lines {
		var lineSlice []string
		row := 0
		if i > 0 {
			lineSlice = append(lineSlice, lines[i-1])
			row = 1
		}
		lineSlice = append(lineSlice, line)
		if i < len(lines)-1 {
			lineSlice = append(lineSlice, lines[i+1])
		}
		for number, ints := range extractNumbers(line) {
			for k := 0; k < len(ints); k += 2 {
				surround := surroundLine(lineSlice, ints[k+1], ints[k+1]+len(number)-1, row)
				if notContains(surround, excluded) {
					sum += ints[k]
					break
				}
			}
		}
	}
	fmt.Println(sum)

	sum = 0
	for i, line := range lines {
		if contains(line, included) {
			row := 0
			var lineSlice []string
			if i > 0 {
				lineSlice = append(lineSlice, lines[i-1])
				row = 1
			}
			lineSlice = append(lineSlice, line)
			if i < len(lines)-1 {
				lineSlice = append(lineSlice, lines[i+1])
			}
			for j, l := range line {
				if string(l) == "*" {
					numbers := extractNumbers(surroundLine(lineSlice, j, j, row))
					if lengthOfMap(numbers) == 2 {
						pow := 1
						for _, ln := range lineSlice {
							for number, ints := range extractNumbers(ln) {
								for k := 0; k < len(ints); k += 2 {
									if ints[k+1]-1 <= j && j <= ints[k+1]+len(number) {
										pow *= ints[k]
									}
								}
							}
						}
						sum += pow
					}
				}
			}
		}
	}
	fmt.Println(sum)
}
