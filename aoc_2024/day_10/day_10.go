package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) [][]int {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close() // Ensure the file is closed after reading

	fscanner := bufio.NewScanner(file)

	var grid [][]int

	for fscanner.Scan() {
		line := strings.TrimSpace(fscanner.Text())
		var row []int
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	return grid
}

func findTrailheads(grid [][]int) []struct{ x, y int } {
	var trailheads []struct{ x, y int }
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				trailheads = append(trailheads, struct{ x, y int }{i, j})
			}
		}
	}
	return trailheads
}

func pathToString(path []struct{ x, y int }) string {
	var builder strings.Builder
	for _, p := range path {
		builder.WriteString(fmt.Sprintf("(%d,%d)", p.x, p.y))
	}
	return builder.String()
}

// Cheeky interface to collect different result from part 1 / 2 8-)
func dfs(grid [][]int, x, y int, visited [][]bool, path []struct{ x, y int }, results interface{}) {

	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return // Out of bounds
	}
	if visited[x][y] { // Visited
		return
	}

	visited[x][y] = true
	path = append(path, struct{ x, y int }{x, y})

	if grid[x][y] == 9 { // Reached a valid end
		switch res := results.(type) {
		case *map[string]bool: // For paths
			pathStr := pathToString(path)
			(*res)[pathStr] = true
		case *map[struct{ x, y int }]bool: // For unique `9`s
			endPoint := struct{ x, y int }{x, y}
			(*res)[endPoint] = true
		}
	} else {
		// Explore neighbors in four directions
		directions := []struct{ dx, dy int }{
			{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		}
		for _, d := range directions {
			nx, ny := x+d.dx, y+d.dy
			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
				if grid[nx][ny] == grid[x][y]+1 { // Valid height increase
					dfs(grid, nx, ny, visited, path, results)
				}
			}
		}
	}
	visited[x][y] = false
}

func calculateTrailheadScores(grid [][]int) int {
	trailheads := findTrailheads(grid)
	totalScore := 0

	for _, trailhead := range trailheads {
		visited := make([][]bool, len(grid))
		for i := range visited {
			visited[i] = make([]bool, len(grid[0]))
		}
		results := make(map[struct{ x, y int }]bool)
		dfs(grid, trailhead.x, trailhead.y, visited, nil, &results)
		totalScore += len(results)
	}

	return totalScore
}

func calculateTrailheadRatings(grid [][]int) int {
	trailheads := findTrailheads(grid)
	totalRating := 0

	for _, trailhead := range trailheads {
		visited := make([][]bool, len(grid))
		for i := range visited {
			visited[i] = make([]bool, len(grid[0]))
		}
		results := make(map[string]bool)
		dfs(grid, trailhead.x, trailhead.y, visited, nil, &results)
		totalRating += len(results)
	}

	return totalRating
}

func main() {
	grid := readFile("real_input.txt")

	totalScore := calculateTrailheadScores(grid)
	fmt.Println("Part one:", totalScore)

	totalRating := calculateTrailheadRatings(grid)
	fmt.Println("Part Two: ", totalRating)
}
