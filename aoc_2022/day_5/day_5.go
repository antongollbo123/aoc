package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type moverStruct struct {
	Quantity    int
	From        int
	Destination int
}

func firstTask(crates [][]string, moves []moverStruct) string {
	columns := make(map[int][]string)
	fmt.Println(crates)
	for i, row := range crates {
		columns[i+1] = row
	}
	for _, move := range moves {
		for i := 0; i < move.Quantity; i++ {
			if len(crates[move.From]) > 0 {
				crate := crates[move.From][0]
				crates[move.From] = crates[move.From][1:]
				crates[move.Destination] = append([]string{crate}, crates[move.Destination]...)
			}
		}
	}
	return "hello"
}

func parseInput(file *os.File) ([][]string, []moverStruct) {
	fscanner := bufio.NewScanner(file)

	var crates [][]string
	var moves []moverStruct

	for fscanner.Scan() {
		line := fscanner.Text()
		if line == "" {
			break
		}
		// Split the line into columns
		columns := strings.Fields(line)
		// Handle columns with missing items
		row := make([]string, len(columns))
		for _, col := range columns {
			fmt.Println(col)
			if col == " " {

				row = append(row, "FUCKOFF")
			} else {
				row = append(row, col)
			}
		}
		crates = append(crates, row)
	}

	for fscanner.Scan() {
		line := fscanner.Text()
		// Example: move 1 from 2 to 1
		parts := strings.Fields(line)
		if len(parts) == 6 {
			quantity, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			destination, _ := strconv.Atoi(parts[5])
			move := moverStruct{Quantity: quantity, From: from, Destination: destination}
			moves = append(moves, move)
		}
	}
	return crates, moves
}

func main() {

	fileName := "sample_input.txt"

	file, _ := os.Open(fileName)
	crates, moves := parseInput(file)

	firstTask(crates, moves)

	//fmt.Println(res)

	// fmt.Println("TASK 1: ", firstSum)

	/*
		file, _ = os.Open(fileName)
		secondSum := secondTask(file)

		fmt.Println("TASK 2: ", secondSum)
	*/
}
