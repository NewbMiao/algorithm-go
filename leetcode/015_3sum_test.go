package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	ds := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 0, 0},
	}
	wants := [][][]int{
		{{-1, -1, 2}, {-1, 0, 1}},
		{{0, 0, 0}},
	}
	for k, v := range ds {
		t.Logf("threeSum, input arr %v, want: %v", v, wants[k])
		r := threeSum(v)
		if fmt.Sprint(r) == fmt.Sprint(wants[k]) {
			t.Log("threeSum is ok")
		} else {
			t.Errorf("threeSum is not ok, result:%v", r)
		}
	}
}
