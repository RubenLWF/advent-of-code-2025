package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc/utils"
)

func main() {
	// Read input
	file, err := utils.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Solve parts
	part1 := solvePart1(file)
	part2 := solvePart2(file)

	// Print results
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func solvePart1(file string) int {
	count := 0

	ranges := strings.Split(file, ",")

	for i := 0; i < len(ranges); i++ {
		bounds := strings.Split(ranges[i], "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		for j := start; j <= end; j++ {
			str := strconv.Itoa(j)
			len := len(str)
			if str[0] == '0' {
				continue
			}
			if len%2 == 0 {
				mid := len / 2
				left := str[:mid]
				right := str[mid:]
				if left == right {
					count += j
				}
			}
		}
	}
	return count
}

func solvePart2(file string) int {
	count := 0

	ranges := strings.Split(file, ",")

	for i := 0; i < len(ranges); i++ {
		bounds := strings.Split(ranges[i], "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		for j := start; j <= end; j++ {
			str := strconv.Itoa(j)
			if str[0] == '0' {
				continue
			}

			length := len(str)
			found := false

			for partLength := 1; partLength <= length/2; partLength++ {
				if length%partLength != 0 {
					continue
				}

				pattern := str[:partLength]
				isRepeating := true

				for k := partLength; k < length; k += partLength {
					if str[k:k+partLength] != pattern {
						isRepeating = false
						break
					}
				}

				if isRepeating {
					count += j
					found = true
					break
				}
			}

			if found {
				continue
			}
		}
	}
	return count
}
