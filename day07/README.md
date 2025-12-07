# Day 07: Laboratories

## Solution

### Part 1

First build the grid from the input and locate the start position.

Simulate beams that move down one row at a time.

When a beam encounters a splitter: The splitter is counted the first time a beam hits it and the beam creates new beams to the left and right (if those columns exist).

Beams that don't hit a splitter simply continue straight down.

Continue the simulation until all beams leave the grid. The answer is the number of unique splitters that were used.

### Part 2

Count how many distinct timelines (paths) reach the bottom. 

I implemented a recursive DFS that moves straight down normally and, when it finds a splitter, creates new branches to the left and right.

Use a memory using the grid position as the key to cache the number of timelines from that cell and avoid recomputation.

Sum the reachable timelines from the start position and return the total.
