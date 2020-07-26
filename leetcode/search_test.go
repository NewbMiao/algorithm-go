package leetcode

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	ds := [][]int{
		{5, 1, 3},
		{4, 5, 6, 7, 8, 1, 2, 3},
		{0},
		{4, 5, 1, 2, 3},
	}
	ts := []int{
		0,
		5,
		1,
		1,
	}
	ws := []int{
		-1,
		1,
		-1,
		2,
	}

	for k, v := range ds {
		t.Log(fmt.Sprintf("search, input arr %v, target %d,want: %v", v, ts[k], ws[k]))
		r := search(v, ts[k])
		if fmt.Sprint(r) == fmt.Sprint(ws[k]) {
			t.Log("search is ok")
		} else {
			t.Error("search is not ok, result:", fmt.Sprintf("%v", r))
		}
	}
}
