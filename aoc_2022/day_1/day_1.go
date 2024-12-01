package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumSlice(calories []int) int {
	sumValue := 0
	for _, s := range calories {
		sumValue = sumValue + s
	}
	return sumValue
}

func getMaxSlice(slice []int) int {
	maxVal := 0
	for _, s := range slice {
		if s > maxVal {
			maxVal = s
		}
	}
	return maxVal
}

func findTop3Values(values []int) []int {
	var top3 [3]int

	for _, v := range values {
		if v > top3[0] {
			top3[2] = top3[1]
			top3[1] = top3[0]
			top3[0] = v
		} else if v > top3[1] {
			top3[2] = top3[1]
			top3[1] = v
		} else if v > top3[2] {
			top3[2] = v
		}
	}

	result := []int{}
	for _, v := range top3 {
		if v != 0 {
			result = append(result, v)
		}
	}

	return result
}

func readFile(fileName string) []int {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close() // Ensure the file is closed after reading

	fscanner := bufio.NewScanner(file)
	lineNum := 1

	var calories []int
	var totalCalories []int

	for fscanner.Scan() {
		line := fscanner.Text()
		if len(line) > 0 {
			intValue, _ := strconv.Atoi(fscanner.Text())
			calories = append(calories, intValue)

		} else if len(line) == 0 {
			totalCalories = append(totalCalories, sumSlice(calories))
			calories = nil
		}
		lineNum++
	}
	if len(calories) > 0 {
		totalCalories = append(totalCalories, sumSlice(calories))
	}
	//fmt.Println("total calories", totalCalories)
	maxSliceValue := getMaxSlice(totalCalories)
	top3 := findTop3Values(totalCalories)
	top3Sum := sumSlice(top3)
	result := []int{}
	result = append(result, maxSliceValue)
	result = append(result, top3Sum)

	return result
}

func main() {
	fileName := "real_input.txt"

	result := readFile(fileName)

	fmt.Println(result[0], result[1])

}
