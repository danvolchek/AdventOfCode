package main

import (
	"bytes"
	"fmt"
	"github.com/danvolchek/AdventOfCode/cmd/internal"
	"os"
	"os/exec"
)

func main() {
	years := internal.GetLocalSolutionInfo(".")

	for _, year := range years {
		if !year.Exists() {
			continue
		}

		if year.Name == "2022" {
			continue
		}

		for _, day := range year.Days {
			if !day.Exists() {
				continue
			}

			for _, typ := range day.Types {
				if !typ.Exists() {
					continue
				}

				for _, part := range typ.Parts {
					if !part.Exists() {
						continue
					}

					fmt.Println("Checking", part.Path)
					if err := check(part.Path); err != nil {
						fmt.Printf("%v: failed: %s\n", part.Path, err)
					}
				}
			}
		}
	}
}

func check(path string) error {
	tmp, err := os.CreateTemp("", "*")
	if err != nil {
		return err
	}

	defer os.Remove(tmp.Name())

	cmd := fmt.Sprintf("go build -o %s %s && %s", tmp.Name(), path, tmp.Name())

	command := exec.Command("/bin/sh", "-c", cmd)

	result, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, result)
	}

	if bytes.Contains(result, []byte("fail")) {
		return fmt.Errorf("%s", result)
	}

	return nil
}
