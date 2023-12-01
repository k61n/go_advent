package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getNumber(word string) int {
	number := ""
	for i := 0; i < len(word); i++ {
		_, err := strconv.Atoi(string(word[i]))
		if err == nil {
			number += string(word[i])
		}
	}
	output, _ := strconv.Atoi(string(number[0]) + string(number[len(number)-1]))
	return output
}

func updateWord(word string, dict map[string]string) string {
	for key, value := range dict {
		word = strings.Replace(word, key, value, -1)
	}
	return word
}

func main() {
	exe, _ := os.Executable()
	dir := filepath.Dir(exe)
	input := "01.txt"

	file, _ := os.Open(filepath.Join(dir, input))
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	sum := 0
	for _, word := range words {
		sum += getNumber(word)
	}
	fmt.Println(sum)

	digits := make(map[string]string)
	digits["one"] = "o1e"
	digits["two"] = "t2o"
	digits["three"] = "t3e"
	digits["four"] = "4"
	digits["five"] = "5e"
	digits["six"] = "6"
	digits["seven"] = "7n"
	digits["eight"] = "e8t"
	digits["nine"] = "n9e"
	sum = 0
	for _, word := range words {
		sum += getNumber(updateWord(word, digits))
	}
	fmt.Println(sum)
}
