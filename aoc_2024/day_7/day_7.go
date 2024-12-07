package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) map[int64][]int64 {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	numMap := make(map[int64][]int64)
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		line := strings.TrimSpace(fscanner.Text())
		if line == "" {
			continue
		}
		val := strings.Split(line, ":")
		if len(val) != 2 {
			fmt.Println("Skipping invalid line:", line)
			continue
		}
		key, err := strconv.ParseInt(strings.TrimSpace(val[0]), 10, 64)
		if err != nil {
			fmt.Println("Invalid key:", val[0])
			continue
		}
		valueStrings := strings.Fields(val[1])
		var values []int64
		for _, v := range valueStrings {
			intVal, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				fmt.Println("Invalid value:", v)
				continue
			}
			values = append(values, intVal)
		}
		numMap[key] = values
	}
	if err := fscanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return numMap
}

// Concatenate two integers as strings and return the result as an integer
func concat(a, b int) int {
	return atoi(fmt.Sprintf("%d%d", a, b))
}

// Convert string to integer safely (like Python's int())
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func canConstruct(numbers []int64, current int64, target int64) bool {
	if len(numbers) == 0 {
		return current == target
	}
	nextNum := numbers[0]
	remainingNumbers := numbers[1:]
	if canConstruct(remainingNumbers, current+nextNum, target) {
		return true
	}

	if canConstruct(remainingNumbers, current*nextNum, target) {
		return true
	}

	//if canConstruct(remainingNumbers, concat(current, nextNum), target) {
	//	return true
	//}
	return false
}

func checkValues(numMap map[int64][]int64) int64 {
	var sum int64 = 0
	for target, numbers := range numMap {
		if len(numbers) == 0 {
			continue
		}
		if canConstruct(numbers[1:], numbers[0], target) {
			sum += target
		}
	}
	return sum
}

func main() {
	fileName := "real_input.txt" // Replace with your file name
	parsedMap := readFile(fileName)

	partOne := checkValues(parsedMap)
	fmt.Println("PART ONE: ", partOne)
}
