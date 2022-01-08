package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Solution struct {
	Year, Day int

	Leaderboard bool
}

func (s Solution) Path(root string) (string, bool) {
	year, day, solType := strconv.Itoa(s.Year), strconv.Itoa(s.Day), s.typeName()

	solPath := filepath.Join(root, year, day, solType)

	return solPath, Exists(solPath)
}

func (s Solution) PartOne(root string) (string, bool) {
	return s.partPath(root, "1")
}

func (s Solution) PartTwo(root string) (string, bool) {
	return s.partPath(root, "2")
}

func (s Solution) partPath(root, part string) (string, bool) {
	base, ok := s.Path(root)

	partPath := filepath.Join(base, part, "main.go")

	return partPath, ok && Exists(partPath)
}

func (s Solution) String() string {
	solType := s.typeName()
	solType = string(solType[0]-32) + solType[1:]
	return fmt.Sprintf("{Year: %v, Day: %v, Type: %v}", s.Year, s.Day, solType)
}

func (s Solution) typeName() string {
	solutionType := "optimized"
	if s.Leaderboard {
		solutionType = "leaderboard"
	}

	return solutionType
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		panic(err)
	}

	return true
}
