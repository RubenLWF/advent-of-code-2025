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

type Point struct {
	row, col int
}

type Beam struct {
	pos Point
}

func buildGrid(lines []string) ([][]rune, Point) {
	grid := make([][]rune, len(lines))
	var start Point
	for i, line := range lines {
		grid[i] = []rune(line)
		for j, ch := range line {
			if ch == 'S' {
				start = Point{i, j}
			}
		}
	}
	return grid, start
}

func solvePart1(lines []string) int {
	var start Point
	var grid [][]rune
	grid, start = buildGrid(lines)

	usedSplitters := make(map[Point]bool)

	beams := []Beam{{pos: start}}

	for len(beams) > 0 {
		var nextBeams []Beam

		for _, beam := range beams {
			nextRow := beam.pos.row + 1

			if nextRow >= len(grid) {
				continue
			}

			nextPos := Point{nextRow, beam.pos.col}

			if grid[nextPos.row][nextPos.col] == '^' {
				if !usedSplitters[nextPos] {
					usedSplitters[nextPos] = true

					if nextPos.col > 0 {
						nextBeams = append(nextBeams, Beam{pos: Point{nextPos.row, nextPos.col - 1}})
					}
					if nextPos.col < len(grid[nextPos.row])-1 {
						nextBeams = append(nextBeams, Beam{pos: Point{nextPos.row, nextPos.col + 1}})
					}
				}
			} else {
				nextBeams = append(nextBeams, Beam{pos: nextPos})
			}
		}

		beams = nextBeams
	}

	return len(usedSplitters)
}

func solvePart2(lines []string) int {
	var start Point
	var grid [][]rune
	grid, start = buildGrid(lines)

	memo := make(map[Point]int)
	return dfs(grid, start.row, start.col, memo)
}

func dfs(grid [][]rune, row, col int, memo map[Point]int) int {
	nextRow := row + 1

	if nextRow >= len(grid) {
		return 1
	}

	if col < 0 || col >= len(grid[nextRow]) {
		return 0
	}

	pos := Point{nextRow, col}

	if count, exists := memo[pos]; exists {
		return count
	}

	nextCell := grid[nextRow][col]
	var timelines int

	if nextCell == '^' {
		timelines = 0

		if col > 0 {
			timelines += dfs(grid, nextRow, col-1, memo)
		}

		if col < len(grid[nextRow])-1 {
			timelines += dfs(grid, nextRow, col+1, memo)
		}
	} else {
		timelines = dfs(grid, nextRow, col, memo)
	}

	memo[pos] = timelines
	return timelines
}
