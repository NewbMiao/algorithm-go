package leetcode

import (
	"testing"
	"fmt"
)

func TestTwoSum(t *testing.T) {
	ds := [][]int{
		{2, 7, 11, 15},
		{0, 4, 3, 0},
		{-1, -2, -3, -4, -5},
	}
	ts := []int{
		9, 0, -8,
	}
	wants := [][]int{
		{0, 1},
		{0, 3},
		{2, 4},
	}
	for k, v := range ds {
		t.Log(fmt.Sprintf("twoSum, input arr %v, target %v, want: %v", v, ts[k], wants[k]))
		r := twoSum(v, ts[k])
		if fmt.Sprint(r) == fmt.Sprint(wants[k]) {
			t.Log("twoSum is ok")
		} else {
			t.Error("twoSum is not ok, result:", fmt.Sprintf("%v", r))
		}

	}

}
