package main

import (
	"testing"
)

const example = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {
	lines := parseExample(example)
	got := solvePart1(lines)
	want := 357

	if got != want {
		t.Errorf("solvePart1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines := parseExample(example)
	got := solvePart2(lines)
	want := 3121910778619

	if got != want {
		t.Errorf("solvePart2() = %d, want %d", got, want)
	}
}

func parseExample(s string) []string {
	lines := []string{}
	for _, line := range splitLines(s) {
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func splitLines(s string) []string {
	result := []string{}
	current := ""
	for _, ch := range s {
		if ch == '\n' {
			result = append(result, current)
			current = ""
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}
