# Day 03: Lobby

## Solution

### Part 1

For both parts, first convert to an array of integers.

Part 1 is easily solved by a nested loop where the outer pointer `i` is the left-most bound of the inner pointer `j`.

Store the highest number found so far. If `digits[i]*10 + digits[j]` is greater then the current stored max, replace it.

Return the sum of the max for each line.

### Part 2

To find the max number with 12 digits, greedily pick digits from left to right.

At each step, find highest digit in valid range. We can't pick a digit too late in the number, since this would prevents us from selecting 12 digits in total.

## Running

```bash
# Run solution
cd ./day01

go run .

# Run tests
cd ./day01

go test .
```
