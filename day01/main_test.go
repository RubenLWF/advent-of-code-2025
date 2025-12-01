package main

import (
	"testing"
)

const example = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
R1000`

func TestPart1(t *testing.T) {
	lines := parseExample(example)
	got := solvePart1(lines)
	want := 3

	if got != want {
		t.Errorf("solvePart1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines := parseExample(example)
	got := solvePart2(lines)
	want := 16

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
