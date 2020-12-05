# Advent of Code

This repo contains my solutions to the [Advent of Code](https://adventofcode.com/) puzzles.

# Format

```
<year>
└───<day>
│   ├───initial
│   │   ├───1
│   │   └───2
│   └───optimized
│       ├───1
│       └───2
template
```

| Directory   | Meaning                                                                                                                    |
|-------------|----------------------------------------------------------------------------------------------------------------------------|
| `year`      | The year the puzzle was released.                                                                                          |
| `day`       | The day the puzzle was released.                                                                                           |
| `initial`   | My first attempt at solving the puzzle; usually as quickly as possible (in real time) to get on leaderboards.              |
| `optimized` | The end result of optimizing the solution (for readability, and usually time/space complexity) and may include benchmarks. |
| `1`         | The first part of the puzzle.                                                                                              |
| `2`         | The second part of the puzzle.                                                                                             |
|`template`   | Contains a script to generate new day folders.                                                                             |


# Notes

## Running

Set your working directory to the root folder, i.e. `AdventOfCode`, before running a solution or benchmark.

## Organization

Each solution is self-contained so there's intentional duplication between the initial/optimized first/second solution files.

## Benchmarks

  - Benchmarks exclude the time to load and parse the input because I don't find that particularly interesting.
    - This usually means early day puzzles (which are usually parsing focused) run very fast.
  - The benchmarks measure time/space complexity, but not all solutions have been optimized for time and space complexity.
  - The benchmarks exist mainly because they're fun to look at, not to measure solution quality.

# Completion

## 2020

|           | 1                                                               | 2                                                               | 3                                                               | 4                                                               | 5                                                               | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 | 25 |
|-----------|-----------------------------------------------------------------|-----------------------------------------------------------------|-----------------------------------------------------------------|-----------------------------------------------------------------|-----------------------------------------------------------------|---|---|---|---|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|
| initial   | [1](2020/1/initial/1/main.go),[2](2020/1/initial/2/main.go)     | [1](2020/2/initial/1/main.go),[2](2020/2/initial/2/main.go)     | [1](2020/3/initial/1/main.go),[2](2020/3/initial/2/main.go)     | [1](2020/4/initial/1/main.go),[2](2020/4/initial/2/main.go)     | [1](2020/5/initial/1/main.go),[2](2020/5/initial/2/main.go)     |   |   |   |   |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |
| optimized | [1](2020/1/optimized/1/main.go),[2](2020/1/optimized/2/main.go) | [1](2020/2/optimized/1/main.go),[2](2020/2/optimized/2/main.go) | [1](2020/3/optimized/1/main.go),[2](2020/3/optimized/2/main.go) | [1](2020/4/optimized/1/main.go),[2](2020/4/optimized/2/main.go) | [1](2020/5/optimized/1/main.go),[2](2020/5/optimized/2/main.go) |   |   |   |   |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |

## 2019

|           | 1                                                           | 2                                                           | 3                                                           | 4                                                               | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 | 25 |
|-----------|-------------------------------------------------------------|-------------------------------------------------------------|-------------------------------------------------------------|-----------------------------------------------------------------|---|---|---|---|---|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|
| initial   | [1](2019/1/initial/1/main.go),[2](2019/1/initial/2/main.go) | [1](2019/2/initial/1/main.go),[2](2019/2/initial/2/main.go) | [1](2019/3/initial/1/main.go),[2](2019/3/initial/2/main.go) | [1](2019/4/initial/1/main.go),[2](2019/4/initial/2/main.go)     |   |   |   |   |   |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |
| optimized |                                                             |                                                             |                                                             | [1](2019/4/optimized/1/main.go),[2](2019/4/optimized/2/main.go) |   |   |   |   |   |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |
