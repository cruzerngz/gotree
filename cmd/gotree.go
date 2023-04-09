package main

import (
	"flag"
	"fmt"
	"gotree/pkg/treewalker"
)

var directory = flag.String("path", ".", "Start path")

var show_hidden_files = flag.Bool("a", false, "Show all files")
var show_file_sizes = flag.Bool("s", false, "Show file sizes")
var show_dirs_only = flag.Bool("d", false, "Show directories only")
var max_recurse_level = flag.Int("L", 1 << 31, "Maximum recursion level")

// var humanreadable = flag.Bool("h", false, "Show directory sizes in a human readable format.")

// Main entry point for gotree.
//
// gotree is my implementation of the linux `tree` command,
// which has been sort of implemented in rust as a recursive
// data structure.
func main() {

	flag.Parse() // parse command-line args

	control_args := treewalker.ControlArgs{
		ShowHiddenFiles:   *show_hidden_files,
		ShowFileSizes:     *show_file_sizes,
		ShowDirsOnly:      *show_dirs_only,
		MaxRecursionLevel: *max_recurse_level,
	}

	err := treewalker.Tree(*directory, []string{}, true, &control_args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
