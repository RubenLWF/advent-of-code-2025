package main

import (
	"testing"
)

const example = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func TestPart1(t *testing.T) {
	lines := parseExample(example)
	got := solvePart1(lines, 10)
	want := 40

	if got != want {
		t.Errorf("solvePart1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines := parseExample(example)
	got := solvePart2(lines)
	want := 25272

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
