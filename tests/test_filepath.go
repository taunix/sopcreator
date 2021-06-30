package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	pages, _ := filepath.Glob("./pkg/templates/*.page.tmpl")
	fmt.Println("pages", pages)
}
