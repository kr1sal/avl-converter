package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// IterFiles recursively iterates through the files in a directory and its subdirectories, returning a slice of file paths.
// It takes the following parameters:
// path: The root directory path to start the iteration.
// name: The name of the files to match (optional).
// extension: The file extension to match (optional).
// depth: The maximum depth to traverse (-1 for unlimited).
func IterFiles(path string, name string, extension string, depth int) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var filepaths []string
	for _, entry := range entries {
		if entry.IsDir() && (depth == -1 || depth > 0) {
			subPaths, err := IterFiles(filepath.Join(path, entry.Name()), name, extension, depth-1)
			if err != nil {
				return nil, err
			}
			filepaths = append(filepaths, subPaths...)
		} else if entry.Type().IsRegular() {
			if (name == "" || entry.Name() == name) && (extension == "" || strings.EqualFold(filepath.Ext(entry.Name()), extension)) {
				filepaths = append(filepaths, filepath.Join(path, entry.Name()))
			}
		} else {
			//fmt.Printf("Skipped entry %v", filepath.Join(path, entry.Name()))
		}
	}
	return filepaths, nil
}
