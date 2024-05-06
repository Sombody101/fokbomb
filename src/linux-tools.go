package main

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

// ? Get a random directory
func getRandDir(dir string) string {
	var selectedDir string

	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {

		// Skip device files
		if strings.HasPrefix(path, "/dev") {
			return filepath.SkipDir
		}

		if rand.Intn(5) == 4 {
			verbose("grd", "Checking dir:", path)
			if __checkDir(path) {
				verbose("grd", "Using dir:", path)

				// 1/5 odds, and the directory must be writable
				selectedDir = path
				return filepath.SkipAll
			}
		}

		return nil
	})

	return selectedDir
}

func getRandFile() string {
	var dirs = []string{
		"/bin",
		"/dev",
		"/var",
	}

	for _, dir := range dirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			if os.IsNotExist(err) {
				continue // Ignore non-existent directory
			}
			panic(fmt.Errorf("error reading directory: %w", err)) // Unexpected error
		}

		if len(files) > 0 {
			return files[rand.Intn(len(files))].Name() // Choose a random file
		}
	}

	return ""
}

func __checkDir(path string) bool {
	path = filepath.Join(path, "test.txt"+fmt.Sprint(rand.Int()))
	_, err := os.Create(path)
	return err == nil
}
