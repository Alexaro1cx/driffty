package utils

import (
	"os"
	"path/filepath"
)

func Glob(path, pattern string) []string {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if matched, _ := filepath.Match(pattern, info.Name()); matched {
				files = append(files, path)
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}
