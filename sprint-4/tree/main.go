package main

/*
Course `Web services on Go`, week 1, homework, `tree` program.
See: week_01\materials.zip\week_1\99_hw\tree

mkdir -p week01_homework/tree
pushd week01_homework/tree
go mod init tree
go mod tidy
pushd ..
go work init
go work use ./tree/
go vet tree
gofmt -w tree
go test -v tree
go run tree . -f
go run tree ./tree/testdata
cd tree && docker build -t mailgo_hw1 .

https://en.wikipedia.org/wiki/Tree_(command)
https://mama.indstate.edu/users/ice/tree/
https://stackoverflow.com/questions/32151776/visualize-tree-in-bash-like-the-output-of-unix-tree

*/

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

/*
	Example output:

	├───project
	│	└───gopher.png (70372b)
	├───static
	│	├───a_lorem
	│	│	├───dolor.txt (empty)
	│	├───css
	│	│	└───body.css (28b)
	...
	│			└───gopher.png (70372b)

	- path should point to a directory,
	- output all dir items in sorted order, w/o distinction file/dir
	- last element prefix is `└───`
	- other elements prefix is `├───`
	- nested elements aligned with one tab `	` for each level
*/

const (
	EOL             = "\n"
	BRANCHING_TRUNK = "├───"
	LAST_BRANCH     = "└───"
	TRUNC_TAB       = "│\t"
	LAST_TAB        = "\t"
	EMPTY_FILE      = "empty"
	ROOT_PREFIX     = ""
	LEVEL_SPACE     = "    "
	LEVEL_PIPE      = "|   "

	USE_RECURSION_ENV_KEY = "RECURSIVE_TREE"
	USE_RECURSION_ENV_VAL = "YES"
)

func main() {
	// This code is given
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage: go run main.go . [-f]")
	}

	out := os.Stdout
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

// dirTree: `tree` program implementation, top-level function, signature is fixed.
// Write `path` dir listing to `out`. If `prinFiles` is set, files is listed along with directories.
func dirTree(out io.Writer, path string, printFiles bool) error {
	// Function to implement, signature is given, don't touch it.

	err := recursiveLookupDir(out, path, printFiles, ROOT_PREFIX)
	if err != nil {
		return err
	}

	return nil
}

func recursiveLookupDir(out io.Writer, path string, printFile bool, prefix string) error {
	dirEntres, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf(
			"Error Read Dir: %s \n Error: %w", path, err,
		)
	}

	if !printFile {
		dirEntres = excludeFiles(dirEntres)
	}

	sort.Slice(
		dirEntres,
		func(i, j int) bool {
			return dirEntres[i].Name() < dirEntres[j].Name()
		},
	)

	lastEntry := len(dirEntres) - 1

	for i, obj := range dirEntres {
		objPath := filepath.Join(path, obj.Name())
		objName := obj.Name()
		isLastEntry := i == lastEntry

		printTrunk(out, isLastEntry, prefix)
		levelPrefix := getLevelPrefix(isLastEntry, prefix)

		if obj.IsDir() {
			printlnDir(out, objName)
			recursiveLookupDir(out, objPath, printFile, levelPrefix)
			continue
		}

		printlnFile(out, obj, objName)

	}
	return nil
}

func excludeFiles(dirEntres []os.DirEntry) []os.DirEntry {
	var objs []os.DirEntry
	for _, entry := range dirEntres {
		if entry.IsDir() {
			objs = append(objs, entry)
		}
	}
	return objs
}

func getLevelPrefix(isLastEntry bool, prefix string) string {
	if isLastEntry {
		return prefix + LAST_TAB
	}
	return prefix + TRUNC_TAB
}

func printTrunk(out io.Writer, isLastEntry bool, levelPrefix string) {
	if isLastEntry {
		out.Write([]byte(levelPrefix + LAST_BRANCH))
	} else {
		out.Write([]byte(levelPrefix + BRANCHING_TRUNK))
	}
}

func printlnDir(out io.Writer, dirName string) {
	out.Write([]byte(dirName))
	out.Write([]byte(EOL))
}

func printlnFile(out io.Writer, file os.DirEntry, fileName string) error {
	out.Write([]byte(fileName))
	fileinfo, err := file.Info()
	if err != nil {
		return fmt.Errorf(
			"Error Get Stat File: %s \n Error: %w", fileName, err,
		)
	}
	sizeFile := fileinfo.Size()
	if sizeFile == 0 {
		fmt.Fprintf(out, " (%s)", EMPTY_FILE)
	} else {
		fmt.Fprintf(out, " (%db)", sizeFile)
	}
	out.Write([]byte(EOL))

	return nil
}
