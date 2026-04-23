package parser

import (
	"log"
	"os"
	"path/filepath"
)

// WalkDir walks the path passed, and return a list
// of all the markdown files in the directory
func WalkDir(path string) []string {
	// 1. check that the passed path is a valid directory
	if file, err := os.Stat(path); err != nil {
		log.Fatalf("Failed to fetch the file stats of path:'%v', err:%v", path, err)
	} else {
		if !file.IsDir() {
			log.Fatalf("Path:'%v' is not a directory!", path)
		}
	}

	// 2. iterate through the directory list
	//    find all the markdown files
	//    and store the list
	markdownFiles := []string{}
	if files, err := os.ReadDir(path); err != nil {
		log.Fatalf("Failed to read directory: '%v', error: %v", path, err)
	} else {
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if ext := filepath.Ext(filepath.Join(path, f.Name())); ext == ".md" || ext == ".markdown" {
				markdownFiles = append(markdownFiles, filepath.Join(path, f.Name()))
			}
		}
	}

	return markdownFiles
}
