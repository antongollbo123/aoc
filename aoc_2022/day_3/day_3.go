package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func getNumericValue(s1 rune) int {
	var numericValue int
	if unicode.IsUpper(s1) {
		numericValue = int(string(s1)[0]) + 58 - 96
	} else {
		numericValue = int(string(s1)[0] - 96)
	}
	return numericValue
}

func findOverlappingChars(s1, s2 string) []rune {
	charMap := make(map[rune]bool)
	overlap := []rune{}

	for _, char := range s1 {
		charMap[char] = true
	}
	for _, char := range s2 {
		if charMap[char] {
			overlap = append(overlap, char)
			delete(charMap, char)
		}
	}
	return overlap
}

func firstTask(file *os.File) int {
	fscanner := bufio.NewScanner(file)
	var sum int
	// First task
	for fscanner.Scan() {
		line := fscanner.Text()
		stringLength := len(line)
		firstCompartment, secondCompartment := line[:stringLength/2], line[stringLength/2:]

		overlap := findOverlappingChars(firstCompartment, secondCompartment)
		sum += getNumericValue(overlap[0])
	}

	return sum
}

func secondTask(file *os.File) int {
	fscanner := bufio.NewScanner(file)
	lineCounter := 0
	var tripleLineMap []string
	fscanner = bufio.NewScanner(file)
	var sum int
	for fscanner.Scan() {
		line := fscanner.Text()
		tripleLineMap = append(tripleLineMap, line)
		lineCounter++

		if lineCounter%3 == 0 {
			res := findOverlappingChars(tripleLineMap[0], tripleLineMap[1])
			res2 := findOverlappingChars(string(res), tripleLineMap[2])
			sum += getNumericValue(res2[0])
			tripleLineMap = nil

		}
	}
	return sum
}

func main() {

	fileName := "real_input.txt"

	file, _ := os.Open(fileName)
	firstSum := firstTask(file)

	fmt.Println("TASK 1: ", firstSum)

	file, _ = os.Open(fileName)
	secondSum := secondTask(file)

	fmt.Println("TASK 2: ", secondSum)

}
