package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	figuresCodes = map[string]string{
		"A": "R",
		"B": "P",
		"C": "S",
		"X": "R",
		"Y": "P",
		"Z": "S",
	}
	figureScore = map[string]int{
		"R": 1,
		"P": 2,
		"S": 3,
	}

	ruleMap = map[string][2]string{
		"R": {"S", "P"},
		"P": {"R", "S"},
		"S": {"P", "R"},
	}
)

func readFile(fileName string) (int, int) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close() // Ensure the file is closed after reading

	fscanner := bufio.NewScanner(file)
	totScore1 := 0
	totScore2 := 0

	for fscanner.Scan() {
		res := strings.Split(fscanner.Text(), " ")
		oppChoice, myChoice := res[0], res[1]
		oppType, myType := figuresCodes[oppChoice], figuresCodes[myChoice]
		score := figureScore[myType]

		newStrategyScore := 0

		// First task
		if oppType == myType {
			score += 3
		} else if oppType == "R" && myType == "S" || oppType == "S" && myType == "P" || oppType == "P" && myType == "R" {
			score += 0
		} else {
			score += 6
		}

		// Second task
		switch myChoice {
		case "X":
			newStrategyScore += 0 + figureScore[ruleMap[oppType][0]]
		case "Y":
			newStrategyScore += 3 + figureScore[oppType]
		case "Z":
			newStrategyScore += 6 + figureScore[ruleMap[oppType][1]]
		}
		totScore1 += score
		totScore2 += newStrategyScore

	}
	return totScore1, totScore2
}

func main() {
	fileName := "real_input.txt"
	res1, res2 := readFile(fileName)

	fmt.Println(res1, res2)

}
