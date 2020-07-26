package leetcode

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	ds := [][]int{
		{2, 3, 6, 7},
		{2, 3, 5},
	}
	ts := []int{
		7,
		8,
	}
	ws := [][][]int{
		{{7}, {2, 2, 3}},
		{{3, 5}, {2, 3, 3}, {2, 2, 2, 2}},
	}

	for k, v := range ds {
		t.Log(fmt.Sprintf("combinationSum, input arr %v, target %d,want: %v", v, ts[k], ws[k]))
		r := combinationSum(v, ts[k])
		if fmt.Sprint(r) == fmt.Sprint(ws[k]) {
			t.Log("combinationSum is ok")
		} else {
			t.Error("combinationSum is not ok, result:", fmt.Sprintf("%v", r))
		}
	}
}
