# Day 05: Cafeteria

## Solution

### Part 1

For both parts, first find the index the input switches from the ranges to the specific IDs, this can be found by looking for an empty string.

Parse the ranges by splitting each line on `"-"`, take the upper and lower bound and put them in an array.

Parse the specific Ids as integers.

For each ID, check if they are in a range, if this is the case, increment the count by 1.

Return the count.

### Part 2

Find the ranges in the same way as for part 1.

Sort the ranges on start index.

Merge all overlapping ranges, this is done by checking whether the start index of the current range is lower than the end index of the last range added to the list of merged ranges.
- If this is not the case, add the current range to the list of merged ranges.
- If this is the case, set the end index of the range it collides with to the max of the existing end index and the end index of the current range.

For each range, add the amount of indices between the start and end to the count.

Return the count.