# Advent of Code

This repo contains solutions to the [Advent of Code](https://adventofcode.com/) puzzles.

# Format

```
year
└───puzzle number
    ├───initial
    │   ├───part 1
    │   └───part 2
    ├───final
    │   ├───part 1
    │   └───part 2
    └───input.txt
```

Where:
- `initial` is my first attempt at solving the puzzle; usually as quickly as possible to get on leaderboards
- `final` is the end result of cleaning up code/optimizing the actual solution runtime and may include benchmarks
  - not all puzzles have a `final` folder (if I didn't want to clean up/optimize it)
- `input.txt` is my puzzle input, if the test has any large enough that it's easier to put it in a file rather than in code
  - to run solutions that have puzzle input files, set your working directory to the puzzle root, e.g. `2019/1`.