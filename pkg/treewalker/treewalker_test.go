package treewalker

import (
	"fmt"
	"testing"
)

// Test the correctness of the fileSize function
func TestFileSize(t *testing.T) {
	var num uint = 123
	for i := 0; i < 13; i++ {
		_size := fileSize(num)
		if len(_size) != 5 {
			t.Fail()
		}
		fmt.Println(_size)
		num = num * 10
	}
}
