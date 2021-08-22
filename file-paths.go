package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	pln := fmt.Println

	p := filepath.Join("dir1", "dir2", "filename")
	pln("p:",p)

	pln(filepath.Join("dir1//","filename"))
	pln(filepath.Join("dir1/../dir1","filename"))

	pln("Dir(p):", filepath.Dir(p))
	pln("Base(p):", filepath.Base(p))

	pln(filepath.IsAbs("dir/file"))
	pln(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	pln(ext)

	pln(strings.TrimSuffix(filename, ext))
	
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	pln(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	pln(rel)
}
