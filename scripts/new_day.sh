#!/bin/bash

# Script to create a new day folder for Advent of Code 2025

if [ -z "$1" ]; then
    echo "Usage: ./new_day.sh <day>"
    echo "Example: ./new_day.sh 01"
    exit 1
fi

DAY=$1

# Ensure day is zero-padded (01, 02, etc.)
DAY=$(printf "%02d" $((10#$DAY)))

DIR="day${DAY}"

if [ -d "$DIR" ]; then
    echo "Error: Directory $DIR already exists"
    exit 1
fi

echo "Creating directory structure for Day $DAY..."

# Create directory
mkdir -p "$DIR"

# Copy template files
cp template/main.go "$DIR/main.go"
cp template/main_test.go "$DIR/main_test.go"
cp template/input.txt "$DIR/input.txt"
cp template/README.md "$DIR/README.md"

# Update README with day number
sed -i "s/Day X/Day $DAY/" "$DIR/README.md"
sed -i "s/DAY=XX/DAY=$DAY/" "$DIR/README.md"

echo "✓ Created $DIR/"
