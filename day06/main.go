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

func parseProblems(lines []string, width int) [][]string {
	problems := [][]string{}
	currentProblem := []string{}

	for col := 0; col < width; col++ {
		separator := true
		for row := 0; row < len(lines); row++ {
			if col < len(lines[row]) && lines[row][col] != ' ' {
				separator = false
				break
			}
		}

		if separator {
			problems = append(problems, currentProblem)
			currentProblem = []string{}
		} else {
			columnData := ""
			for row := 0; row < len(lines); row++ {
				if col < len(lines[row]) {
					columnData += string(lines[row][col])
				}
			}
			currentProblem = append(currentProblem, columnData)
		}
	}

	problems = append(problems, currentProblem)

	return problems
}

func getResult(numbers []int, operator string) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		switch operator {
		case "+":
			result += numbers[i]
		case "*":
			result *= numbers[i]
		}
	}
	return result
}

func solvePart1(lines []string) int {
	width := len(lines[0])

	problems := parseProblems(lines, width)

	count := 0
	numRows := len(problems[0][0])
	for _, problem := range problems {
		rows := []string{}
		for row := 0; row < numRows; row++ {
			rowStr := ""
			for col := 0; col < len(problem); col++ {
				rowStr += string(problem[col][row])
			}
			rows = append(rows, strings.TrimSpace(rowStr))
		}

		operator := rows[len(rows)-1]
		numbers := []int{}
		for i := 0; i < len(rows)-1; i++ {
			num, _ := strconv.Atoi(rows[i])
			numbers = append(numbers, num)
		}

		count += getResult(numbers, operator)
	}

	return count
}

func solvePart2(lines []string) int {
	width := len(lines[0])

	problems := parseProblems(lines, width)

	count := 0
	numRows := len(problems[0][0])
	for _, problem := range problems {
		operator := string(problem[0][numRows-1])

		numbers := []int{}
		for col := len(problem) - 1; col >= 0; col-- {
			numStr := ""
			for row := 0; row < numRows-1; row++ {
				ch := problem[col][row]
				if ch != ' ' {
					numStr += string(ch)
				}
			}
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}

		count += getResult(numbers, operator)
	}

	return count
}
