package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	x int
	y int
}

func readFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, ch := range line {
			row = append(row, string(ch))
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return grid
}

func findAntennas(grid [][]string) map[string][]coord {
	antennas := make(map[string][]coord)
	for y, row := range grid {
		for x, val := range row {
			if val != "." {
				antennas[val] = append(antennas[val], coord{x, y})
			}
		}
	}
	return antennas
}

func isValidCoord(c coord, grid [][]string) bool {
	return c.x >= 0 && c.x < len(grid[0]) && c.y >= 0 && c.y < len(grid)
}

func calculateAntinodes(grid [][]string, antennas map[string][]coord) map[coord]bool {
	antinodes := make(map[coord]bool)
	for _, coords := range antennas {
		if len(coords) < 2 {
			continue
		}
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				a1 := coords[i]
				a2 := coords[j]

				dx := a2.x - a1.x
				dy := a2.y - a1.y

				n1 := coord{x: a1.x - dx, y: a1.y - dy}
				n2 := coord{x: a2.x + dx, y: a2.y + dy}

				//Check validity
				if isValidCoord(n1, grid) {
					antinodes[n1] = true
				}
				if isValidCoord(n2, grid) {
					antinodes[n2] = true
				}
			}
		}
	}
	return antinodes
}

func calculateAlignedAntinodes(grid [][]string, antennas map[string][]coord) map[coord]bool {
	alignedAntinodes := make(map[coord]bool)

	for _, coords := range antennas {
		// One antenna only --> Antinode
		if len(coords) == 1 {
			alignedAntinodes[coords[0]] = true
			continue
		}
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				a1 := coords[i]
				a2 := coords[j]

				// Antennas become antinodes
				alignedAntinodes[a1] = true
				alignedAntinodes[a2] = true

				//Difference in x and y
				dx := a2.x - a1.x
				dy := a2.y - a1.y
				gcd := greatestCommonDivisor(abs(dx), abs(dy))

				// Normalize step sizes
				stepX := dx / gcd
				stepY := dy / gcd

				// Explore in both directions from a1 and a2
				directions := []int{-1, 1}
				for _, dir := range directions {
					for k := 1; k <= len(grid[0])+len(grid); k++ {
						point := coord{
							x: a1.x + dir*k*stepX,
							y: a1.y + dir*k*stepY,
						}
						if !isValidCoord(point, grid) {
							break
						}

						alignedAntinodes[point] = true
					}
				}
			}
		}
	}
	return alignedAntinodes
}

// Recursively find GCD through modulo giving the remainder, when b == 0, we have the GCD
func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	grid := readFile("sample_input.txt")
	antennas := findAntennas(grid)

	antinodes := calculateAntinodes(grid, antennas)
	fmt.Println("Part one:", len(antinodes))

	alignedAntinodes := calculateAlignedAntinodes(grid, antennas)
	fmt.Println("Part two:", len(alignedAntinodes))
}
