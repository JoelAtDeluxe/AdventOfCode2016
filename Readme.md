# Advent of Code 2016

Solved using Golang (1.10 / 1.11, if versions come up)

## Puzzle Commentary

### Day 1

Part 1 changed substantially after I had to start working on part 2.
Part 2 was a pain

## Day 2

Part 1 is easy enough, but the challenge is that the array is sort of backwards -- probably my own fault with the indexes
Part 2 is still pretty easy, but it nearly had me rewrite this. One solution I didn't really try much was something along these lines

In python...

```python
movement = {
  1: {'U': 1, 'R': 2, 'D': 4, 'L': 1},
  # ...
  5: {'U': 2, 'R': 6, 'D': 8, 'L': 4},
}
position = 5
for c in directions:
  position = movement[position][direction]

```

Doing this solution would have made this pretty trivial, but might be hard to maintain as the grid grew from 3x3 to, say, 5x5

## Day 3

Go regex isn't as straight forward as I'd like

I also like how the core logic didn't change -- just the parsing logic changed, which helped my code.