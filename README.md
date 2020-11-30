# Advent of Code

This repo contains my solutions to the [Advent of Code](https://adventofcode.com/) puzzles.

# Format

```
<year>
└───<day>
    ├───initial
    │   ├───1
    │   └───2
    ├───final
    │   ├───1
    │   └───2
    └───input.txt
```

Where:
- `initial` is my first attempt at solving the puzzle; usually as quickly as possible to get on leaderboards
- `final` is the end result of cleaning up code/optimizing the actual solution runtime and may include benchmarks
  - not all puzzles have one; i.e. if I didn't want to continue work on the puzzle
- `1`/`2` indicate the first/second puzzle for that day
- `input.txt` is my puzzle input
  - not all puzzles have one; i.e. if the input is small it's instead embedded in the code directly
  - to run solutions that do: set your working directory to the containing folder, e.g. `AdventOfCode/2019/1`.