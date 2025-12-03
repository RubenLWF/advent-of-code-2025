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

func solvePart1(lines []string) int {
	count := 0

	for _, line := range lines {
		digits := []int{}
		for _, ch := range line {
			digits = append(digits, int(ch-'0'))
		}

		max := 0
		for i := 0; i < len(digits)-1; i++ {
			for j := i + 1; j < len(digits); j++ {
				value := digits[i]*10 + digits[j]
				if value > max {
					max = value
				}
			}
		}
		count += max
	}
	return count
}

func solvePart2(lines []string) int {
	count := 0

	for _, line := range lines {
		digits := []int{}
		for _, ch := range line {
			digits = append(digits, int(ch-'0'))
		}

		selected := []int{}
		pos := 0
		need := 12

		for need > 0 && pos < len(digits) {
			maxDigit := -1
			maxPos := -1

			end := len(digits) - need + 1

			for i := pos; i < end; i++ {
				if digits[i] > maxDigit {
					maxDigit = digits[i]
					maxPos = i
				}
			}

			selected = append(selected, maxDigit)
			pos = maxPos + 1
			need--
		}

		value := 0
		for _, d := range selected {
			value = value*10 + d
		}

		count += value
	}
	return count
}
