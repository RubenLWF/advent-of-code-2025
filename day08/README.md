# Day 08: Playground

## Solution

### Part 1

For both parts, first get the points and find all edges between these points. Next, initialize an instance of UnionFind. This implementation of UnionFind not only keeps track of the parent, but also of the size of the current circuit.

Try to make the first 10 connections. To do this, first get the root of both circuits the junction boxes are in.
If these are the same, no new connection has to be made.
Add the circuits together, meaning all boxes in the circuit now have the same root.

Find the sizes of the 3 largest circuits and multiply them together to get the result for part 1.

### Part 2

First, like part 1, parse the points, get the sorted edges, and create an instance of UnionFind.

Add the shortest edge which adds a new circuit box to the circuit until the size of the largest circuit is equal to the amount of circuit boxes, meaning all boxes are now in a single network.

Multiply the x coordinate of the two circuit boxes that were connected last to find the result of part 2.
