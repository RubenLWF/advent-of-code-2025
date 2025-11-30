# Advent of Code 2025

This year, I will use Advent of Code to learn Go.

## Project Structure

```
advent-of-code-2025/
├── go.mod                  # Go module definition
├── Makefile                # Build and run commands
├── README.md               # This file
├── utils/                  # Reusable utility functions
│   ├── input.go            # File reading and parsing
│   ├── math.go             # Math utilities 
│   ├── grid.go             # 2D grid operations and Point type
│   └── collections.go      # Set, Queue, Stack implementations
├── template/               # Template files for new days
│   ├── main.go
│   ├── main_test.go
│   ├── README.md
│   └── input.txt
├── scripts/
│   └── new_day.sh          # Script to create new day
└── dayXX/                  # Individual day folders
    ├── main.go
    ├── main_test.go
    ├── input.txt
    └── README.md
```

## Getting Started

### Prerequisites

- Go 1.23 or later
- Make (optional, but recommended)

### Creating a New Day

Use the provided script to quickly scaffold a new day:

```bash
make new DAY=01
# or directly:
./scripts/new_day.sh 01
```

This creates a new `dayXX` folder with:

- `main.go` - Main solution file with part 1 and part 2 functions
- `main_test.go` - Test file for examples
- `input.txt` - Your puzzle input (add manually from AoC website)
- `README.md` - Documentation for your solution

### Running Solutions

```bash
# Run a specific day
make run DAY=01

# Or use go directly
go run ./day01
```

### Running Tests

```bash
# Test all days
make test

# Test a specific day
make test-day DAY=01

# Or use go directly
go test ./day01
```

## Typical Workflow

1. **Create a new day**: `make new DAY=01`
2. **Get your input**: Copy your puzzle input from [adventofcode.com](https://adventofcode.com/2025) to `day01/input.txt`
3. **Add test cases**: Update `day01/main_test.go` with the example from the puzzle description
4. **Implement solution**: Write your solution in `day01/main.go`
5. **Test**: `make test-day DAY=01`
6. **Run**: `make run DAY=01`
7. **Submit**: Copy the answer and submit on AoC website

## Make Commands

```bash
make help       # Show all available commands
make new DAY=XX # Create new day folder
make run DAY=XX # Run solution for specific day
make test       # Run all tests
make test-day DAY=XX # Run tests for specific day
make clean      # Clean build artifacts
```
