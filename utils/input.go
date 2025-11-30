package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a file and returns all lines as a slice of strings
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// ReadFile reads entire file content as a string
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ReadInts reads a file and returns all lines parsed as integers
func ReadInts(filename string) ([]int, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	var nums []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

// SplitOnEmptyLines splits content into groups separated by empty lines
func SplitOnEmptyLines(content string) []string {
	return strings.Split(strings.TrimSpace(content), "\n\n")
}

// ParseGrid converts lines into a 2D grid of runes
func ParseGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

// ParseIntGrid converts lines into a 2D grid of integers
func ParseIntGrid(lines []string) ([][]int, error) {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, ch := range line {
			num, err := strconv.Atoi(string(ch))
			if err != nil {
				return nil, err
			}
			grid[i][j] = num
		}
	}
	return grid, nil
}
