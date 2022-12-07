package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strconv"
	"strings"
)

type cdInstr struct {
	arg string
}

type lsEntry struct {
	dir  bool
	name string
	size int
}

type lsInstr struct {
	entries []lsEntry
}

type instruction struct {
	cd *cdInstr
	ls *lsInstr
}

func parse(line string) []instruction {
	lines := strings.Split(line, "\n")

	var instructions []instruction

	i := 0
	for i < len(lines) {
		if strings.Index(lines[i], "$ cd ") == 0 {
			instructions = append(instructions, instruction{
				cd: &cdInstr{arg: lines[i][len("$ cd "):]},
				ls: nil,
			})
			i += 1
		} else if strings.Index(lines[i], "$ ls") == 0 {
			var lsEntries []lsEntry

			j := i + 1
			for {
				ints := lib.Ints(lines[j])
				if len(ints) != 0 {
					lsEntries = append(lsEntries, lsEntry{
						name: strings.TrimSpace(strings.ReplaceAll(lines[j], strconv.Itoa(ints[0]), "")),
						size: ints[0],
					})
				} else if strings.Index(lines[j], "dir") == 0 {
					lsEntries = append(lsEntries, lsEntry{
						dir:  true,
						name: lines[j][len("dir "):],
					})
				} else {
					break
				}
				j += 1

				if j == len(lines) {
					break
				}
			}

			instructions = append(instructions, instruction{ls: &lsInstr{lsEntries}})
			i = j
		} else if lines[i] != "" {
			panic(lines[i])
		} else {
			break
		}
	}

	return instructions
}

type file struct {
	name string
	size int
}

func (f file) Size() int {
	return f.size
}

type directory struct {
	name string

	files       []file
	directories []*directory
	parent      *directory

	size int
}

func (d *directory) Size() int {
	if d.size != 0 {
		return d.size
	}

	mySize := lib.SumSlice(lib.Map(d.files, file.Size)) + lib.SumSlice(lib.Map(d.directories, func(dd *directory) int {
		return dd.Size()
	}))

	d.size = mySize

	return mySize
}

func getTotalSize(d *directory) int {
	childSizes := lib.SumSlice(lib.Map(d.directories, getTotalSize))

	mySize := d.Size()

	if mySize <= 100000 {
		return mySize + childSizes
	}

	return childSizes
}

func solve(lines []instruction) int {
	var root = &directory{
		name: "/",
	}

	var cwd = root

outer:
	for _, instr := range lines {
		if instr.cd != nil {
			if instr.cd.arg == ".." {
				cwd = cwd.parent
				continue outer
			}

			for _, dirs := range cwd.directories {
				if dirs.name == instr.cd.arg {
					cwd = dirs
					continue outer
				}
			}

			if cwd != root {
				panic("bad")
			}
		} else {
			for _, entry := range instr.ls.entries {
				if !entry.dir {
					cwd.files = append(cwd.files, file{
						name: entry.name,
						size: entry.size,
					})
				} else {
					newDir := &directory{
						name:   entry.name,
						parent: cwd,
					}

					cwd.directories = append(cwd.directories, newDir)
				}
			}
		}
	}

	ret := getTotalSize(root)

	return ret
}

func main() {
	solver := lib.Solver[[]instruction, int]{
		ParseF: lib.ParseStringFunc(parse),
		SolveF: solve,
	}

	solver.Expect("$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k", 95437)
	solver.Verify(1583951)
}
