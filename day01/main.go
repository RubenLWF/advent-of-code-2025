package main

import (
	"fmt"
	"log"
	"strconv"

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

func solvePart1(lines []string) int {
	var count int = 0
	var pointer int = 50

	for i := 0; i < len(lines); i++ {
		pos := len(lines[i]) > 0 && lines[i][:1] == "R"
		v, _ := strconv.Atoi(lines[i][1:])

		if !pos {
			v = -v
		}

		pointer = utils.Mod(pointer+v, 100)

		if pointer == 0 {
			count++
		}
		
	}
	return count
}

func solvePart2(lines []string) int {
	var count int = 0
	var pointer int = 50

	for i := 0; i < len(lines); i++ {
		pos := len(lines[i]) > 0 && lines[i][:1] == "R"
		v, _ := strconv.Atoi(lines[i][1:])

		quotient, v := utils.Divmod(v, 100)

		count += quotient

		if !pos {
			v = -v
		}

		if pointer != 0 && (pointer + v < 0 || pointer + v > 100) {
			count++
		}

		pointer = utils.Mod(pointer+v, 100)

		if pointer == 0 {
			count++
		}
		
	}
	return count
}
