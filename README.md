# Advent of Code

This repo contains my solutions to the [Advent of Code](https://adventofcode.com/) puzzles.

# Solution format

```
<year>
 └───<day>
     ├───leaderboard
     │   ├───1
     │   └───2
     └───optimized
         ├───1
         └───2
```

| Directory     | Meaning                                                                                                      |
|---------------|--------------------------------------------------------------------------------------------------------------|
| `year`        | Holds solutions for all puzzles in that year.                                                                |
| `day`         | Holds solutions for all puzzles on that day.                                                                 |
| `leaderboard` | A quick attempt at the puzzle to try to get onto the leaderboards.                                           |
| `optimized`   | The end result of optimizing the solution. Primarily for readability, secondarily for time/space complexity. |
| `1`           | Solution to the first part of the puzzle.                                                                    |
| `2`           | Solution to the second part of the puzzle.                                                                   |

# Notes

## Running

The working directory needs to be the root folder of the repository, i.e. `AdventOfCode`, to run both solutions and scripts.

## Organization

Each day is partitioned into leaderboard and optimized solutions. The optimized solution usually builds on top of the leaderboard solution.

Each part of each solution is partitioned into its own folder, where only that part is solved.

# Scripts

## Generating new solution directories

`go run cmd/template/main.go` without arguments (or `make template`) will infer which directory to generate by looking at the last completed puzzle and generating the next one.
E.g. if 2021/5/leaderboard exists, it'll generate 2021/5/optimized. Once that exists, it'll generate 2021/6/leaderboard, and so on.

When inferring arguments, solutions to skip can be specified through `skip.txt`.
Each line is of the format `year/day # optional comment`, and if day is omitted the entire year is skipped.
`year` or `day` can either be a single number `num`, or a range like `num-num` to exclude multiple values.

The arguments can be provided through the command line as well. See `go run cmd/template/main.go --help` for argument info.

## Updating the completion tables

`go run cmd/readme/main.go` (or `make readme`) will update the completion tables below based on the directory structure of the repository. It takes no arguments.

## Doing both

`make` will infer a new day, create it, and then update the readme.

## Running tests

`make test` will run tests on the command scripts.

# Completion

## 2022

|             | 1                                                                   | 2                                                                   | 3                                                                   | 4                                                                   | 5                                                                   | 6                                                                   | 7                                                                   | 8                                                                   | 9                                                                   | 10                                                                    | 11                                                                    | 12                                                                    | 13                                                                    | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 | 25 |
|-------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|----|----|----|----|----|----|----|----|----|----|----|----|
| leaderboard | [1](2022/1/leaderboard/1/main.go),[2](2022/1/leaderboard/2/main.go) | [1](2022/2/leaderboard/1/main.go),[2](2022/2/leaderboard/2/main.go) | [1](2022/3/leaderboard/1/main.go),[2](2022/3/leaderboard/2/main.go) | [1](2022/4/leaderboard/1/main.go),[2](2022/4/leaderboard/2/main.go) | [1](2022/5/leaderboard/1/main.go),[2](2022/5/leaderboard/2/main.go) | [1](2022/6/leaderboard/1/main.go),[2](2022/6/leaderboard/2/main.go) | [1](2022/7/leaderboard/1/main.go),[2](2022/7/leaderboard/2/main.go) | [1](2022/8/leaderboard/1/main.go),[2](2022/8/leaderboard/2/main.go) | [1](2022/9/leaderboard/1/main.go),[2](2022/9/leaderboard/2/main.go) | [1](2022/10/leaderboard/1/main.go),[2](2022/10/leaderboard/2/main.go) | [1](2022/11/leaderboard/1/main.go),[2](2022/11/leaderboard/2/main.go) | [1](2022/12/leaderboard/1/main.go),[2](2022/12/leaderboard/2/main.go) | [1](2022/13/leaderboard/1/main.go),[2](2022/13/leaderboard/2/main.go) |    |    |    |    |    |    |    |    |    |    |    |    |
| optimized   | [1](2022/1/optimized/1/main.go),[2](2022/1/optimized/2/main.go)     | [1](2022/2/optimized/1/main.go),[2](2022/2/optimized/2/main.go)     | [1](2022/3/optimized/1/main.go),[2](2022/3/optimized/2/main.go)     | [1](2022/4/optimized/1/main.go),[2](2022/4/optimized/2/main.go)     | [1](2022/5/optimized/1/main.go),[2](2022/5/optimized/2/main.go)     | [1](2022/6/optimized/1/main.go),[2](2022/6/optimized/2/main.go)     | [1](2022/7/optimized/1/main.go),[2](2022/7/optimized/2/main.go)     | [1](2022/8/optimized/1/main.go),[2](2022/8/optimized/2/main.go)     |                                                                     | [1](2022/10/optimized/1/main.go),[2](2022/10/optimized/2/main.go)     | [1](2022/11/optimized/1/main.go),[2](2022/11/optimized/2/main.go)     | [1](2022/12/optimized/1/main.go),[2](2022/12/optimized/2/main.go)     |                                                                       |    |    |    |    |    |    |    |    |    |    |    |    |

## 2021

|             | 1                                                                   | 2                                                                   | 3                                                                   | 4                                                                   | 5                                                                   | 6                                                                   | 7                                                                   | 8                                                                   | 9                                                                   | 10                                                                    | 11                                                                    | 12                                                                    | 13                                                                    | 14                                                                    | 15                                                                    | 16                                                                    | 17                                                                    | 18                                                                    | 19                                                                    | 20                                                                    | 21                                                                    | 22                                                                    | 23                                                                    | 24                                                                    | 25                                 |
|-------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|------------------------------------|
| leaderboard | [1](2021/1/leaderboard/1/main.go),[2](2021/1/leaderboard/2/main.go) | [1](2021/2/leaderboard/1/main.go),[2](2021/2/leaderboard/2/main.go) | [1](2021/3/leaderboard/1/main.go),[2](2021/3/leaderboard/2/main.go) | [1](2021/4/leaderboard/1/main.go),[2](2021/4/leaderboard/2/main.go) | [1](2021/5/leaderboard/1/main.go),[2](2021/5/leaderboard/2/main.go) | [1](2021/6/leaderboard/1/main.go),[2](2021/6/leaderboard/2/main.go) | [1](2021/7/leaderboard/1/main.go),[2](2021/7/leaderboard/2/main.go) | [1](2021/8/leaderboard/1/main.go),[2](2021/8/leaderboard/2/main.go) | [1](2021/9/leaderboard/1/main.go),[2](2021/9/leaderboard/2/main.go) | [1](2021/10/leaderboard/1/main.go),[2](2021/10/leaderboard/2/main.go) | [1](2021/11/leaderboard/1/main.go),[2](2021/11/leaderboard/2/main.go) | [1](2021/12/leaderboard/1/main.go),[2](2021/12/leaderboard/2/main.go) | [1](2021/13/leaderboard/1/main.go),[2](2021/13/leaderboard/2/main.go) | [1](2021/14/leaderboard/1/main.go),[2](2021/14/leaderboard/2/main.go) | [1](2021/15/leaderboard/1/main.go),[2](2021/15/leaderboard/2/main.go) | [1](2021/16/leaderboard/1/main.go),[2](2021/16/leaderboard/2/main.go) | [1](2021/17/leaderboard/1/main.go),[2](2021/17/leaderboard/2/main.go) | [1](2021/18/leaderboard/1/main.go),[2](2021/18/leaderboard/2/main.go) | [1](2021/19/leaderboard/1/main.go),[2](2021/19/leaderboard/2/main.go) | [1](2021/20/leaderboard/1/main.go),[2](2021/20/leaderboard/2/main.go) | [1](2021/21/leaderboard/1/main.go),[2](2021/21/leaderboard/2/main.go) | [1](2021/22/leaderboard/1/main.go),[2](2021/22/leaderboard/2/main.go) | [1](2021/23/leaderboard/1/main.go),[2](2021/23/leaderboard/2/main.go) | [1](2021/24/leaderboard/1/main.go),[2](2021/24/leaderboard/2/main.go) | [1](2021/25/leaderboard/1/main.go) |
| optimized   | [1](2021/1/optimized/1/main.go),[2](2021/1/optimized/2/main.go)     | [1](2021/2/optimized/1/main.go),[2](2021/2/optimized/2/main.go)     | [1](2021/3/optimized/1/main.go),[2](2021/3/optimized/2/main.go)     | [1](2021/4/optimized/1/main.go),[2](2021/4/optimized/2/main.go)     | [1](2021/5/optimized/1/main.go),[2](2021/5/optimized/2/main.go)     | [1](2021/6/optimized/1/main.go),[2](2021/6/optimized/2/main.go)     | [1](2021/7/optimized/1/main.go),[2](2021/7/optimized/2/main.go)     | [1](2021/8/optimized/1/main.go),[2](2021/8/optimized/2/main.go)     | [1](2021/9/optimized/1/main.go),[2](2021/9/optimized/2/main.go)     | [1](2021/10/optimized/1/main.go),[2](2021/10/optimized/2/main.go)     | [1](2021/11/optimized/1/main.go),[2](2021/11/optimized/2/main.go)     | [1](2021/12/optimized/1/main.go),[2](2021/12/optimized/2/main.go)     | [1](2021/13/optimized/1/main.go),[2](2021/13/optimized/2/main.go)     | [1](2021/14/optimized/1/main.go),[2](2021/14/optimized/2/main.go)     |                                                                       | [1](2021/16/optimized/1/main.go),[2](2021/16/optimized/2/main.go)     | [1](2021/17/optimized/1/main.go),[2](2021/17/optimized/2/main.go)     |                                                                       |                                                                       |                                                                       |                                                                       |                                                                       |                                                                       | [1](2021/24/optimized/1/main.go),[2](2021/24/optimized/2/main.go)     | [1](2021/25/optimized/1/main.go)   |

## 2020

|             | 1                                                                   | 2                                                                   | 3                                                                   | 4                                                                   | 5                                                                   | 6                                                                   | 7                                                                   | 8                                                                   | 9                                                                   | 10                                                                    | 11                                                                    | 12                                                                    | 13                                                                    | 14                                                                    | 15                                                                    | 16                                                                    | 17                                                                    | 18                                                                    | 19                                                                    | 20                                                                    | 21                                                                    | 22                                                                    | 23                                                                    | 24                                                                    | 25                                 |
|-------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|------------------------------------|
| leaderboard | [1](2020/1/leaderboard/1/main.go),[2](2020/1/leaderboard/2/main.go) | [1](2020/2/leaderboard/1/main.go),[2](2020/2/leaderboard/2/main.go) | [1](2020/3/leaderboard/1/main.go),[2](2020/3/leaderboard/2/main.go) | [1](2020/4/leaderboard/1/main.go),[2](2020/4/leaderboard/2/main.go) | [1](2020/5/leaderboard/1/main.go),[2](2020/5/leaderboard/2/main.go) | [1](2020/6/leaderboard/1/main.go),[2](2020/6/leaderboard/2/main.go) | [1](2020/7/leaderboard/1/main.go),[2](2020/7/leaderboard/2/main.go) | [1](2020/8/leaderboard/1/main.go),[2](2020/8/leaderboard/2/main.go) | [1](2020/9/leaderboard/1/main.go),[2](2020/9/leaderboard/2/main.go) | [1](2020/10/leaderboard/1/main.go),[2](2020/10/leaderboard/2/main.go) | [1](2020/11/leaderboard/1/main.go),[2](2020/11/leaderboard/2/main.go) | [1](2020/12/leaderboard/1/main.go),[2](2020/12/leaderboard/2/main.go) | [1](2020/13/leaderboard/1/main.go),[2](2020/13/leaderboard/2/main.go) | [1](2020/14/leaderboard/1/main.go),[2](2020/14/leaderboard/2/main.go) | [1](2020/15/leaderboard/1/main.go),[2](2020/15/leaderboard/2/main.go) | [1](2020/16/leaderboard/1/main.go),[2](2020/16/leaderboard/2/main.go) | [1](2020/17/leaderboard/1/main.go),[2](2020/17/leaderboard/2/main.go) | [1](2020/18/leaderboard/1/main.go),[2](2020/18/leaderboard/2/main.go) | [1](2020/19/leaderboard/1/main.go),[2](2020/19/leaderboard/2/main.go) | [1](2020/20/leaderboard/1/main.go),[2](2020/20/leaderboard/2/main.go) | [1](2020/21/leaderboard/1/main.go),[2](2020/21/leaderboard/2/main.go) | [1](2020/22/leaderboard/1/main.go),[2](2020/22/leaderboard/2/main.go) | [1](2020/23/leaderboard/1/main.go),[2](2020/23/leaderboard/2/main.go) | [1](2020/24/leaderboard/1/main.go),[2](2020/24/leaderboard/2/main.go) | [1](2020/25/leaderboard/1/main.go) |
| optimized   | [1](2020/1/optimized/1/main.go),[2](2020/1/optimized/2/main.go)     | [1](2020/2/optimized/1/main.go),[2](2020/2/optimized/2/main.go)     | [1](2020/3/optimized/1/main.go),[2](2020/3/optimized/2/main.go)     | [1](2020/4/optimized/1/main.go),[2](2020/4/optimized/2/main.go)     | [1](2020/5/optimized/1/main.go),[2](2020/5/optimized/2/main.go)     | [1](2020/6/optimized/1/main.go),[2](2020/6/optimized/2/main.go)     | [1](2020/7/optimized/1/main.go),[2](2020/7/optimized/2/main.go)     | [1](2020/8/optimized/1/main.go),[2](2020/8/optimized/2/main.go)     | [1](2020/9/optimized/1/main.go),[2](2020/9/optimized/2/main.go)     | [1](2020/10/optimized/1/main.go),[2](2020/10/optimized/2/main.go)     | [1](2020/11/optimized/1/main.go),[2](2020/11/optimized/2/main.go)     | [1](2020/12/optimized/1/main.go),[2](2020/12/optimized/2/main.go)     | [1](2020/13/optimized/1/main.go),[2](2020/13/optimized/2/main.go)     |                                                                       |                                                                       | [1](2020/16/optimized/1/main.go),[2](2020/16/optimized/2/main.go)     | [1](2020/17/optimized/1/main.go)                                      | [1](2020/18/optimized/1/main.go)                                      |                                                                       |                                                                       |                                                                       |                                                                       | [1](2020/23/optimized/1/main.go),[2](2020/23/optimized/2/main.go)     |                                                                       | [1](2020/25/optimized/1/main.go)   |

## 2019

|             | 1                                                                   | 2                                                                   | 3                                                                   | 4                                                                   | 5                               | 6 | 7 | 8 | 9                               | 10 | 11 | 12 | 13                               | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 | 25 |
|-------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------|---|---|---|---------------------------------|----|----|----|----------------------------------|----|----|----|----|----|----|----|----|----|----|----|----|
| leaderboard | [1](2019/1/leaderboard/1/main.go),[2](2019/1/leaderboard/2/main.go) | [1](2019/2/leaderboard/1/main.go),[2](2019/2/leaderboard/2/main.go) | [1](2019/3/leaderboard/1/main.go),[2](2019/3/leaderboard/2/main.go) | [1](2019/4/leaderboard/1/main.go),[2](2019/4/leaderboard/2/main.go) |                                 |   |   |   |                                 |    |    |    |                                  |    |    |    |    |    |    |    |    |    |    |    |    |
| optimized   |                                                                     |                                                                     |                                                                     | [1](2019/4/optimized/1/main.go),[2](2019/4/optimized/2/main.go)     | [2](2019/5/optimized/2/main.go) |   |   |   | [2](2019/9/optimized/2/main.go) |    |    |    | [2](2019/13/optimized/2/main.go) |    |    |    |    |    |    |    |    |    |    |    |    |

## 2016

|             | 1                                                                   | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 | 25 |
|-------------|---------------------------------------------------------------------|---|---|---|---|---|---|---|---|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|
| leaderboard | [1](2016/1/leaderboard/1/main.go),[2](2016/1/leaderboard/2/main.go) |   |   |   |   |   |   |   |   |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |
| optimized   |                                                                     |   |   |   |   |   |   |   |   |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |

## 2015

|             | 1                                                                   | 2                                                                   | 3                                                                   | 4                                                                   | 5                                                                   | 6                                                                   | 7                                                                   | 8                                                                   | 9                                                                   | 10                                                                    | 11                                                                    | 12                                                                    | 13                                                                    | 14                                                                    | 15                                                                    | 16                                                                    | 17                                                                    | 18                                                                    | 19                                                                    | 20                                                                    | 21 | 22 | 23 | 24 | 25 |
|-------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|----|----|----|----|----|
| leaderboard | [1](2015/1/leaderboard/1/main.go),[2](2015/1/leaderboard/2/main.go) | [1](2015/2/leaderboard/1/main.go),[2](2015/2/leaderboard/2/main.go) | [1](2015/3/leaderboard/1/main.go),[2](2015/3/leaderboard/2/main.go) | [1](2015/4/leaderboard/1/main.go),[2](2015/4/leaderboard/2/main.go) | [1](2015/5/leaderboard/1/main.go),[2](2015/5/leaderboard/2/main.go) | [1](2015/6/leaderboard/1/main.go),[2](2015/6/leaderboard/2/main.go) | [1](2015/7/leaderboard/1/main.go),[2](2015/7/leaderboard/2/main.go) | [1](2015/8/leaderboard/1/main.go),[2](2015/8/leaderboard/2/main.go) | [1](2015/9/leaderboard/1/main.go),[2](2015/9/leaderboard/2/main.go) | [1](2015/10/leaderboard/1/main.go),[2](2015/10/leaderboard/2/main.go) | [1](2015/11/leaderboard/1/main.go),[2](2015/11/leaderboard/2/main.go) | [1](2015/12/leaderboard/1/main.go),[2](2015/12/leaderboard/2/main.go) | [1](2015/13/leaderboard/1/main.go),[2](2015/13/leaderboard/2/main.go) | [1](2015/14/leaderboard/1/main.go),[2](2015/14/leaderboard/2/main.go) | [1](2015/15/leaderboard/1/main.go),[2](2015/15/leaderboard/2/main.go) | [1](2015/16/leaderboard/1/main.go),[2](2015/16/leaderboard/2/main.go) | [1](2015/17/leaderboard/1/main.go),[2](2015/17/leaderboard/2/main.go) | [1](2015/18/leaderboard/1/main.go),[2](2015/18/leaderboard/2/main.go) | [1](2015/19/leaderboard/1/main.go),[2](2015/19/leaderboard/2/main.go) | [1](2015/20/leaderboard/1/main.go),[2](2015/20/leaderboard/2/main.go) |    |    |    |    |    |
| optimized   |                                                                     |                                                                     |                                                                     |                                                                     |                                                                     |                                                                     |                                                                     |                                                                     |                                                                     |                                                                       |                                                                       |                                                                       |                                                                       |                                                                       |                                                                       |                                                                       |                                                                       |                                                                       |                                                                       |                                                                       |    |    |    |    |    |
