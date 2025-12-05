package main

import (
	"testing"
)

const example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

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
	want := 14

	if got != want {
		t.Errorf("solvePart2() = %d, want %d", got, want)
	}
}

func parseExample(s string) []string {
	return splitLines(s)
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
