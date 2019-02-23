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

## Day 4

A pretty interesting problem, but sort of hard to do. FOr the 2nd half, I really just wanted to use sql:
`select letter from table order by count desc, letter asc`

## Day 5

So, Go's md5 implmenetation returns back a `[16]byte`, but Go's hex takes `[]byte`, which is super annoying
But, there's an easy way to do this conversaion: use `rtn[:]` to turn it into a byte slick pretty easily

## Day 9

This one is a bit weird, and focuses on scanning and keeping track of where you are. This ended up not being as hard as I figured it would be. The big challenge for me here was making sure that I was keep track of which character I was in, while factoring what the next loop iteration was going to do. 
Also, I need a way to keep that regex global, or relatively global so that I don't need to redefine it on every call.

I hate the while syntax here. Seems pretty terrible, really. I suppose that's why most other languages opt for a while style and a for style? 
Also, I'm pretty tired of using `:= range` for `for` loops. I don't really get why they opted to use `:=` instead of `in`, or alternatively why `range` is required. I nearly always write it without the range, then I need to go back and add it.

## Day 10

This was a pain. I'm not super happy with my code, so I may make some revisions later on.

The basic flow is this:
  * Figure out eah of the relationships for each bot/output and their station positions
  * Encode that logic somewhere (A struct here, with one reciever method -- a class would probably work better here)
  * Find any bot that has 2 elements -- thankfully there's only one
  * Execute the logic (recursively) until everything ends up in an output. 

The big problem I faced here was actually with references. I wanted some shorthand variable to use over and over
without using `*` and `&` (kind of a pain to type on this keyboard), but I ended up making copies without realizing it
Soo, no luck there. I ended up going back and replacing all of the values with references, and then my code worked perfectly

As a result though, there's a fair bit of debug code here that needs to be cleaned out.