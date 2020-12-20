# Advent of Code

This repo contains my solutions to the [Advent of Code](https://adventofcode.com/) puzzles.

# Format

```
┌───<year>
│   └───<day>
│       ├───naughty
│       │   ├───1
│       │   └───2
│       └───nice
│           ├───1
│           └───2
└───template
```

| Directory  | Meaning                                                                                                       |
|------------|---------------------------------------------------------------------------------------------------------------|
| `year`     | The year the puzzle was released.                                                                             |
| `day`      | The day the puzzle was released.                                                                              |
| `naughty`  | My first attempt at solving the puzzle; usually as quickly as possible (in real time) to get on leaderboards. |
| `nice`     | The end result of optimizing the solution. Primarily for readability, secondarily for time/space complexity.  |
| `1`        | The first part of the puzzle.                                                                                 |
| `2`        | The second part of the puzzle.                                                                                |
| `template` | Contains a script to generate new day folders.                                                                |

# Notes

## Running

Set your working directory to the root folder, i.e. `AdventOfCode`, before running a solution.

## Organization

Each solution is self-contained so there's intentional duplication between the initial/optimized first/second solution files.

## Naming

The naughty and nice folders reference the Christmas song `Santa Claus Is Coming to Town`; AoC is also Christmas themed.
 
The naughty solutions are hard to read and inefficient, hence naughty. The nice solutions read well and are efficient, hence nice. 

# Completion

## 2020

|         | 1                                                           | 2                                                           | 3                                                           | 4                                                           | 5                                                           | 6                                                           | 7                                                           | 8                                                           | 9                                                           | 10                                                            | 11                                                            | 12                                                            | 13                                                            | 14                                                            | 15                                                            | 16                                                            | 17                                                            | 18                                                            | 19                                                            | 20                                                            | 21 | 22 | 23 | 24 | 25 |
|---------|-------------------------------------------------------------|-------------------------------------------------------------|-------------------------------------------------------------|-------------------------------------------------------------|-------------------------------------------------------------|-------------------------------------------------------------|-------------------------------------------------------------|-------------------------------------------------------------|-------------------------------------------------------------|---------------------------------------------------------------|---------------------------------------------------------------|---------------------------------------------------------------|---------------------------------------------------------------|---------------------------------------------------------------|---------------------------------------------------------------|---------------------------------------------------------------|---------------------------------------------------------------|---------------------------------------------------------------|---------------------------------------------------------------|---------------------------------------------------------------|----|----|----|----|----|
| naughty | [1](2020/1/naughty/1/main.go),[2](2020/1/naughty/2/main.go) | [1](2020/2/naughty/1/main.go),[2](2020/2/naughty/2/main.go) | [1](2020/3/naughty/1/main.go),[2](2020/3/naughty/2/main.go) | [1](2020/4/naughty/1/main.go),[2](2020/4/naughty/2/main.go) | [1](2020/5/naughty/1/main.go),[2](2020/5/naughty/2/main.go) | [1](2020/6/naughty/1/main.go),[2](2020/6/naughty/2/main.go) | [1](2020/7/naughty/1/main.go),[2](2020/7/naughty/2/main.go) | [1](2020/8/naughty/1/main.go),[2](2020/8/naughty/2/main.go) | [1](2020/9/naughty/1/main.go),[2](2020/9/naughty/2/main.go) | [1](2020/10/naughty/1/main.go),[2](2020/10/naughty/2/main.go) | [1](2020/11/naughty/1/main.go),[2](2020/11/naughty/2/main.go) | [1](2020/12/naughty/1/main.go),[2](2020/12/naughty/2/main.go) | [1](2020/13/naughty/1/main.go),[2](2020/13/naughty/2/main.go) | [1](2020/14/naughty/1/main.go),[2](2020/14/naughty/2/main.go) | [1](2020/15/naughty/1/main.go),[2](2020/15/naughty/2/main.go) | [1](2020/16/naughty/1/main.go),[2](2020/16/naughty/2/main.go) | [1](2020/17/naughty/1/main.go),[2](2020/17/naughty/2/main.go) | [1](2020/18/naughty/1/main.go),[2](2020/18/naughty/2/main.go) | [1](2020/19/naughty/1/main.go),[2](2020/19/naughty/2/main.go) | [1](2020/20/naughty/1/main.go),[2](2020/20/naughty/2/main.go) |    |    |    |    |    |
| nice    | [1](2020/1/nice/1/main.go),[2](2020/1/nice/2/main.go)       | [1](2020/2/nice/1/main.go),[2](2020/2/nice/2/main.go)       | [1](2020/3/nice/1/main.go),[2](2020/3/nice/2/main.go)       | [1](2020/4/nice/1/main.go),[2](2020/4/nice/2/main.go)       | [1](2020/5/nice/1/main.go),[2](2020/5/nice/2/main.go)       | [1](2020/6/nice/1/main.go),[2](2020/6/nice/2/main.go)       | [1](2020/7/nice/1/main.go),[2](2020/7/nice/2/main.go)       | [1](2020/8/nice/1/main.go),[2](2020/8/nice/2/main.go)       | [1](2020/9/nice/1/main.go),[2](2020/9/nice/2/main.go)       | [1](2020/10/nice/1/main.go),[2](2020/10/nice/2/main.go)       | [1](2020/11/nice/1/main.go),[2](2020/11/nice/2/main.go)       | [1](2020/12/nice/1/main.go),[2](2020/12/nice/2/main.go)       | [1](2020/13/nice/1/main.go),[2](2020/13/nice/2/main.go)       |                                                               |                                                               | [1](2020/16/nice/1/main.go),[2](2020/16/nice/2/main.go)       |                                                               |                                                               |                                                               |                                                               |    |    |    |    |    |
## 2019

|         | 1                                                           | 2                                                           | 3                                                           | 4                                                           | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 | 25 |
|---------|-------------------------------------------------------------|-------------------------------------------------------------|-------------------------------------------------------------|-------------------------------------------------------------|---|---|---|---|---|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|
| naughty | [1](2019/1/naughty/1/main.go),[2](2019/1/naughty/2/main.go) | [1](2019/2/naughty/1/main.go),[2](2019/2/naughty/2/main.go) | [1](2019/3/naughty/1/main.go),[2](2019/3/naughty/2/main.go) | [1](2019/4/naughty/1/main.go),[2](2019/4/naughty/2/main.go) |   |   |   |   |   |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |
| nice    |                                                             |                                                             |                                                             | [1](2019/4/nice/1/main.go),[2](2019/4/nice/2/main.go)       |   |   |   |   |   |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |