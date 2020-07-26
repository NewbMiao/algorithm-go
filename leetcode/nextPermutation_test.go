package leetcode

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	ds := [][][]int{
		{{1, 2, 3}, {1, 3, 2}},
		{{3, 2, 1}, {1, 2, 3}},
		{{1, 1, 5}, {1, 5, 1}},
		{{1, 3, 2}, {2, 1, 3}},
	}

	for _, v := range ds {
		t.Log(fmt.Sprintf("nextPermutation, input arr %v, want: %v", v[0], v[1]))
		r := nextPermutation(v[0])
		if fmt.Sprint(r) == fmt.Sprint(v[1]) {
			t.Log("nextPermutation is ok")
		} else {
			t.Error("nextPermutation is not ok, result:", fmt.Sprintf("%v", r))
		}
	}
}
