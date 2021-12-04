# Advent of Code

This repo contains my solutions to the [Advent of Code](https://adventofcode.com/) puzzles.

# Format

```
┌───<year>
│   └───<day>
│       ├───leaderboard
│       │   ├───1
│       │   └───2
│       └───optimized
│           ├───1
│           └───2
└───cmd
```

| Directory     | Meaning                                                                                                      |
|---------------|--------------------------------------------------------------------------------------------------------------|
| `year`        | The year the puzzle was released.                                                                            |
| `day`         | The day the puzzle was released.                                                                             |
| `leaderboard` | A quick attempt at the puzzle to try to get onto the leaderboards.                                           |
| `optimized`   | The end result of optimizing the solution. Primarily for readability, secondarily for time/space complexity. |
| `1`           | The first part of the puzzle.                                                                                |
| `2`           | The second part of the puzzle.                                                                               |
| `cmd`         | Contains useful scripts (e.g. generating new day directories.)                                               |

# Notes

## Running

The working directory needs to be the root folder of the repository, i.e. `AdventOfCode`, to run a solution.

## Organization

Each day is partitioned into leaderboard and optimized solutions. The optimized solution usually builds on top of the leaderboard solution.

Each part of each solution is partitioned into its own folder, where only that part is solved.

# Completion

## 2021

|             | 1                                                                   | 2                                                                   | 3                                                                   | 4                                                                   | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 |
|-------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---|---|---|---|---|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|
| leaderboard | [1](2021/1/leaderboard/1/main.go),[2](2021/1/leaderboard/2/main.go) | [1](2021/2/leaderboard/1/main.go),[2](2021/2/leaderboard/2/main.go) | [1](2021/3/leaderboard/1/main.go),[2](2021/3/leaderboard/2/main.go) | [1](2021/4/leaderboard/1/main.go),[2](2021/4/leaderboard/2/main.go) |   |   |   |   |   |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |
| optimized   | [1](2021/1/optimized/1/main.go),[2](2021/1/optimized/2/main.go)     | [1](2021/2/optimized/1/main.go),[2](2021/2/optimized/2/main.go)     | [1](2021/3/optimized/1/main.go),[2](2021/3/optimized/2/main.go)     | [1](2021/4/optimized/1/main.go),[2](2021/4/optimized/2/main.go)     |   |   |   |   |   |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |

## 2020

|             | 1                                                                   | 2                                                                   | 3                                                                   | 4                                                                   | 5                                                                   | 6                                                                   | 7                                                                   | 8                                                                   | 9                                                                   | 10                                                                    | 11                                                                    | 12                                                                    | 13                                                                    | 14                                                                    | 15                                                                    | 16                                                                    | 17                                                                    | 18                                                                    | 19                                                                    | 20                                                                    | 21                                                                    | 22                                                                    | 23                                                                    | 24                                                                    |
|-------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------------------------|
| leaderboard | [1](2020/1/leaderboard/1/main.go),[2](2020/1/leaderboard/2/main.go) | [1](2020/2/leaderboard/1/main.go),[2](2020/2/leaderboard/2/main.go) | [1](2020/3/leaderboard/1/main.go),[2](2020/3/leaderboard/2/main.go) | [1](2020/4/leaderboard/1/main.go),[2](2020/4/leaderboard/2/main.go) | [1](2020/5/leaderboard/1/main.go),[2](2020/5/leaderboard/2/main.go) | [1](2020/6/leaderboard/1/main.go),[2](2020/6/leaderboard/2/main.go) | [1](2020/7/leaderboard/1/main.go),[2](2020/7/leaderboard/2/main.go) | [1](2020/8/leaderboard/1/main.go),[2](2020/8/leaderboard/2/main.go) | [1](2020/9/leaderboard/1/main.go),[2](2020/9/leaderboard/2/main.go) | [1](2020/10/leaderboard/1/main.go),[2](2020/10/leaderboard/2/main.go) | [1](2020/11/leaderboard/1/main.go),[2](2020/11/leaderboard/2/main.go) | [1](2020/12/leaderboard/1/main.go),[2](2020/12/leaderboard/2/main.go) | [1](2020/13/leaderboard/1/main.go),[2](2020/13/leaderboard/2/main.go) | [1](2020/14/leaderboard/1/main.go),[2](2020/14/leaderboard/2/main.go) | [1](2020/15/leaderboard/1/main.go),[2](2020/15/leaderboard/2/main.go) | [1](2020/16/leaderboard/1/main.go),[2](2020/16/leaderboard/2/main.go) | [1](2020/17/leaderboard/1/main.go),[2](2020/17/leaderboard/2/main.go) | [1](2020/18/leaderboard/1/main.go),[2](2020/18/leaderboard/2/main.go) | [1](2020/19/leaderboard/1/main.go),[2](2020/19/leaderboard/2/main.go) | [1](2020/20/leaderboard/1/main.go),[2](2020/20/leaderboard/2/main.go) | [1](2020/21/leaderboard/1/main.go),[2](2020/21/leaderboard/2/main.go) | [1](2020/22/leaderboard/1/main.go),[2](2020/22/leaderboard/2/main.go) | [1](2020/23/leaderboard/1/main.go),[2](2020/23/leaderboard/2/main.go) | [1](2020/24/leaderboard/1/main.go),[2](2020/24/leaderboard/2/main.go) |
| optimized   | [1](2020/1/optimized/1/main.go),[2](2020/1/optimized/2/main.go)     | [1](2020/2/optimized/1/main.go),[2](2020/2/optimized/2/main.go)     | [1](2020/3/optimized/1/main.go),[2](2020/3/optimized/2/main.go)     | [1](2020/4/optimized/1/main.go),[2](2020/4/optimized/2/main.go)     | [1](2020/5/optimized/1/main.go),[2](2020/5/optimized/2/main.go)     | [1](2020/6/optimized/1/main.go),[2](2020/6/optimized/2/main.go)     | [1](2020/7/optimized/1/main.go),[2](2020/7/optimized/2/main.go)     | [1](2020/8/optimized/1/main.go),[2](2020/8/optimized/2/main.go)     | [1](2020/9/optimized/1/main.go),[2](2020/9/optimized/2/main.go)     | [1](2020/10/optimized/1/main.go),[2](2020/10/optimized/2/main.go)     | [1](2020/11/optimized/1/main.go),[2](2020/11/optimized/2/main.go)     | [1](2020/12/optimized/1/main.go),[2](2020/12/optimized/2/main.go)     | [1](2020/13/optimized/1/main.go),[2](2020/13/optimized/2/main.go)     |                                                                       |                                                                       | [1](2020/16/optimized/1/main.go),[2](2020/16/optimized/2/main.go)     | [1](2020/17/optimized/1/main.go)                                      | [1](2020/18/optimized/1/main.go)                                      |                                                                       |                                                                       |                                                                       |                                                                       | [1](2020/23/optimized/1/main.go),[2](2020/23/optimized/2/main.go)     |                                                                       |

## 2019

|             | 1                                                                   | 2                                                                   | 3                                                                   | 4                                                                   | 5                                       | 6 | 7 | 8 | 9                                       | 10 | 11 | 12 | 13                                        | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 |
|-------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|---------------------------------------------------------------------|-----------------------------------------|---|---|---|-----------------------------------------|----|----|----|-------------------------------------------|----|----|----|----|----|----|----|----|----|----|----|
| leaderboard | [1](2019/1/leaderboard/1/main.go),[2](2019/1/leaderboard/2/main.go) | [1](2019/2/leaderboard/1/main.go),[2](2019/2/leaderboard/2/main.go) | [1](2019/3/leaderboard/1/main.go),[2](2019/3/leaderboard/2/main.go) | [1](2019/4/leaderboard/1/main.go),[2](2019/4/leaderboard/2/main.go) |                                         |   |   |   |                                         |    |    |    |                                           |    |    |    |    |    |    |    |    |    |    |    |
| optimized   |                                                                     |                                                                     |                                                                     | [1](2019/4/optimized/1/main.go),[2](2019/4/optimized/2/main.go)     | [1](2019/5/main.go),[2](2019/5/main.go) |   |   |   | [1](2019/9/main.go),[2](2019/9/main.go) |    |    |    | [1](2019/13/main.go),[2](2019/13/main.go) |    |    |    |    |    |    |    |    |    |    |    |
