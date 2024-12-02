package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) [][]string {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close() // Ensure the file is closed after reading

	fscanner := bufio.NewScanner(file)

	var lineList [][]string

	for fscanner.Scan() {
		line := strings.TrimSpace(fscanner.Text())
		innerList := strings.Split(line, " ")
		lineList = append(lineList, innerList)
	}
	return lineList
}

func isSafeReport(list []int) bool {
	var direction int

	for j := 0; j < len(list)-1; j++ {
		diff := list[j+1] - list[j]
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return false
		}
		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else if diff < 0 {
				direction = -1
			}
		}
		if (direction == 1 && diff < 0) || (direction == -1 && diff > 0) {
			return false
		}
	}
	return true
}
func partOne(reportList [][]string) int {
	var numSafe int
	for _, list := range reportList {
		intList := make([]int, len(list))
		for i, s := range list {
			intList[i], _ = strconv.Atoi(s)
		}
		if isSafeReport(intList) {
			numSafe++
			continue
		}
	}
	return numSafe
}

func listToStringKey(list []int) string {
	var sb strings.Builder
	for _, num := range list {
		sb.WriteString(strconv.Itoa(num) + ",")
	}
	return sb.String()
}

func partTwo(reportList [][]string) int {
	var numSafe int
	for _, list := range reportList {
		intList := make([]int, len(list))
		for i, s := range list {
			intList[i], _ = strconv.Atoi(s)
		}
		// Check if the report is inherently safe
		if isSafeReport(intList) {
			numSafe++
			continue
		}

		// Try removing each element and check if the list becomes safe
		for i := 0; i < len(intList); i++ {
			// Create a modified list without the i-th element
			modifiedList := append([]int{}, intList[:i]...)       // Create a new list up to the index
			modifiedList = append(modifiedList, intList[i+1:]...) // Append the rest after the index

			// Check if the modified list is safe
			if isSafeReport(modifiedList) {
				numSafe++
				break // We can stop after finding one valid removal
			}
		}
	}
	return numSafe
}

func main() {

	lineList := readFile("real_input.txt")

	partOne := partOne(lineList)
	partTwo := partTwo(lineList)
	fmt.Println(partOne, partTwo)
}
