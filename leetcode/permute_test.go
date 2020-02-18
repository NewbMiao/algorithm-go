package leetcode

import (
	"testing"
	"fmt"
)

func TestPermute(t *testing.T) {
	ds := [][]int{
		{1, 2, 3},
	}

	ws := [][][]int{
		{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 2, 1},
			{3, 1, 2},
		},
	}

	for k, v := range ds {
		t.Log(fmt.Sprintf("permute, input arr %v, want: %v", v, ws[k]))
		r := permute(v)
		if fmt.Sprint(r) == fmt.Sprint(ws[k]) {
			t.Log("permute is ok")
		} else {
			t.Error("permute is not ok, result:", r)
		}
	}
}
