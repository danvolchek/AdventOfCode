package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

// CdCommand represents a shell command that ran cd.
type CdCommand struct {
	// The directory argument.
	directory string
}

// LsCommand represents a shell command that ran ls.
type LsCommand struct {
	// The output directories.
	dirs []LsDir

	// The output files.
	files []LsFile
}

// LsDir represents an output of an LsCommand that is a directory.
type LsDir struct {
	name string
}

// LsFile represents an output of an LsCommand that is a file.
type LsFile struct {
	name string
	size int
}

// Command represents a command.
type Command interface {
	// Run runs the command given the current working directory. It should return the new working directory.
	Run(cwd *Directory) *Directory
}

func parse(input string) []Command {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var commands []Command

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if _, dir, ok := strings.Cut(line, "$ cd "); ok {
			commands = append(commands, CdCommand{
				directory: dir,
			})
		} else if strings.Index(line, "$ ls") == 0 {
			var lsInstr LsCommand

			for _, lsLine := range lines[i+1:] {
				if strings.Index(lsLine, "$ ") == 0 {
					break
				}

				if _, dir, ok := strings.Cut(lsLine, "dir "); ok {
					lsInstr.dirs = append(lsInstr.dirs, LsDir{
						name: dir,
					})
				} else {
					lsInstr.files = append(lsInstr.files, LsFile{
						name: dir,
						size: lib.Int(lsLine),
					})
				}

				i += 1
			}

			commands = append(commands, lsInstr)
		} else {
			panic(line)
		}
	}

	return commands
}

// File represents a file in a Directory.
type File struct {
	name string
	size int
}

// Size returns the file's size.
func (f File) Size() int {
	return f.size
}

// Directory represents a directory.
type Directory struct {
	name string

	files          []File
	subdirectories []*Directory
	parent         *Directory

	// Only computed after a call to Size.
	size int
}

// Size returns the directory's size, which is the sum of all file sizes and subdirectory sizes.
func (d *Directory) Size() int {
	// If already computed and cached, return that.
	if d.size != 0 {
		return d.size
	}

	totalFileSize := lib.SumSlice(lib.Map(d.files, File.Size))

	// For some reason, using the method reference Directory.Size doesn't work here
	totalSubdirectorySize := lib.SumSlice(lib.Map(d.subdirectories, func(sd *Directory) int {
		return sd.Size()
	}))

	size := totalFileSize + totalSubdirectorySize

	d.size = size
	return size
}

// Run returns the new working directory based on the CdCommand's argument.
func (c CdCommand) Run(cwd *Directory) *Directory {
	switch c.directory {
	case "..":
		return cwd.parent
	case "/":
		return cwd // assume '/' only happens when the cwd is already root
	default:
		// Note: as is present in the input but not explicitly stated,
		// this assumes the directory appeared in an ls before a cd into it
		//
		// To not make this assumption, the cd should always create a new subdirectory and add it to cwd
		// And then the subdirectory part of LsCommand.Run should be removed/dirs not even parsed for ls output
		//
		// I like how this is organized better, though
		subdir := lib.Filter(cwd.subdirectories, func(d *Directory) bool {
			return d.name == c.directory
		})[0]

		return subdir
	}
}

// Run adds the directories and files seen by the ls command to the current working directory.
func (l LsCommand) Run(cwd *Directory) *Directory {
	for _, file := range l.files {
		cwd.files = append(cwd.files, File{
			name: file.name,
			size: file.size,
		})
	}

	for _, dir := range l.dirs {
		cwd.subdirectories = append(cwd.subdirectories, &Directory{
			name:   dir.name,
			parent: cwd,
		})
	}

	return cwd
}

// buildDirectoryTree builds a directory tree based on the commands and returns the root of the tree.
func buildDirectoryTree(commands []Command) *Directory {
	root := &Directory{
		name: "/",
	}

	cwd := root
	for _, command := range commands {
		cwd = command.Run(cwd)
	}

	return root
}

const maxSize = 100000

// totalSizeSmallerThanMax returns the sum of the size of every directory rooted at d that is less than maxSize.
func totalSizeSmallerThanMax(d *Directory) int {
	totalSubdirectorySize := lib.SumSlice(lib.Map(d.subdirectories, totalSizeSmallerThanMax))

	size := d.Size()

	// Include current directory because it's under the size limit
	if size <= maxSize {
		return size + totalSubdirectorySize
	}

	// Otherwise exclude it
	return totalSubdirectorySize
}

func solve(commands []Command) int {
	return totalSizeSmallerThanMax(buildDirectoryTree(commands))
}

func main() {
	solver := lib.Solver[[]Command, int]{
		ParseF: lib.ParseStringFunc(parse),
		SolveF: solve,
	}

	solver.Expect("$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k", 95437)
	solver.Verify(1583951)
}
