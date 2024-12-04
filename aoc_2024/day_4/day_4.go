package main

import (
	"bufio"
	"fmt"
	"os"
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
		innerList := strings.Split(line, "\n")

		lineList = append(lineList, innerList)
	}
	return lineList
}

func makeGrid(lineList [][]string) [][]string {
	grid := make([][]string, len(lineList))

	for i, row := range lineList {
		grid[i] = strings.Split(row[0], "")
	}

	return grid
}

func gridTraversal(grid [][]string, word string) (int, int) {
	allowedDirections := [][]int{
		{0, 1},   // Right
		{0, -1},  // Left
		{1, 0},   // Down
		{-1, 0},  // Up
		{1, 1},   // Down-Right
		{-1, -1}, // Up-Left
		{1, -1},  // Down-Left
		{-1, 1},  // Up-Right
	}

	wordCount := 0 // To keep track of the number of matches
	countXMAS := 0
	// Loop through every cell in the grid
	for row := range grid {
		for col := range grid[row] {
			// Check all directions from the current cell
			for _, dir := range allowedDirections {
				if checkWord(grid, row, col, word, dir) {
					wordCount++
				}
			}
		}
	}
	// Loop again, but for second task :D
	for row := range grid {
		for col := range grid[row] {
			// Check if out-of-bounds
			if row > 0 && row < len(grid)-1 && col > 0 && col < len(grid[row])-1 {
				if isXmas(grid, row, col) {
					countXMAS++
				}
			}

		}
	}

	return wordCount, countXMAS
}

func checkWord(grid [][]string, startX, startY int, word string, direction []int) bool {
	wordLength := len(word)
	for i := 0; i < wordLength; i++ {
		newX := startX + i*direction[0]
		newY := startY + i*direction[1]
		if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[0]) {
			return false
		}
		if grid[newX][newY] != string(word[i]) {
			return false
		}
	}
	return true
}

func isXmas(grid [][]string, row, col int) bool {
	// If input is not A, we do not continue. We are looking for the centerpieces
	if grid[row][col] != "A" {
		return false
	}
	// Given that A is the centerpiece, we wan't to check the diagonals only.
	topLeft := grid[row-1][col-1]
	topRight := grid[row-1][col+1]
	bottomLeft := grid[row+1][col-1]
	bottomRight := grid[row+1][col+1]
	// Check if isMas, i.e. topLeft is S, bottomRight is M -> topRight is M bottomLeft is S
	// OR topLeft is M, bottomRight is S -> topRight is M bottomLeft is S
	return (isMas(topLeft, bottomRight) && isMas(topRight, bottomLeft))
}

func isMas(start, end string) bool {
	// We check both possible combinations of crosses, either start with S - M or ends with S - M
	return (start == "M" && end == "S") || (start == "S" && end == "M")
}

func main() {

	lineList := readFile("real_input.txt")
	grid := makeGrid(lineList)
	word := "XMAS"

	wc, xmas := gridTraversal(grid, word)

	fmt.Println(wc, xmas)
}
