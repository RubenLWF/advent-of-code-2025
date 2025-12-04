# Day 01: Secret Entrance

## Solution

### Part 1

Use to first character of each line to see if the rotation is positive or negative.

Based in this, either take the value as-is, or make it negative.

Add the value to the current pointer and take `value mod 100` to get the real value of the pointer.

If the pointer is at 0 at the end of a move, increase count by 1.

### Part 2

Use the same method as part 1 to decide whether a move is positive er negative.

Divide the total rotation amount by 100, taking the quotient and the remainder.

Add the quotient to the count directly.

Use the first character of the string to decide whether the remainder should be used as-is, or whether it should be negated.

Perform some logic to see whether the pointer passes over 0 in this move. This is done by checking if `pointer == 0` and if the result of adding the value would put it out of bounds.

Apply `value mod 100`.

See if the pointer ends up at 0. If so, increase the count by 1.
