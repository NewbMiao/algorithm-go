package leetcode

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	ds := [][]int{
		{2, 2},
	}
	ts := []int{
		2,
	}
	ws := [][]int{
		{0, 1},
	}
	for k, v := range ds {
		t.Log(fmt.Sprintf("searchRange, input arr %v, target %v, want: %v", v, ts[k], ws[k]))
		r := searchRange(v, ts[k])
		if fmt.Sprint(r) == fmt.Sprint(ws[k]) {
			t.Log("searchRange is ok")
		} else {
			t.Error("searchRange is not ok, result:", fmt.Sprintf("%v", r))
		}
	}
}
