package main

import (
	"bytes"
	"os"
	tree "pets/tree/utils"
)

func main() {
	// out := os.Stdout
	out := new(bytes.Buffer)
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out *bytes.Buffer, path string, printFiles bool) error {
	err := tree.DirTree(out, path, printFiles, true, 0)
	if err != nil {
		return err
	}

	return nil
}
