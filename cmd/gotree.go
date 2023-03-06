package main

import (
	"flag"
	"fmt"
	"gotree/pkg/treewalker"
)

var directory = flag.String("path", ".", "Start path")

// var humanreadable = flag.Bool("h", false, "Show directory sizes in a human readable format.")

// Main entry point for gotree.
//
// gotree is my implementation of the linux `tree` command,
// which has been sort of implemented in rust as a recursive
// data structure.
func main() {

	flag.Parse() // parse command-line args

	err := treewalker.Tree(*directory, []string{}, true)
	if err != nil {
		fmt.Println(err)
		return
	}

}
