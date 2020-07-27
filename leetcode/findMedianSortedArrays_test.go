package leetcode

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	ds := [][2][]int{
		{{1, 3}, {2, 5, 7}},
		{{1, 2}, {3, 4}},
	}
	ws := []float64{
		3,
		2.5,
	}
	for k, v := range ds {
		t.Log(fmt.Sprintf("FindMedianSortedArrays, input arr A %v, B %v, want: %v", v[0], v[1], ws[k]))
		r := findMedianSortedArrays(v[0], v[1])
		if r == ws[k] {
			t.Log("FindMedianSortedArrays is ok")
		} else {
			t.Error("FindMedianSortedArrays is not ok, result:", r)
		}
	}
}
