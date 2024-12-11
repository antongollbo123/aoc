package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Cache = make(map[int][]int)

func ReadFile(fileName string) map[int]int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	frequencies := make(map[int]int)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		for _, num := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(num)
			frequencies[n]++
		}
	}

	return frequencies
}

func transformStone(stone int) []int {
	// Check if stone exists in cache
	if result, exists := Cache[stone]; exists {
		return result
	}
	var result []int
	switch {
	case stone == 0:
		result = []int{1}
	case len(strconv.Itoa(stone))%2 == 0:
		digits := strconv.Itoa(stone)
		mid := len(digits) / 2
		left, _ := strconv.Atoi(digits[:mid])
		right, _ := strconv.Atoi(digits[mid:])
		result = []int{left, right}
	default:
		result = []int{stone * 2024}
	}
	// Store in cache
	Cache[stone] = result
	return result
}

func Transform(frequencies map[int]int) map[int]int {
	newFrequencies := make(map[int]int)
	for stone, count := range frequencies {
		transformed := transformStone(stone)
		for _, t := range transformed {
			newFrequencies[t] += count
		}
	}

	return newFrequencies
}

func main() {
	stones := ReadFile("real_input.txt")
	for i := 0; i < 25; i++ {
		stones = Transform(stones)
	}
	totalStones := 0
	for _, count := range stones {
		totalStones += count
	}
	fmt.Println("Part One:", totalStones)

	// Reset Cache
	Cache = make(map[int][]int)
	stones = ReadFile("real_input.txt")
	for i := 0; i < 75; i++ {
		stones = Transform(stones)
	}
	totalStones = 0
	for _, count := range stones {
		totalStones += count
	}
	fmt.Println("Part Two:", totalStones)
}
