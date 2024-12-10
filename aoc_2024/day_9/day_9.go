package main

import (
	"bufio"
	"fmt"
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
		innerList := strings.Split(line, "\n")
		lineList = append(lineList, innerList)
	}
	return lineList
}

func readDiscMap(line []string) []int {

	input := []rune(line[0])
	var fileBlock []int
	indexer := 0
	for idx, elem := range input {
		stringRune := string(elem)
		intRune, _ := strconv.Atoi(stringRune)
		if idx%2 == 0 {
			for i := range intRune {
				if i >= 0 {
					fileBlock = append(fileBlock, indexer)
				}
			}
			indexer++
		} else {
			for i := range intRune {
				fileBlock = append(fileBlock, -1)
				fmt.Println(i)
			}
		}
	}
	return fileBlock
}

func compactDisk(diskMap []int) []int {
	steps := 0

	for {
		moved := false
		rightmostFileIndex := -1

		for i := len(diskMap) - 1; i >= 0; i-- {
			if diskMap[i] != -1 {
				rightmostFileIndex = i
				break
			}
		}

		if rightmostFileIndex == -1 {
			break
		}

		leftmostFreeSpaceIndex := -1
		for i := 0; i < rightmostFileIndex; i++ {
			if diskMap[i] == -1 {
				leftmostFreeSpaceIndex = i
				break
			}
		}

		if leftmostFreeSpaceIndex == -1 {
			break
		}
		diskMap[leftmostFreeSpaceIndex] = diskMap[rightmostFileIndex]
		diskMap[rightmostFileIndex] = -1
		moved = true
		steps++
		if steps%10000 == 0 {
			fmt.Printf("Step %d: Current progress\n", steps)
		}
		if !moved {
			break
		}
	}

	return diskMap
}

func calculateChecksum(finalState []int) int64 {
	var checksum int64 = 0
	realPosition := 0

	for _, block := range finalState {
		if block != -1 {
			fileID := block
			checksum += int64(realPosition * fileID)
			realPosition++
		}
	}

	return checksum
}

func main() {

	line := readFile("real_input.txt")[0]
	fileBlock := readDiscMap(line)

	fmt.Println(fileBlock)
	fmt.Println("NORMAL: ", fileBlock)
	rearranged := compactDisk(fileBlock)
	fmt.Println("REARRANGED: ", rearranged)
	partOne := calculateChecksum(rearranged)
	fmt.Println("PART ONE: ", partOne)

}
