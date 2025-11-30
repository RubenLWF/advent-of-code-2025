package utils

// Point represents a 2D coordinate
type Point struct {
	X, Y int
}

// Add returns the sum of two points
func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

// Sub returns the difference of two points
func (p Point) Sub(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}

// ManhattanDistance returns the Manhattan distance between two points
func (p Point) ManhattanDistance(other Point) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

// Neighbors returns the 4 adjacent points (up, down, left, right)
func (p Point) Neighbors() []Point {
	return []Point{
		{X: p.X, Y: p.Y - 1}, // Up
		{X: p.X, Y: p.Y + 1}, // Down
		{X: p.X - 1, Y: p.Y}, // Left
		{X: p.X + 1, Y: p.Y}, // Right
	}
}

// Neighbors8 returns all 8 adjacent points (including diagonals)
func (p Point) Neighbors8() []Point {
	return []Point{
		{X: p.X - 1, Y: p.Y - 1}, // Top-left
		{X: p.X, Y: p.Y - 1},     // Top
		{X: p.X + 1, Y: p.Y - 1}, // Top-right
		{X: p.X - 1, Y: p.Y},     // Left
		{X: p.X + 1, Y: p.Y},     // Right
		{X: p.X - 1, Y: p.Y + 1}, // Bottom-left
		{X: p.X, Y: p.Y + 1},     // Bottom
		{X: p.X + 1, Y: p.Y + 1}, // Bottom-right
	}
}

// Direction vectors for common movements
var (
	North = Point{X: 0, Y: -1}
	South = Point{X: 0, Y: 1}
	East  = Point{X: 1, Y: 0}
	West  = Point{X: -1, Y: 0}
)

// Grid represents a 2D grid with bounds checking
type Grid struct {
	Data   [][]rune
	Width  int
	Height int
}

// NewGrid creates a new grid from lines
func NewGrid(lines []string) *Grid {
	data := ParseGrid(lines)
	return &Grid{
		Data:   data,
		Width:  len(data[0]),
		Height: len(data),
	}
}

// InBounds checks if a point is within the grid
func (g *Grid) InBounds(p Point) bool {
	return p.X >= 0 && p.X < g.Width && p.Y >= 0 && p.Y < g.Height
}

// Get returns the value at a point (if in bounds)
func (g *Grid) Get(p Point) (rune, bool) {
	if !g.InBounds(p) {
		return 0, false
	}
	return g.Data[p.Y][p.X], true
}

// Set sets the value at a point (if in bounds)
func (g *Grid) Set(p Point, value rune) bool {
	if !g.InBounds(p) {
		return false
	}
	g.Data[p.Y][p.X] = value
	return true
}

// Find returns all points containing a specific value
func (g *Grid) Find(target rune) []Point {
	var points []Point
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.Data[y][x] == target {
				points = append(points, Point{X: x, Y: y})
			}
		}
	}
	return points
}
