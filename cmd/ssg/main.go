package main

import (
	"fmt"

	"github.com/biplobmanna/gossg/internal/parser"
)

func main() {
	markdownFiles := parser.WalkDir("./site/content")

	for _, f := range markdownFiles {
		fmt.Println(f)
	}
}
