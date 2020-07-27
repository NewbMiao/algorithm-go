package stack

import (
	"fmt"
	"testing"
)

func TestGetNum(t *testing.T) {
	arrs := [][]int{
		{9, 7, 4, 6, 3, 1, 7, 7, 5, 6, 5, 8, 0, 0, 6, 6, 2, 3, 2, 3, 0, 0, 6, 2, 9, 7, 0, 3, 9, 4},
		{9, 6, 0, 7, 7, 3, 6, 7, 9, 3, 7, 2, 1, 9, 9, 1, 2, 0, 0, 6, 3, 2, 6, 4, 9, 5, 6, 2, 1, 3},
		{4, 8, 9, 9, 2, 8, 3, 0, 4, 5, 5, 8, 4, 6, 7, 2, 7, 9, 2, 7, 0, 5, 6, 1, 0, 5, 5, 0, 1, 4},
		{4, 3, 1, 3, 5, 8, 2, 2, 8, 2, 3, 7, 5, 6, 8, 5, 5, 2, 5, 5, 4, 5, 8, 4, 8, 5, 1, 9, 6, 4},
	}
	num := 5
	wants := []int{86, 82, 99, 114}

	for k := range arrs {
		r := getNum(arrs[k], num)
		t.Log(fmt.Sprintf("GetNum, input %v, want: %v", arrs[k], wants[k]))

		if r == wants[k] {
			t.Log("GetNum is ok")
		} else {
			t.Error("GetNum is not ok, result:", r)
		}
	}
}
