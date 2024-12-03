package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close() // Ensure the file is closed after reading

	fscanner := bufio.NewScanner(file)

	var lineList []string

	for fscanner.Scan() {
		line := strings.TrimSpace(fscanner.Text())
		lineList = append(lineList, line)
	}
	return lineList
}

func findPattern(input string, enabled bool) int {
	if !enabled {
		return 0 // Skip processing if disabled
	}
	regex := regexp.MustCompile(`mul\s*\(\s*(\d{1,3})\s*,\s*(\d{1,3})\s*\)`)
	match := regex.FindAllStringSubmatch(input, -1)

	totSum := 0
	for _, match := range match {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		res := x * y
		totSum += res
	}

	return totSum
}

func findNextKeyword(input string) (nextIndex int, keyword string, remaining string, currentSegment string) {
	indexDo := strings.Index(input, "do()")
	indexDont := strings.Index(input, "don't()")

	if indexDo != -1 && (indexDont == -1 || indexDo < indexDont) {
		nextIndex = indexDo
		keyword = "do()"
	} else if indexDont != -1 {
		nextIndex = indexDont
		keyword = "don't()"
	} else {
		nextIndex = -1
		return
	}

	nextIndexEnd := nextIndex + len(keyword)
	remaining = input[nextIndexEnd:]
	currentSegment = input[:nextIndex]
	return
}

func processInstructions(input string) int {
	enabled := true
	totalSum := 0
	for {
		nextIndex, keyword, remaining, currentSegment := findNextKeyword(input)
		if nextIndex == -1 {
			break
		}
		totalSum += findPattern(currentSegment, enabled)
		if keyword == "do()" {
			enabled = true
		} else if keyword == "don't()" {
			enabled = false
		}
		input = remaining
	}
	totalSum += findPattern(input, enabled)
	return totalSum
}

func main() {

	fileName := "real_input.txt"
	lineList := readFile(fileName)

	input := strings.Join(lineList[:], ",")
	partOne := findPattern(input, true)

	partTwo := processInstructions(input)
	fmt.Println(partOne, partTwo)
}
