# Day 04: Printing Department

## Solution

### Part 1

I defined an array containing the relative indices of the 8 directions

For both parts, first convert the input to a grid of runes, since this is way easier to work with than an array of strings.

For each cell in the grid, check whether it contains a roll. If this is the case, take the sum of all, non out of bounds, neighboring cells that contain a roll. If this number < 4, the roll can be picked up.

Take the count of all rolls where this is the case, this is the answer for part 1.

### Part 2

Perform the following steps until we break out of the for loop.

Create a new array, `toRemove`, this stores the coordinates of all the rolls tha should be removed after this iteration.

Run the same checks as for part 1, only now, rather than incrementing the count if the number of adjecent cells containing a roll < 4, add the coordinates of the cell to `toRemove`.

Check if `toRemove` contains any items.

- If this is not the case, we are done and we break out of the loop.
- If this is the case, add the length of `toRemove` to the count, and remove the rolls on all coordinates in `toRemove`. Start a new iteration of the loop.

Return the count.
