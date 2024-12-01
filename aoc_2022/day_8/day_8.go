package main

import (
	"bufio"
	"os"
)

func firstTask(file *os.File) int {
	fscanner := bufio.NewScanner(file)

}

func main() {

	fileName := "sample_input.txt"

	file, _ := os.Open(fileName)

	firstTask(file)

}
