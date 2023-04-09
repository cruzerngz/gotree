package treewalker

import (
	"fmt"
	"gotree/pkg/consts"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"strings"

	colour "github.com/TwiN/go-color"
)

// Array-map of sizes
var SIZE_SUFFIX = [...]uint8{
	'B', // first element not used
	'K',
	'M',
	'G',
	'T',
	'P',
}

// This function is called at the top level with an empty prefix array.
// During recursion, whatever new prefix is gets passed down to the inner function.
// If the directory is the last in the array, the last flag is set.
// This is done so that the prefix is cleared correctly.
func Tree(dirpath string, prefix []string, last bool, control_args *ControlArgs) error {
	if len(prefix) > control_args.MaxRecursionLevel {
		return nil
	}

	dirs, err := os.ReadDir(dirpath)
	if err != nil {
		return err
	}

	// filtering from control_args here
	var filtered []fs.DirEntry = []fs.DirEntry{}
	for idx := range dirs {
		to_append := true

		if !control_args.ShowHiddenFiles {
			if dirs[idx].Name()[0] == '.' {
				to_append = false
			}
		}

		if control_args.ShowDirsOnly {
			if !dirs[idx].IsDir() {
				to_append = false
			}
		}

		if to_append {
			filtered = append(filtered, dirs[idx])
		}
	}

	dirs = filtered

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
				err := Tree(filepath.Join(dirpath, dirs[idx].Name()), prefix, true, control_args)
				if err != nil {
					return err
				}
			} else {
				err := Tree(filepath.Join(dirpath, dirs[idx].Name()), prefix, false, control_args)
				if err != nil {
					return err
				}
				setLastElement(prefix, consts.T_BLOCK) // revert the change in last element
			}

		} else {
			// special case for child files
			if len(prefix)-1 < control_args.MaxRecursionLevel {
				_file, err := dirs[idx].Info()
				if err != nil {
					return err
				}
				displayLineNew(prefix, _file, control_args)

				// displayLine(prefix, dirs[idx].Name())
			}
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

func displayLineNew(prefix []string, file fs.FileInfo, control_args *ControlArgs) {

	// _file, err := file.Info()
	// if err != nil {
	// 	return er
	// }

	prefix_str := strings.Join(prefix, "")

	if control_args.ShowFileSizes {
		prefix_str += fmt.Sprintf("[%s]  ", fileSize(uint(file.Size())))
	}

	var name_str string
	if file.IsDir() {
		name_str = colour.InBlue(file.Name())
	} else {
		name_str = file.Name()
	}

	fmt.Printf("%s%s\n", prefix_str, name_str)
}

// Formats the size in bytes to a left padded 5-char long string.
func fileSize(size_bytes uint) string {
	if size_bytes < 1000 {
		return fmt.Sprintf("%5d", size_bytes)
	}

	// actually 1 less than num_digits
	powers_of_ten := uint(math.Log10(float64(size_bytes)))
	// get the nearest byte size
	suffix := SIZE_SUFFIX[powers_of_ten/3]

	nearest_size := powers_of_ten - powers_of_ten%3

	result := float32(size_bytes) / float32(math.Pow10(int(nearest_size)))
	result_num_str := fmt.Sprintf("%.3g", result)
	if len(result_num_str) > 3 {
		result_num_str = result_num_str[:3]
		result_num_str = strings.TrimSuffix(result_num_str, ".")
	}

	return fmt.Sprintf("%4s%c", result_num_str, suffix)
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
