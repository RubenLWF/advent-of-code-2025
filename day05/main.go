package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

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
	splitIndex := -1
	for i, line := range lines {
		if line == "" {
			splitIndex = i
			break
		}
	}

	ranges := [][2]int{}
	for _, line := range lines[:splitIndex] {
		bounds := strings.Split(line, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, [2]int{start, end})
	}

	ids := []int{}
	for _, line := range lines[splitIndex+1:] {
		id, _ := strconv.Atoi(line)
		ids = append(ids, id)
	}

	count := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				count++
				break
			}
		}
	}
	return count
}

func solvePart2(lines []string) int {
	splitIndex := -1
	for i, line := range lines {
		if line == "" {
			splitIndex = i
			break
		}
	}

	ranges := [][2]int{}
	for _, line := range lines[:splitIndex] {
		bounds := strings.Split(line, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, [2]int{start, end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := [][2]int{}
	for _, r := range ranges {
		length := len(merged)
		if length == 0 || merged[length-1][1] < r[0] {
			merged = append(merged, r)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], r[1])
		}
	}

	count := 0
	for _, r := range merged {
		count += r[1] - r[0] + 1
	}
	return count
}
