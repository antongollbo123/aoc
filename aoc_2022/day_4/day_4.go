package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkContainment(firstRange, secondRange []string) bool {
	firstStart, _ := strconv.Atoi(firstRange[0])
	firstEnd, _ := strconv.Atoi(firstRange[1])

	secondStart, _ := strconv.Atoi(secondRange[0])
	secondEnd, _ := strconv.Atoi(secondRange[1])

	return firstStart <= secondStart && firstEnd >= secondEnd ||
		secondStart <= firstStart && secondEnd >= firstEnd
}

func checkPartialContainment(firstRange, secondRange []string) bool {
	firstStart, _ := strconv.Atoi(firstRange[0])
	firstEnd, _ := strconv.Atoi(firstRange[1])

	secondStart, _ := strconv.Atoi(secondRange[0])
	secondEnd, _ := strconv.Atoi(secondRange[1])

	return firstStart <= secondEnd && firstEnd >= secondStart

}

func solution(file *os.File) (int, int) {

	fscanner := bufio.NewScanner(file)
	var firstTask int
	var secondTask int
	for fscanner.Scan() {
		line := fscanner.Text()
		ranges := strings.Split(line, ",")

		firstRange := strings.Split(ranges[0], "-")
		secondRange := strings.Split(ranges[1], "-")

		if checkContainment(firstRange, secondRange) {
			firstTask++
		}
		if checkPartialContainment(firstRange, secondRange) {
			secondTask++
		}
	}
	return firstTask, secondTask
}

func main() {

	fileName := "real_input.txt"

	file, _ := os.Open(fileName)
	firstSum, secondSum := solution(file)

	fmt.Println("TASK 1: ", firstSum)

	fmt.Println("TASK 2: ", secondSum)

}
