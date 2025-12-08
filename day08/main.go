package main

import (
	"fmt"
	"log"
	"math"
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
	part1 := solvePart1(lines, 1000)
	part2 := solvePart2(lines)

	// Print results
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

type Point3D struct {
	X, Y, Z int
}

type Edge struct {
	Distance float64
	I, J     int
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	uf.parent[rootY] = rootX
	uf.size[rootX] += uf.size[rootY]
	return true
}

func (uf *UnionFind) GetComponentSizes() []int {
	componentSizes := make(map[int]int)

	for i := range uf.parent {
		root := uf.Find(i)
		componentSizes[root] = uf.size[root]
	}

	sizes := make([]int, 0, len(componentSizes))
	for _, size := range componentSizes {
		sizes = append(sizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes
}

func distance(p1, p2 Point3D) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	dz := float64(p1.Z - p2.Z)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func parsePoint(line string) Point3D {
	coordinates := strings.Split(line, ",")

	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	z, _ := strconv.Atoi(coordinates[2])

	return Point3D{X: x, Y: y, Z: z}
}

func parsePoints(lines []string) []Point3D {
	var points []Point3D
	for _, line := range lines {
		point := parsePoint(line)
		points = append(points, point)
	}
	return points
}

func getSortedEdges(points []Point3D) []Edge {
	n := len(points)
	var edges []Edge
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := distance(points[i], points[j])
			edges = append(edges, Edge{Distance: dist, I: i, J: j})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Distance < edges[j].Distance
	})

	return edges
}

func solvePart1(lines []string, numConnections int) int {
	points := parsePoints(lines)

	n := len(points)

	edges := getSortedEdges(points)

	uf := NewUnionFind(n)

	for i := 0; i < numConnections; i++ {
		uf.Union(edges[i].I, edges[i].J)
	}

	circuitSizes := uf.GetComponentSizes()

	return circuitSizes[0] * circuitSizes[1] * circuitSizes[2]
}

func solvePart2(lines []string) int {
	points := parsePoints(lines)

	n := len(points)

	edges := getSortedEdges(points)

	uf := NewUnionFind(n)

	var lastEdge Edge
	for _, edge := range edges {
		if uf.Union(edge.I, edge.J) {
			lastEdge = edge
			if uf.size[uf.Find(edge.I)] == n {
				break
			}
		}
	}

	return points[lastEdge.I].X * points[lastEdge.J].X
}
