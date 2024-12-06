package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) (map[int][]int, [][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	fscanner := bufio.NewScanner(file)
	rules := make(map[int][]int)
	var pages [][]int
	isPagesSection := false

	for fscanner.Scan() {
		line := strings.TrimSpace(fscanner.Text())
		if line == "" {
			isPagesSection = true
			continue
		}

		if !isPagesSection {
			// Parse rules.
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				return nil, nil, fmt.Errorf("invalid rule format: %s", line)
			}
			left, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, nil, err
			}
			right, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, err
			}
			rules[left] = append(rules[left], right)
		} else {
			// Parse pages.
			parts := strings.Split(line, ",")
			var update []int
			for _, part := range parts {
				num, err := strconv.Atoi(part)
				if err != nil {
					return nil, nil, err
				}
				update = append(update, num)
			}
			pages = append(pages, update)
		}
	}

	if err := fscanner.Err(); err != nil {
		return nil, nil, err
	}

	return rules, pages, nil
}

func checkOccurence(currentNum int, seen []int, rules map[int][]int) bool {
	vals, ok := rules[currentNum]
	if !ok {
		return true
	}

	for _, requiredAfter := range vals {
		for _, val := range seen {
			if val == currentNum {
				break
			}
			if val == requiredAfter {
				return false
			}
		}
	}
	return true
}

func sort(rules map[int][]int, input []int) []int {
	visited := make(map[int]bool) // Tracks visited nodes
	stack := []int{}              // Stack to hold the topologically sorted nodes

	// Helper function to check if a node is part of the input list
	isInInput := func(node int) bool {
		for _, page := range input {
			if page == node {
				return true
			}
		}
		return false
	}

	// Recursive DFS
	var dfs func(node int)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		for _, dependent := range rules[node] {
			if isInInput(dependent) {
				dfs(dependent)
			}
		}
		stack = append(stack, node)
	}
	for _, num := range input {
		if !visited[num] {
			dfs(num)
		}
	}
	result := []int{}
	for i := len(stack) - 1; i >= 0; i-- {
		result = append(result, stack[i])
	}

	return result
}

func iteratePages(pages [][]int, rules map[int][]int) (int, int) {
	sumMidValues := 0
	sumReorderedMidValues := 0
	for _, update := range pages {
		var seen []int
		isValid := true

		for _, num := range update {
			if !checkOccurence(num, seen, rules) {
				isValid = false
				break
			}
			seen = append(seen, num)
		}

		if isValid {
			sumMidValues += update[len(update)/2]
		} else {
			// Reorder the incorrect update
			reordered := sort(rules, update)
			sumReorderedMidValues += reordered[len(reordered)/2]
		}
	}
	return sumMidValues, sumReorderedMidValues

}

func main() {
	rules, pages, err := readFile("real_input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	partOne, partTwo := iteratePages(pages, rules)

	fmt.Printf("Part one: %d\n", partOne)
	fmt.Printf("Part two: %d\n", partTwo)
}
