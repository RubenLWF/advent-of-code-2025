# Day 02: Gift Shop

## Solution

### Part 1

First split on commas and dashes to get lower and upper bound of each range.

For each item in the range, check if splitting the number in half results in 2 equal parts. If so, add the number to the count.

Strings that start with 0 can be ignored.

### Part 2

Same method to get lower and upper bounds as part 1.

Check for all numbers `i` between `0` and `len/2` if `len%i == 0`. If this is the case, the number can be split up into multiple parts of length `i`.

Take the first `i` characters of the string as the pattern. Match each interval if `i` characters to the pattern, if all intervals match the pattern, add the number to the count.

## Running

```bash
# Run solution
cd ./day01

go run .

# Run tests
cd ./day01

go test .
```
