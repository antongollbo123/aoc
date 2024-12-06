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

	fmt.Println(vals)
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

func iteratePages(pages [][]int, rules map[int][]int) {
	sumMidValues := 0

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
		}
	}

	fmt.Printf("Sum of middle values: %d\n", sumMidValues)
}

func main() {
	rules, pages, err := readFile("real_input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	iteratePages(pages, rules)
}
