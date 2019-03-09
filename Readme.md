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

## Day 11

No clue what to do here.

## Day 12

I coded this up correctly on the first pass, but it seemed like it didn't work, since it was pretty slow

I ended up doing two things to make this work faster:

1. Cheat the number identification. I originally used a regular expression (not even a compiled one), but figured it would be faster to simply check to see that the first character looked like it could be a number, and not worry about if it was error prone. I didn't test, but this had a large impact on performance. it could probably be further optimized here as well
2. Pre-pare out the steps. Initially I split the inputs on each pass, which ended up being needless.

A 3rd improvement would be to memoize the relative results from a (or maybe each) cell, up until a jump. This one is more complicated, but could maybe save a number of iterations through the loop, and focus only on the parts that matter (namely, jumps)

## Day 15

I think there's a faster way to do this, possibly having to use the LCM of the discs. The idea of the quicker way is this:

1. You've figured out how to drop it through n holes. You can repeat this by only choosing LCMs of the n discs, with each multiple being the number of positions in contains (in my example, the first 3 discs: `17 * 19 * 7`). 
2. When you use multiples of the LCM, you are guarenteed to pass through the same n discs. 
3. Each LCM is going to advance the next disc some number of values. You need to figure out what that next number is. Once you have that, you can advance the guess count until you find a match
4. once you find that match, you should now be able to drop the capsule n+1 steps (the first n pass, because we always increment by lcms, the next, because the modulous finally works)


## Day 18

This one was actually pretty straight forward -- I appreciate the puzzlesa in which they give explicit rules on how to transform the data. 
However, part 2 is kind of weird. Normally part 2 requires some kind of code change. Technically, this did require a code change -- a 40 to 400000
but this isn't very significant. But, I guess what they were going for was maybe having some people run into problems with memory consumption? As a psuedo code solutino for that (actually just python), you can instead not build out the grid, and instead look at only two rows. the logic looks something like this:

```python
def calc_safe_squares(start_row, num_rows):
  def calc_safe_in_row(row):
    return sum([1 for c in s if c =='.'])

  last_row = start_row
  row_sum = 0
  for i in range(num_rows-1):
    row_sum += calc_safe_in_row(last_row)
    last_row = build_row(last_row)  # same logic as before -- look at the previous row based off of the established rules
  return row_sum
```

## Day 19

I actually saw a youtube video on this: https://www.youtube.com/watch?v=uCsD3ZGzMgE
Short version: we can wrtie the number as `2^n + l` (the biggest power of 2 less than the number, plus the remainder)
If we do this, then the "winning" seat will be `2l + 1`
Funny note: this video was published on Oct 28 2016, for the 2016 puzzle, so maybe a puzzle creator took inspiration from this?

Part 2 was a bit more difficult. I ended up having to implement a brute force solution, but that was going to take far too long to execute.
I ended up looking at the numbers again and noticed the following pattern:
for every n that can be evenly represted as 3^n, the winning seat is the last seat (i.e. `W(n) = n`)
From there, solutions progress at slope = 1 for some period of time, before switching to slope=2. These switches occur
on the "middle" multiple of 3, or expressed differently, at the midpoint between `3^n` and `3^(n-1)`. My math capabitiles have
apparently deteriated to the point that I couldn't merge these into a single equation, nor am I sure that a single equation exists.
So, in my solution, I instead just did a quick check based on where the midpoint is, relative to the encapsulating powers of 3, and checked
which side was closer, and determined the result from there.