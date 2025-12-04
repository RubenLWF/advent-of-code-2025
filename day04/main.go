package main

import (
	"fmt"
	"log"

	"aoc/utils"
)

func main() {
	// Read input
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Solve parts
	part1 := solvePart1(lines)
	part2 := solvePart2(lines)

	// Print results
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

var DIRS = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func convertToGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i := range lines {
		grid[i] = []rune(lines[i])
	}
	return grid
}

func countAdjacent(grid [][]rune, i, j int, dirs [][2]int) int {
	adjacent := 0
	for _, dir := range dirs {
		ni, nj := i+dir[0], j+dir[1]
		if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) {
			if grid[ni][nj] == '@' {
				adjacent++
			}
		}
	}
	return adjacent
}

func solvePart1(lines []string) int {
	count := 0

	grid := convertToGrid(lines)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '@' {
				if countAdjacent(grid, i, j, DIRS) < 4 {
					count++
				}
			}
		}
	}
	return count
}

func solvePart2(lines []string) int {
	count := 0

	grid := convertToGrid(lines)

	for {
		toRemove := make([][2]int, 0)

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == '@' {
					if countAdjacent(grid, i, j, DIRS) < 4 {
						toRemove = append(toRemove, [2]int{i, j})
					}
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		count += len(toRemove)

		for _, pos := range toRemove {
			grid[pos[0]][pos[1]] = '.'
		}
	}
	return count
}
