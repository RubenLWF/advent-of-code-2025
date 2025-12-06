# Day 06: Trash Compactor

## Solution

### Part 1

I ended up overcomplicating part 1 a bit, which ended up helping me for part 2.

For both parts, first split up into separate problems.

This is done by iterating over the columns, and checking if all cells in the column have the value `' '`. While traversing the columns, keep track of the current problem, which we add to the `problems` list as soon as we hit a seperator.

Now we are going to actually compute the result of all problems. For this, we loop over the problems, then over the rows, and then over the colums. This will let us build the numbers like how a human would read them, since they have a strange format in `problems`. (`[1  * 24   356 ]` -> `123 * 45 * 6`)

Next, we take the opator, which is in the final row, and we convert the numbers to integers so we can actually perform computations with them.

Based on the operator, we take the sum or the product of all numbers.

### Part 2

Again, first get the seperate problems.

For each problem, first take the operator.

Loop over each column. Check each cell, if the cell does not contain `' '`, append the value of the cell to the current number. Once all rows have been checked, convert the current number to an integer.

Based on the operator, we take the sum or the product of all numbers.
