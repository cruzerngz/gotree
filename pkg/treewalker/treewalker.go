package treewalker

import (
	"fmt"
	"gotree/pkg/consts"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	colour "github.com/TwiN/go-color"
)

// This function is called at the top level with an empty prefix array.
// During recursion, whatever new prefix is gets passed down to the inner function.
// If the directory is the last in the array, the last flag is set.
// This is done so that the prefix is cleared correctly.
func Tree(dirpath string, prefix []string, last bool) error {
	dirs, err := os.ReadDir(dirpath)
	if err != nil {
		return err
	}

	last_dir := lastDir(dirs)
	last_index := len(dirs) - 1

	// print the top level name
	if last {
		setLastElement(prefix, consts.L_BLOCK)
	} else {
		setLastElement(prefix, consts.T_BLOCK)
	}
	if len(prefix) == 0 {
		displayLine(prefix, dirColour(dirpath)) // dirs are displayed in green
	} else {
		displayLine(prefix, dirColour(filepath.Base(dirpath))) // dirs are displayed in green
	}

	// some prefix manip
	setLastElement(prefix, consts.V_BLOCK)

	if last {
		setLastElement(prefix, consts.E_BLOCK)
	}

	prefix = append(prefix, consts.T_BLOCK) // push block

	for idx := 0; idx < len(dirs); idx++ {

		if idx == last_index && last_index != last_dir {
			setLastElement(prefix, consts.L_BLOCK)
		}

		if dirs[idx].IsDir() {

			if idx == last_dir && last_dir == last_index {
				Tree(filepath.Join(dirpath, dirs[idx].Name()), prefix, true)
			} else {
				Tree(filepath.Join(dirpath, dirs[idx].Name()), prefix, false)
				setLastElement(prefix, consts.T_BLOCK) // revert the change in last element
			}

		} else {
			displayLine(prefix, dirs[idx].Name())
		}
	}

	return nil
}

func dirColour(dirpath string) string {
	return colour.InBold(colour.InBlue(dirpath))
}

// Helper function to set the last element of an array.
func setLastElement(string_arr []string, item string) {
	if len(string_arr) > 0 {
		string_arr[len(string_arr)-1] = item
	}
}

// Helper function to print a single line to stdout.
// The prefix are the "branches" in the tree; the branch is the new extension,
// and the contents are the names of the files/directories.
func displayLine(prefix []string, contents string) {
	fmt.Printf("%s%s\n", strings.Join(prefix, ""), contents)
}

// Returns the index of the last directory in the array.
// Returns the length of the array if no directories are found.
func lastDir(dirs []fs.DirEntry) int {
	index := len(dirs)

	for _index, subdir := range dirs {
		if subdir.IsDir() {
			index = _index
		}
	}
	return index
}
