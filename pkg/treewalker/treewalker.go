package treewalker

import (
	"io/ioutil"
	"strings"
)

func ShowTree(path string) string {

	lines := walk_tree(path)

	return strings.Join(lines, "\n")
}

func walk_tree(dirpath string) []string {
	println("asdasdas")
	files, err := ioutil.ReadDir(".");



}
