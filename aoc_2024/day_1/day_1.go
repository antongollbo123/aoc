package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(fileName string) ([]int, []int) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close() // Ensure the file is closed after reading

	fscanner := bufio.NewScanner(file)

	var leftList []int
	var rightList []int

	for fscanner.Scan() {
		line := strings.TrimSpace(fscanner.Text())
		stringSlice := strings.Split(line, "   ")

		valOne, _ := strconv.Atoi(stringSlice[0])
		valTwo, _ := strconv.Atoi(stringSlice[1])

		leftList = append(leftList, valOne)
		rightList = append(rightList, valTwo)

	}
	return leftList, rightList
}

func subtractLists(listOne []int, listTwo []int) int {
	var sum int
	sort.Ints(listOne)
	sort.Ints(listTwo)
	for i := range listOne {
		diff := listOne[i] - listTwo[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	return sum
}

func createMap(intSlice []int) map[int]int {
	intMap := map[int]int{}

	for _, value := range intSlice {
		intMap[value]++
	}

	return intMap
}

func getSimilarityScore(listOne []int, mapTwo map[int]int) int {
	var sum int
	for _, value := range listOne {
		sum += value * mapTwo[value]
	}
	return sum
}

func main() {

	leftList, rightList := readFile("real_input.txt")

	intMap := createMap(rightList)

	partOne := subtractLists(leftList, rightList)

	partTwo := getSimilarityScore(leftList, intMap)

	fmt.Println("First task: ", partOne, "\nSecond Task: ", partTwo)

}
